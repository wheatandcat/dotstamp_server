package models

import "github.com/jinzhu/gorm"

// ContributionTotalFollows ユーザー投稿フォロー総数
type ContributionTotalFollows struct {
	BaseModel
	UserContributionID int `json:"user_contribution_id"`
	Count              int
}

// Add 追加する
func (c *ContributionTotalFollows) Add() error {
	return Create(c)
}

// Save 保存する
func (c *ContributionTotalFollows) Save() error {
	return Save(c)
}

// GetListByUserContributionID ユーザー投稿IDからリストを取得する
func (c *ContributionTotalFollows) GetListByUserContributionID(uID []int) (contributionTotalFollows []ContributionTotalFollows, db *gorm.DB, err error) {
	whereList := []map[string]interface{}{
		{"UserContributionID": uID},
	}
	option := make(map[string]interface{})

	db, err = GetListWhere(&contributionTotalFollows, "User_contribution_ID IN :UserContributionID", whereList, option)

	return
}

// Truncate 空にする
func (c *ContributionTotalFollows) Truncate() error {
	return Truncate("contribution_total_follows")
}
