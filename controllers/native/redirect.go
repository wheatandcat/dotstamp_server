package controllersNative

import (
	"github.com/wheatandcat/dotstamp_server/controllers"
)

// RedirectController コールバックコントローラ
type RedirectController struct {
	controllers.BaseController
}

// RedirectRequest コールバックリクエスト
type RedirectRequest struct {
	AccessToken string `form:"access_token"`
}

// Get コールバックする
func (c *RedirectController) Get() {
	request := RedirectRequest{}
	if err := c.ParseForm(&request); err != nil {
		c.RedirectError(err, 0)
		return
	}

	url := "exp://exp.host/@wheatandcat/dotstamp_native/?access_token=" + request.AccessToken
	c.Redirect(url, 302)
}
