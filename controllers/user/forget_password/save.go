package controllersForgetPassword

import (
	"dotstamp_server/controllers"
	"dotstamp_server/models"
	"dotstamp_server/utils"
	"dotstamp_server/utils/user"

	validator "gopkg.in/go-playground/validator.v9"
)

// SaveController パスワード保存コントローラー
type SaveController struct {
	controllers.BaseController
}

// SaveRequest パスワード保存リクエスト
type SaveRequest struct {
	Email    string `form:"email"`
	Keyword  string `form:"keyword"`
	Password string `form:"password" validate:"min=8,max=100"`
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
		c.ServerError(err, controllers.ErrCodeCommon, 0)
		return
	}

	if err := c.ParseForm(&request); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, 0)
		return
	}

	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, 0)
		return
	}

	email, err := utils.Urldecode(request.Email)
	if err != nil {
		c.ServerError(err, controllers.ErrParameter, 0)
		return
	}
	e, err := utils.Decrypter([]byte(email))
	if err != nil {
		c.ServerError(err, controllers.ErrParameter, 0)
		return
	}

	keyword, err := utils.Urldecode(request.Keyword)
	if err != nil {
		c.ServerError(err, controllers.ErrParameter, 0)
		return
	}
	k, err := utils.Decrypter([]byte(keyword))
	if err != nil {
		c.ServerError(err, controllers.ErrParameter, 0)
		return
	}

	r, err := user.IsUpdatePassword(e, k)
	if err != nil {
		c.ServerError(err, controllers.ErrParameter, 0)
		return
	}

	if r == false {
		c.ServerError(err, controllers.ErrParameter, 0)
		return
	}

	tx := models.Begin()

	if err := user.UpadateToPassword(e, request.Password); err != nil {
		models.Rollback(tx)
		c.ServerError(err, controllers.ErrParameter, 0)
		return
	}

	models.Commit(tx)

	c.Data["json"] = CheckResponse{
		Warning: false,
		Message: "",
	}

	c.ServeJSON()
}
