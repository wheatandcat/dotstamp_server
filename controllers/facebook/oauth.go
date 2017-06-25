package controllersFacebook

import (
	"github.com/wheatandcat/dotstamp_server/utils"
	"github.com/wheatandcat/dotstamp_server/utils/oauth/facebook"

	"github.com/astaxie/beego"
)

// OauthController Oauth2コントローラー
type OauthController struct {
	beego.Controller
}

// Get 認証する
func (c *OauthController) Get() {
	config := facebook.GetConnect()

	state := utils.GetRandString(10)
	c.SetSession("facebookOauthState", state)

	url := config.AuthCodeURL(state)

	c.Redirect(url, 302)
}
