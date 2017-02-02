package contributions

import (
	"encoding/json"
	"dotstamp_server/models"
	"dotstamp_server/utils/character"
)

// SaveBody 本本
type SaveBody struct {
	Priority      int
	Body          string
	DirectionType int
	TalkType      int
	Character     SaveCharacter
}

// SaveCharacter 保存キャラクター
type SaveCharacter struct {
	ID int
}

// GetBody 取得本文
type GetBody struct {
	Priority      int
	Body          string
	DirectionType int
	TalkType      int
	Character     GetCharacter
}

// GetCharacter 取得キャラクター
type GetCharacter struct {
	ID       int
	FileName string
}

// SaveDetail 詳細を保存する
func SaveDetail(userContributionID int, body string) error {
	ucd := GetDetailByUserContributionID(userContributionID)

	data, err := StirngToSaveBody(body)
	if err != nil {
		return err
	}

	st, err := json.Marshal(data)
	if err != nil {
		return err
	}

	ucd.Body = string(st)

	ucd.Save()

	return nil
}

// StirngToSaveBody 保存本文に変換する
func StirngToSaveBody(body string) (b []SaveBody, err error) {
	bytes := []byte(body)
	err = json.Unmarshal(bytes, &b)
	if err != nil {
		return b, err
	}

	return b, err
}

// StirngToGetBody 取得本文に変換する
func StirngToGetBody(body string) (b []GetBody, err error) {
	bytes := []byte(body)
	err = json.Unmarshal(bytes, &b)
	if err != nil {
		return b, err
	}

	for k, v := range b {
		b[k].Character.FileName = characters.GetImageName(v.Character.ID)
	}

	return b, err
}

// GetDetailByUserContributionID 投稿IDから投稿詳細を取得する
func GetDetailByUserContributionID(uID int) models.UserContributionDetail {
	u := &models.UserContributionDetail{}

	return u.GetByUserContributionID(uID)
}

// GetBodyByUserContributionID 投稿IDから本文を取得する
func GetBodyByUserContributionID(uID int) ([]GetBody, error) {
	u := GetDetailByUserContributionID(uID)

	b, err := StirngToGetBody(u.Body)
	if err != nil {
		return b, err
	}

	return b, nil
}
