package models

import (
	"github.com/jinzhu/gorm"
)

const (
	// TalkTypeText 会話タイプ:テキスト
	TalkTypeText = 1
	// TalkTypeImage 会話タイプ:画像
	TalkTypeImage = 2
)

// UserContributionSoundDital ユーザ投稿音声
type UserContributionSoundDital struct {
	gorm.Model
	UserContributionID int `json:"user_contribution_id"`
	Priority           int
	TalkType           int `json:"talk_type"`
	Body               string
	BodySound          string `json:"body_sound"`
	VoiceType          int    `json:"voice_type"`
}

// Add 追加する
func (u *UserContributionSoundDital) Add() error {
	return Create(u)
}

// Save 保存する
func (u *UserContributionSoundDital) Save() error {
	return Save(u)
}

// GetListByUserContributionID 投稿IDからリスト取得する
func (u *UserContributionSoundDital) GetListByUserContributionID(uID int) (userContributionSound []UserContributionSound, db *gorm.DB, err error) {
	whereList := []map[string]interface{}{
		{"UserContributionID": uID},
	}
	option := make(map[string]interface{})

	db, err = GetWhere(&userContributionSound, "User_contribution_ID = :UserContributionID", whereList, option)

	return
}
