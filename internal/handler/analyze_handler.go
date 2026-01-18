package handler

import (
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
	var input struct {
		Data string `json:"data"`
	}
	c.ShouldBindJSON(&input)
	result := h.service.AnalyzeData(input.Data) // 调用业务层
	c.JSON(200, gin.H{"result": result})
}
