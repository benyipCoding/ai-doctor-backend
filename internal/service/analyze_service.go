package service

import (
	"fmt"

	"ai-doctor-backend/config"
	"ai-doctor-backend/internal/dto"
	helpers "ai-doctor-backend/pkg"
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
func (s *AnalyzeService) AnalyzeData(payload dto.AnalyzePayload) (string, error) {
	imageBytes, err := helpers.Base64ToBytes(payload.Data)
	if err != nil {
		return "", fmt.Errorf("invalid base64 data: %w", err)
	}

	// 目前示例仅打印长度并返回占位结果，实际应调用 AI 接口完成分析
	fmt.Printf("Decoded image bytes length: %d\n", len(imageBytes))

	// TODO: 使用 s.apiKey 与 LLM/AI 服务交互获取真实分析结果
	result := "analysis started"
	return result, nil
}
