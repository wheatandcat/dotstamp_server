package models

import "github.com/jinzhu/gorm"

// UserProfileImage ユーザープロフィール画像
type UserProfileImage struct {
	BaseModel
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
func (u *UserProfileImage) GetListByUserID(uID int) (userProfileImage []UserProfileImage, db *gorm.DB, err error) {
	whereList := []map[string]interface{}{
		{"UserID": uID},
	}
	option := make(map[string]interface{})

	db, err = GetListWhere(&userProfileImage, "User_ID = :UserID", whereList, option)

	return
}

// GetScanByUserID ユーザーIDからスキャン取得する
func (u *UserProfileImage) GetScanByUserID(uID int, dest interface{}) error {
	whereList := []map[string]interface{}{
		{"UserID": uID},
	}
	option := make(map[string]interface{})

	return GeScanWhere(dest, "user_profile_images", "User_ID = :UserID", whereList, option)
}
