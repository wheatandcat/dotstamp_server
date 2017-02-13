package models

import "github.com/jinzhu/gorm"

// UserContributionDetail ユーザー投稿詳細
type UserContributionDetail struct {
	gorm.Model
	UserContributionID int `json:"user_contribution_id"`
	Body               string
}

// Add 追加する
func (u *UserContributionDetail) Add() error {
	return Create(u)
}

// Save 保存する
func (u *UserContributionDetail) Save() error {
	return Save(u)
}

// Delete 削除する
func (u *UserContributionDetail) Delete() error {
	return Delete(u)
}

// GetByUserContributionID 投稿IDから取得する
func (u *UserContributionDetail) GetByUserContributionID(uID int) (userContributionDetail UserContributionDetail) {
	whereList := []map[string]interface{}{
		{"UserContributionID": uID},
	}
	option := make(map[string]interface{})

	GetWhere(&userContributionDetail, "User_contribution_ID = :UserContributionID", whereList, option)

	return
}
