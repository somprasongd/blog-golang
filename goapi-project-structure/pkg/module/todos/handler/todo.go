package handler

import (
	"goapi-project-structure/pkg/common"
	"goapi-project-structure/pkg/module/todos/core/dto"
	"goapi-project-structure/pkg/module/todos/core/ports"

	"github.com/gofiber/fiber/v2"
)

type TodoHandler struct {
	serv ports.TodoService
}

func NewTodoHandler(serv ports.TodoService) *TodoHandler {
	return &TodoHandler{serv}
}

func (h TodoHandler) CreateTodo(c *fiber.Ctx) error {
	// แปลง JSON เป็น struct
	todoForm := new(dto.NewTodoForm)
	if err := c.BodyParser(todoForm); err != nil {
		return common.ResponseError(c, err)

	}
	// ส่งต่อไปให้ service ทำงาน
	todo, err := h.serv.Create(*todoForm)
	if err != nil {
		// error จะถูกจัดการมาจาก service แล้ว
		return common.ResponseError(c, err)
	}

	// คืนค่า todo ที่เพิ่งบันทึกเสร็จกลับไปในรูปแบบ JSON
	return common.ResponseCreated(c, "todo", todo)
}

func (h TodoHandler) ListTodo(c *fiber.Ctx) error {
	return c.SendString("List Todo")
}

func (h TodoHandler) GetTodo(c *fiber.Ctx) error {
	return c.SendString("Get Todo")
}

func (h TodoHandler) UpdateTodoStatus(c *fiber.Ctx) error {
	return c.SendString("Update Todo Status")
}

func (h TodoHandler) DeleteTodo(c *fiber.Ctx) error {
	return c.SendString("Delete Todo")
}
