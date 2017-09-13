package controllersNative

import (
	"github.com/wheatandcat/dotstamp_server/controllers"
)

// DevCallbackController コールバックコントローラ
type DevCallbackController struct {
	controllers.BaseController
}

// DevCallbackRequest コールバックリクエスト
type DevCallbackRequest struct {
	Code  string `form:"code"`
	State string `form:"state"`
}

// Get コールバックする
func (c *DevCallbackController) Get() {
	request := DevCallbackRequest{}
	if err := c.ParseForm(&request); err != nil {
		c.RedirectError(err, 0)
		return
	}

	url := "exp://ts-xwe.wheatandcat.dotstamp-native.exp.direct:80/?code=" + request.Code
	c.Redirect(url, 302)
}
