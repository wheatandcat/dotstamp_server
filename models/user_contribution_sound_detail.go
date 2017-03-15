package models

import (
	"github.com/jinzhu/gorm"
)

const (
	// TalkTypeText 会話タイプ:テキスト
	TalkTypeText = 1
	// TalkTypeImage 会話タイプ:画像
	TalkTypeImage = 2
	// MakeStatusUncreated 状態:未作成
	MakeStatusUncreated = 0
	// MakeStatusMade 状態:作成済み
	MakeStatusMade = 1
)

// UserContributionSoundDetail ユーザ投稿音声
type UserContributionSoundDetail struct {
	gorm.Model
	UserContributionID int `json:"user_contribution_id"`
	Priority           int
	TalkType           int `json:"talk_type"`
	Body               string
	BodySound          string `json:"body_sound"`
	VoiceType          int    `json:"voice_type"`
	MakeStatus         int    `json:"make_status"`
}

// Add 追加する
func (u *UserContributionSoundDetail) Add() error {
	return Create(u)
}

// Save 保存する
func (u *UserContributionSoundDetail) Save() error {
	return Save(u)
}

// GetListByUserContributionID 投稿IDからリスト取得する
func (u *UserContributionSoundDetail) GetListByUserContributionID(uID int) (userContributionSoundDetail []UserContributionSoundDetail, db *gorm.DB, err error) {
	whereList := []map[string]interface{}{
		{"UserContributionID": uID},
	}
	option := make(map[string]interface{})

	db, err = GetListWhere(&userContributionSoundDetail, "User_contribution_ID = :UserContributionID", whereList, option)

	return
}

// UpdateToMakeStatusByUserContributionID 投稿IDから作成状態を更新する
func (u *UserContributionSoundDetail) UpdateToMakeStatusByUserContributionID(uID int, makeStatus int) (err error) {
	userContributionSoundDetail := []UserContributionSoundDetail{}

	whereList := []map[string]interface{}{
		{"UserContributionID": uID},
	}
	option := make(map[string]interface{})
	update := []interface{}{
		"make_status",
		makeStatus,
	}

	_, err = Update(&userContributionSoundDetail, update, "User_contribution_ID = :UserContributionID", whereList, option)

	return
}

// GetByID IDから取得する
func (u *UserContributionSoundDetail) GetByID(id uint) (userContributionSoundDetail UserContributionSoundDetail, db *gorm.DB, err error) {
	whereList := []map[string]interface{}{
		{"ID": id},
	}
	option := make(map[string]interface{})

	db, err = GetWhere(&userContributionSoundDetail, "ID = :ID", whereList, option)

	return
}
