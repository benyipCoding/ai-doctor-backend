package model

type Todo struct {
	BaseModel
	Title  string `json:"title"`
	Status bool   `json:"status"`
}
