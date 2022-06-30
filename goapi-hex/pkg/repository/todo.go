package repository

import (
	"goapi-hex/pkg/model"
	"goapi-hex/pkg/ports"

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
