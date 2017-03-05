package main

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"dotstamp_server/models"
	"dotstamp_server/utils"
	"dotstamp_server/utils/contribution"
	"dotstamp_server/utils/follow"

	"github.com/astaxie/beego"
)

var followMap map[int]int
var contributionIDList []int
var err error
var logfile *os.File

// getAppPath アプリケーションパスを取得する
func getAppPath() string {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, "../../.."+string(filepath.Separator))))

	return apppath
}

func init() {
	apppath := getAppPath()

	if os.Getenv("ENV_CONF") == "prod" {
		err = beego.LoadAppConfig("ini", apppath+"/dotstamp_server/conf/app_prod.conf")
	} else {
		err = beego.LoadAppConfig("ini", apppath+"/dotstamp_server/conf/app_dev.conf")
	}
	if err != nil {
		panic(err)
	}

}

func main() {
	logfile, err = utils.LogFile("contributionTotalFollows")
	if err != nil {
		panic(err)
	}
	defer logfile.Close()
	log.SetOutput(io.MultiWriter(logfile))
	log.SetFlags(log.Ldate | log.Ltime)

	log.Println("start!")

	tx := models.Begin()

	if err = AddContributionTotalFollows(); err != nil {
		models.Rollback(tx)
		panic(err)
	}

	if err = SaveUserContributionSearchToFollowCount(); err != nil {
		models.Rollback(tx)
		panic(err)
	}

	models.Commit(tx)

	log.Println("finish!")
}

// AddContributionTotalFollows フォロー数を追加する
func AddContributionTotalFollows() error {
	contributionIDList, err = contributions.GetViewStatusPublicIDList()
	if err != nil {
		return err
	}

	followList, err := follows.GetListByUserContributionIDList(contributionIDList)
	if err != nil {
		return err
	}

	followMap = follows.GetFollowCountMap(followList)

	follows.TruncateTotal()

	if err = follows.AddTotalMap(followMap); err != nil {
		return err
	}

	return nil
}

// SaveUserContributionSearchToFollowCount 検索のフォロー数を更新する
func SaveUserContributionSearchToFollowCount() error {
	search, err := contributions.GetSearchListByUserContributionIDList(contributionIDList)
	if err != nil {
		return err
	}

	return contributions.SaveToFollowCount(search, followMap)
}
