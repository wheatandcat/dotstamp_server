package database

import (
	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	_ "github.com/ziutek/mymysql/godrv"
)

var db *gorm.DB
var transactionDB *gorm.DB

// GormConnect gorm接続を取得する
func GormConnect() (*gorm.DB, error) {
	if transactionDB != nil {
		return transactionDB, nil
	}

	if db != nil {
		return db, nil
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
		return db, err
	}

	if beego.AppConfig.String("runmode") == "dev" {
		db.LogMode(true)
	}

	return db, nil
}

// Transaction トランザクション
func Transaction(db *gorm.DB) {
	transactionDB = db
}
