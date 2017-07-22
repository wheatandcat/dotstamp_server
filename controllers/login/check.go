package controllersLogin

import (
	"encoding/json"

	"github.com/wheatandcat/dotstamp_server/controllers"
	"github.com/wheatandcat/dotstamp_server/utils/user"
)

// CheckController 登録確認コントローラ
type CheckController struct {
	controllers.BaseController
}

// CheckRequest 確認リクエスト
type CheckRequest struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

// Post ログイン
func (c *CheckController) Post() {
	request := CheckRequest{}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &request); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, 0)
		return
	}

	u, err := user.GetByEmailAndPassword(request.Email, request.Password)
	if err != nil {
		c.ServerError(err, controllers.ErrUserOrPasswordDifferent, 0)
		return
	}

	c.SetSession("user_id", u.ID)

	c.Data["json"] = true

	c.ServeJSON()
}
