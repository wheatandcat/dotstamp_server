package csvModels

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
	"time"
	"dotstamp_server/utils"

	"github.com/astaxie/beego"
)

func failOnError(e error) {
	if e != nil {
		panic(e)
	}
}

// StringToInt 文字列を数値に変換する
func StringToInt(s string) int {
	i, e := strconv.Atoi(s)
	if e != nil {
		panic(e)
	}

	return i
}

// StringToDate 文字列を日付に変換する
func StringToDate(s string) time.Time {
	t, e := time.Parse("2006-01-02", s)
	if e != nil {
		panic(e)
	}

	return t
}

// GetMapAll 全てマップ取得する
func GetMapAll(f string) (r []map[string]string) {
	dir := beego.AppConfig.String("resourcesDir")

	file, e := os.Open(dir + "csv/" + f)

	failOnError(e)
	defer file.Close()

	reader := csv.NewReader(file)

	count := 0

	var columnList []string
	for {
		record, e := reader.Read() // 1行読み出す
		if e == io.EOF {
			return r
		}

		failOnError(e)

		// 一行目はカラム取得
		if count == 0 {
			columnList = record
			count++
			continue
		}

		// 項目追加
		list, _ := utils.GetArrayCombile(columnList, record)
		r = append(r, list)
	}
}
