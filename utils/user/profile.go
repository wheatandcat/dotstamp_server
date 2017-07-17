package user

import (
	"time"

	"github.com/wheatandcat/dotstamp_server/models"
)

// Profile プロフィール
type Profile struct {
	ID      uint      `json:"id"`
	UserID  int       `json:"userID"`
	Created time.Time `json:"created"`
}

// GetProfileImageListByUserID ユーザIDからプロフィール画像リストを取得する
func GetProfileImageListByUserID(uID int) ([]Profile, error) {

	profile := []Profile{}
	u := models.UserProfileImage{}

	err := u.GetScanByUserID(uID, &profile)
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
