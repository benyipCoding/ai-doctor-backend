package handler

import (
	"ai-doctor-backend/internal/dto"
	"ai-doctor-backend/internal/service"

	"github.com/gin-gonic/gin"
)

type AnalyzeHandler struct {
	service *service.AnalyzeService
}

func NewAnalyzeHandler() *AnalyzeHandler {
	return &AnalyzeHandler{
		service: service.NewAnalyzeService(),
	}
}

func (h *AnalyzeHandler) Analyze(c *gin.Context) {
	// 获取请求体数据
	var input dto.AnalyzePayload
	c.ShouldBindJSON(&input)
	h.service.AnalyzeData(input) // 调用业务层
	c.JSON(200, gin.H{"message": "analysis started"})
}
