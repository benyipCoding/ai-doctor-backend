package model

import (
	"gorm.io/gorm"
)

// 定义通用结构体模型
type BaseModel struct {
	gorm.Model
	// 其他通用字段可以在这里添加，例如 CreatedAt, UpdatedAt 等
}
