package main

import (
	"path/filepath"
	"runtime"

	"dotstamp_server/utils/contribution"
	"dotstamp_server/utils/follow"

	"github.com/astaxie/beego"
)

// getAppPath アプリケーションパスを取得する
func getAppPath() string {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, "../../.."+string(filepath.Separator))))

	return apppath
}

func init() {
	apppath := getAppPath()

	err := beego.LoadAppConfig("ini", apppath+"/dotstamp_server/conf/app.conf")
	if err != nil {
		panic(err)
	}
}

func main() {
	if err := Exec(); err != nil {
		panic(err)
	}
}

// Exec 実行する
func Exec() error {
	contributionIDList, err := contributions.GetViewStatusPublicIDList()
	if err != nil {
		return err
	}

	followList, err := follows.GetListByUserContributionIDList(contributionIDList)
	if err != nil {
		return err
	}

	followMap := follows.GetFollowCountMap(followList)

	follows.TruncateTotal()

	if err := follows.AddTotalMap(followMap); err != nil {
		return err
	}

	return nil
}
