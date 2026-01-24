package main

import (
	"context"

	"ai-doctor-backend/config"
	"ai-doctor-backend/internal/handler"
	"ai-doctor-backend/pkg/database"
	clients "ai-doctor-backend/pkg/sdk-clients"

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

	// 4. 初始化 GenAI 客户端
	clients.InitGenAIClient(ctx)

	// 5. 启动 Gin 服务器
	r := gin.Default()

	analyze := handler.NewAnalyzeHandler()

	// 注册路由
	api := r.Group("/api_v1")
	api.POST("/analyzeImage", analyze.Analyze)

	r.Run(":8080")
}
