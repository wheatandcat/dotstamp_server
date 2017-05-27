package controllersFacebook

import (
	"dotstamp_server/utils"
	"example-golang-oauth2/lib/facebook"

	"github.com/astaxie/beego"
)

// Oauth2Controller Oauth2コントローラー
type Oauth2Controller struct {
	beego.Controller
}

// Get 認証する
func (c *Oauth2Controller) Get() {
	config := facebook.GetConnect()

	state := utils.GetRandString(10)
	c.SetSession("facebookOauthState", state)

	url := config.AuthCodeURL(state)

	c.Redirect(url, 302)
}
