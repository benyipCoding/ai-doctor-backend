package clients

import (
	"ai-doctor-backend/config"
	"context"
	"fmt"

	"google.golang.org/genai"
)

var (
	GenAIClient *genai.Client
	err         error
)

func InitGenAIClient(ctx context.Context) {
	GenAIClient, err = genai.NewClient(ctx, &genai.ClientConfig{
		APIKey: config.GlobalConfig.AI.APIKey,
		HTTPOptions: genai.HTTPOptions{
			BaseURL: config.GlobalConfig.AI.BaseURL,
		},
	})
	if err != nil {
		panic(fmt.Errorf("连接 GenAI 失败: %w", err))
	}
}
