package mapper

import (
	"goapi-testing/pkg/module/todo/core/dto"
	"goapi-testing/pkg/module/todo/core/model"
)

func CreateTodoFormToModel(dto dto.NewTodoForm) *model.Todo {
	return &model.Todo{
		Text: dto.Text,
	}
}

func TodoToDto(m *model.Todo) *dto.TodoResponse {
	return &dto.TodoResponse{
		ID:        m.ID.String(),
		Text:      m.Text,
		Completed: m.Status.String() == "done",
	}
}

func TodosToDto(Todos model.Todos) dto.TodoResponses {
	dtos := make([]*dto.TodoResponse, len(Todos))
	for i, t := range Todos {
		dtos[i] = TodoToDto(t)
	}

	return dtos
}
