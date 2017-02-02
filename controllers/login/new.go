package controllersLogin

import (
	"dotstamp_server/controllers"
	"dotstamp_server/utils/user"
)

// NewController 新規登録コントローラ
type NewController struct {
	controllers.BaseController
}

// Post 新規ログイン
func (t *NewController) Post() {
	email := t.GetString("email")
	pass := t.GetString("password")

	userID, err := user.Add(email, email, pass)
	if err != nil {
		t.ServerError(err, controllers.ErrCreateUser)
		return
	}

	t.SetSession("user_id", userID)
	t.Data["json"] = true

	t.ServeJSON()
}
