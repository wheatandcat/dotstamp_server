package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// UserContributionSearch ユーザー投稿検索
type UserContributionSearch struct {
	ID                 uint   `gorm:"primary_key"`
	UserContributionID int    `json:"user_contribution_id"`
	Search             string `gorm:"index:search"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

// Add 追加する
func (u *UserContributionSearch) Add() error {
	return Create(u)
}

// Save 保存する
func (u *UserContributionSearch) Save() error {
	return Save(u)
}

// Delete 削除する
func (u *UserContributionSearch) Delete() error {
	return Delete(u)
}

// GetByUserContributionID 投稿IDから取得する
func (u *UserContributionSearch) GetByUserContributionID(id int) (userContributionSearch UserContributionSearch, db *gorm.DB, err error) {
	whereList := []map[string]interface{}{
		{"UserContributionID": id},
	}
	option := make(map[string]interface{})

	db, err = GetWhere(&userContributionSearch, "User_contribution_ID = :UserContributionID", whereList, option)

	return
}

// GetListBySearch 検索からリストを取得する
func (u *UserContributionSearch) GetListBySearch(search string) (userContributionSearch []UserContributionSearch, db *gorm.DB, err error) {
	whereList := []map[string]interface{}{
		{"Search": "+" + search},
	}
	option := make(map[string]interface{})

	db, err = GetListWhere(&userContributionSearch, "MATCH(`search`) AGAINST( :Search IN BOOLEAN MODE)", whereList, option)

	return
}
