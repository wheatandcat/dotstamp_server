package controllersLogin

import (
	"dotstamp_server/controllers"
	"dotstamp_server/utils/user"
)

// Checkontroller 登録確認コントローラ
type Checkontroller struct {
	controllers.BaseController
}

// Post ログイン
func (t *Checkontroller) Post() {
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
