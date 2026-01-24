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

func (h *AnalyzeHandler) Analyze(ctx *gin.Context) {
	// 获取请求体数据
	var payload dto.AnalyzePayload
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	h.service.AnalyzeData(payload) // 调用业务层
	ctx.JSON(200, gin.H{"message": "analysis started"})
}
