package graphql

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/muhfaris/skeleton-hexagonal/internal/config"
	"github.com/muhfaris/skeleton-hexagonal/transport/http/graphql/handlers"
	"github.com/muhfaris/skeleton-hexagonal/transport/http/graphql/resolver"
	"github.com/muhfaris/skeleton-hexagonal/transport/http/graphql/schema"
)

var cApp *config.ConfigApp

type Graphql struct {
	port      int
	schema    string
	router    *fiber.App
	resolvers *resolver.Resolver
}

func init() {
	cApp = config.CreateConfigApp()
}

func NewGraphql(port int) *Graphql {
	r := fiber.New()

	// initialize default config
	r.Use(cors.New())
	r.Use(recover.New())
	r.Use(requestid.New())
	r.Use(logger.New())

	// graceful shutdown when interupt signal detected
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		_ = r.Shutdown()
	}()

	rslvr := resolver.NewResolver(cApp.DB)

	g := &Graphql{
		port:      port,
		schema:    schema.LoadGraphqlSchemas(),
		router:    r,
		resolvers: rslvr,
	}

	handlers.InitRouter(g, r)
	return g
}

func (r *Graphql) Serve() {
	port := fmt.Sprintf(":%d", r.port)
	if err := r.router.Listen(port); err != nil {
		panic(err)
	}
}

func (r *Graphql) GetResolvers() *resolver.Resolver {
	return r.resolvers
}

func (r *Graphql) GetSchema() string {
	return r.schema
}

// GraphQLHandler is to wrap graphql server relay
func (app *Graphql) GraphQLHandler(h http.Handler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// TODO: middleware authentication / authorization

		// server graphql relay
		return c.Next()
	}
}
