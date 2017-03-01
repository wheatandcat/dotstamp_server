package contributions

import "dotstamp_server/models"

// SearchValue 検索値
type SearchValue struct {
	UserContributionID int
	Search             string
	Order              int
}

// SearchWord 検索文
type SearchWord struct {
	Title string
	Body  string
	Tag   string
}

// JoinSearchWord 検索文を連結する
func JoinSearchWord(s SearchWord) string {
	return s.Title + "/" + s.Body + "/" + s.Tag
}

// GetSearchWordBody 検索本文を取得する
func GetSearchWordBody(body string) (s string, err error) {
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

// GetSearchListByUserContributionIDList 投稿IDリストからリスト取得する
func GetSearchListByUserContributionIDList(uID []int) ([]models.UserContributionSearch, error) {
	u := models.UserContributionSearch{}
	r, _, err := u.GetListByUserContributionIDList(uID)

	return r, err
}

// AddOrSaveSearch 検索を追加か保存する
func AddOrSaveSearch(uID int, s string) error {
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

// GetSearchValueListBySearch 検索から検索値リストを取得する
func GetSearchValueListBySearch(search string, order string, limit int, offset int) ([]SearchValue, error) {
	s := []SearchValue{}

	u := models.UserContributionSearch{}
	user, _, err := u.GetListBySearch(search, order, limit, offset)
	if err != nil {
		return s, err
	}

	if len(user) == 0 {
		return s, nil
	}

	for key, v := range user {
		tmp := SearchValue{
			UserContributionID: v.UserContributionID,
			Search:             v.Search,
			Order:              key,
		}

		s = append(s, tmp)
	}

	return s, nil
}

// SaveToFollowCount フォロー数を保存する
func SaveToFollowCount(u []models.UserContributionSearch, m map[int]int) error {
	for _, v := range u {
		if v.FollowCount != m[v.UserContributionID] {
			v.FollowCount = m[v.UserContributionID]
			if err := v.Save(); err != nil {
				return err
			}
		}
	}

	return nil
}

// GetCountBySearch 検索から数を取得する
func GetCountBySearch(search string, order string) (int, error) {
	u := models.UserContributionSearch{}

	return u.GetCountBySearch(search, order)
}
