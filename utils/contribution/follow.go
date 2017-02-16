package contributions

import "dotstamp_server/models"

// AddFollow 詳細を保存する
func AddFollow(uID int, cID int) error {
	u := models.UserContributionFollow{
		UserID:             uID,
		UserContributionID: cID,
	}

	return u.Add()
}

// DeleteFollow フォローを削除する
func DeleteFollow(id uint) error {
	u := models.UserContributionFollow{}
	user, _, err := u.GetByID(id)
	if err != nil {
		return err
	}

	return user.Delete()
}

// GetFollowListByUserContributionID 投稿IDからフォローリストを取得する
func GetFollowListByUserContributionID(cID int) ([]models.UserContributionFollow, error) {
	u := models.UserContributionFollow{}
	r, _, err := u.GetListByUserContributionID(cID)

	return r, err
}
