package models

import "github.com/jinzhu/gorm"

// UserCharacter ユーザーキャラクター
type UserCharacter struct {
	gorm.Model
	UserID   int `json:"user_id"`
	Name     string
	Info     string
	Priority int
}

// Add 追加する
func (u *UserCharacter) Add() error {
	return Create(u)
}

// GetListByUserID ユーザーIDからリストを取得する
func (u *UserCharacter) GetListByUserID(uID int) (userCharacter []UserCharacter, db *gorm.DB, err error) {
	whereList := []map[string]interface{}{
		{"UserID": uID},
	}
	option := make(map[string]interface{})

	db, err = GetListWhere(&userCharacter, "User_ID = :UserID", whereList, option)

	return
}

// GetListByIDList IDリストからリスト取得する
func (u *UserCharacter) GetListByIDList(id []int) (userCharacter []UserCharacter, db *gorm.DB, err error) {
	whereList := []map[string]interface{}{
		{"IDList": id},
	}
	option := make(map[string]interface{})

	db, err = GetListWhere(&userCharacter, "ID IN :IDList", whereList, option)

	return
}
