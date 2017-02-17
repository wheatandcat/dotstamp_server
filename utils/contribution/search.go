package contributions

import "dotstamp_server/models"

// AddSearch 検索を追加する
func AddSearch(uID int, search string) error {
	u := models.UserContributionSearch{
		UserContributionID: uID,
		Search:             search,
	}

	return u.Add()
}

// GetSearchByUserContributionID 投稿IDから取得する
func GetSearchByUserContributionID(uID int) (models.UserContributionSearch, error) {
	u := models.UserContributionSearch{}
	r, _, err := u.GetByUserContributionID(uID)

	return r, err
}
