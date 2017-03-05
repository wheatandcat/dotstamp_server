package main

import (
	_ "dotstamp_server/routers"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/redis"
	_ "github.com/garyburd/redigo/redis"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/ziutek/mymysql/godrv"
)

func main() {
	apppath := getAppPath()

	var err error
	if os.Getenv("ENV_CONF") == "PROD_BLUE" {
		err = beego.LoadAppConfig("ini", apppath+"/dotstamp_server/conf/app_prod_blue.conf")
	} else if os.Getenv("ENV_CONF") == "PROD_GREEN" {
		log.Println("PROD_GREEN")
		err = beego.LoadAppConfig("ini", apppath+"/dotstamp_server/conf/app_prod_green.conf")
	} else {
		err = beego.LoadAppConfig("ini", apppath+"/dotstamp_server/conf/app_dev.conf")
	}

	if err != nil {
		panic(err)
	}

	beego.Run()
}

// getAppPath アプリケーションパスを取得する
func getAppPath() string {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))

	return apppath
}
