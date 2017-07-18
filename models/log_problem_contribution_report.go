package models

const (
	// ProblemTypeSpam 問題タイプ：スパム
	ProblemTypeSpam = 1
	// ProblemTypeInappropriate 問題タイプ：不適切
	ProblemTypeInappropriate = 2
)

// LogProblemContributionReport 投稿画像ログ
type LogProblemContributionReport struct {
	BaseModel
	UserID             int `json:"user_id"`
	Type               int `json:"type"`
	UserContributionID int `json:"user_contribution_id"`
}

// Add 追加する
func (l *LogProblemContributionReport) Add() error {
	return Create(l)
}
