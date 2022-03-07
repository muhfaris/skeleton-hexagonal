package rest

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/muhfaris/skeleton-hexagonal/internal/app"
	"github.com/muhfaris/skeleton-hexagonal/internal/config"
	"github.com/muhfaris/skeleton-hexagonal/transport/http/rest/handlers"
)

var cApp *config.ConfigApp

type Rest struct {
	port     int
	router   *fiber.App
	services *app.ServiceApp
}

func init() {
	cApp = config.CreateConfigApp()
}

func NewRest(port int) *Rest {
	r := fiber.New()

	// initialize default config
	r.Use(cors.New())
	r.Use(recover.New())
	r.Use(requestid.New())

	// graceful shutdown when interupt signal detected
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		_ = r.Shutdown()
	}()

	servicesApp := app.NewServiceApp(cApp.DB)

	rest := &Rest{
		port:     port,
		router:   r,
		services: servicesApp,
	}

	handlers.InitRouter(rest, r)
	return rest
}

func (r *Rest) Serve() {
	port := fmt.Sprintf(":%d", r.port)
	if err := r.router.Listen(port); err != nil {
		panic(err)
	}
}
func (r *Rest) GetServices() *app.ServiceApp {
	return r.services
}
