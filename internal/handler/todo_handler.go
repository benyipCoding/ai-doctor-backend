package handler

import (
	"ai-doctor-backend/internal/service"

	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	service *service.TodoService
}

func (h *TodoHandler) Create(c *gin.Context) {
	var input struct {
		Title string `json:"title"`
	}
	c.ShouldBindJSON(&input)

	h.service.AddTask(input.Title) // 调用业务层
	c.JSON(200, gin.H{"status": "ok"})
}
