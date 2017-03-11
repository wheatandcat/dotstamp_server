package main

import (
	_ "dotstamp_server/routers"
	"os"
	"path/filepath"

	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/redis"
	_ "github.com/garyburd/redigo/redis"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/ziutek/mymysql/godrv"
)

func main() {
	var err error
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}

	if os.Getenv("ENV_CONF") == "prod_blue" {
		err = beego.LoadAppConfig("ini", dir+"/conf/app_prod_blue.conf")
	} else if os.Getenv("ENV_CONF") == "prod_green" {
		err = beego.LoadAppConfig("ini", dir+"/conf/app_prod_green.conf")
	} else {
		err = beego.LoadAppConfig("ini", dir+"/conf/app_dev.conf")
	}

	beego.BConfig.WebConfig.StaticExtensionsToGzip = []string{".css", ".js"}

	if err != nil {
		panic(err)
	}

	beego.Run()
}
