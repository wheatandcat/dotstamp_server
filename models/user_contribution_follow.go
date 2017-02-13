package models

import "github.com/jinzhu/gorm"

// UserContributionFollow ユーザー投稿フォロー
type UserContributionFollow struct {
	gorm.Model
	UserID             int `json:"user_id"`
	UserContributionID int `json:"user_contribution_id"`
}

// Add 追加する
func (u *UserContributionFollow) Add(uID int, ucID int) error {
	u.UserID = uID
	u.UserContributionID = ucID

	return Save(u)
}

// Delete 削除する
func (u *UserContributionFollow) Delete() error {
	return Delete(u)
}

// GetListByUserContributionID 投稿IDからフォローを取得する
func (u *UserContributionFollow) GetListByUserContributionID(ucID int) (userContributionFollow []UserContributionFollow, db *gorm.DB, err error) {
	whereList := []map[string]interface{}{
		{"UserContributionID": ucID},
	}
	option := make(map[string]interface{})

	db, err = GetWhere(&userContributionFollow, "User_contribution_ID = :UserContributionID", whereList, option)

	return
}

// GetListByUserContributionIDList 投稿IDリストからフォローを取得する
func (u *UserContributionFollow) GetListByUserContributionIDList(ucID []int) (userContributionFollow []UserContributionFollow, db *gorm.DB, err error) {
	whereList := []map[string]interface{}{
		{"UserContributionID": ucID},
	}
	option := make(map[string]interface{})

	db, err = GetWhere(&userContributionFollow, "User_contribution_ID IN :UserContributionID", whereList, option)

	return
}
