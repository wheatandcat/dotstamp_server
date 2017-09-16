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
	Code  string `form:"code"`
	State string `form:"state"`
}

// Get コールバックする
func (c *DevRedirectController) Get() {
	request := DevRedirectRequest{}
	if err := c.ParseForm(&request); err != nil {
		c.RedirectError(err, 0)
		return
	}

	url := "exp://ts-xwe.wheatandcat.dotstamp-native.exp.direct:80/?code=" + request.Code
	c.Redirect(url, 302)
}
