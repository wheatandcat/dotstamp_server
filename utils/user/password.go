package user

import (
	"dotstamp_server/models"
	"dotstamp_server/utils"
)

// AddForgetPassword 忘れたパスワードを追加する
func AddForgetPassword(email string) error {
	rand := utils.GetRandString(50)

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
