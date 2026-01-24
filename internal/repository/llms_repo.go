package repository

import (
	"context"

	"ai-doctor-backend/internal/model"
	"ai-doctor-backend/pkg/database"

	"gorm.io/gorm"
)

type LLMRepository struct {
	db *gorm.DB
}

// NewLLMRepository returns a new repository using the global DB instance from pkg/database
func NewLLMRepository() *LLMRepository {
	return &LLMRepository{db: database.GetDB()}
}

// Create 插入一条 LLM 记录
func (r *LLMRepository) Create(ctx context.Context, llm *model.LLM) error {
	return r.db.WithContext(ctx).Create(llm).Error
}

// GetByID 根据主键 ID 查询
func (r *LLMRepository) GetByID(ctx context.Context, id uint) (*model.LLM, error) {
	var llm model.LLM
	if err := r.db.WithContext(ctx).First(&llm, id).Error; err != nil {
		return nil, err
	}
	return &llm, nil
}

// GetByKey 根据 Key 查询（Key 在模型中使用 uniqueIndex）
func (r *LLMRepository) GetByKey(ctx context.Context, key string) (*model.LLM, error) {
	var llm model.LLM
	if err := r.db.WithContext(ctx).Where("key = ?", key).First(&llm).Error; err != nil {
		return nil, err
	}
	return &llm, nil
}

// List 列表查询，支持分页（limit, offset）
func (r *LLMRepository) List(ctx context.Context, limit, offset int) ([]model.LLM, error) {
	var list []model.LLM
	if err := r.db.WithContext(ctx).Limit(limit).Offset(offset).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

// Update 保存更新的 LLM（根据主键）
func (r *LLMRepository) Update(ctx context.Context, llm *model.LLM) error {
	return r.db.WithContext(ctx).Save(llm).Error
}

// Delete 根据主键删除
func (r *LLMRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.LLM{}, id).Error
}
