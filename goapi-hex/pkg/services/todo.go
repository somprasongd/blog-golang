package services

import (
	"errors"
	"goapi-hex/pkg/dto"
	"goapi-hex/pkg/model"
	"goapi-hex/pkg/ports"
)

type todoService struct {
	repo ports.TodoRepository
}

func NewTodoService(repo ports.TodoRepository) ports.TodoService {
	return &todoService{repo}
}

func (s todoService) Create(form dto.NewTodoForm) (*dto.TodoResponse, error) {
	if form.Text == "" {
		return nil, errors.New("text is required")
	}

	todo := model.Todo{
		Text: form.Text,
	}

	err := s.repo.Create(&todo)
	if err != nil {
		return nil, errors.New("database error while insert new todo")
	}

	serializedTodo := dto.TodoResponse{
		ID:   todo.ID,
		Text: todo.Text,
		Done: todo.Done,
	}

	return &serializedTodo, nil
}
