package handler

import (
	"net/http"
	"strconv"

	"ai-doctor-backend/internal/service"

	"github.com/gin-gonic/gin"
)

type LLMHandler struct {
	svc *service.LLMService
}

func NewLLMHandler() *LLMHandler {
	return &LLMHandler{svc: service.NewLLMService()}
}

// List 返回 llms 列表，支持 ?limit=&offset=
func (h *LLMHandler) List(c *gin.Context) {
	limit := 10
	offset := 0

	if s := c.Query("limit"); s != "" {
		if v, err := strconv.Atoi(s); err == nil && v > 0 {
			limit = v
		}
	}
	if s := c.Query("offset"); s != "" {
		if v, err := strconv.Atoi(s); err == nil && v >= 0 {
			offset = v
		}
	}

	list, err := h.svc.List(c.Request.Context(), limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, list)
}
