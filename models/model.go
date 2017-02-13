package models

import (
	"strings"

	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// gormConnect gorm接続を取得する
func gormConnect() *gorm.DB {
	dbms := "mysql"
	user := beego.AppConfig.String("mysqluser")
	pass := beego.AppConfig.String("mysqlpass")
	protocol := "tcp(127.0.0.1:3306)"
	database := beego.AppConfig.String("mysqldb")

	connect := user + ":" + pass + "@" + protocol + "/" + database + "?parseTime=true&loc=Asia%2FTokyo"
	db, err := gorm.Open(dbms, connect)

	if err != nil {
		panic(err.Error())
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

	return db
}

// GetWhere 条件から取得する
func GetWhere(dbModel interface{}, where string, bindList []map[string]interface{}, option map[string]interface{}) error {
	db := getDbOption(where, bindList, option)

	return db.First(dbModel).Error
}

// GetListWhere 条件からリストを取得する
func GetListWhere(dbModel interface{}, where string, bindList []map[string]interface{}, option map[string]interface{}) error {
	db := getDbOption(where, bindList, option)

	return db.Find(dbModel).Error
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
