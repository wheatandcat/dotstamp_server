package models

import (
	"time"
)

// UserCharacterImage ユーザーキャラクター画像
type UserCharacterImage struct {
	ID          int `beedb:"PK"`
	UserID      int `sql:"user_id"`
	CharacterID int `sql:"character_id"`
	Priority    int
	DeleteFlag  int `sql:"delete_flag"`
	Created     time.Time
	Updated     time.Time
}

// Add 追加する
func (u *UserCharacterImage) Add() error {
	u.DeleteFlag = DeleteFlagOff
	u.Created = time.Now()
	u.Updated = time.Now()

	return Save(u)
}

// GetListByUserID ユーザーIDからリストを取得する
func (u *UserCharacterImage) GetListByUserID(uID int) (userCharacterImage []UserCharacterImage) {
	whereList := []map[string]interface{}{
		{"UserID": uID},
		{"DeleteFlag": DeleteFlagOff},
	}
	option := make(map[string]interface{})

	GetListWhere(&userCharacterImage, "User_ID = :UserID AND Delete_flag = :DeleteFlag", whereList, option)

	return
}

// GetByID IDから取得する
func (u *UserCharacterImage) GetByID(id int) (userCharacterImage UserCharacterImage) {
	whereList := []map[string]interface{}{
		{"ID": id},
		{"DeleteFlag": DeleteFlagOff},
	}
	option := make(map[string]interface{})

	GetWhere(&userCharacterImage, "ID = :ID AND Delete_flag = :DeleteFlag", whereList, option)

	return
}

// Delete 削除する
func (u *UserCharacterImage) Delete() error {
	u.DeleteFlag = DeleteFlagOn
	u.Updated = time.Now()

	return Save(u)
}
