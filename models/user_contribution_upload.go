package models

import (
	"github.com/jinzhu/gorm"
)

// UserContributionUpload ユーザ投稿アップロード
type UserContributionUpload struct {
	gorm.Model
	UserID             int    `json:"user_id"`
	UserContributionID int    `json:"user_contribution_id"`
	Token              string `json:"token"`
}

// Add 追加する
func (u *UserContributionUpload) Add() error {
	return Create(u)
}

// Save 保存する
func (u *UserContributionUpload) Save() error {
	return Save(u)
}

// GetByUserContributionID 投稿IDから取得する
func (u *UserContributionUpload) GetByUserContributionID(uID int) (userContributionUpload UserContributionUpload, db *gorm.DB, err error) {
	whereList := []map[string]interface{}{
		{"UserContributionID": uID},
	}
	option := make(map[string]interface{})

	db, err = GetWhere(&userContributionUpload, "User_contribution_ID = :UserContributionID", whereList, option)

	return
}
