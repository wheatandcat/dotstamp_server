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
	uc := GetByUserContributionID(userContributionID)
	if uc.UserID != userID {
		return errors.New("difference UserID")
	}

	uc.Title = title

	return uc.Save()
}

// DeleteByID 削除する
func DeleteByID(userContributionID int, userID int) error {
	uc := GetByUserContributionID(userContributionID)

	if uc.UserID != userID {
		return errors.New("difference UserID")
	}

	if e := uc.Delete(); e != nil {
		return e
	}

	ucd := GetDetailByUserContributionID(userContributionID)

	return ucd.Delete()
}

// GetByUserContributionID 投稿IDから取得する
func GetByUserContributionID(userContributionID int) models.UserContribution {
	userContribution := &models.UserContribution{}

	return userContribution.GetByID(userContributionID)
}

// GetListByUserID ユーザIDから取得する
func GetListByUserID(userID int) []models.UserContribution {
	userContribution := &models.UserContribution{}

	return userContribution.GetListByUserID(userID)
}

// GetContributionByUserContributionID 投稿IDから取得する
func GetContributionByUserContributionID(userContributionID int) (c Contribution, err error) {
	uc := GetByUserContributionID(userContributionID)
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
func GetByTop(offset int, size int) (contributionList []Contribution, err error) {
	uc := &models.UserContribution{}
	userContribution := uc.GetByTop(offset, size)

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
