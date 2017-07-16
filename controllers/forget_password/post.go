package controllersForgetPassword

import (
	"github.com/wheatandcat/dotstamp_server/controllers"
	"github.com/wheatandcat/dotstamp_server/models"
	"github.com/wheatandcat/dotstamp_server/utils"
	"github.com/wheatandcat/dotstamp_server/utils/mail"
	"github.com/wheatandcat/dotstamp_server/utils/user"

	validator "gopkg.in/go-playground/validator.v9"

	"github.com/astaxie/beego"
)

// PostRequest パスワードを忘れ追加リクエスト
type PostRequest struct {
	Email string `form:"email" validate:"required,email"`
}

// PostResponse パスワードを忘れ追加レスポンス
type PostResponse struct {
	Warning bool   `json:"warning"`
	Message string `json:"message"`
}

// Post ユーザー投稿一覧を取得する
func (c *MainController) Post() {
	request := PostRequest{}
	if err := c.ParseForm(&request); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, 0)
		return
	}

	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, 0)
		return
	}

	u, err := user.GetByEmail(request.Email)
	if err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, 0)
		return
	}

	if u.ID == uint(0) {
		c.Data["json"] = PostResponse{
			Warning: true,
			Message: "メールアドレスが見つかりませんでした",
		}
		c.ServeJSON()
		return
	}

	if err = user.DeleteByEmail(request.Email); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, 0)
		return
	}

	keyword := utils.GetRandString(50)

	tx := models.Begin()

	if err = user.AddForgetPassword(request.Email, keyword); err != nil {
		models.Rollback(tx)
		c.ServerError(err, controllers.ErrCodeCommon, 0)
		return
	}

	var url string
	url, err = mail.GetForgetpasswordURL(request.Email, keyword)
	if err != nil {
		models.Rollback(tx)
		c.ServerError(err, controllers.ErrCodeCommon, 0)
		return
	}

	models.Commit(tx)

	top := beego.AppConfig.String("topurl")

	f := mail.ForgetpasswordTemplate{
		URL:   top + "password/reset/" + url,
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
		c.ServerError(err, controllers.ErrCodeCommon, 0)
		return
	}

	c.Data["json"] = PostResponse{
		Warning: false,
		Message: "",
	}

	c.ServeJSON()
}
