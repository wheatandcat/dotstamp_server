package models

import (
	"strconv"

	"github.com/jinzhu/gorm"
)

// ViewStatusPublic 公開状態
const ViewStatusPublic = 1

// ViewStatusPrivate プライベート状態
const ViewStatusPrivate = 2

// UserContribution ユーザ投稿
type UserContribution struct {
	gorm.Model
	UserID     int `json:"user_id"`
	Title      string
	ViewStatus int `json:"view_status"`
}

// GetIDAndAdd 投稿してIDを取得する
func (u *UserContribution) GetIDAndAdd() (uint, error) {
	if err := Create(u); err != nil {
		return 0, err
	}

	return u.ID, nil
}

// Save 保存する
func (u *UserContribution) Save() error {
	return Save(u)
}

// Delete 削除する
func (u *UserContribution) Delete() error {
	return Delete(u)
}

// GetByID 投稿IDから取得する
func (u *UserContribution) GetByID(id int) (userContribution UserContribution, db *gorm.DB, err error) {
	whereList := []map[string]interface{}{
		{"ID": id},
	}
	option := make(map[string]interface{})

	db, err = GetWhere(&userContribution, "ID = :ID", whereList, option)
	return
}

// GetListByUserID 投稿IDからリスト取得する
func (u *UserContribution) GetListByUserID(userID int, order string, limit int, offset int) (userContribution []UserContribution, db *gorm.DB, err error) {
	whereList := []map[string]interface{}{
		{"UserID": userID},
	}

	option := map[string]interface{}{
		"order":  order,
		"limit":  limit,
		"offset": offset,
	}

	db, err = GetListWhere(&userContribution, "User_ID = :UserID", whereList, option)
	return
}

// GetCountByUserID 投稿IDから数を取得する
func (u *UserContribution) GetCountByUserID(userID int, order string) (int, error) {
	userContribution := UserContribution{}

	whereList := []map[string]interface{}{
		{"UserID": userID},
	}

	option := make(map[string]interface{})

	return GetCount(&userContribution, "user_contributions", "User_ID = :UserID AND Deleted_at IS NULL", whereList, option)
}

// GetByTop 新着から投稿リスト取得する
func (u *UserContribution) GetByTop(o int, s int) (userContributionList []UserContribution, db *gorm.DB, err error) {
	whereList := []map[string]interface{}{}

	optionMap := map[string]interface{}{
		"order":  "ID desc",
		"limit":  s,
		"offset": o,
	}

	db, err = GetListWhere(&userContributionList, "View_status = "+strconv.Itoa(ViewStatusPublic), whereList, optionMap)
	return
}

// GetListByIDList IDリストから投稿リストを取得する
func (u *UserContribution) GetListByIDList(idList []int) (userContributionList []UserContribution, db *gorm.DB, err error) {
	whereList := []map[string]interface{}{
		{"IDList": idList},
	}
	optionMap := make(map[string]interface{})

	db, err = GetListWhere(&userContributionList, "ID IN :IDList", whereList, optionMap)
	return
}

// GetListByViewStatusPublic IDリストから投稿リストを取得する
func (u *UserContribution) GetListByViewStatusPublic() (userContributionList []UserContribution, db *gorm.DB, err error) {
	whereList := []map[string]interface{}{}

	optionMap := map[string]interface{}{
		"select": "id",
	}

	db, err = GetListWhere(&userContributionList, "View_status = "+strconv.Itoa(ViewStatusPublic), whereList, optionMap)
	return
}
