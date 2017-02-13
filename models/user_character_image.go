package models

import "github.com/jinzhu/gorm"

// UserCharacterImage ユーザーキャラクター画像
type UserCharacterImage struct {
	gorm.Model
	UserID      int `json:"user_id"`
	CharacterID int `json:"character_id"`
	Priority    int
}

// Add 追加する
func (u *UserCharacterImage) Add() error {
	return Create(u)
}

// GetListByUserID ユーザーIDからリストを取得する
func (u *UserCharacterImage) GetListByUserID(uID int) (userCharacterImage []UserCharacterImage) {
	whereList := []map[string]interface{}{
		{"UserID": uID},
	}
	option := make(map[string]interface{})

	GetListWhere(&userCharacterImage, "User_ID = :UserID", whereList, option)

	return
}

// GetByID IDから取得する
func (u *UserCharacterImage) GetByID(id int) (userCharacterImage UserCharacterImage) {
	whereList := []map[string]interface{}{
		{"ID": id},
	}
	option := make(map[string]interface{})

	GetWhere(&userCharacterImage, "ID = :ID", whereList, option)

	return
}

// Delete 削除する
func (u *UserCharacterImage) Delete() error {
	return Delete(u)
}
