package testsDatabase

import (
	"dotstamp_server/models/database"

	"github.com/astaxie/beedb"
)

// GetLink コネクションを取得する
func GetLink() beedb.Model {
	db := database.GetDB()
	orm := beedb.New(db)

	return orm
}

// Execute クエリを実行する
func Execute(query string) error {
	db := GetLink()

	_, err := db.Exec(query)

	return err
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
