package handlers

import (
	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/muhfaris/skeleton-hexagonal/internal/app"
)

func InitRouter(graph app.GraphqlApp, r *fiber.App) {
	schema := graph.GetSchema()
	resolver := graph.GetResolvers()
	s := graphql.MustParseSchema(schema, resolver)

	graphqlGroup := r.Group("")
	graphqlGroup.Post("/graphql", adaptor.HTTPHandler(&relay.Handler{
		Schema: s,
	}))
}
