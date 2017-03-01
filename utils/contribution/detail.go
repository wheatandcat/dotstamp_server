package contributions

import (
	"dotstamp_server/models"
	"dotstamp_server/utils/character"
	"encoding/json"
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
	ID        int
	FileName  string
	VoiceType int
}

// SaveDetail 詳細を保存する
func SaveDetail(userContributionID int, body string) error {
	ucd, err := GetDetailByUserContributionID(userContributionID)
	if err != nil {
		return err
	}

	b, err := StirngToSaveBody(body)
	if err != nil {
		return err
	}

	st, err := json.Marshal(b)
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
		b[k].Character.FileName = characters.GetImageName(uint(v.Character.ID))
	}

	return b, err
}

// GetDetailByUserContributionID 投稿IDから投稿詳細を取得する
func GetDetailByUserContributionID(uID int) (models.UserContributionDetail, error) {
	u := &models.UserContributionDetail{}

	r, _, err := u.GetByUserContributionID(uID)
	if err != nil {
		return r, err
	}

	return r, err
}

// GetBodyByUserContributionID 投稿IDから本文を取得する
func GetBodyByUserContributionID(uID int) ([]GetBody, error) {
	b := []GetBody{}
	u, err := GetDetailByUserContributionID(uID)
	if err != nil {
		return b, err
	}

	b, err = StirngToGetBody(u.Body)
	if err != nil {
		return b, err
	}

	return b, nil
}
