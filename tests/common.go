package test

import (
	"dotstamp_server/tests/database"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/astaxie/beego"
	"gopkg.in/yaml.v2"
)

// getAppPath アプリケーションパスを取得する
func getAppPath() string {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))

	return apppath
}

// Setup テストの初期設定
func Setup() {
	apppath := getAppPath()

	//beego.TestBeegoInit(apppath)
	err := beego.LoadAppConfig("ini", apppath+"/conf/app_test.conf")
	if err != nil {
		panic(err)
	}
}

// UserContributionTag ユーザー投稿タグ
type UserContributionTag struct {
	ID                 int `beedb:"PK"`
	UserContributionID int `sql:"user_contribution_id"`
	Name               string
	DeleteFlag         int `sql:"delete_flag"`
	Created            time.Time
	Updated            time.Time
}

// SetupFixture フィクスチャー設定する
func SetupFixture(tableNameList []string) {
	var err error

	for _, tableName := range tableNameList {
		if err = deleteFixture(tableName); err == nil {
			err = addFixture(tableName)
		}

		if err != nil {
			panic(err)
		}
	}

}

// deleteFixture データを削除する
func deleteFixture(tableName string) error {
	return testsDatabase.Truncate(tableName)
}

// addFixture データを追加する
func addFixture(tableName string) error {

	apppath := getAppPath()
	dir := beego.AppConfig.String("resourcesDir")

	buf, err := ioutil.ReadFile(apppath + "/" + dir + "fixture/" + tableName + ".yml")
	if err != nil {
		return err
	}

	var fixtures map[string]map[string]interface{}
	yaml.Unmarshal(buf, &fixtures)

	data := []map[string]interface{}{}
	for _, fixture := range fixtures {
		data = append(data, fixture)
	}

	err = testsDatabase.InsertBatch(tableName, data)
	if err != nil {
		return err
	}

	return nil
}

// removeLogFile ログファイル削除する
func removeLogFile(file string) error {
	apppath := getAppPath()

	return os.Remove(apppath + "/logs/" + file + ".log")
}
