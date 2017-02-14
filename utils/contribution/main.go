package contributions

import (
	"dotstamp_server/models"
	"dotstamp_server/utils/tag"
	"dotstamp_server/utils/user"
	"errors"
	"time"
)

// Contribution 投稿情報
type Contribution struct {
	ID                uint
	User              user.User
	Title             string
	FewDaysAgoMessage string
	Tag               []tags.Tag
	Body              []GetBody
	UpdatedAt         time.Time
	CreatedAt         time.Time
}

// Add 投稿する
func Add(userID int, title string, body string) (uint, error) {
	userContribution := &models.UserContribution{
		UserID: userID,
		Title:  title,
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
func Save(userContributionID int, userID int, title string) error {
	uc, err := GetByUserContributionID(userContributionID)
	if err != nil {
		return err
	}
	if uc.UserID != userID {
		return errors.New("difference UserID")
	}

	uc.Title = title

	return uc.Save()
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

// GetByTop 新着を取得する
func GetByTop(offset int, size int) ([]Contribution, error) {
	uc := &models.UserContribution{}
	contributionList := []Contribution{}
	userContribution, _, err := uc.GetByTop(offset, size)
	if err != nil {
		return contributionList, err
	}
	if len(userContribution) == 0 {
		return contributionList, nil
	}

	var idList []int
	var userIDList []int
	for _, val := range userContribution {
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

	for _, val := range userContribution {

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
