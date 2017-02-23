package contributions

import (
	"dotstamp_server/models"
	"dotstamp_server/utils/follow"
	"dotstamp_server/utils/tag"
	"dotstamp_server/utils/user"
	"errors"
	"time"
)

// Contribution 投稿
type Contribution struct {
	ID                uint
	User              user.User
	Title             string
	FewDaysAgoMessage string
	Tag               []tags.Tag
	Body              []GetBody
	Search            string
	UpdatedAt         time.Time
	CreatedAt         time.Time
}

// Add 投稿する
func Add(userID int, title string, body string, v int) (uint, error) {
	userContribution := &models.UserContribution{
		UserID:     userID,
		Title:      title,
		ViewStatus: v,
	}

	userContributionID, err := userContribution.GetIDAndAdd()
	if err != nil {
		return 0, err
	}

	userContributionDetail := &models.UserContributionDetail{
		UserContributionID: int(userContributionID),
		Body:               body,
	}
	userContributionDetail.Add()

	return userContributionID, nil
}

// Save 保存する
func Save(userContributionID int, userID int, title string, v int) error {
	u, err := GetByUserContributionID(userContributionID)
	if err != nil {
		return err
	}
	if u.UserID != userID {
		return errors.New("difference UserID")
	}

	u.Title = title
	u.ViewStatus = v

	return u.Save()
}

// DeleteByID 削除する
func DeleteByID(userContributionID int, userID int) error {
	uc, err := GetByUserContributionID(userContributionID)
	if err != nil {
		return err
	}

	if uc.UserID != userID {
		return errors.New("difference UserID")
	}

	if e := uc.Delete(); e != nil {
		return e
	}

	ucd, err := GetDetailByUserContributionID(userContributionID)
	if err != nil {
		return err
	}

	return ucd.Delete()
}

// GetByUserContributionID 投稿IDから取得する
func GetByUserContributionID(userContributionID int) (models.UserContribution, error) {
	userContribution := &models.UserContribution{}

	r, _, err := userContribution.GetByID(userContributionID)

	return r, err
}

// GetListByUserID ユーザIDから取得する
func GetListByUserID(userID int) ([]models.UserContribution, error) {
	userContribution := &models.UserContribution{}

	r, _, err := userContribution.GetListByUserID(userID)

	return r, err
}

// GetContributionByUserContributionID 投稿IDから取得する
func GetContributionByUserContributionID(userContributionID int) (c Contribution, err error) {
	uc, err := GetByUserContributionID(userContributionID)
	if err != nil {
		return c, err
	}
	var u user.User

	if u, err = user.GetByUserID(uc.UserID); err != nil {
		return c, err
	}

	var tag []tags.Tag
	if tag, err = tags.GetListByUserContributionID(userContributionID); err != nil {
		return c, err
	}

	var body []GetBody
	if body, err = GetBodyByUserContributionID(userContributionID); err != nil {
		return c, err
	}

	user := user.User{
		ID:   u.ID,
		Name: u.Name,
	}

	contribution := Contribution{
		ID:        uc.ID,
		User:      user,
		Title:     uc.Title,
		Tag:       tag,
		Body:      body,
		UpdatedAt: uc.UpdatedAt,
		CreatedAt: uc.CreatedAt,
	}

	return contribution, nil
}

// getContributionList 投稿リストを取得する
func getContributionList(u []models.UserContribution) (contributionList []Contribution, err error) {
	if len(u) == 0 {
		return contributionList, nil
	}

	var idList []int
	var userIDList []int
	for _, val := range u {
		idList = append(idList, int(val.ID))
		userIDList = append(idList, int(val.UserID))
	}

	var tagMap map[int][]tags.Tag
	if tagMap, err = tags.GetMapByUserContributionIDList(idList); err != nil {
		return contributionList, err
	}

	var userMap map[int]user.User
	if userMap, err = user.GetMaptByUserIDList(userIDList); err != nil {
		return contributionList, err
	}

	for _, val := range u {
		if len(tagMap[int(val.ID)]) == 0 {
			tagMap[int(val.ID)] = []tags.Tag{}
		}

		c := Contribution{
			ID:                val.ID,
			User:              userMap[val.UserID],
			Title:             val.Title,
			CreatedAt:         val.CreatedAt,
			UpdatedAt:         val.UpdatedAt,
			FewDaysAgoMessage: "",
			Tag:               tagMap[int(val.ID)],
		}
		contributionList = append(contributionList, c)
	}

	return contributionList, nil
}

// GetListByTop 新着を取得する
func GetListByTop(offset int, size int) ([]Contribution, error) {
	u := &models.UserContribution{}
	userContribution, _, err := u.GetByTop(offset, size)
	if err != nil {
		return []Contribution{}, err
	}

	return getContributionList(userContribution)
}

// GetListBySearchValue 検索値からリストを取得する
func GetListBySearchValue(s []SearchValue) ([]Contribution, error) {
	idList := []int{}
	for _, v := range s {
		idList = append(idList, v.UserContributionID)
	}

	u := &models.UserContribution{}
	contributionList := []Contribution{}
	userContribution, _, err := u.GetListByIDList(idList)
	if err != nil {
		return contributionList, err
	}

	m := map[int]models.UserContribution{}
	orderMap := map[int]int{}
	for _, v := range s {
		orderMap[v.UserContributionID] = v.Order
	}

	keyList := []int{}
	for _, v := range userContribution {
		m[orderMap[int(v.ID)]] = v
		keyList = append(keyList, int(v.ID))

	}

	userContributionList := []models.UserContribution{}
	for v := range keyList {
		userContributionList = append(userContributionList, m[v])
	}

	r, err := getContributionList(userContributionList)
	if err != nil {
		return contributionList, err
	}

	for k := range r {
		for _, v := range s {
			if r[k].ID == uint(v.UserContributionID) {
				r[k].Search = v.Search
			}
		}
	}

	return r, nil
}

// GetListByFollowOrderValue フォロー順からリストを取得する
func GetListByFollowOrderValue(f []follows.OrderValue) ([]Contribution, error) {
	idList := []int{}
	for _, v := range f {
		idList = append(idList, v.UserContributionID)
	}

	u := &models.UserContribution{}
	contributionList := []Contribution{}
	userContribution, _, err := u.GetListByIDList(idList)
	if err != nil {
		return contributionList, err
	}

	m := map[int]models.UserContribution{}
	orderMap := map[int]int{}
	for _, v := range f {
		orderMap[v.UserContributionID] = v.Order
	}

	keyList := []int{}
	for _, v := range userContribution {
		m[orderMap[int(v.ID)]] = v
		keyList = append(keyList, int(v.ID))
	}

	userContributionList := []models.UserContribution{}
	for v := range keyList {
		userContributionList = append(userContributionList, m[v])
	}

	return getContributionList(userContributionList)
}

// GetViewStatusPublicIDList 公開状態のIDリストを取得する
func GetViewStatusPublicIDList() ([]int, error) {
	r := []int{}

	u := models.UserContribution{}
	user, _, err := u.GetListByViewStatusPublic()
	if err != nil {
		return r, err
	}

	for _, v := range user {
		r = append(r, int(v.ID))
	}

	return r, nil
}
