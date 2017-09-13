package controllersNative

import (
	"github.com/wheatandcat/dotstamp_server/controllers"
)

// CallbackController コールバックコントローラ
type CallbackController struct {
	controllers.BaseController
}

// CallbackRequest コールバックリクエスト
type CallbackRequest struct {
	Code  string `form:"code"`
	State string `form:"state"`
}

// Get コールバックする
func (c *CallbackController) Get() {
	request := CallbackRequest{}
	if err := c.ParseForm(&request); err != nil {
		c.RedirectError(err, 0)
		return
	}

	url := "exp://exp.host/@wheatandcat/dotstamp_native/?code=" + request.Code
	c.Redirect(url, 302)
}
