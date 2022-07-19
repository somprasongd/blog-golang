package todo

import (
	"goapi-project-structure/pkg/app"
	"goapi-project-structure/pkg/module/todo/core/service"
	"goapi-project-structure/pkg/module/todo/handler"
	"goapi-project-structure/pkg/module/todo/repository"
	"goapi-project-structure/pkg/util"
)

func Init(ctx *app.Context) {
	// สร้าง dependencies ทั้งหมด
	repo := repository.NewTodoRepositoryDB(ctx.DB)
	serv := service.NewTodoService(repo)
	h := handler.NewTodoHandler(serv)

	todos := ctx.Router.Group(ctx.Config.App.BaseUrl + "/todos")

	todos.Post("", util.WrapFiberHandler(h.CreateTodo))
	todos.Get("", util.WrapFiberHandler(h.ListTodo))
	todos.Get("/:id", util.WrapFiberHandler(h.GetTodo))
	todos.Patch("/:id", util.WrapFiberHandler(h.UpdateTodoStatus))
	todos.Delete("/:id", util.WrapFiberHandler(h.DeleteTodo))
}
