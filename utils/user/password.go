package user

import (
	"dotstamp_server/models"
	"errors"
)

// AddForgetPassword 忘れたパスワードを追加する
func AddForgetPassword(email string) error {
	// TODO 乱数にする
	rand := "aaaa"
	u := models.UserForgetPassword{
		Email:   email,
		Keyword: rand,
	}

	return u.Add()
}

// GetForgetPasswordByEmail メールアドレスから忘れたパスワードを取得する
func GetForgetPasswordByEmail(email string) (models.UserForgetPassword, error) {
	u := models.UserForgetPassword{}
	r, _, err := u.GetByEmail(email)

	return r, err
}

// IsUpdatePassword パスワードが更新可能か判定する
func IsUpdatePassword(email string, keyword string) (bool, error) {
	up, err := GetForgetPasswordByEmail(email)
	if err != nil {
		return false, err
	}

	if up.Keyword != keyword {
		return false, errors.New("difference keyword")
	}

	return true, nil
}
