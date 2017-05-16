package controllersGoogle

import (
	"dotstamp_server/utils"
	"dotstamp_server/utils/oauth/google"

	"github.com/astaxie/beego"
)

// OauthController Oauth2コントローラー
type OauthController struct {
	beego.Controller
}

// Get 認証する
func (c *OauthController) Get() {
	config := google.GetConnect()

	state := utils.GetRandString(10)
	c.SetSession("googleOauthState", state)

	url := config.AuthCodeURL(state)

	c.Redirect(url, 302)
}
