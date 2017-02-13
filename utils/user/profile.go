package user

import (
	"dotstamp_server/models"
	"time"
)

// Profile プロフィール
type Profile struct {
	ID      uint
	UserID  int
	Created time.Time
}

// GetProfileImageListByUserID ユーザIDからプロフィール画像リストを取得する
func GetProfileImageListByUserID(uID int) ([]Profile, error) {

	profile := []Profile{}
	u := models.UserProfileImage{}

	_, db, err := u.GetListByUserID(uID)
	if err != nil {
		return profile, err
	}

	err = db.Table("user_profile_images").Scan(&profile).Error
	if err != nil {
		return profile, err
	}

	return profile, nil
}

// GetIDAndAddProfileImage プロフィール画像を追加してIDを取得する
func GetIDAndAddProfileImage(uID int) (uint, error) {
	u := models.UserProfileImage{
		UserID: uID,
	}

	return u.GetIDAndAdd()
}
