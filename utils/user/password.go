package user

import (
	"dotstamp_server/models"
	"dotstamp_server/utils"
	"time"
)

// AddForgetPassword 忘れたパスワードを追加する
func AddForgetPassword(email string, keyword string) error {
	u := models.UserForgetPassword{
		Email:   email,
		Keyword: keyword,
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
		return false, nil
	}

	if up.CreatedAt.Add(1*time.Hour).Unix() < utils.Now().Unix() {
		return false, nil
	}

	return true, nil
}

// DeleteByEmail メールアドレスから削除する
func DeleteByEmail(email string) error {
	u := models.UserForgetPassword{}
	r, _, err := u.GetListByEmail(email)
	if err != nil {
		return err
	}

	if len(r) == 0 {
		return nil
	}

	return u.DeleteList(r)
}
