package handler

import (
	"time"

	"github.com/gin-gonic/gin"
)

// HealthHandler 提供简单的连通性检测接口
type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
		"time":   time.Now().Format(time.RFC3339),
	})
}
