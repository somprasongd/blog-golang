package main

import (
	"goapi-hex/pkg/database"
	"goapi-hex/pkg/handlers"
	"goapi-hex/pkg/repository"
	"goapi-hex/pkg/services"

	"github.com/gofiber/fiber/v2"
)

func main() {
	db, err := database.ConnectDB("gorm.db")
	if err != nil {
		panic(err)
	}
	defer database.CloseDB(db)

	app := fiber.New()

	repo := repository.NewTodoRepositoryDB(db)
	serv := services.NewTodoService(repo)
	h := handlers.NewTodoHandler(serv)

	todos := app.Group("/api/todos")
	todos.Post("", h.CreateTodo)
	todos.Get("", h.ListTodo)
	todos.Get("/:id", h.GetTodo)
	todos.Patch("/:id", h.UpdateTodo)
	todos.Delete("/:id", h.DeleteTodo)

	app.Listen(":8080")
}
