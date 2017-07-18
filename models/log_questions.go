package models

// LogQuestion 問い合わせログ
type LogQuestion struct {
	BaseModel
	UserID int    `json:"user_id"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}

// Add 追加する
func (l *LogQuestion) Add() error {
	return Create(l)
}
