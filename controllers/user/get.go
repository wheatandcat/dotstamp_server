package controllersUser

import (
	"github.com/wheatandcat/dotstamp_server/controllers"
	"github.com/wheatandcat/dotstamp_server/utils/user"
)

// GetResponse 詳細レスポンス
type GetResponse struct {
	User    user.User      `json:"user"`
	Profile []user.Profile `json:"profiles"`
}

// Get ユーザー情報
func (c *MainController) Get() {
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

	c.Data["json"] = GetResponse{
		User:    u,
		Profile: p,
	}
	c.ServeJSON()
}
