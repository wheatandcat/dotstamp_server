package testsDatabase

import (
	"dotstamp_server/models/database"
	"strconv"
	"strings"
)

// ErrRecordeNotFound レコードなし
const ErrRecordeNotFound = "record not found"

// ErrFileTypeUnMatch レコードなし
const ErrFileTypeUnMatch = "file type unmatch"

// Execute クエリを実行する
func Execute(query string) error {
	db, err := database.GormConnect()
	if err != nil {
		return err
	}

	return db.Exec(query).Error
}

func checkError(err error) error {
	if err == nil {
		return nil
	}

	if err.Error() == ErrRecordeNotFound {
		return nil
	}

	if err.Error() == ErrFileTypeUnMatch {
		return nil
	}

	return err
}

// Truncate 空にする
func Truncate(tableName string) error {
	db, err := database.GormConnect()
	if err != nil {
		return err
	}

	err = db.Exec("TRUNCATE TABLE " + tableName).Error

	return checkError(err)
}

// GetFindAll 全て取得する
func GetFindAll(dbModel interface{}) error {
	db, err := database.GormConnect()
	if err != nil {
		return err
	}

	return db.Find(dbModel).Error
}

// InsertBatch 複数挿入する
func InsertBatch(tableName string, add []map[string]interface{}) error {
	db, err := database.GormConnect()
	if err != nil {
		return err
	}

	sql := "INSERT INTO " + tableName + " (`"

	column := []string{}

	for k := range add[0] {
		column = append(column, k)
	}

	val := map[int][]string{}

	for k, v := range add {
		for _, c := range column {
			insert := v[c]
			switch insert := insert.(type) {
			case string:
				val[k] = append(val[k], insert)
			case int:
				val[k] = append(val[k], strconv.Itoa(insert))
			}
		}
	}

	sql += strings.Join(column, "`,`")

	sql += "`) VALUES "

	s := []string{}
	for _, v := range val {
		s = append(s, "('"+strings.Join(v, "','")+"')")
	}

	sql += strings.Join(s, ",")

	return db.Exec(sql).Error
}
