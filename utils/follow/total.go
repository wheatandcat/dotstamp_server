package follows

import "github.com/wheatandcat/dotstamp_server/models"

// TruncateTotal 統計を空にする
func TruncateTotal() error {
	u := models.ContributionTotalFollows{}

	return u.Truncate()
}

// AddTotal 統計を追加する
func AddTotal(uID int, count int) error {
	u := models.ContributionTotalFollows{
		UserContributionID: uID,
		Count:              count,
	}

	return u.Add()
}

// AddTotalMap 統計マップを追加する
func AddTotalMap(m map[int]int) error {
	for id, count := range m {
		if err := AddTotal(id, count); err != nil {
			return err
		}
	}

	return nil
}

// GetTotalListByUserContributionIDList 投稿IDリストから統計を取得する
func GetTotalListByUserContributionIDList(idList []int) ([]models.ContributionTotalFollows, error) {
	u := models.ContributionTotalFollows{}

	r, _, err := u.GetListByUserContributionID(idList)
	return r, err
}

// getToatlMap 統計マップを取得する
func getToatlMap(u []models.ContributionTotalFollows) map[int]int {
	r := map[int]int{}

	for _, v := range u {
		r[v.UserContributionID] = v.Count
	}

	return r
}

// GetTotalMapByUserContributionIDList 投稿IDリストから統計を取得する
func GetTotalMapByUserContributionIDList(idList []int) (map[int]int, error) {
	u, err := GetTotalListByUserContributionIDList(idList)
	if err != nil {
		return map[int]int{}, err
	}

	return getToatlMap(u), nil
}
