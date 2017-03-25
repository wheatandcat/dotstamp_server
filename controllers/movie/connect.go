package controllersMovie

import (
	"dotstamp_server/controllers"
	"dotstamp_server/utils/contribution"
	"dotstamp_server/utils/movie"
	"strconv"

	"github.com/astaxie/beego"

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

// Get 接続する
func (c *ConnectController) Get() {
	userID := c.GetUserID()
	if !c.IsNoLogin(userID) {
		c.Redirect(beego.AppConfig.String("errorUrl"), 302)
		return
	}

	request := ConnectRequest{}
	if err := c.ParseForm(&request); err != nil {
		c.Redirect(beego.AppConfig.String("errorUrl"), 302)
		return
	}

	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		c.Redirect(beego.AppConfig.String("errorUrl"), 302)
		return
	}

	u, err := contributions.GetByUserContributionID(request.UserContributionID)
	if err != nil {
		c.Redirect(beego.AppConfig.String("errorUrl"), 302)
		return
	}

	if userID != u.UserID {
		c.Redirect(beego.AppConfig.String("errorUrl"), 302)
		return
	}

	if !contributions.ExistsMovie(request.UserContributionID) {
		c.Redirect(beego.AppConfig.String("errorUrl"), 302)
		return
	}

	config := movie.GetConnect()
	url := config.AuthCodeURL(strconv.Itoa(request.UserContributionID))

	c.Redirect(url, 302)
}
