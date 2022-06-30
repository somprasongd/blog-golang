package handlers

import (
	"goapi-hex/pkg/dto"
	"goapi-hex/pkg/ports"

	"github.com/gofiber/fiber/v2"
)

type TodoHandler struct {
	serv ports.TodoService
}

func NewTodoHandler(serv ports.TodoService) *TodoHandler {
	return &TodoHandler{serv}
}

func (h TodoHandler) CreateTodo(c *fiber.Ctx) error {
	// bind json to struct
	todoForm := new(dto.NewTodoForm)
	if err := c.BodyParser(todoForm); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})

	}

	todo, err := h.serv.Create(*todoForm)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// Return new todo in json format
	return c.JSON(todo)
}

func (h TodoHandler) ListTodo(c *fiber.Ctx) error {
	return c.SendString("List Todo")
}

func (h TodoHandler) GetTodo(c *fiber.Ctx) error {
	return c.SendString("Get Todo")
}

func (h TodoHandler) UpdateTodo(c *fiber.Ctx) error {
	return c.SendString("Update Todo")
}

func (h TodoHandler) DeleteTodo(c *fiber.Ctx) error {
	return c.SendString("Delete Todo")
}
