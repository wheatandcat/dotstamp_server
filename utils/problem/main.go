package problem

import "github.com/wheatandcat/dotstamp_server/models"

// Add 追加する
func Add(userID int, uID int, programType int) error {
	log := models.LogProblemContributionReport{
		UserID:             userID,
		UserContributionID: uID,
		Type:               programType,
	}

	return log.Add()
}
