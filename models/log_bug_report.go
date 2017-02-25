package models

import "github.com/jinzhu/gorm"

// LogBugReport ユーザー投稿フォロー総数
type LogBugReport struct {
	gorm.Model
	UserID int `json:"user_id"`
	Body   string
}

// Add 追加する
func (c *LogBugReport) Add() error {
	return Create(c)
}
