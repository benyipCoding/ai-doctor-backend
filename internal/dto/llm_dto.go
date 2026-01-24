package dto

import "ai-doctor-backend/internal/model"

// LLMDto 是返回给前端的 LLM 视图对象，仅包含需要的字段
type LLMDto struct {
	Key  string `json:"id"`
	Name string `json:"name"`
	Tag  string `json:"tag"`
	Desc string `json:"desc"`
}

// FromModels 将 model.LLM 列表映射为 DTO 列表
func FromModels(list []model.LLM) []LLMDto {
	out := make([]LLMDto, 0, len(list))
	for _, l := range list {
		out = append(out, LLMDto{
			Key:  l.Key,
			Name: l.Name,
			Tag:  l.Tag,
			Desc: l.Desc,
		})
	}
	return out
}
