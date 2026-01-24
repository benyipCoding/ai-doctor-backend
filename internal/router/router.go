package router

import (
	"github.com/gin-gonic/gin"

	"ai-doctor-backend/internal/handler"
)

// RegisterAPIRoutes 将所有 /api_v1 下的路由集中注册，便于维护
func RegisterAPIRoutes(r *gin.Engine) {
	api := r.Group("/api_v1")

	analyze := handler.NewAnalyzeHandler()
	llms := handler.NewLLMHandler()

	api.POST("/analyzeImage", analyze.Analyze)
	api.GET("/llms", llms.List)
}
