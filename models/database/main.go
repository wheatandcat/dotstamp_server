package database

import (
	"database/sql"

	"github.com/astaxie/beego"
	_ "github.com/ziutek/mymysql/godrv"
)

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
