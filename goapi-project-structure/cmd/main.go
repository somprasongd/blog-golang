package main

import (
	"goapi-project-structure/pkg/app"
	"goapi-project-structure/pkg/config"
	"goapi-project-structure/pkg/module"
	"log"
)

func main() {
	cfg := config.LoadConfig()

	app := app.New(cfg)

	app.InitDB()
	defer app.CloseDB()

	app.InitRouter()

	module.Init(app.Context)

	err := app.Serve()
	if err != nil {
		log.Fatalf("fiber.Listen failed %s", err)
	}
}
