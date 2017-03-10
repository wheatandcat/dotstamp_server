package main

import (
	_ "dotstamp_server/routers"
	"os"

	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/redis"
	_ "github.com/garyburd/redigo/redis"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/ziutek/mymysql/godrv"
)

func main() {
	var err error
	if os.Getenv("ENV_CONF") == "prod_blue" {
		err = beego.LoadAppConfig("ini", "./conf/app_prod_blue.conf")
		beego.SetStaticPath("/static", "blue/static")

	} else if os.Getenv("ENV_CONF") == "prod_green" {
		err = beego.LoadAppConfig("ini", "./conf/app_prod_green.conf")
		beego.SetStaticPath("/static", "green/static")

	} else {
		err = beego.LoadAppConfig("ini", "./conf/app_dev.conf")
	}

	if err != nil {
		panic(err)
	}

	beego.Run()
}
