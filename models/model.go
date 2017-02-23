package models

import (
	"strings"

	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// ErrRecordeNotFound レコードなし
const ErrRecordeNotFound = "record not found"

var db *gorm.DB

// gormConnect gorm接続を取得する
func gormConnect() *gorm.DB {
	if db != nil {
		return db
	}

	var err error
	dbms := "mysql"
	user := beego.AppConfig.String("mysqluser")
	pass := beego.AppConfig.String("mysqlpass")
	protocol := beego.AppConfig.String("mysqlhost")
	database := beego.AppConfig.String("mysqldb")

	connect := user + ":" + pass + "@" + protocol + "/" + database + "?parseTime=true&loc=Asia%2FTokyo"
	db, err = gorm.Open(dbms, connect)

	if err != nil {
		panic(err.Error())
	}

	if beego.AppConfig.String("runmode") == "dev" {
		db.LogMode(true)
	}

	return db
}

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
	db := gormConnect()

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
func GetCount(dbModel interface{}, where string, bindList []map[string]interface{}, option map[string]interface{}) (int, error) {
	db := getDbOption(where, bindList, option)

	count := 0
	err := db.Find(dbModel).Count(&count).Error
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
	db := gormConnect()

	return db.Create(dbModel).Error
}

// Save 更新する
func Save(dbModel interface{}) error {
	db := gormConnect()

	return db.Save(dbModel).Error
}

// Delete 削除する
func Delete(dbModel interface{}) error {
	db := gormConnect()

	return db.Delete(dbModel).Error
}

// Truncate 空にする
func Truncate(tableName string) error {
	db := gormConnect()

	return db.Exec("TRUNCATE TABLE " + tableName).Error
}
