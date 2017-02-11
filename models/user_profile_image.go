package models

import (
	"time"
)

// UserProfileImage ユーザープロフィール画像
type UserProfileImage struct {
	ID         int `beedb:"PK"`
	UserID     int `sql:"user_id"`
	DeleteFlag int `sql:"delete_flag"`
	Created    time.Time
	Updated    time.Time
}

// GetIDAndAdd 追加してIDを取得する
func (u *UserProfileImage) GetIDAndAdd() (int, error) {
	u.DeleteFlag = DeleteFlagOff
	u.Created = time.Now()
	u.Updated = time.Now()

	if err := Save(u); err != nil {
		return 0, err
	}

	return u.ID, nil
}

// GetListByUserID ユーザーIDからリストを取得する
func (u *UserProfileImage) GetListByUserID(uID int) (userProfileImage []UserProfileImage) {
	whereList := []map[string]interface{}{
		{"UserID": uID},
		{"DeleteFlag": DeleteFlagOff},
	}
	option := make(map[string]interface{})

	GetListWhere(&userProfileImage, "User_ID = :UserID AND Delete_flag = :DeleteFlag", whereList, option)

	return
}
