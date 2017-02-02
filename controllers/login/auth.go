package controllersLogin

import (
	"dotstamp_server/controllers"
	"dotstamp_server/utils/user"
)

// AuthController 認証コントローラ
type AuthController struct {
	controllers.BaseController
}

// AuthResponse 削除レスポンス
type AuthResponse struct {
	Login bool
	Name  string
}

// Post ログイン中か判定する
func (t *AuthController) Post() {
	userID := t.GetUserID()

	var response AuthResponse
	if !t.IsNoLogin(userID) {
		response = AuthResponse{
			Login: false,
			Name:  "",
		}
	} else {

		u, err := user.GetByUserID(userID)
		if err != nil {
			t.ServerError(err, controllers.ErrCodeUserNotFound)
		}

		response = AuthResponse{
			Login: true,
			Name:  u.Name,
		}
	}

	t.Data["json"] = response

	t.ServeJSON()
}
