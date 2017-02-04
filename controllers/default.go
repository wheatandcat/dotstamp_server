package controllers

import (
	"log"

	"github.com/astaxie/beego"
)

// MainController メインコントローラ
type MainController struct {
	beego.Controller
}

// Get 取得する
func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"

	c.TplName = "index.tpl"
	path := beego.AppConfig.String("viewspath")
	log.Println(path)
	log.Println(c)
}
