package ports

import "goapi-project-structure/pkg/module/todos/core/model"

type TaskRepository interface {
	Create(t *model.Task) error
	Find(page pagination.PagingRequest, filters dto.ListTaskFilter) (model.Tasks, *pagination.PagingResult, error)
	FindById(id string) (*model.Task, error)
	UpdateById(id string, t *model.Task) error
	DeleteById(id string) error
}
