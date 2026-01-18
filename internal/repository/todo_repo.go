package repository

import (
	"ai-doctor-backend/internal/model"

	"gorm.io/gorm"
)

type TodoRepository struct {
	db *gorm.DB
}

func (r *TodoRepository) CreateTodo(todo *model.Todo) error {
	return r.db.Create(todo).Error
}
