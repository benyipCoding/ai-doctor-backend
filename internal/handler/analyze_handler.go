package handler

import (
	"ai-doctor-backend/internal/dto"
	"ai-doctor-backend/internal/response"
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
		// ctx.JSON(400, gin.H{"error": err.Error()})
		response.Handle(ctx, nil, err)
		return
	}
	// 调用业务层进行分析，依据返回的 err 决定响应
	result, err := h.service.AnalyzeData(payload)
	response.Handle(ctx, gin.H{"message": result}, err)
}
