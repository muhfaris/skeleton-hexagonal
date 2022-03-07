package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhfaris/skeleton-hexagonal/internal/app"
)

func InitRouter(rest app.RestApp, r *fiber.App) {
	healthcheckGroup := r.Group("/health")
	healthcheckGroup.Get("", HealthCheckHandler(rest))

	// initialize v1 handler
	V1Handler(rest, r)
}
