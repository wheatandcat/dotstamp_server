package controllersMovie

import (
	"dotstamp_server/controllers"
	"dotstamp_server/utils/movie"
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

// Get 接続する
func (c *ConnectController) Get() {
	config := movie.GetConnect()

	url := config.AuthCodeURL("st001")

	c.Redirect(url, 302)
}
