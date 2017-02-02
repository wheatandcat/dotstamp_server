package models

import (
	"time"
	"dotstamp_server/utils"
)

// UserContributionTag ユーザー投稿タグ
type UserContributionTag struct {
	ID                 int `beedb:"PK"`
	UserContributionID int `sql:"user_contribution_id"`
	Name               string
	DeleteFlag         int `sql:"delete_flag"`
	Created            time.Time
	Updated            time.Time
}

// AddList リストを保存する
func (uc *UserContributionTag) AddList(u []UserContributionTag) error {
	return InsertBatch("user_contribution_tag", utils.DbStructListToMapList(u))
}

// GetListByUserContributionID 投稿IDから取得する
func (uc *UserContributionTag) GetListByUserContributionID(id int) (userContributionTag []UserContributionTag) {
	whereList := []map[string]interface{}{
		{"UserContributionID": id},
		{"DeleteFlag": DeleteFlagOff},
	}
	option := make(map[string]interface{})

	GetListWhere(&userContributionTag, "User_contribution_ID = :UserContributionID AND Delete_flag = :DeleteFlag", whereList, option)

	return
}

// GetListByUserContributionIDList 投稿IDリストから取得する
func (uc *UserContributionTag) GetListByUserContributionIDList(idList []int) (userContributionTag []UserContributionTag) {
	whereList := []map[string]interface{}{
		{"User_contribution_ID": idList},
		{"Delete_flag": DeleteFlagOff},
	}
	option := make(map[string]interface{})

	GetListWhere(&userContributionTag, "User_contribution_ID IN :User_contribution_ID AND Delete_flag = :Delete_flag", whereList, option)

	return
}

// GetByID IDから取得する
func (uc *UserContributionTag) GetByID(id int) (userContributionTag UserContributionTag) {
	whereList := []map[string]interface{}{
		{"ID": id},
		{"DeleteFlag": DeleteFlagOff},
	}
	option := make(map[string]interface{})

	GetWhere(&userContributionTag, "ID = :ID AND Delete_flag = :DeleteFlag", whereList, option)

	return
}

// Save 保存する
func (uc *UserContributionTag) Save() error {
	uc.Updated = time.Now()
	return Save(uc)
}

// Delete 削除する
func (uc *UserContributionTag) Delete() error {
	uc.DeleteFlag = DeleteFlagOn
	uc.Updated = time.Now()

	return Save(uc)
}

// GetFindAll 全て取得する
func (uc *UserContributionTag) GetFindAll() (u []UserContributionTag) {
	GetFindAll(&u)

	return
}
