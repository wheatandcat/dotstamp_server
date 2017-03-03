package models

import "github.com/jinzhu/gorm"

// LogUserContribution 投稿画像ログ
type LogUserContribution struct {
	gorm.Model
	UserContributionID int `json:"user_contribution_id"`
	UserID             int `json:"user_id"`
}

// Add 追加する
func (l *LogUserContribution) Add() error {
	return Create(l)
}
