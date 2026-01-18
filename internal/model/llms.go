package model

// LLM 表示大模型列表中的一条记录。
// 对应前端示例中的字段：id, name, tag, desc
type LLM struct {
	BaseModel
	// Key 用于存储像 "gemini-flash" 这样的字符串 ID
	Key  string `gorm:"size:100;uniqueIndex" json:"id"`
	Name string `gorm:"size:200" json:"name"`
	Tag  string `gorm:"size:50" json:"tag"`
	Desc string `gorm:"type:text" json:"desc"`
}

// TableName 指定数据库表名为 llms
func (LLM) TableName() string {
	return "llms"
}
