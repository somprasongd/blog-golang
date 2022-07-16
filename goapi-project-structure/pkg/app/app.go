package app

import (
	"fmt"
	"goapi-project-structure/pkg/app/database"
	"goapi-project-structure/pkg/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"gorm.io/gorm"
)

type Context struct {
	Config *config.Config
	Router *fiber.App
	DB     *gorm.DB
}

type app struct {
	*Context
}

func New(cfg *config.Config) *app {
	return &app{Context: &Context{
		Config: cfg,
	}}
}

func (a *app) InitDB() {
	db, err := database.New(a.Config)
	if err != nil {
		panic(err)
	}
	a.DB = db
}

func (a *app) CloseDB() {
	database.CloseDB(a.DB)
}

func (a *app) InitRouter() {
	r := fiber.New()
	// Default middleware config
	r.Use(cors.New())
	r.Use(logger.New())
	r.Use(requestid.New())

	a.Router = r
}

func (a *app) Serve() error {
	return a.Router.Listen(fmt.Sprintf(":%v", a.Config.Server.Port))
}
