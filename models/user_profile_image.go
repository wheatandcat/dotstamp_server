package models

import "github.com/jinzhu/gorm"

// UserProfileImage ユーザープロフィール画像
type UserProfileImage struct {
	gorm.Model
	UserID int `json:"user_id"`
}

// GetIDAndAdd 追加してIDを取得する
func (u *UserProfileImage) GetIDAndAdd() (uint, error) {
	if err := Create(u); err != nil {
		return 0, err
	}

	return u.ID, nil
}

// GetListByUserID ユーザーIDからリストを取得する
func (u *UserProfileImage) GetListByUserID(uID int) (userProfileImage []UserProfileImage) {
	whereList := []map[string]interface{}{
		{"UserID": uID},
	}
	option := make(map[string]interface{})

	GetListWhere(&userProfileImage, "User_ID = :UserID", whereList, option)

	return
}
