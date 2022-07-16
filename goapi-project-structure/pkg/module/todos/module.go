package todos

import (
	"goapi-project-structure/pkg/app"
	"goapi-project-structure/pkg/module/todos/core/services"
	"goapi-project-structure/pkg/module/todos/handler"
	"goapi-project-structure/pkg/module/todos/repository"
)

func Init(ctx *app.Context) {
	// สร้าง dependencies ทั้งหมด
	repo := repository.NewTodoRepositoryDB(ctx.DB)
	serv := services.NewTodoService(repo)
	h := handler.NewTodoHandler(serv)

	todos := ctx.Router.Group(ctx.Config.App.BaseUrl + "/todos")

	todos.Post("", h.CreateTodo)
	todos.Get("", h.ListTodo)
	todos.Get("/:id", h.GetTodo)
	todos.Patch("/:id", h.UpdateTodoStatus)
	todos.Delete("/:id", h.DeleteTodo)
}
