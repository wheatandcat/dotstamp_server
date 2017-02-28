package models

import "github.com/jinzhu/gorm"

const (
	// ProblemTypeSpam 問題タイプ：スパム
	ProblemTypeSpam = 1
	// ProblemTypeInappropriate 問題タイプ：不適切
	ProblemTypeInappropriate = 2
)

// LogProblemContributionReport 投稿画像ログ
type LogProblemContributionReport struct {
	gorm.Model
	UserID             int `json:"user_id"`
	Type               int
	UserContributionID int `json:"user_contribution_id"`
}

// Add 追加する
func (l *LogProblemContributionReport) Add() error {
	return Create(l)
}
