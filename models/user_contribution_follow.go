package models

import "github.com/jinzhu/gorm"

// UserContributionFollow ユーザー投稿フォロー
type UserContributionFollow struct {
	BaseModel
	UserID             int `json:"user_id"`
	UserContributionID int `json:"user_contribution_id"`
}

// Add 追加する
func (u *UserContributionFollow) Add() error {
	return Create(u)
}

// Delete 削除する
func (u *UserContributionFollow) Delete() error {
	return Delete(u)
}

// GetByID IDから取得する
func (u *UserContributionFollow) GetByID(id uint) (userContributionFollow UserContributionFollow, db *gorm.DB, err error) {
	whereList := []map[string]interface{}{
		{"ID": id},
	}
	option := make(map[string]interface{})

	db, err = GetWhere(&userContributionFollow, "ID = :ID", whereList, option)

	return
}

// GetListByUserContributionID 投稿IDから取得する
func (u *UserContributionFollow) GetListByUserContributionID(ucID int) (userContributionFollow []UserContributionFollow, db *gorm.DB, err error) {
	whereList := []map[string]interface{}{
		{"UserContributionID": ucID},
	}
	option := make(map[string]interface{})

	db, err = GetListWhere(&userContributionFollow, "User_contribution_ID = :UserContributionID", whereList, option)

	return
}

// GetCountByUserContributionID 投稿IDから数を取得する
func (u *UserContributionFollow) GetCountByUserContributionID(ucID int) (int, error) {
	userContributionFollow := []UserContributionFollow{}

	whereList := []map[string]interface{}{
		{"UserContributionID": ucID},
	}
	option := make(map[string]interface{})

	return GetCount(&userContributionFollow, "user_contribution_follows", "User_contribution_ID = :UserContributionID AND Deleted_at IS NULL", whereList, option)
}

// GetListByUserContributionIDList 投稿IDリストから取得する
func (u *UserContributionFollow) GetListByUserContributionIDList(ucID []int) (userContributionFollow []UserContributionFollow, db *gorm.DB, err error) {
	whereList := []map[string]interface{}{
		{"UserContributionID": ucID},
	}
	optionMap := map[string]interface{}{
		"select": "user_contribution_id",
	}

	db, err = GetListWhere(&userContributionFollow, "User_contribution_ID IN :UserContributionID", whereList, optionMap)

	return
}

// GetByUserIDAndUserContributionID ユーザIDと投稿IDから取得する
func (u *UserContributionFollow) GetByUserIDAndUserContributionID(uID int, ucID int) (userContributionFollow UserContributionFollow, db *gorm.DB, err error) {
	whereList := []map[string]interface{}{
		{"UserID": uID},
		{"UserContributionID": ucID},
	}
	option := make(map[string]interface{})

	db, err = GetWhere(&userContributionFollow, "User_ID = :UserID AND User_contribution_ID = :UserContributionID", whereList, option)

	return
}

// GetCountByUserIDAndUserContributionID ユーザIDと投稿IDから数を取得する
func (u *UserContributionFollow) GetCountByUserIDAndUserContributionID(uID int, ucID int) (int, error) {
	userContributionFollow := []UserContributionFollow{}

	whereList := []map[string]interface{}{
		{"UserID": uID},
		{"UserContributionID": ucID},
	}
	option := make(map[string]interface{})

	return GetCount(&userContributionFollow, "user_contribution_follows", "User_ID = :UserID AND User_contribution_ID = :UserContributionID AND Deleted_at IS NULL", whereList, option)
}

// GetCountByUserID ユーザIDから数を取得する
func (u *UserContributionFollow) GetCountByUserID(uID int, order string) (int, error) {
	userContributionFollow := []UserContributionFollow{}

	whereList := []map[string]interface{}{
		{"UserID": uID},
	}
	option := make(map[string]interface{})

	return GetCount(&userContributionFollow, "user_contribution_follows", "User_ID = :UserID AND Deleted_at IS NULL", whereList, option)
}

// GetListByUserID ユーザIDからリストを取得する
func (u *UserContributionFollow) GetListByUserID(uID int, order string, limit int, offset int) (userContributionFollow []UserContributionFollow, db *gorm.DB, err error) {
	whereList := []map[string]interface{}{
		{"UserID": uID},
	}
	option := map[string]interface{}{
		"order":  order,
		"limit":  limit,
		"offset": offset,
	}

	db, err = GetListWhere(&userContributionFollow, "User_ID = :UserID", whereList, option)

	return
}
