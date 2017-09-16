package controllersNative

import (
	"github.com/wheatandcat/dotstamp_server/controllers"
)

// DevRedirectController コールバックコントローラ
type DevRedirectController struct {
	controllers.BaseController
}

// DevRedirectRequest コールバックリクエスト
type DevRedirectRequest struct {
	AccessToken string `form:"access_token"`
}

// Get コールバックする
func (c *DevRedirectController) Get() {
	request := DevRedirectRequest{}
	if err := c.ParseForm(&request); err != nil {
		c.RedirectError(err, 0)
		return
	}

	url := "exp://ts-xwe.wheatandcat.dotstamp-native.exp.direct:80/?access_token=" + request.AccessToken
	c.Redirect(url, 302)
}
