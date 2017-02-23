package database

import (
	"database/sql"

	"github.com/astaxie/beedb"
	"github.com/astaxie/beego"
	_ "github.com/ziutek/mymysql/godrv"
)

var db *sql.DB

// GetDB DB取得する
func GetDB() *sql.DB {
	if db != nil {
		return db
	}

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
