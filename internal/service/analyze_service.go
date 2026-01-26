package service

import (
	"fmt"

	"ai-doctor-backend/config"
	"ai-doctor-backend/internal/dto"
	prompt "ai-doctor-backend/internal/prompt/analyzeImage"
	helpers "ai-doctor-backend/pkg"
	clients "ai-doctor-backend/pkg/sdk-clients"

	"github.com/gin-gonic/gin"
	"google.golang.org/genai"
)

type AnalyzeService struct {
	apiKey string
}

func NewAnalyzeService() *AnalyzeService {
	return &AnalyzeService{
		apiKey: config.GlobalConfig.AI.APIKey,
	}
}

// AnalyzeData 解码并分析传入的数据，返回分析结果或错误
func (s *AnalyzeService) AnalyzeData(payload dto.AnalyzePayload, ctx *gin.Context) (string, error) {
	imageBytes, err := helpers.Base64ToBytes(payload.Data)
	if err != nil {
		return "", fmt.Errorf("invalid base64 data: %w", err)
	}
	// 生成提示语
	promptText := prompt.GeneratePrompt(payload.ExplanationStyle)
	// 构建内容部分
	parts := []*genai.Part{
		genai.NewPartFromBytes(imageBytes, payload.MimeType),
		genai.NewPartFromText(promptText),
	}
	// 调用 GenAI 客户端生成内容
	contents := []*genai.Content{
		genai.NewContentFromParts(parts, genai.RoleUser),
	}
	// 配置生成参数
	// cfg := &genai.GenerateContentConfig{
	// 	ResponseMIMEType: "application/json",
	// }
	// 调用生成接口
	resp, err := clients.GenAIClient.Models.GenerateContent(ctx, payload.LLMKey, contents, nil)

	if err != nil {
		return "", fmt.Errorf("AI content generation failed: %w", err)
	}
	fmt.Println(resp)

	// 目前示例仅打印长度并返回占位结果，实际应调用 AI 接口完成分析
	fmt.Printf("Decoded image bytes length: %d\n", len(imageBytes))

	// TODO: 使用 s.apiKey 与 LLM/AI 服务交互获取真实分析结果
	result := "analysis started"
	return result, nil
}
