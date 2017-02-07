package controllersUser

import (
	"dotstamp_server/controllers"
	"dotstamp_server/utils/user"
)

// ShowController 詳細確認
type ShowController struct {
	controllers.BaseController
}

// Post ユーザー情報
func (c *ShowController) Post() {
	userID := c.GetUserID()
	if !c.IsNoLogin(userID) {
		c.ServerLoginNotFound()
		return
	}

	u, err := user.GetByUserID(userID)
	if err != nil {
		c.ServerError(err, controllers.ErrCodeUserNotFound)
	}

	c.Data["json"] = u
	c.ServeJSON()
}
