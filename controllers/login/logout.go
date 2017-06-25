package controllersLogin

import "github.com/wheatandcat/dotstamp_server/controllers"

// LogoutController ログアウトコントローラ
type LogoutController struct {
	controllers.BaseController
}

// Post ログアウト
func (c *LogoutController) Post() {

	c.DelSession("user_id")

	c.Data["json"] = true

	c.ServeJSON()
}
