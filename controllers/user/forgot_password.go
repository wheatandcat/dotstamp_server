package controllersUser

import (
	"dotstamp_server/controllers"
	"dotstamp_server/utils/user"
)

// ForgotPasswordController パスワードを忘れた
type ForgotPasswordController struct {
	controllers.BaseController
}

// ForgotPasswordRequest パスワードを忘れたリクエスト
type ForgotPasswordRequest struct {
	Email string `form:"email"`
}

// ForgotPasswordResponse パスワードを忘れたレスポンス
type ForgotPasswordResponse struct {
	Warning bool
	Message string
}

// Post ユーザー投稿一覧を取得する
func (c *ForgotPasswordController) Post() {
	request := ForgotPasswordRequest{}
	if err := c.ParseForm(&request); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon)
		return
	}

	u, err := user.GetByEmail(request.Email)
	if err != nil {
		c.ServerError(err, controllers.ErrCodeCommon)
		return
	}

	if u.ID == uint(0) {
		c.Data["json"] = ForgotPasswordResponse{
			Warning: true,
			Message: "メールアドレスが見つかりませんでした",
		}
	}

	c.Data["json"] = ForgotPasswordResponse{
		Warning: false,
		Message: "",
	}

	c.ServeJSON()
}
