package models

// LogContributionImage 投稿画像ログ
type LogContributionImage struct {
	BaseModel
	UserContributionID int `json:"user_contribution_id"`
}

// GetIDAndAdd 追加してIDを取得する
func (l *LogContributionImage) GetIDAndAdd() (uint, error) {
	if err := Create(l); err != nil {
		return 0, err
	}

	return l.ID, nil
}
