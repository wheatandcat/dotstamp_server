package controllersNative

import (
	"github.com/wheatandcat/dotstamp_server/controllers"
)

// DevCallbackController コールバックコントローラ
type DevCallbackController struct {
	controllers.BaseController
}

// Get コールバックする
func (c *DevCallbackController) Get() {

	c.TplName = "dev_oauth.tpl"
}
