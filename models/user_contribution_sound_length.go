package models

import (
	"github.com/jinzhu/gorm"
)

// UserContributionSoundLength ユーザ投稿音声長さ
type UserContributionSoundLength struct {
	gorm.Model
	UserContributionID int `json:"user_contribution_id"`
	Second             int `json:"second"`
	Length             int `json:"length"`
}

// Add 追加する
func (u *UserContributionSoundLength) Add() error {
	return Create(u)
}

// Save 保存する
func (u *UserContributionSoundLength) Save() error {
	return Save(u)
}

// GetByUserContributionID 投稿IDから取得する
func (u *UserContributionSoundLength) GetByUserContributionID(uID int) (userContributionSoundLength UserContributionSoundLength, db *gorm.DB, err error) {
	whereList := []map[string]interface{}{
		{"UserContributionID": uID},
	}
	option := make(map[string]interface{})

	db, err = GetWhere(&userContributionSoundLength, "User_contribution_ID = :UserContributionID", whereList, option)

	return
}

// GetByTop 新着から投稿リスト取得する
func (u *UserContributionSoundLength) GetByTop(o int, s int) (userContributionSoundLength []UserContributionSoundLength, db *gorm.DB, err error) {
	whereList := []map[string]interface{}{}

	optionMap := map[string]interface{}{
		"order":  "ID desc",
		"limit":  s,
		"offset": o,
	}

	db, err = GetListWhere(&userContributionSoundLength, "", whereList, optionMap)
	return
}
