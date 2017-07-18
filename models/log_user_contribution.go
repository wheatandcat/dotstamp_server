package models

// LogUserContribution 投稿画像ログ
type LogUserContribution struct {
	BaseModel
	UserContributionID int `json:"user_contribution_id"`
	UserID             int `json:"user_id"`
}

// Add 追加する
func (l *LogUserContribution) Add() error {
	return Create(l)
}
