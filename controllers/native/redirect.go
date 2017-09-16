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
	Code  string `form:"code"`
	State string `form:"state"`
}

// Get コールバックする
func (c *RedirectController) Get() {
	request := RedirectRequest{}
	if err := c.ParseForm(&request); err != nil {
		c.RedirectError(err, 0)
		return
	}

	url := "exp://exp.host/@wheatandcat/dotstamp_native/?code=" + request.Code
	c.Redirect(url, 302)
}
