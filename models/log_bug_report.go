package models

// LogBugReport バグレポート
type LogBugReport struct {
	BaseModel
	UserID int    `json:"user_id"`
	Body   string `json:"body"`
}

// Add 追加する
func (c *LogBugReport) Add() error {
	return Create(c)
}
