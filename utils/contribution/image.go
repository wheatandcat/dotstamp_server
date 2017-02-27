package contributions

import "dotstamp_server/models"

// GetImageIDAndAdd 追加して画像IDを取得する
func GetImageIDAndAdd(userContributionID int) (uint, error) {
	l := models.LogContributionImage{
		UserContributionID: userContributionID,
	}

	return l.GetIDAndAdd()
}
