package controllersForgetPassword

import (
	"dotstamp_server/controllers"
	"dotstamp_server/models"
	"dotstamp_server/utils"
	"dotstamp_server/utils/mail"
	"dotstamp_server/utils/user"

	validator "gopkg.in/go-playground/validator.v9"

	"github.com/astaxie/beego"
)

// AddController パスワード忘れ追加
type AddController struct {
	controllers.BaseController
}

// AddRequest パスワードを忘れ追加リクエスト
type AddRequest struct {
	Email string `form:"email" validate:"required,email"`
}

// AddResponse パスワードを忘れ追加レスポンス
type AddResponse struct {
	Warning bool
	Message string
}

// Post ユーザー投稿一覧を取得する
func (c *AddController) Post() {
	request := AddRequest{}
	if err := c.ParseForm(&request); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon)
		return
	}

	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon)
		return
	}

	u, err := user.GetByEmail(request.Email)
	if err != nil {
		c.ServerError(err, controllers.ErrCodeCommon)
		return
	}

	if u.ID == uint(0) {
		c.Data["json"] = AddResponse{
			Warning: true,
			Message: "メールアドレスが見つかりませんでした",
		}
		c.ServeJSON()
		return
	}

	if err = user.DeleteByEmail(request.Email); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon)
		return
	}

	keyword := utils.GetRandString(50)

	tx := models.Begin()

	if err = user.AddForgetPassword(request.Email, keyword); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon)
		return
	}

	var url string
	url, err = mail.GetForgetpasswordURL(request.Email, keyword)
	if err != nil {
		models.Rollback(tx)
		c.ServerError(err, controllers.ErrCodeCommon)
		return
	}

	models.Commit(tx)

	top := beego.AppConfig.String("topurl")

	f := mail.ForgetpasswordTemplate{
		URL:   top + "#/password/reset/" + url,
		Host:  top,
		Email: beego.AppConfig.String("email"),
	}
	m := mail.GetForgetpasswordBody(f)
	b := mail.Body{
		From:    beego.AppConfig.String("email"),
		To:      request.Email,
		Subject: "[dotstamp]パスワード再設定",
		Message: string(m),
	}

	err = mail.Send(request.Email, mail.GetBody(b))
	if err != nil {
		c.ServerError(err, controllers.ErrCodeCommon)
		return
	}

	c.Data["json"] = AddResponse{
		Warning: false,
		Message: "",
	}

	c.ServeJSON()
}
