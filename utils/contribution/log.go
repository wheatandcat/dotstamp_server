package contributions

import (
	"dotstamp_server/models"
)

// AddLog ログを追加する
func AddLog(userID int, uID int) error {
	u := models.LogUserContribution{
		UserID:             userID,
		UserContributionID: uID,
	}

	return u.Add()
}
