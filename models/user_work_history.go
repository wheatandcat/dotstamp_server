package models

import (
	"time"
)

// UserWorkHistory ユーザー投稿履歴
type UserWorkHistory struct {
	ID         int `beedb:"PK"`
	UserID     int `sql:"user_id"`
	WorkID     int `sql:"work_id"`
	DeleteFlag int `sql:"delete_flag"`
	Created    time.Time
	Updated    time.Time
}

// GetListByUserID ユーザーIDからリストを取得する
func (u *UserWorkHistory) GetListByUserID(uID int) (userWorkHistory []UserWorkHistory) {
	w := []map[string]interface{}{
		{"UserID": uID},
		{"DeleteFlag": DeleteFlagOff},
	}

	o := map[string]interface{}{}

	GetListWhere(&userWorkHistory, "User_ID = :UserID AND Delete_flag = :DeleteFlag", w, o)
	return
}
