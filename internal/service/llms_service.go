package service

import (
	"context"

	"ai-doctor-backend/internal/model"
	"ai-doctor-backend/internal/repository"
)

type LLMService struct {
	repo *repository.LLMRepository
}

func NewLLMService() *LLMService {
	return &LLMService{repo: repository.NewLLMRepository()}
}

// List 调用仓库的 List 方法获取 llms 表数据，支持分页
func (s *LLMService) List(ctx context.Context, limit, offset int) ([]model.LLM, error) {
	return s.repo.List(ctx, limit, offset)
}
