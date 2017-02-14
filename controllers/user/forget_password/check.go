package controllersForgetPassword

import (
	"dotstamp_server/controllers"
	"dotstamp_server/utils/user"
)

// CheckController パスワード忘れ確認コントローラー
type CheckController struct {
	controllers.BaseController
}

// CheckRequest パスワードを忘れ確認リクエスト
type CheckRequest struct {
	Email   string `form:"email"`
	Keyword string `form:"key"`
}

// CheckResponse パスワードを忘れ確認レスポンス
type CheckResponse struct {
	Warning bool
	Message string
}

// Post パスワード忘れ確認
func (c *CheckController) Post() {
	request := CheckRequest{}
	if err := c.ParseForm(&request); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon)
		return
	}

	r, err := user.IsUpdatePassword(request.Email, request.Keyword)
	if err != nil {
		c.ServerError(err, controllers.ErrParameter)
		return
	}

	if r == false {
		c.Data["json"] = CheckResponse{
			Warning: true,
			Message: "不正なURLです",
		}
		c.ServeJSON()
		return
	}

	c.Data["json"] = CheckResponse{
		Warning: false,
		Message: "",
	}

	c.ServeJSON()
}
