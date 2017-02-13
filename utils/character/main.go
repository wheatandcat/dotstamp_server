package characters

import (
	"dotstamp_server/models"
)

// Character キャラクター
type Character struct {
	ID       uint
	Name     string
	Info     string
	Priority int
}

// Add 追加する
func Add(uID int, name string, info string, p int) (uint, error) {
	u := &models.UserCharacter{
		UserID:   uID,
		Name:     name,
		Info:     info,
		Priority: p,
	}

	err := u.Add()

	return u.ID, err
}

// GetListByUserID ユーザーIDからリストを取得する
func GetListByUserID(uID int) ([]Character, error) {
	u := models.UserCharacter{}

	character := []Character{}
	_, db, err := u.GetListByUserID(uID)

	db.Table("user_characters").Scan(&character)

	return character, err
}
