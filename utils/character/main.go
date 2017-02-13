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
func GetListByUserID(uID int) (character []Character) {
	u := models.UserCharacter{}

	r := u.GetListByUserID(uID)

	for _, v := range r {
		c := Character{
			ID:       v.ID,
			Name:     v.Name,
			Info:     v.Info,
			Priority: v.Priority,
		}

		character = append(character, c)
	}

	return
}
