package models

import (
	"time"
)

// ViewStatusPublic 公開状態
const ViewStatusPublic = 1

// ViewStatusPrivate プライベート状態
const ViewStatusPrivate = 2

// UserContribution ユーザー投稿
type UserContribution struct {
	ID         int `beedb:"PK"`
	UserID     int `sql:"user_id"`
	Title      string
	ViewStatus int `sql:"view_status"`
	DeleteFlag int `sql:"delete_flag"`
	Created    time.Time
	Updated    time.Time
}

// GetIDAndAdd 投稿してIDを取得する
func (u *UserContribution) GetIDAndAdd() (int, error) {
	u.DeleteFlag = DeleteFlagOff
	u.Created = time.Now()
	u.Updated = time.Now()

	if err := Save(u); err != nil {
		return 0, err
	}

	return u.ID, nil
}

// Save 保存する
func (u *UserContribution) Save() error {
	u.DeleteFlag = DeleteFlagOff
	u.Updated = time.Now()

	return Save(u)
}

// Delete 削除する
func (u *UserContribution) Delete() error {
	u.DeleteFlag = DeleteFlagOn
	u.Updated = time.Now()

	return Save(u)
}

// GetByID 投稿IDから取得する
func (u *UserContribution) GetByID(id int) (userContribution UserContribution) {
	whereList := []map[string]interface{}{
		{"ID": id},
		{"DeleteFlag": DeleteFlagOff},
	}
	option := make(map[string]interface{})

	GetWhere(&userContribution, "ID = :ID AND Delete_flag = :DeleteFlag", whereList, option)
	return
}

// GetListByUserID 投稿IDから取得する
func (u *UserContribution) GetListByUserID(userID int) (userContribution []UserContribution) {
	whereList := []map[string]interface{}{
		{"UserID": userID},
		{"DeleteFlag": DeleteFlagOff},
	}

	option := map[string]interface{}{
		"order": "ID desc",
		"limit": map[string]int{
			"size":   20,
			"offset": 0,
		},
	}

	GetListWhere(&userContribution, "User_ID = :UserID AND Delete_flag = :DeleteFlag", whereList, option)
	return
}

// GetByTop 新着から投稿リスト取得する
func (u *UserContribution) GetByTop(o int, s int) (userContributionList []UserContribution) {
	whereList := []map[string]interface{}{
		{"DeleteFlag": DeleteFlagOff},
	}

	optionMap := map[string]interface{}{
		"order": "ID desc",
		"limit": map[string]int{
			"offset": o,
			"size":   s,
		},
	}

	GetListWhere(&userContributionList, "Delete_flag = :DeleteFlag", whereList, optionMap)
	return
}
