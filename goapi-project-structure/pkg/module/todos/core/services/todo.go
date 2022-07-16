package services

import (
	"goapi-project-structure/pkg/common/errs"
	"goapi-project-structure/pkg/module/todos/core/dto"
	"goapi-project-structure/pkg/module/todos/core/model"
	"goapi-project-structure/pkg/module/todos/core/ports"
)

type todoService struct {
	repo ports.TodoRepository
}

func NewTodoService(repo ports.TodoRepository) ports.TodoService {
	return &todoService{repo}
}

func (s todoService) Create(form dto.NewTodoForm) (*dto.TodoResponse, error) {
	// การตรวจสอบ
	if form.Text == "" {
		return nil, errs.NewBadRequestError("text is required")
	}

	todo := model.Todo{
		Text: form.Text,
	}
	// เรียกใช้ repo เพื่อบันทึกข้อมูลใหม่
	err := s.repo.Create(&todo)
	if err != nil {
		return nil, errs.NewUnexpectedError("database error while insert new todo")
	}

	// สร้าง struct ที่ต้องการให้ handler ส่งกลับไปหา client
	serializedTodo := dto.TodoResponse{
		ID:   todo.ID,
		Text: todo.Text,
		Done: todo.Done,
	}

	return &serializedTodo, nil
}

func (s todoService) List(completed string) ([]dto.TodoResponse, error) {
	return nil, nil
}
func (s todoService) Get(id int) (*dto.TodoResponse, error) {
	return nil, nil
}
func (s todoService) Update(id int, updateTodo dto.UpdateTodoForm) error {
	return nil
}
func (s todoService) Delete(id int) error {
	return nil
}
