package controllersMovie

import (
	"dotstamp_server/controllers"

	validator "gopkg.in/go-playground/validator.v9"
)

// ConnectController 接続コントローラ
type ConnectController struct {
	controllers.BaseController
}

// ConnectRequest 接続リクエスト
type ConnectRequest struct {
	UserContributionID int `form:"userContributionId" validate:"min=1"`
}

// ConnectResponse 接続レスポンス
type ConnectResponse struct {
	URL     string
	Warning bool
	Message string
}

// Post 接続する
func (c *ConnectController) Post() {
	userID := c.GetUserID()
	if !c.IsNoLogin(userID) {
		c.ServerLoginNotFound()
		return
	}

	request := ConnectRequest{}
	if err := c.ParseForm(&request); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	c.Data["json"] = ConnectResponse{
		Warning: false,
		Message: "",
	}

	c.ServeJSON()

}
