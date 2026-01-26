package main

import (
	"context"
	"fmt"

	"ai-doctor-backend/config"
	"ai-doctor-backend/internal/middleware"
	"ai-doctor-backend/internal/router"
	"ai-doctor-backend/pkg/database"
	"ai-doctor-backend/pkg/logger"
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

	// 6. 注册中间件
	// 全局异常处理中间件
	log := logger.New()
	defer log.Sync()

	r.Use(
		middleware.TraceMiddleware(),
		middleware.LoggerMiddleware(log),
		middleware.RecoveryMiddleware(log),
	)

	// 注册路由（集中在 internal/router）
	router.RegisterAPIRoutes(r)

	r.Run(fmt.Sprintf(":%d", config.GlobalConfig.Server.Port))
}
