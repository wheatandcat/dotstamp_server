package models

import (
	"time"
)

// UserContributionDetail ユーザー投稿詳細
type UserContributionDetail struct {
	ID                 int `beedb:"PK"`
	UserContributionID int `sql:"user_contribution_id"`
	Body               string
	DeleteFlag         int `sql:"delete_flag"`
	Created            time.Time
	Updated            time.Time
}

// Add 追加する
func (u *UserContributionDetail) Add() error {
	u.DeleteFlag = DeleteFlagOff
	u.Created = time.Now()
	u.Updated = time.Now()

	return u.Save()
}

// Save 保存する
func (u *UserContributionDetail) Save() error {
	u.DeleteFlag = DeleteFlagOff
	u.Updated = time.Now()

	return Save(u)
}

// Delete 削除する
func (u *UserContributionDetail) Delete() error {
	u.DeleteFlag = DeleteFlagOn
	u.Updated = time.Now()

	return Save(u)
}

// GetByUserContributionID 投稿IDから取得する
func (u *UserContributionDetail) GetByUserContributionID(uID int) (userContributionDetail UserContributionDetail) {
	whereList := []map[string]interface{}{
		{"UserContributionID": uID},
		{"DeleteFlag": DeleteFlagOff},
	}
	option := make(map[string]interface{})

	GetWhere(&userContributionDetail, "User_contribution_ID = :UserContributionID AND Delete_flag = :DeleteFlag", whereList, option)

	return
}
