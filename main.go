package main

import (
	_ "dotstamp_server/routers"
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

	err := beego.LoadAppConfig("ini", apppath+"/dotstamp_server/conf/app.conf")
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
