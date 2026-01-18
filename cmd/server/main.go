package main

import (
	"context"
	"fmt"

	"ai-doctor-backend/config"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

var (
	db  *gorm.DB
	rdb *redis.Client
	ctx = context.Background()
)

func main() {
	// 1. 加载配置
	config.InitConfig()

	// 2. 使用配置初始化数据库
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		config.GlobalConfig.Database.Host,
		config.GlobalConfig.Database.User,
		config.GlobalConfig.Database.Password,
		config.GlobalConfig.Database.Dbname,
		config.GlobalConfig.Database.Port,
	)

	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// 自动迁移数据库模式
	db.AutoMigrate(&Todo{})

	// 3. 使用配置初始化 Redis
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
	// 4. 启动 Gin 服务器
	r := gin.Default()
	r.Run(":8080")
}
