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
func (u *UserMaster) GetByEmail(email string) (userMaster UserMaster, db *gorm.DB, err error) {
	whereList := []map[string]interface{}{
		{"Email": email},
	}
	option := make(map[string]interface{})

	db, err = GetWhere(&userMaster, "Email LIKE :Email", whereList, option)

	return
}

// GetByID ユーザ IDから取得する
func (u *UserMaster) GetByID(id int) (userMaster UserMaster, db *gorm.DB, err error) {
	whereList := []map[string]interface{}{
		{"ID": id},
	}
	option := make(map[string]interface{})

	db, err = GetWhere(&userMaster, "ID = :ID", whereList, option)

	return
}

// GetScanByID ユーザIDからスキャン取得する
func (u *UserMaster) GetScanByID(id int, dest interface{}) error {
	whereList := []map[string]interface{}{
		{"ID": id},
	}
	option := make(map[string]interface{})

	return GeScanWhere(dest, "user_masters", "ID = :ID", whereList, option)
}

// GetListByIDList ユーザ IDリストからリストを取得する
func (u *UserMaster) GetListByIDList(idList []int) (userMaster []UserMaster, db *gorm.DB, err error) {
	whereList := []map[string]interface{}{
		{"ID": idList},
	}
	option := make(map[string]interface{})

	db, err = GetListWhere(&userMaster, "ID IN :ID", whereList, option)

	return
}

// GetScanByIDList ユーザIDリストからスキャン取得する
func (u *UserMaster) GetScanByIDList(idList []int, dest interface{}) error {
	whereList := []map[string]interface{}{
		{"ID": idList},
	}
	option := make(map[string]interface{})

	return GeScanWhere(dest, "user_masters", "ID IN :ID", whereList, option)
}
