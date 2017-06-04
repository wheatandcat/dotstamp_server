package models

import "github.com/jinzhu/gorm"

// LogBugReport バグレポート
type LogBugReport struct {
	gorm.Model
	UserID int `json:"user_id"`
	Body   string
}

// Add 追加する
func (c *LogBugReport) Add() error {
	return Create(c)
}
