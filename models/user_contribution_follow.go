package models

import (
	"time"
)

// UserContributionFollow ユーザー投稿フォロー
type UserContributionFollow struct {
	ID                 int `beedb:"PK"`
	UserID             int `sql:"user_id"`
	UserContributionID int `sql:"user_contribution_id"`
	DeleteFlag         int `sql:"delete_flag"`
	Created            time.Time
	Updated            time.Time
}

// Add 追加する
func (u *UserContributionFollow) Add(uID int, ucID int) error {
	u.UserID = uID
	u.UserContributionID = ucID
	u.DeleteFlag = DeleteFlagOff
	u.Created = time.Now()
	u.Updated = time.Now()

	return Save(u)
}

// Delete 削除する
func (u *UserContributionFollow) Delete() error {
	u.DeleteFlag = DeleteFlagOn
	u.Updated = time.Now()

	return Save(u)
}

// GetListByUserContributionID 投稿IDからフォローを取得する
func (u *UserContributionFollow) GetListByUserContributionID(ucID int) (userContributionFollow []UserContributionFollow) {
	whereList := []map[string]interface{}{
		{"UserContributionID": ucID},
		{"DeleteFlag": DeleteFlagOff},
	}
	option := make(map[string]interface{})

	GetWhere(&userContributionFollow, "User_contribution_ID = :UserContributionID AND Delete_flag = :DeleteFlag", whereList, option)

	return
}

// GetListByUserContributionIDList 投稿IDリストからフォローを取得する
func (u *UserContributionFollow) GetListByUserContributionIDList(ucID []int) (userContributionFollow []UserContributionFollow) {
	whereList := []map[string]interface{}{
		{"UserContributionID": ucID},
		{"DeleteFlag": DeleteFlagOff},
	}
	option := make(map[string]interface{})

	GetWhere(&userContributionFollow, "User_contribution_ID IN :UserContributionID AND Delete_flag = :DeleteFlag", whereList, option)

	return
}
