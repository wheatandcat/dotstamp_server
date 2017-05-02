package controllers

import (
	"github.com/astaxie/beego"
)

// ErrorController エラーコントローラー
type ErrorController struct {
	beego.Controller
}

// Error404 エラー:404
func (c *ErrorController) Error404() {
	c.Redirect(beego.AppConfig.String("topUrl"), 200)
}
