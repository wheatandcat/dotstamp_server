package controllersUser

import (
	"github.com/wheatandcat/dotstamp_server/controllers"
	"github.com/wheatandcat/dotstamp_server/utils/user"
)

// ShowController 詳細確認
type ShowController struct {
	controllers.BaseController
}

// ShowResponse 詳細レスポンス
type ShowResponse struct {
	User    user.User
	Profile []user.Profile
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
		c.ServerError(err, controllers.ErrCodeUserNotFound, userID)
	}

	p, err := user.GetProfileImageListByUserID(userID)
	if err != nil {
		c.ServerError(err, controllers.ErrCodeUserNotFound, userID)
	}

	c.Data["json"] = ShowResponse{
		User:    u,
		Profile: p,
	}
	c.ServeJSON()
}
