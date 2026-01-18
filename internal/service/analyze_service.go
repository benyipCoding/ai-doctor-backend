package service

import (
	"ai-doctor-backend/config"
	"ai-doctor-backend/internal/dto"
	"fmt"
)

type AnalyzeService struct {
	apiKey string
}

func NewAnalyzeService() *AnalyzeService {
	return &AnalyzeService{
		apiKey: config.GlobalConfig.AI.APIKey,
	}
}

func (s *AnalyzeService) AnalyzeData(payload dto.AnalyzePayload) {
	// 使用 s.apiKey 调用 AI 服务进行数据分析的逻辑
	// 这里是一个示例实现，实际逻辑会根据具体需求进行编写
	// result := "Analyzed result for data: " + data
	// return result
	fmt.Println()
}
