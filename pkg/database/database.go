package database

import (
	"ai-doctor-backend/config"
	"ai-doctor-backend/internal/model"
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	rdb *redis.Client
)

func InitPostgres() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		config.GlobalConfig.Database.Host,
		config.GlobalConfig.Database.User,
		config.GlobalConfig.Database.Password,
		config.GlobalConfig.Database.Dbname,
		config.GlobalConfig.Database.Port,
	)
	db, _ = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// 自动迁移数据库模式
	db.AutoMigrate(&model.Todo{})
	db.AutoMigrate(&model.LLM{})

}

func InitRedis(ctx context.Context) {
	// Redis 初始化逻辑（如果需要）
	rdb = redis.NewClient(&redis.Options{
		Addr:     config.GlobalConfig.Redis.Addr,
		Password: config.GlobalConfig.Redis.Password,
		DB:       config.GlobalConfig.Redis.DB,
	})

	// 测试 Redis 连接
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Errorf("连接 Redis 失败: %w", err))
	}

}
