package models

import "github.com/jinzhu/gorm"

// UserWorkHistory ユーザー投稿履歴
type UserWorkHistory struct {
	gorm.Model
	UserID int `json:"user_id"`
	WorkID int `json:"work_id"`
}

// GetListByUserID ユーザーIDからリストを取得する
func (u *UserWorkHistory) GetListByUserID(uID int) (userWorkHistory []UserWorkHistory, db *gorm.DB, err error) {
	w := []map[string]interface{}{
		{"UserID": uID},
	}

	o := map[string]interface{}{}

	db, err = GetListWhere(&userWorkHistory, "User_ID = :UserID", w, o)
	return
}
