package models

import "github.com/jinzhu/gorm"

// UserForgetPassword ユーザの忘れたパスワード
type UserForgetPassword struct {
	gorm.Model
	Email   string
	Keyword string
}

// Add 追加する
func (u *UserForgetPassword) Add() error {
	return Create(u)
}

// Delete 削除する
func (u *UserForgetPassword) Delete() error {
	return Delete(u)
}

// GetByEmail メールアドレスから取得する
func (u *UserForgetPassword) GetByEmail(email string) (userForgetPassword UserForgetPassword, db *gorm.DB, err error) {
	whereList := []map[string]interface{}{
		{"Email": email},
	}
	option := make(map[string]interface{})

	db, err = GetWhere(&userForgetPassword, "Email LIKE :Email", whereList, option)

	return
}
