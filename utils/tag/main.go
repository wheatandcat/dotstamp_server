package tags

import (
	"dotstamp_server/models"
	"dotstamp_server/utils"
	"strings"

	"github.com/mitchellh/mapstructure"
)

// Tag タグ
type Tag struct {
	ID                 uint
	UserContributionID int
	Name               string
}

// Save 保存する
func Save(id int, n string) error {
	u := models.UserContributionTag{}
	u = u.GetByID(id)
	u.Name = n

	return u.Save()
}

// DeleteByID IDから削除する
func DeleteByID(id int) error {
	u := models.UserContributionTag{}
	u = u.GetByID(id)

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
	userContributionTag := u.GetListByUserContributionID(uID)

	tag := []Tag{}
	if err := mapstructure.Decode(utils.StructListToMapList(userContributionTag), &tag); err != nil {
		return tag, err
	}

	return tag, nil
}

// GetMapByUserContributionIDList 投稿IDからマップを取得する
func GetMapByUserContributionIDList(uIDList []int) (tagMap map[int][]Tag, err error) {
	u := &models.UserContributionTag{}
	userContributionTag := u.GetListByUserContributionIDList(uIDList)

	tagList := []Tag{}
	if err = mapstructure.Decode(utils.StructListToMapList(userContributionTag), &tagList); err != nil {
		return map[int][]Tag{}, err
	}

	tagMap = map[int][]Tag{}
	for _, tag := range tagList {
		tagMap[tag.UserContributionID] = append(tagMap[tag.UserContributionID], tag)
	}

	return tagMap, nil
}
