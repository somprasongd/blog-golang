package users

import (
	"goapi/pkg/app"
	"goapi/pkg/module/users/core/ports"
	"goapi/pkg/module/users/core/service"
	"goapi/pkg/module/users/handler"
	"goapi/pkg/module/users/repository"
	"goapi/pkg/util"

	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	BaseURL     string
	Router      *fiber.App
	UserService ports.UserService
}

func Init(ctx *app.Context) {
	// สร้าง dependencies ทั้งหมด
	repo := repository.NewUserRepositoryDB(ctx.DB.DB)
	svc := service.NewUserService(repo)

	cfg := RouteConfig{
		BaseURL:     ctx.Config.App.BaseUrl,
		Router:      ctx.Router,
		UserService: svc,
	}

	SetupRoutes(cfg)
}

func SetupRoutes(cfg RouteConfig) {
	h := handler.NewUserHandler(cfg.UserService)

	userss := cfg.Router.Group(cfg.BaseURL + "/users")

	userss.Post("", util.WrapFiberHandler(h.CreateUser))
	userss.Get("", util.WrapFiberHandler(h.ListUser))
	userss.Get("/:id", util.WrapFiberHandler(h.GetUser))
	userss.Patch("/:id", util.WrapFiberHandler(h.UpdateUserPassword))
	userss.Delete("/:id", util.WrapFiberHandler(h.DeleteUser))
}
