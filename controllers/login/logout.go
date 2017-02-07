package controllersLogin

import "dotstamp_server/controllers"

// LogoutController ログアウトコントローラ
type LogoutController struct {
	controllers.BaseController
}

// Post ログアウト
func (t *LogoutController) Post() {

	t.DelSession("user_id")

	t.Data["json"] = true

	t.ServeJSON()
}
