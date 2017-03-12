package controllers

import (
	"github.com/astaxie/beego"
)

// MainController メインコントローラ
type MainController struct {
	beego.Controller
}

// Get 取得する
func (c *MainController) Get() {
	c.Data["StaticUrl"] = beego.AppConfig.String("staticUrl")

	c.TplName = "index.tpl"
}
