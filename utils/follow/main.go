package follows

import "github.com/wheatandcat/dotstamp_server/models"

// OrderValue 順番値
type OrderValue struct {
	UserContributionID int
	Order              int
}

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

// GetCountByUserContributionID 投稿IDからフォロー数を取得する
func GetCountByUserContributionID(cID int) (int, error) {
	u := models.UserContributionFollow{}

	return u.GetCountByUserContributionID(cID)
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

	return u.GetCountByUserIDAndUserContributionID(uID, ucID)
}

// GetListByUserID ユーザIDからリストを取得する
func GetListByUserID(uID int, order string, limit int, offset int) ([]models.UserContributionFollow, error) {
	u := models.UserContributionFollow{}
	r, _, err := u.GetListByUserID(uID, order, limit, offset)

	return r, err
}

// GetOrderValueListByUserID ユーザIDから順番値リストを取得する
func GetOrderValueListByUserID(uID int, order string, limit int, offset int) (o []OrderValue, err error) {
	u, err := GetListByUserID(uID, order, limit, offset)
	if err != nil {
		return o, err
	}

	if len(u) == 0 {
		return o, nil
	}

	for key, v := range u {
		tmp := OrderValue{
			UserContributionID: v.UserContributionID,
			Order:              key,
		}

		o = append(o, tmp)
	}

	return o, nil
}

// GetListByUserContributionIDList 投稿IDリストからフォローリストを取得する
func GetListByUserContributionIDList(cID []int) ([]models.UserContributionFollow, error) {
	u := models.UserContributionFollow{}
	r, _, err := u.GetListByUserContributionIDList(cID)

	return r, err
}

// GetFollowCountMap フォロー数マップを取得する
func GetFollowCountMap(u []models.UserContributionFollow) map[int]int {
	m := map[int]int{}

	for _, v := range u {
		id := int(v.UserContributionID)
		if _, ok := m[id]; !ok {
			m[id] = 0
		}

		m[id]++
	}

	return m
}

// GetCountByUserID ユーザIDから数を取得する
func GetCountByUserID(uID int, order string) (int, error) {
	u := models.UserContributionFollow{}

	return u.GetCountByUserID(uID, order)
}
