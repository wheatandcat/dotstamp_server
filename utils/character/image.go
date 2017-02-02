package characters

import (
	"errors"
	"dotstamp_server/models"
	"dotstamp_server/utils"
)

// Image 画像
type Image struct {
	ID          int
	CharacterID int
	Priority    int
	FileName    string
}

// AddImage 画像を追加する
func AddImage(uID int, cID int, p int) (int, error) {
	u := &models.UserCharacterImage{
		UserID:      uID,
		CharacterID: cID,
		Priority:    p,
	}

	e := u.Add()

	return u.ID, e
}

// GetImageListByUserID ユーザーIDからリストを取得する
func GetImageListByUserID(uID int) (image []Image) {
	u := models.UserCharacterImage{}
	r := u.GetListByUserID(uID)

	for _, v := range r {
		st := Image{
			ID:          v.ID,
			CharacterID: v.CharacterID,
			Priority:    v.Priority,
			FileName:    GetImageName(v.ID),
		}

		image = append(image, st)
	}

	return
}

// DeleteByID IDから削除する
func DeleteByID(id int, userID int) error {
	u := models.UserCharacterImage{}
	userCharacterImage := u.GetByID(id)

	if userCharacterImage.UserID != userID {
		return errors.New("User_ID is wrong")
	}

	return userCharacterImage.Delete()
}

// GetImageName 画像名を取得する
func GetImageName(id int) string {
	return utils.IntToEncryption(id) + ".jpg"
}
