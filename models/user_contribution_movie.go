package models

import (
	"strconv"

	"github.com/jinzhu/gorm"
)

const (
	// MovieTypeYoutube 動画タイプ：Youtube
	MovieTypeYoutube = 1
)

// UserContributionMovie ユーザ投稿動画
type UserContributionMovie struct {
	gorm.Model
	UserContributionID int    `json:"user_contribution_id"`
	MovieType          int    `json:"movie_type"`
	MovieID            string `json:"movie_id"`
	MovieStatus        int    `json:"movie_status"`
}

// Add 追加する
func (u *UserContributionMovie) Add() error {
	return Create(u)
}

// Save 保存する
func (u *UserContributionMovie) Save() error {
	return Save(u)
}

// GetByUserContributionID 投稿IDから取得する
func (u *UserContributionMovie) GetByUserContributionID(uID int, t int) (userContributionMovie UserContributionMovie, db *gorm.DB, err error) {
	whereList := []map[string]interface{}{
		{"UserContributionID": uID},
		{"MovieType": t},
	}
	option := make(map[string]interface{})

	db, err = GetWhere(&userContributionMovie, "User_contribution_ID = :UserContributionID AND movie_type = :MovieType", whereList, option)

	return
}

// GetListByUserContributionIDList 投稿IDリストからリスト取得する
func (u *UserContributionMovie) GetListByUserContributionIDList(uID []int, t int) (userContributionMovie []UserContributionMovie, db *gorm.DB, err error) {
	whereList := []map[string]interface{}{
		{"UserContributionID": uID},
		{"MovieType": t},
	}
	option := make(map[string]interface{})

	db, err = GetListWhere(&userContributionMovie, "User_contribution_ID IN :UserContributionID AND movie_type = :MovieType AND movie_status = "+strconv.Itoa(StatusPublic), whereList, option)

	return
}
