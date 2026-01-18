package service

import (
	"ai-doctor-backend/internal/model"
	"ai-doctor-backend/internal/repository"
)

type TodoService struct {
	repo *repository.TodoRepository
}

func (s *TodoService) AddTask(title string) error {
	todo := &model.Todo{Title: title}
	// 逻辑处理...
	return s.repo.CreateTodo(todo)
}
