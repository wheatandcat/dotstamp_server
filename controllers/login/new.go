package controllersLogin

import (
	"dotstamp_server/controllers"
	"dotstamp_server/utils/user"
)

// NewController 新規登録コントローラ
type NewController struct {
	controllers.BaseController
}

// NewRequest 新規リクエスト
type NewRequest struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

// Post 新規ログイン
func (c *NewController) Post() {
	request := NewRequest{}

	if err := c.ParseForm(&request); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon)
		return
	}

	userID, err := user.Add(request.Email, request.Email, request.Password)
	if err != nil {
		c.ServerError(err, controllers.ErrCreateUser)
		return
	}

	c.SetSession("user_id", userID)

	c.Data["json"] = userID

	c.ServeJSON()
}
