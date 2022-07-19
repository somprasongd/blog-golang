package mocks

import (
	"goapi-project-structure/pkg/common"
	"goapi-project-structure/pkg/module/todo/core/dto"
	"goapi-project-structure/pkg/module/todo/core/model"
	"goapi-project-structure/pkg/module/todo/core/ports"

	"github.com/stretchr/testify/mock"
)

type todoRepositoryMock struct {
	mock.Mock
}

var _ ports.TodoRepository = &todoRepositoryMock{}

func NewTodoRepositoryMock() *todoRepositoryMock {
	return &todoRepositoryMock{}
}

func (m *todoRepositoryMock) Create(t *model.Todo) error {
	args := m.Called(t)
	return args.Error(0)
}

func (m *todoRepositoryMock) Find(page common.PagingRequest, filters dto.ListTodoFilter) (model.Todos, *common.PagingResult, error) {
	args := m.Called(page, filters)
	return args.Get(0).(model.Todos), args.Get(1).(*common.PagingResult), args.Error(2)
}

func (m *todoRepositoryMock) FindById(id string) (*model.Todo, error) {
	args := m.Called()
	return args.Get(0).(*model.Todo), args.Error(1)
}

func (m *todoRepositoryMock) UpdateStatusById(id string, status bool) (*model.Todo, error) {
	args := m.Called()
	return args.Get(0).(*model.Todo), args.Error(1)
}

func (m *todoRepositoryMock) DeleteById(id string) error {
	args := m.Called()
	return args.Error(0)
}
