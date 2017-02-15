package controllersForgetPassword

import (
	"dotstamp_server/controllers"
	"dotstamp_server/utils"
	"dotstamp_server/utils/user"
)

// SaveController パスワード保存コントローラー
type SaveController struct {
	controllers.BaseController
}

// SaveRequest パスワード保存リクエスト
type SaveRequest struct {
	Email    string `form:"email"`
	Keyword  string `form:"keyword"`
	Password string `form:"password"`
}

// SaveResponse パスワード保存レスポンス
type SaveResponse struct {
	Warning bool
	Message string
}

// Post パスワード保存
func (c *SaveController) Post() {
	request := SaveRequest{}
	if err := c.ParseForm(&request); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon)
		return
	}

	e, err := utils.Decrypter([]byte(request.Email))
	if err != nil {
		c.ServerError(err, controllers.ErrParameter)
		return
	}

	k, err := utils.Decrypter([]byte(request.Keyword))
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
		c.ServerError(err, controllers.ErrParameter)
		return
	}

	if err := user.UpadateToPassword(e, request.Password); err != nil {
		c.ServerError(err, controllers.ErrParameter)
		return
	}

	c.Data["json"] = CheckResponse{
		Warning: false,
		Message: "",
	}

	c.ServeJSON()
}
