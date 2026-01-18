package main

import (
	"context"

	"ai-doctor-backend/config"
	"ai-doctor-backend/internal/handler"
	"ai-doctor-backend/pkg/database"

	"github.com/gin-gonic/gin"
)

func main() {
	// 创建可取消的根 Context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 1. 加载配置
	config.InitConfig()

	// 2. 使用配置初始化数据库
	database.InitPostgres()

	// 3. 使用配置初始化 Redis
	database.InitRedis(ctx)

	// 4. 启动 Gin 服务器
	r := gin.Default()

	health := handler.NewHealthHandler()
	analyze := handler.NewAnalyzeHandler()

	// 注册路由
	api := r.Group("/api_v1")
	api.GET("/ping", health.Ping)
	api.POST("/analyzeImage", analyze.Analyze)

	r.Run(":8080")
}
