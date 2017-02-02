package models

import (
	"time"
)

// UserCharacter ユーザーキャラクター
type UserCharacter struct {
	ID         int `beedb:"PK"`
	UserID     int `sql:"user_id"`
	Name       string
	Info       string
	Priority   int
	DeleteFlag int `sql:"delete_flag"`
	Created    time.Time
	Updated    time.Time
}

// Add 追加する
func (u *UserCharacter) Add() error {
	u.DeleteFlag = DeleteFlagOff
	u.Created = time.Now()
	u.Updated = time.Now()

	return Save(u)
}

// GetListByUserID ユーザーIDからリストを取得する
func (u *UserCharacter) GetListByUserID(uID int) (userCharacter []UserCharacter) {
	whereList := []map[string]interface{}{
		{"UserID": uID},
		{"DeleteFlag": DeleteFlagOff},
	}
	option := make(map[string]interface{})

	GetListWhere(&userCharacter, "User_ID = :UserID AND Delete_flag = :DeleteFlag", whereList, option)

	return
}

// GetListByIDList IDリストからリスト取得する
func (u *UserCharacter) GetListByIDList(id []int) (userCharacter []UserCharacter) {
	whereList := []map[string]interface{}{
		{"IDList": id},
		{"DeleteFlag": DeleteFlagOff},
	}
	option := make(map[string]interface{})

	GetListWhere(&userCharacter, "ID IN :IDList AND Delete_flag = :DeleteFlag", whereList, option)

	return
}
