package controllersNative

import (
	"github.com/wheatandcat/dotstamp_server/controllers"
)

// CallbackController コールバックコントローラ
type CallbackController struct {
	controllers.BaseController
}

// Get コールバックする
func (c *CallbackController) Get() {

	c.TplName = "oauth.tpl"
}
