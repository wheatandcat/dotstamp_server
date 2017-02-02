package models

import (
	"database/sql"
	"strings"

	"github.com/astaxie/beedb"
	"github.com/astaxie/beego"
	_ "github.com/ziutek/mymysql/godrv"
)

// DeleteFlagOn 削除ON
const DeleteFlagOn = 1

// DeleteFlagOff 削除OFF
const DeleteFlagOff = 0

// GetDB DB取得する
func GetDB() *sql.DB {
	username := beego.AppConfig.String("mysqluser")
	password := beego.AppConfig.String("mysqlpass")
	database := beego.AppConfig.String("mysqldb")

	db, err := sql.Open("mymysql", database+"/"+username+"/"+password)
	if err != nil {
		panic(err)
	}

	return db
}

// GetLink コネクションを取得する
func GetLink() beedb.Model {
	db := GetDB()
	orm := beedb.New(db)

	return orm
}

// Execute クエリを実行する
func Execute(query string) error {
	db := GetLink()

	_, err := db.Exec(query)

	return err
}

// Begin トランザクション
func Begin() error {
	return Execute("BEGIN")
}

// Commit コミット
func Commit() error {
	return Execute("COMMIT")
}

// Rollback ロールバック
func Rollback() error {
	return Execute("ROLLBACK")
}

// Truncate テーブルデータを空にする
func Truncate(tableName string) {
	Execute("TRUNCATE TABLE " + tableName)
}

// InsertBatch 挿入する(複数)
func InsertBatch(tableName string, add []map[string]interface{}) error {
	db := GetLink()

	_, err := db.SetTable(tableName).InsertBatch(add)

	return err
}

// Save 保存する
func Save(dbModel interface{}) error {
	db := GetLink()

	return db.Save(dbModel)
}

// GetFindAll 全て取得する
func GetFindAll(dbModel interface{}) error {
	db := GetLink()

	return db.FindAll(dbModel)
}

// GetBindAndPlaceHolder バインドとプレースホルダの結果を取得する
func GetBindAndPlaceHolder(where string, bindList []map[string]interface{}) (string, []interface{}) {
	bind := []interface{}{}
	var holder string

	for _, list := range bindList {
		for key, value := range list {

			switch value := value.(type) {
			// 配列対応
			case []int:
				holder = " ("

				i := 0
				for _, data := range value {
					if i > 0 {
						holder += ", "
					}
					holder += "?"
					bind = append(bind, data)
					i++
				}

				holder += ") "
			default:
				holder = "?"
				bind = append(bind, value)
			}

			where = strings.Replace(where, ":"+key, holder, 1)
		}
	}

	return where, bind
}

// GetDbOption DBオプションを取得する
func GetDbOption(where string, bindList []map[string]interface{}, option map[string]interface{}) beedb.Model {
	db := GetLink()

	if where != "" {
		w, bind := GetBindAndPlaceHolder(where, bindList)
		db.Where(w, bind...)
	}

	if order, ok := option["order"].(string); ok {
		db.OrderBy(order)
	}

	if limit, ok := option["limit"].(map[string]int); ok {
		db.Limit(limit["size"], limit["offset"])
	}

	return db
}

// GetListWhere 条件からリストを取得する
func GetListWhere(dbModel interface{}, where string, bindList []map[string]interface{}, option map[string]interface{}) interface{} {
	db := GetDbOption(where, bindList, option)

	db.FindAll(dbModel)

	return dbModel
}

// GetWhere 条件から取得する
func GetWhere(dbModel interface{}, where string, bindList []map[string]interface{}, option map[string]interface{}) interface{} {
	db := GetDbOption(where, bindList, option)

	db.Find(dbModel)

	return dbModel
}
