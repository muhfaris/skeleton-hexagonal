package cmd

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/muhfaris/adsrobot/gateway/handler"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func serve(ctx context.Context, app *handler.App) {
	cors := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "OPTIONS", "PATCH", "DELETE"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	router := mux.NewRouter()
	h := handler.NewHandler(app)
	// Internal use
	api := router.PathPrefix("/api").Subrouter()
	v1 := api.PathPrefix("/v1").Subrouter()

	v1.Use(app.Middleware.InternalAuthenticationMiddleware)
	h.InternalHandler(app, v1)

	s := &http.Server{
		Addr:        fmt.Sprintf(":%d", app.Config.Port),
		Handler:     cors(router),
		ReadTimeout: time.Duration(app.Config.HTTP.ReadTimeout) * time.Minute,
	}

	done := make(chan struct{})
	go func() {
		<-ctx.Done()
		if err := s.Shutdown(context.Background()); err != nil {
			logrus.Error(err)
		}
		close(done)
	}()

	logrus.Infof("serving api at http://127.0.0.1:%d", app.Config.Port)
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		logrus.Error(err)
	}
	<-done
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "serves the api",
	RunE: func(cmd *cobra.Command, args []string) error {
		// init config
		app := handler.NewApp(
			dbPool,
			cachePool,
		)

		ctx, cancel := context.WithCancel(context.Background())
		go func() {
			ch := make(chan os.Signal, 1)
			signal.Notify(ch, os.Interrupt)
			<-ch
			logrus.Info("signal caught. shutting down...")
			cancel()
		}()

		var wg sync.WaitGroup

		wg.Add(1)
		go func() {
			defer wg.Done()
			defer cancel()
			serve(ctx, app)
		}()

		wg.Wait()
		return nil
	},
}
