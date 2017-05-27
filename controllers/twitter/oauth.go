package controllersTwitter

import (
	"dotstamp_server/utils/oauth/twitter"

	"github.com/astaxie/beego"
)

// OauthController Oauth2コントローラー
type OauthController struct {
	beego.Controller
}

// Get 認証する
func (c *OauthController) Get() {
	c.StartSession()

	config := twitter.GetConnect()
	rt, err := config.RequestTemporaryCredentials(nil, beego.AppConfig.String("callBackUrl")+"api/twitter/callback", nil)
	if err != nil {
		panic(err)
	}

	c.CruSession.Set("request_token", rt.Token)
	c.CruSession.Set("request_token_secret", rt.Secret)

	url := config.AuthorizationURL(rt, nil)

	c.Redirect(url, 302)
}
