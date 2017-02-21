package follows

import "dotstamp_server/models"

// Add 詳細を保存する
func Add(uID int, cID int) error {
	u := models.UserContributionFollow{
		UserID:             uID,
		UserContributionID: cID,
	}

	return u.Add()
}

// Delete フォローを削除する
func Delete(id uint) error {
	u := models.UserContributionFollow{}
	user, _, err := u.GetByID(id)
	if err != nil {
		return err
	}

	return user.Delete()
}

// GetListByUserContributionID 投稿IDからフォローリストを取得する
func GetListByUserContributionID(cID int) ([]models.UserContributionFollow, error) {
	u := models.UserContributionFollow{}
	r, _, err := u.GetListByUserContributionID(cID)

	return r, err
}

// GetByUserIDAndUserContributionID ユーザIDと投稿IDから取得する
func GetByUserIDAndUserContributionID(uID int, ucID int) (models.UserContributionFollow, error) {
	u := models.UserContributionFollow{}
	r, _, err := u.GetByUserIDAndUserContributionID(uID, ucID)

	return r, err
}

// GetCountByUserIDAndUserContributionID ユーザIDと投稿IDから件数を取得する
func GetCountByUserIDAndUserContributionID(uID int, ucID int) (int, error) {
	u := models.UserContributionFollow{}
	_, db, err := u.GetByUserIDAndUserContributionID(uID, ucID)
	if err != nil {
		return 0, err
	}

	c := 0
	db.Table("user_contribution_follows").Count(&c)

	return c, err
}
