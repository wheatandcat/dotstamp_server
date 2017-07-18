package models

// LogQuestion 問い合わせログ
type LogQuestion struct {
	BaseModel
	UserID int `json:"user_id"`
	Email  string
	Body   string
}

// Add 追加する
func (l *LogQuestion) Add() error {
	return Create(l)
}
