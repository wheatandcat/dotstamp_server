package question

import "github.com/wheatandcat/dotstamp_server/models"

// Add 追加する
func Add(uID int, body string, email string) error {
	log := models.LogQuestion{
		UserID: uID,
		Body:   body,
		Email:  email,
	}

	return log.Add()
}
