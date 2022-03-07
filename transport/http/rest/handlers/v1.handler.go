package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhfaris/skeleton-hexagonal/internal/app"
)

func V1Handler(rest app.RestApp, r *fiber.App) {
	v1Group := r.Group("/v1")
	v1Group.Post("/login", SignInHandler(rest))
	v1Group.Post("/signup", SignUpHandler(rest))
}
