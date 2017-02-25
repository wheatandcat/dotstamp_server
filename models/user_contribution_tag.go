package models

import "github.com/jinzhu/gorm"

// UserContributionTag ユーザー投稿タグ
type UserContributionTag struct {
	gorm.Model
	UserContributionID int `json:"user_contribution_id"`
	Name               string
}

// Add 追加する
func (uc *UserContributionTag) Add() (err error) {
	return Create(uc)
}

// AddList リストを追加する
func (uc *UserContributionTag) AddList(u []UserContributionTag) (err error) {
	for _, user := range u {
		if err = Create(&user); err != nil {
			return err
		}
	}

	return nil
}

// GetListByUserContributionID 投稿IDから取得する
func (uc *UserContributionTag) GetListByUserContributionID(id int) (userContributionTag []UserContributionTag, db *gorm.DB, err error) {
	whereList := []map[string]interface{}{
		{"UserContributionID": id},
	}
	option := make(map[string]interface{})

	db, err = GetListWhere(&userContributionTag, "User_contribution_ID = :UserContributionID", whereList, option)

	return
}

// GetScanListByUserContributionID 投稿IDからスキャン取得する
func (uc *UserContributionTag) GetScanListByUserContributionID(id int, dest interface{}) error {
	whereList := []map[string]interface{}{
		{"UserContributionID": id},
	}
	option := make(map[string]interface{})

	return GeScanWhere(dest, "user_contribution_tags", "User_contribution_ID = :UserContributionID", whereList, option)
}

// GetListByUserContributionIDList 投稿IDリストから取得する
func (uc *UserContributionTag) GetListByUserContributionIDList(idList []int) (userContributionTag []UserContributionTag, db *gorm.DB, err error) {
	whereList := []map[string]interface{}{
		{"UserContributionID": idList},
	}
	option := make(map[string]interface{})

	db, err = GetListWhere(&userContributionTag, "User_contribution_ID IN :UserContributionID", whereList, option)

	return
}

// GetScanListByUserContributionIDList 投稿IDリストからスキャン取得する
func (uc *UserContributionTag) GetScanListByUserContributionIDList(idList []int, dest interface{}) error {
	whereList := []map[string]interface{}{
		{"UserContributionID": idList},
	}
	option := make(map[string]interface{})

	return GeScanWhere(dest, "user_contribution_tags", "User_contribution_ID IN :UserContributionID", whereList, option)
}

// GetByID IDから取得する
func (uc *UserContributionTag) GetByID(id int) (userContributionTag UserContributionTag, db *gorm.DB, err error) {
	whereList := []map[string]interface{}{
		{"ID": id},
	}
	option := make(map[string]interface{})

	db, err = GetWhere(&userContributionTag, "ID = :ID", whereList, option)

	return
}

// Save 保存する
func (uc *UserContributionTag) Save() error {
	return Save(uc)
}

// Delete 削除する
func (uc *UserContributionTag) Delete() error {
	return Delete(uc)
}
