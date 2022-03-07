package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// HealthCheckHandler is handler for health check
func HealthCheckHandler(rest interface{}) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		return ctx.SendStatus(http.StatusOK)
	}
}
