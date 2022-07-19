package todo

import (
	"goapi-testing/pkg/app"
	"goapi-testing/pkg/module/todo/core/ports"
	"goapi-testing/pkg/module/todo/core/service"
	"goapi-testing/pkg/module/todo/handler"
	"goapi-testing/pkg/module/todo/repository"
	"goapi-testing/pkg/util"

	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	BaseURL     string
	Router      *fiber.App
	TodoService ports.TodoService
}

func Init(ctx *app.Context) {
	// สร้าง dependencies ทั้งหมด
	repo := repository.NewTodoRepositoryDB(ctx.DB)
	svc := service.NewTodoService(repo)

	cfg := RouteConfig{
		BaseURL:     ctx.Config.App.BaseUrl,
		Router:      ctx.Router,
		TodoService: svc,
	}

	SetupRoutes(cfg)
	// h := handler.NewTodoHandler(serv)

	// todos := ctx.Router.Group(ctx.Config.App.BaseUrl + "/todos")

	// todos.Post("", util.WrapFiberHandler(h.CreateTodo))
	// todos.Get("", util.WrapFiberHandler(h.ListTodo))
	// todos.Get("/:id", util.WrapFiberHandler(h.GetTodo))
	// todos.Patch("/:id", util.WrapFiberHandler(h.UpdateTodoStatus))
	// todos.Delete("/:id", util.WrapFiberHandler(h.DeleteTodo))
}

func SetupRoutes(cfg RouteConfig) {
	h := handler.NewTodoHandler(cfg.TodoService)

	todos := cfg.Router.Group(cfg.BaseURL + "/todos")

	todos.Post("", util.WrapFiberHandler(h.CreateTodo))
	todos.Get("", util.WrapFiberHandler(h.ListTodo))
	todos.Get("/:id", util.WrapFiberHandler(h.GetTodo))
	todos.Patch("/:id", util.WrapFiberHandler(h.UpdateTodoStatus))
	todos.Delete("/:id", util.WrapFiberHandler(h.DeleteTodo))
}
