package handlers

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/muhfaris/skeleton-hexagonal/internal/app"
	"github.com/muhfaris/skeleton-hexagonal/internal/respond"
	"github.com/muhfaris/skeleton-hexagonal/transport/structures"
)

func SignInHandler(rest app.RestApp) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		params := new(structures.LoginRead)

		if err := ctx.BodyParser(params); err != nil {
			return respond.Fail(
				ctx,
				respond.WithStatus(http.StatusBadRequest),
				respond.WithError(err),
			)
		}

		services := rest.GetServices()
		if services == nil {
			return respond.Fail(
				ctx,
				respond.WithStatus(http.StatusInternalServerError),
				respond.WithError(fmt.Errorf("services signin not started")),
			)
		}

		data, err := services.UserPublicService.Login(ctx.UserContext(), params)
		if err != nil {
			return respond.Fail(
				ctx,
				respond.WithStatus(http.StatusInternalServerError),
				respond.WithError(err),
			)
		}

		return respond.Success(ctx, respond.WithData(data))
	}
}
