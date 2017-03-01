package models

import (
	"github.com/jinzhu/gorm"
)

const (
	// SoundStatusPublic 音声公開状態
	SoundStatusPublic = 1
	// SoundStatusPrivate 音声非公開状態
	SoundStatusPrivate = 2
)

// UserContributionSound ユーザ投稿音声
type UserContributionSound struct {
	gorm.Model
	UserContributionID int `json:"user_contribution_id"`
	SoundStatus        int `json:"sound_status"`
}

// Add 追加する
func (u *UserContributionSound) Add() error {
	return Create(u)
}

// Save 保存する
func (u *UserContributionSound) Save() error {
	return Save(u)
}

// GetByUserContributionID 投稿IDから取得する
func (u *UserContributionSound) GetByUserContributionID(uID int) (userContributionSound UserContributionSound, db *gorm.DB, err error) {
	whereList := []map[string]interface{}{
		{"UserContributionID": uID},
	}
	option := make(map[string]interface{})

	db, err = GetWhere(&userContributionSound, "User_contribution_ID = :UserContributionID", whereList, option)

	return
}
