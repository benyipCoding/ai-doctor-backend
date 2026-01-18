package dto

type InlineData struct {
	MimeType string `json:"mimeType"`
	Data     string `json:"data"` // base64 字符串
}

type InlineDataPart struct {
	InlineData *InlineData `json:"inlineData,omitempty"`
}

type TextPart struct {
	Text string `json:"text,omitempty"`
}

type ContentPart struct {
	Text       *string     `json:"text,omitempty"`
	InlineData *InlineData `json:"inlineData,omitempty"`
}

type Content struct {
	Role  string        `json:"role"`
	Parts []ContentPart `json:"parts"`
}

type GenerationConfig map[string]interface{}

type AnalyzePayload struct {
	ExplanationStyle string           `json:"explanationStyle" binding:"required"`
	Contents         []Content        `json:"contents" binding:"required"`
	GenerationConfig GenerationConfig `json:"generationConfig,omitempty"`
}
