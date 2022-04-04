package app

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/muhfaris/skeleton-hexagonal/transport/http/graphql/resolver"
)

type GraphqlApp interface {
	GetSchema() string
	GetResolvers() *resolver.Resolver
	GraphQLHandler(h http.Handler) fiber.Handler
}
