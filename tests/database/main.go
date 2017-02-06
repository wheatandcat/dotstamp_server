package testsDatabase

import "dotstamp_server/models/database"

// Execute クエリを実行する
func Execute(query string) error {
	db := database.GetLink()

	_, err := db.Exec(query)

	return err
}

// Truncate テーブルデータを空にする
func Truncate(tableName string) error {
	err := Execute("TRUNCATE TABLE " + tableName)

	return err
}

// GetFindAll 全て取得する
func GetFindAll(dbModel interface{}) error {
	db := database.GetLink()

	return db.FindAll(dbModel)
}

// InsertBatch 挿入する(複数)
func InsertBatch(tableName string, add []map[string]interface{}) error {
	db := database.GetLink()

	_, err := db.SetTable(tableName).InsertBatch(add)

	return err
}
