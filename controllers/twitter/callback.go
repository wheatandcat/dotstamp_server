package controllersTwitter

import (
	"dotstamp_server/controllers"
	"dotstamp_server/utils/oauth/twitter"
	"net/url"

	"github.com/astaxie/beego"
	"github.com/garyburd/go-oauth/oauth"
)

// CallbackController コールバックコントローラ
type CallbackController struct {
	controllers.BaseController
}

// CallbackRequest コールバックリクエスト
type CallbackRequest struct {
	Token    string `form:"oauth_token"`
	Verifier string `form:"oauth_verifier"`
}

// Get コールバックする
func (c *CallbackController) Get() {
	c.StartSession()

	request := CallbackRequest{}
	if err := c.ParseForm(&request); err != nil {
		c.RedirectError(err, 0)
		return
	}

	at, err := twitter.GetAccessToken(
		&oauth.Credentials{
			Token:  c.CruSession.Get("request_token").(string),
			Secret: c.CruSession.Get("request_token_secret").(string),
		},
		request.Verifier,
	)
	if err != nil {
		c.RedirectError(err, 0)
	}

	c.CruSession.Set("oauth_secret", at.Secret)
	c.CruSession.Set("oauth_token", at.Token)

	account := twitter.Account{}
	if err = twitter.GetMe(at, &account); err != nil {
		c.RedirectError(err, 0)
	}

	url := beego.AppConfig.String("topurl") + "oauth/?email=" + url.QueryEscape(account.Email)
	c.Redirect(url, 302)
}
