package module

import (
	"goapi-project-structure/pkg/app"
	"goapi-project-structure/pkg/module/todos"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func Init(ctx *app.Context) {
	todos.Init(ctx)

	ctx.Router.Get("/healthz", healthCheckHandler)
}

func healthCheckHandler(c *fiber.Ctx) error {
	return c.SendStatus(http.StatusOK)
}
