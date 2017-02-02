package main

import (
	_ "dotstamp_server/routers"

	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/redis"
	_ "github.com/garyburd/redigo/redis"
)

func main() {
	beego.Run()
}
