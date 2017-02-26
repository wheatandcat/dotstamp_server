package models

import (
	"strconv"
	"strings"

	"dotstamp_server/models/database"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// ErrRecordeNotFound レコードなし
const ErrRecordeNotFound = "record not found"

// getBindAndPlaceHolder バインドとプレースホルダの結果を取得する
func getBindAndPlaceHolder(where string, bindList []map[string]interface{}) (string, []interface{}) {
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

// getDbOption DBオプションを取得する
func getDbOption(where string, bindList []map[string]interface{}, option map[string]interface{}) *gorm.DB {
	db := database.GormConnect()

	if where != "" {
		w, bind := getBindAndPlaceHolder(where, bindList)
		db = db.Where(w, bind...)
	}

	if order, ok := option["order"].(string); ok {
		db = db.Order(order)
	}

	if limit, ok := option["limit"].(int); ok {
		db = db.Limit(limit)
	}

	if offset, ok := option["offset"].(int); ok {
		db = db.Offset(offset)
	}

	if sel, ok := option["select"].(string); ok {
		db = db.Select(sel)
	}

	return db
}

// GetWhere 条件から取得する
func GetWhere(dbModel interface{}, where string, bindList []map[string]interface{}, option map[string]interface{}) (*gorm.DB, error) {
	db := getDbOption(where, bindList, option)

	err := db.First(dbModel).Error
	if err != nil && err.Error() == ErrRecordeNotFound {
		return db, nil
	}

	return db, err
}

// GeScanWhere 条件から置き換えリストを取得する
func GeScanWhere(dest interface{}, name string, where string, bindList []map[string]interface{}, option map[string]interface{}) error {
	where += " AND Deleted_at IS NULL"

	db := getDbOption(where, bindList, option)

	err := db.Table(name).Scan(dest).Error
	if err != nil && err.Error() == ErrRecordeNotFound {
		return nil
	}

	return err
}

// GetCount 条件から数を取得する
func GetCount(dbModel interface{}, name string, where string, bindList []map[string]interface{}, option map[string]interface{}) (int, error) {
	db := getDbOption(where, bindList, option)

	count := 0
	err := db.Table(name).Count(&count).Error
	if err != nil && err.Error() == ErrRecordeNotFound {
		return 0, nil
	}

	return count, err
}

// GetListWhere 条件からリストを取得する
func GetListWhere(dbModel interface{}, where string, bindList []map[string]interface{}, option map[string]interface{}) (*gorm.DB, error) {
	db := getDbOption(where, bindList, option)

	err := db.Find(dbModel).Error
	if err != nil && err.Error() == ErrRecordeNotFound {
		return db, nil
	}

	return db, err
}

// Create 作成する
func Create(dbModel interface{}) error {
	db := database.GormConnect()

	return db.Create(dbModel).Error
}

// Save 更新する
func Save(dbModel interface{}) error {
	db := database.GormConnect()

	return db.Save(dbModel).Error
}

// Delete 削除する
func Delete(dbModel interface{}) error {
	db := database.GormConnect()

	return db.Delete(dbModel).Error
}

// Truncate 空にする
func Truncate(tableName string) error {
	db := database.GormConnect()

	return db.Exec("TRUNCATE TABLE " + tableName).Error
}

// InsertBatch 複数挿入する
func InsertBatch(tableName string, add []map[string]interface{}) error {
	db := database.GormConnect()

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

// Begin トランザクションを貼る
func Begin() *gorm.DB {
	db := database.GormConnect()
	tx := db.Begin()
	database.Transaction(tx)

	return tx
}

// Rollback ロールバックする
func Rollback(db *gorm.DB) {
	db.Rollback()
	database.Transaction(nil)
}

// Commit コミットする
func Commit(db *gorm.DB) {
	db.Commit()
	database.Transaction(nil)
}
