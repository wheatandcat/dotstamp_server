package models

import "github.com/jinzhu/gorm"

// UserMaster ユーザー情報
type UserMaster struct {
	gorm.Model     `model:"true"`
	Name           string
	Email          string `validate:"required,email"`
	Password       string
	ProfileImageID int `json:"profile_image_id"`
}

// GetIDAndAdd 投稿してIDを取得する
func (u *UserMaster) GetIDAndAdd() (uint, error) {
	if err := Create(u); err != nil {
		return 0, err
	}

	return u.ID, nil
}

// Save 保存する
func (u *UserMaster) Save() error {
	return Save(u)
}

// GetByEmail メールアドレスから取得する
func (u *UserMaster) GetByEmail(email string) (userMaster UserMaster) {
	whereList := []map[string]interface{}{
		{"Email": email},
	}
	option := make(map[string]interface{})

	GetWhere(&userMaster, "Email LIKE :Email", whereList, option)

	return
}

// GetByID ユーザ IDから取得する
func (u *UserMaster) GetByID(id int) (userMaster UserMaster) {
	whereList := []map[string]interface{}{
		{"ID": id},
	}
	option := make(map[string]interface{})

	GetWhere(&userMaster, "ID = :ID", whereList, option)

	return
}

// GetListByIDList ユーザ IDリストからリストを取得する
func (u *UserMaster) GetListByIDList(idList []int) (userMaster []UserMaster) {
	whereList := []map[string]interface{}{
		{"ID": idList},
	}
	option := make(map[string]interface{})

	GetListWhere(&userMaster, "ID IN :ID", whereList, option)

	return
}
