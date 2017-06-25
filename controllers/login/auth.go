package controllersLogin

import (
	"github.com/wheatandcat/dotstamp_server/controllers"
	"github.com/wheatandcat/dotstamp_server/utils/user"
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
func (c *AuthController) Post() {
	userID := c.GetUserID()

	var response AuthResponse
	if !c.IsNoLogin(userID) {
		response = AuthResponse{
			Login: false,
			Name:  "",
		}
	} else {

		u, err := user.GetByUserID(userID)
		if err != nil {
			c.ServerError(err, controllers.ErrCodeUserNotFound, userID)
		}

		response = AuthResponse{
			Login: true,
			Name:  u.Name,
		}
	}

	c.Data["json"] = response

	c.ServeJSON()
}
