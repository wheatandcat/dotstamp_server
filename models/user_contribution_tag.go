package models

import "github.com/jinzhu/gorm"

// UserContributionTag ユーザー投稿タグ
type UserContributionTag struct {
	gorm.Model
	UserContributionID int `json:"user_contribution_id"`
	Name               string
}

// AddList リストを保存する
func (uc *UserContributionTag) AddList(u []UserContributionTag) (err error) {
	for _, user := range u {
		if err = Create(&user); err != nil {
			return err
		}
	}

	return nil
}

// GetListByUserContributionID 投稿IDから取得する
func (uc *UserContributionTag) GetListByUserContributionID(id int) (userContributionTag []UserContributionTag) {
	whereList := []map[string]interface{}{
		{"UserContributionID": id},
	}
	option := make(map[string]interface{})

	GetListWhere(&userContributionTag, "User_contribution_ID = :UserContributionID", whereList, option)

	return
}

// GetListByUserContributionIDList 投稿IDリストから取得する
func (uc *UserContributionTag) GetListByUserContributionIDList(idList []int) (userContributionTag []UserContributionTag) {
	whereList := []map[string]interface{}{
		{"User_contribution_ID": idList},
	}
	option := make(map[string]interface{})

	GetListWhere(&userContributionTag, "User_contribution_ID IN :User_contribution_ID", whereList, option)

	return
}

// GetByID IDから取得する
func (uc *UserContributionTag) GetByID(id int) (userContributionTag UserContributionTag) {
	whereList := []map[string]interface{}{
		{"ID": id},
	}
	option := make(map[string]interface{})

	GetWhere(&userContributionTag, "ID = :ID", whereList, option)

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
