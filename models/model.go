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

// ErrFileTypeUnMatch レコードなし
const ErrFileTypeUnMatch = "file type unmatch"

const (
	// StatusPublic 状態；公開
	StatusPublic = 1
	// StatusPrivate 状態；非公開
	StatusPrivate = 2
	// StatusError 状態：エラー
	StatusError = 3
	// StatusRunning 状態：実行中
	StatusRunning = 4
	// StatusReMeake 状態：作り直し
	StatusReMeake = 5
	// StatusUploading 状態：アップロード中
	StatusUploading = 6
	// StatusMade 状態：作成済み
	StatusMade = 7
)

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
func getDbOption(where string, bindList []map[string]interface{}, option map[string]interface{}) (*gorm.DB, error) {
	db, err := database.GormConnect()
	if err != nil {
		return db, err
	}

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

	return db, nil
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

// GetWhere 条件から取得する
func GetWhere(dbModel interface{}, where string, bindList []map[string]interface{}, option map[string]interface{}) (*gorm.DB, error) {
	db, err := getDbOption(where, bindList, option)
	if err != nil {
		return db, err
	}

	err = db.First(dbModel).Error
	if err = checkError(err); err != nil {
		return db, nil
	}

	return db, err
}

// GeScanWhere 条件から置き換えリストを取得する
func GeScanWhere(dest interface{}, name string, where string, bindList []map[string]interface{}, option map[string]interface{}) error {
	where += " AND Deleted_at IS NULL"

	db, err := getDbOption(where, bindList, option)
	if err != nil {
		return err
	}

	err = db.Table(name).Scan(dest).Error
	if err = checkError(err); err != nil {
		return nil
	}

	return err
}

// GetCount 条件から数を取得する
func GetCount(dbModel interface{}, name string, where string, bindList []map[string]interface{}, option map[string]interface{}) (int, error) {
	db, err := getDbOption(where, bindList, option)
	if err != nil {
		return 0, err
	}

	count := 0
	err = db.Table(name).Count(&count).Error
	if err = checkError(err); err != nil {
		return 0, nil
	}

	return count, err
}

// GetListWhere 条件からリストを取得する
func GetListWhere(dbModel interface{}, where string, bindList []map[string]interface{}, option map[string]interface{}) (*gorm.DB, error) {
	db, err := getDbOption(where, bindList, option)
	if err != nil {
		return db, err
	}

	err = db.Find(dbModel).Error
	if err = checkError(err); err != nil {
		return db, nil
	}

	return db, err
}

// Update 条件から更新する
func Update(dbModel interface{}, s []interface{}, where string, bindList []map[string]interface{}, option map[string]interface{}) (*gorm.DB, error) {
	db, err := getDbOption(where, bindList, option)
	if err != nil {
		return db, err
	}

	err = db.Model(dbModel).Update(s...).Error
	if err = checkError(err); err != nil {
		return db, nil
	}

	return db, err
}

// Updates 条件から複数更新する
func Updates(dbModel interface{}, s interface{}, where string, bindList []map[string]interface{}, option map[string]interface{}) (*gorm.DB, error) {
	db, err := getDbOption(where, bindList, option)
	if err != nil {
		return db, err
	}

	err = db.Model(dbModel).Updates(s).Error
	if err = checkError(err); err != nil {
		return db, nil
	}

	return db, err
}

// Create 作成する
func Create(dbModel interface{}) error {
	db, err := database.GormConnect()
	if err != nil {
		return err
	}

	return db.Create(dbModel).Error
}

// Save 更新する
func Save(dbModel interface{}) error {
	db, err := database.GormConnect()
	if err != nil {
		return err
	}

	return db.Save(dbModel).Error
}

// Delete 削除する
func Delete(dbModel interface{}) error {
	db, err := database.GormConnect()
	if err != nil {
		return err
	}

	return db.Delete(dbModel).Error
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

// Begin トランザクションを貼る
func Begin() *gorm.DB {
	db, err := database.GormConnect()
	if err != nil {
		panic(err)
	}

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

// Lock ロックする
func Lock(tableName string, id int) error {
	db, err := database.GormConnect()
	if err != nil {
		return err
	}

	return db.Exec("select * from " + tableName + " where id = " + strconv.Itoa(id) + " for update").Error
}
