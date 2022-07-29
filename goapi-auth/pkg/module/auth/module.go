package auth

import (
	"goapi/pkg/app"
	"goapi/pkg/module/auth/core/ports"
	"goapi/pkg/module/auth/core/service"
	"goapi/pkg/module/auth/handler"
	"goapi/pkg/module/users/repository"
	"goapi/pkg/util"

	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	BaseURL     string
	Router      *fiber.App
	AuthService ports.AuthService
}

func Init(ctx *app.Context) {
	// สร้าง dependencies ทั้งหมด
	repo := repository.NewUserRepositoryDB(ctx.DB.DB)
	svc := service.NewUserService(repo)

	cfg := RouteConfig{
		BaseURL:     ctx.Config.App.BaseUrl,
		Router:      ctx.Router,
		AuthService: svc,
	}

	SetupRoutes(cfg)
}

func SetupRoutes(cfg RouteConfig) {
	h := handler.NewAuthHandler(cfg.AuthService)

	auth := cfg.Router.Group(cfg.BaseURL + "/auth")

	auth.Post("/register", util.WrapFiberHandler(h.Register))
	auth.Post("/login", util.WrapFiberHandler(h.Login))
	auth.Post("/logout", util.WrapFiberHandler(h.Logout))
	auth.Post("/refresh", util.WrapFiberHandler(h.Refresh))
	auth.Get("/profile", util.WrapFiberHandler(h.Profile))
}
