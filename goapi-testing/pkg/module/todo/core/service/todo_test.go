package service_test

import (
	"errors"
	"goapi-project-structure/pkg/common"
	"goapi-project-structure/pkg/module/todo/core/dto"
	"goapi-project-structure/pkg/module/todo/core/mapper"
	"goapi-project-structure/pkg/module/todo/core/model"
	"goapi-project-structure/pkg/module/todo/core/service"
	"goapi-project-structure/pkg/module/todo/mocks"
	"testing"

	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

func TestTodo(t *testing.T) {

	t.Run("Add Todo Service", func(t *testing.T) {
		t.Run("Success", func(t *testing.T) {
			// Arrage
			mockForm := dto.NewTodoForm{
				Text: "Test new todo",
			}
			mockModel := mapper.CreateTodoFormToModel(mockForm)
			mockResp := mapper.TodoToDto(mockModel)

			repo := mocks.NewTodoRepositoryMock()

			repo.On("Create", mockModel).Return(nil)

			svc := service.NewTodoService(repo)

			// Act
			got, err := svc.Create(mockForm, "")

			// Assert
			assert.NoError(t, err)
			assert.Equal(t, mockResp, got)

		})
		t.Run("Invalid JSON Boby", func(t *testing.T) {
			// Arrage
			mockForm := dto.NewTodoForm{
				Text: "",
			}
			repo := mocks.NewTodoRepositoryMock()
			svc := service.NewTodoService(repo)

			// Act
			_, err := svc.Create(mockForm, "")

			// Assert
			assert.ErrorIs(t, err, common.NewInvalidError("text: text is a required field"))
			repo.AssertNotCalled(t, "Create")

		})
		t.Run("Error", func(t *testing.T) {
			// Arrage
			mockForm := dto.NewTodoForm{
				Text: "Test new todo",
			}

			mockModel := mapper.CreateTodoFormToModel(mockForm)
			mockResp := mapper.TodoToDto(mockModel)

			repo := mocks.NewTodoRepositoryMock()
			repo.On("Create", mockModel).Return(errors.New("Some error down call chain"))

			svc := service.NewTodoService(repo)

			// Act
			got, err := svc.Create(mockForm, "")
			assert.NotEqual(t, mockResp, got)
			assert.ErrorIs(t, err, common.ErrDbInsert)
		})
	})

	t.Run("List Todo Service", func(t *testing.T) {
		t.Run("Success", func(t *testing.T) {
			// Arrage
			page := common.PagingRequest{
				Page:  1,
				Limit: 10,
				Order: "",
			}

			mockFilters := dto.ListTodoFilter{}
			mockFilters.Term = "1"
			b := false
			mockFilters.Completed = &b

			mockTodo := model.Todo{
				ID:     uuid.FromStringOrNil("7bce9463-37ce-4413-8f2f-31f3c643e1d5"),
				Text:   "Todo 1",
				Status: model.OPEN,
			}
			mockTodos := model.Todos{&mockTodo}
			mockResp := mapper.TodosToDto(mockTodos)

			mockPageRes := &common.PagingResult{
				Page:      1,
				Limit:     10,
				PrevPage:  0,
				NextPage:  2,
				Count:     20,
				TotalPage: 2,
			}

			repo := mocks.NewTodoRepositoryMock()

			repo.On("Find", page, mockFilters).Return(mockTodos, mockPageRes, nil)

			svc := service.NewTodoService(repo)

			// Act
			got, gotPage, err := svc.List(page, mockFilters, "")

			// Assert
			assert.NoError(t, err)
			assert.Equal(t, mockResp, got)
			assert.Equal(t, mockPageRes, gotPage)
		})
		t.Run("Error", func(t *testing.T) {
			// Arrage
			page := common.PagingRequest{
				Page:  1,
				Limit: 10,
				Order: "",
			}

			mockFilters := dto.ListTodoFilter{}
			mockFilters.Term = "1"
			b := false
			mockFilters.Completed = &b

			mockTodo := model.Todo{
				ID:     uuid.FromStringOrNil("7bce9463-37ce-4413-8f2f-31f3c643e1d5"),
				Text:   "Todo 1",
				Status: model.OPEN,
			}
			mockTodos := model.Todos{&mockTodo}
			mockResp := mapper.TodosToDto(mockTodos)

			mockPageRes := &common.PagingResult{
				Page:      1,
				Limit:     10,
				PrevPage:  0,
				NextPage:  2,
				Count:     20,
				TotalPage: 2,
			}

			repo := mocks.NewTodoRepositoryMock()

			repo.On("Find", page, mockFilters).Return(nil, nil, errors.New("Some error down call chain"))

			svc := service.NewTodoService(repo)

			// Act
			got, gotPage, err := svc.List(page, mockFilters, "")

			// Assert
			assert.Error(t, err)
			assert.NotEqual(t, mockResp, got)
			assert.NotEqual(t, mockPageRes, gotPage)
			assert.ErrorIs(t, err, common.ErrDbQuery)
		})
	})
}
