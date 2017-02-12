package user

import (
	"dotstamp_server/models"
	"dotstamp_server/utils"
	"time"

	"github.com/mitchellh/mapstructure"
)

// Profile プロフィール
type Profile struct {
	ID      int
	UserID  int
	Created time.Time
}

// GetProfileImageListByUserID ユーザIDからプロフィール画像リストを取得する
func GetProfileImageListByUserID(uID int) (profile []Profile, err error) {
	u := models.UserProfileImage{}

	p := u.GetListByUserID(uID)

	Profile := []Profile{}
	if err = mapstructure.Decode(utils.StructListToMapList(p), &Profile); err != nil {
		return Profile, err
	}

	return Profile, nil
}

// GetIDAndAddProfileImage プロフィール画像を追加してIDを取得する
func GetIDAndAddProfileImage(uID int) (int, error) {
	u := models.UserProfileImage{
		UserID: uID,
	}

	return u.GetIDAndAdd()
}
