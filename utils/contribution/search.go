package contributions

import "dotstamp_server/models"

// GetSearchWordByBody 本文から検索文を取得する
func GetSearchWordByBody(body string) (s string, err error) {
	b, err := StirngToGetBody(body)
	if err != nil {
		return "", err
	}

	for _, v := range b {
		s += v.Body
	}

	return s, nil
}

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

// AddOrSaveSearch 検索を追加か保存する
func AddOrSaveSearch(uID int, b string) error {
	s, err := GetSearchWordByBody(b)
	if err != nil {
		return err
	}

	u, err := GetSearchByUserContributionID(uID)
	if err != nil {
		return err
	}

	if u.ID == uint(0) {
		return AddSearch(uID, s)
	}

	u.Search = s
	return u.Save()
}

// DeleteSearchByUserContributionID 投稿IDから削除する
func DeleteSearchByUserContributionID(uID int) error {
	u, err := GetSearchByUserContributionID(uID)
	if err != nil {
		return err
	}

	if u.ID == uint(0) {
		return nil
	}

	return u.Delete()
}
