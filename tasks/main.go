package tasks

import (
	"os"

	"dotstamp_server/utils"
	"dotstamp_server/utils/log"

	"github.com/astaxie/beego"
)

// SetConfig コンフィグを設定する
func SetConfig() (err error) {
	apppath := utils.GetAppPath()

	if os.Getenv("ENV_CONF") == "prod" {
		err = beego.LoadAppConfig("ini", apppath+"/conf/app_prod.conf")
	} else {
		err = beego.LoadAppConfig("ini", apppath+"/conf/app_dev.conf")
	}

	return err
}

// Err エラーにする
func Err(err error, file string) {
	logs.Err("["+file+"]"+err.Error(), 0)

	panic(err)
}
