package test

import (
	"dotstamp_server/tests/database"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
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

// CopyTestFile テストファイルをコピーする
func CopyTestFile(id int) {
	apppath := getAppPath()

	src := apppath + "/tests/files/sound/1.mp3"
	dist := apppath + "/static/files/sound/" + strconv.Itoa(id) + ".mp3"

	cmd := "cp " + src + " " + dist
	exec.Command("sh", "-c", cmd).Output()

	src = apppath + "/tests/files/sound/1.m4a"
	dist = apppath + "/static/files/tmp/sound/" + strconv.Itoa(id) + ".m4a"

	cmd = "cp " + src + " " + dist
	exec.Command("sh", "-c", cmd).Output()

	src = apppath + "/tests/files/sound/1.wav"
	dist = apppath + "/static/files/tmp/sound/" + strconv.Itoa(id) + ".wav"

	cmd = "cp " + src + " " + dist
	exec.Command("sh", "-c", cmd).Output()

	src = apppath + "/tests/files/sound/1.wav"
	dist = apppath + "/static/files/tmp/sound/" + strconv.Itoa(id) + "_1.wav"

	cmd = "cp " + src + " " + dist
	exec.Command("sh", "-c", cmd).Output()

	src = apppath + "/tests/files/movie/1.mp4"
	dist = apppath + "/static/files/tmp/movie/" + strconv.Itoa(id) + ".mp4"

	cmd = "cp " + src + " " + dist
	exec.Command("sh", "-c", cmd).Output()

	src = apppath + "/tests/files/movie/1.mp4"
	dist = apppath + "/static/files/movie/" + strconv.Itoa(id) + ".mp4"

	cmd = "cp " + src + " " + dist
	exec.Command("sh", "-c", cmd).Output()
}
