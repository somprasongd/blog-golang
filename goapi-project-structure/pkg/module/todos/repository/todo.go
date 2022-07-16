package repository

import (
	"goapi-project-structure/pkg/module/todos/core/model"
	"goapi-project-structure/pkg/module/todos/core/ports"

	"gorm.io/gorm"
)

type todoRepositoryDB struct {
	db *gorm.DB
}

func NewTodoRepositoryDB(db *gorm.DB) ports.TodoRepository {
	return &todoRepositoryDB{db}
}

func (r todoRepositoryDB) Create(todo *model.Todo) error {
	return r.db.Create(&todo).Error
}

func (r todoRepositoryDB) Find() ([]model.Todo, error) {
	return nil, nil
}
func (r todoRepositoryDB) FindById(id int) (*model.Todo, error) {
	return nil, nil
}
func (r todoRepositoryDB) UpdateStatusById(id int, isDone bool) error {
	return nil
}
func (r todoRepositoryDB) DeleteById(id int) error {
	return nil
}
