package tags

import (
	"dotstamp_server/models"
	"strings"
)

// Tag タグ
type Tag struct {
	ID                 uint
	UserContributionID int
	Name               string
}

// Save 保存する
func Save(id int, n string) (err error) {
	u := models.UserContributionTag{}
	u, _, err = u.GetByID(id)
	if err != nil {
		return err
	}

	u.Name = n

	return u.Save()
}

// DeleteByID IDから削除する
func DeleteByID(id int) (err error) {
	u := models.UserContributionTag{}
	u, _, err = u.GetByID(id)
	if err != nil {
		return err
	}

	return u.Delete()
}

// AddList 追加する
func AddList(uID int, n string) error {
	n = strings.Replace(strings.TrimSpace(n), "　", " ", -1)
	namelist := strings.Split(n, " ")

	userContributionTag := []models.UserContributionTag{}

	for _, name := range namelist {
		u := models.UserContributionTag{
			UserContributionID: uID,
			Name:               name,
		}

		userContributionTag = append(userContributionTag, u)
	}

	uct := models.UserContributionTag{}

	return uct.AddList(userContributionTag)
}

// GetListByUserContributionID 投稿IDからリストを取得する
func GetListByUserContributionID(uID int) ([]Tag, error) {
	u := &models.UserContributionTag{}
	tag := []Tag{}

	_, db, err := u.GetListByUserContributionID(uID)
	if err != nil {
		return tag, err
	}

	err = db.Table("user_contribution_tags").Scan(&tag).Error
	if err != nil {
		return tag, err
	}

	return tag, nil
}

// GetMapByUserContributionIDList 投稿IDからマップを取得する
func GetMapByUserContributionIDList(uIDList []int) (map[int][]Tag, error) {
	tagMap := map[int][]Tag{}

	u := &models.UserContributionTag{}
	_, db, err := u.GetListByUserContributionIDList(uIDList)
	if err != nil {
		return tagMap, err
	}

	tagList := []Tag{}
	err = db.Table("user_contribution_tags").Scan(&tagList).Error
	if err != nil {
		return tagMap, err
	}

	for _, tag := range tagList {
		tagMap[tag.UserContributionID] = append(tagMap[tag.UserContributionID], tag)
	}

	return tagMap, nil
}

// GetTagNameJoin 連結したタグ名を取得する
func GetTagNameJoin(uID int) (string, error) {
	t, err := GetListByUserContributionID(uID)
	if err != nil {
		return "", err
	}

	list := []string{}
	for _, v := range t {
		list = append(list, v.Name)
	}

	return strings.Join(list, ","), nil
}
