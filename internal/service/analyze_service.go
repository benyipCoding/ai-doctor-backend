package service

import (
	"ai-doctor-backend/config"
	"ai-doctor-backend/internal/dto"
	helpers "ai-doctor-backend/pkg"
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
	imageBytes, err := helpers.Base64ToBytes(string(payload.Data))
	if err != nil {
		// 解析异常则抛出错误
		panic("Error decoding base64 image: " + err.Error())
	}
	fmt.Printf("Decoded image bytes length: %d\n", len(imageBytes))
	panic("Error decoding base64 image: ")

	// 使用 s.apiKey 调用 AI 服务进行数据分析的逻辑
	// 这里是一个示例实现，实际逻辑会根据具体需求进行编写
	// result := "Analyzed result for data: " + data
	// return result
}
