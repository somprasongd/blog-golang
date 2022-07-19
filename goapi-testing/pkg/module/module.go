package module

import (
	"goapi-testing/pkg/app"
	"goapi-testing/pkg/module/todo"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func Init(ctx *app.Context) {
	todo.Init(ctx)

	ctx.Router.Get("/healthz", healthCheckHandler)
}

func healthCheckHandler(c *fiber.Ctx) error {
	return c.SendStatus(http.StatusOK)
}
