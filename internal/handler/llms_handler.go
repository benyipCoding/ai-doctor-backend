package handler

import (
	"ai-doctor-backend/internal/dto"
	"ai-doctor-backend/internal/response"
	"ai-doctor-backend/internal/service"
	helpers "ai-doctor-backend/pkg"

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
	// 使用通用 helper 解析分页参数
	limit, offset := helpers.ParseLimitOffset(c.Query("limit"), c.Query("offset"), 10)
	// 调用服务层获取 LLM 列表
	list, err := h.svc.List(c.Request.Context(), limit, offset)
	if err != nil {
		response.Handle(c, nil, err)
		return
	}
	// 转换为 DTO 并返回（使用 dto 包提供的映射函数）
	response.Handle(c, dto.FromModels(list), nil)
}
