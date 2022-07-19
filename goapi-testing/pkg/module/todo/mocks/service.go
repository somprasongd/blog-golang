package mocks

import (
	"goapi-project-structure/pkg/common"
	"goapi-project-structure/pkg/module/todo/core/dto"
	"goapi-project-structure/pkg/module/todo/core/ports"

	"github.com/stretchr/testify/mock"
)

type taskServiceMock struct {
	mock.Mock
}

var _ ports.TodoService = &taskServiceMock{}

func NewTaskServiceMock() *taskServiceMock {
	return &taskServiceMock{}
}

func (m *taskServiceMock) Create(newTask dto.NewTodoForm, reqId string) (*dto.TodoResponse, error) {
	args := m.Called(newTask, reqId)
	return args.Get(0).(*dto.TodoResponse), args.Error(1)
}

func (m *taskServiceMock) List(page common.PagingRequest, filters dto.ListTodoFilter, reqId string) (dto.TodoResponses, *common.PagingResult, error) {
	args := m.Called(page, filters, reqId)
	return args.Get(0).(dto.TodoResponses), args.Get(1).(*common.PagingResult), args.Error(2)
}

func (m *taskServiceMock) Get(id string, reqId string) (*dto.TodoResponse, error) {
	args := m.Called(id, reqId)
	return args.Get(0).(*dto.TodoResponse), args.Error(1)
}

func (m *taskServiceMock) UpdateStatus(id string, updateTodo dto.UpdateTodoForm, reqId string) (*dto.TodoResponse, error) {
	args := m.Called(id, updateTodo, reqId)
	return args.Get(0).(*dto.TodoResponse), args.Error(1)
}

func (m *taskServiceMock) Delete(id string, reqId string) error {
	args := m.Called(id, reqId)
	return args.Error(0)
}
