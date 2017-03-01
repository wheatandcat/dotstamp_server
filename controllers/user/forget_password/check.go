package controllersForgetPassword

import (
	"dotstamp_server/controllers"
	"dotstamp_server/utils"
	"dotstamp_server/utils/user"
)

// CheckController パスワード忘れ確認コントローラー
type CheckController struct {
	controllers.BaseController
}

// CheckResponse パスワードを忘れ確認レスポンス
type CheckResponse struct {
	Warning bool
	Message string
}

// Post パスワード忘れ確認
func (c *CheckController) Post() {
	email, err := utils.Urldecode(c.Ctx.Input.Param(":email"))
	if err != nil {
		c.ServerError(err, controllers.ErrParameter)
		return
	}
	e, err := utils.Decrypter([]byte(email))
	if err != nil {
		c.ServerError(err, controllers.ErrParameter)
		return
	}

	keyword, err := utils.Urldecode(c.Ctx.Input.Param(":keyword"))
	if err != nil {
		c.ServerError(err, controllers.ErrParameter)
		return
	}
	k, err := utils.Decrypter([]byte(keyword))
	if err != nil {
		c.ServerError(err, controllers.ErrParameter)
		return
	}

	r, err := user.IsUpdatePassword(e, k)
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
