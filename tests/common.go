package test

import (
	"io/ioutil"
	"path/filepath"
	"runtime"
	"dotstamp_server/models"

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

// SetupFixture フィクスチャー設定する
func SetupFixture(tableNameList []string) {
	for _, tableName := range tableNameList {

		deleteFixture(tableName)

		addFixture(tableName)
	}

}

// deleteFixture データを削除する
func deleteFixture(tableName string) {
	models.Truncate(tableName)
}

// addFixture データを追加する
func addFixture(tableName string) {
	apppath := getAppPath()
	dir := beego.AppConfig.String("resourcesDir")

	buf, err := ioutil.ReadFile(apppath + "/" + dir + "fixture/" + tableName + ".yml")
	if err != nil {
		panic(err)
	}

	var fixtures map[string]map[string]interface{}
	yaml.Unmarshal(buf, &fixtures)

	data := []map[string]interface{}{}
	for _, fixture := range fixtures {
		data = append(data, fixture)
	}

	err = models.InsertBatch(tableName, data)
	if err != nil {
		panic(err)
	}
}
