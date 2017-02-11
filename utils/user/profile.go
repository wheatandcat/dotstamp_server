package user

import (
	"dotstamp_server/models"
)

// GetProfileImageListByUserID ユーザIDからプロフィール画像リストを取得する
func GetProfileImageListByUserID(uID int) []models.UserProfileImage {
	u := models.UserProfileImage{}

	return u.GetListByUserID(uID)
}

// GetIDAndAddProfileImage プロフィール画像を追加してIDを取得する
func GetIDAndAddProfileImage(uID int) (int, error) {
	u := models.UserProfileImage{
		UserID: uID,
	}

	return u.GetIDAndAdd()
}
