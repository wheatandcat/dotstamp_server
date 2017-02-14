package controllers

import (
	"log"

	"github.com/astaxie/beego"
)

var session = make(map[string]interface{})

// isSession セッションを使用するか判定する
func isSession() bool {
	if beego.AppConfig.String("runmode") == "test" {
		return false
	}

	return true
}

// GetUserID ユーザIDを取得する
func (c *BaseController) GetUserID() int {
	if !isSession() {
		r, _ := c.GetInt("user_id", 1)
		return r
	}

	uID := c.GetSession("user_id")
	log.Println(uID)
	if uID, ok := uID.(int); ok {
		return uID
	}
	if uintID, ok := uID.(uint); ok {
		return int(uintID)
	}

	return noUserID
}

// GetSession セッションを取得する
func (c *BaseController) GetSession(name string) interface{} {
	if !isSession() {
		return session[name]
	}

	if c.CruSession == nil {
		c.StartSession()
	}
	return c.CruSession.Get(name)
}

// SetSession セッションを設定する
func (c *BaseController) SetSession(name string, value interface{}) {
	log.Println("SetSession")
	if !isSession() {
		session[name] = value
		return
	}

	if c.CruSession == nil {
		c.StartSession()
	}

	log.Println(name)
	log.Println(value)
	c.CruSession.Set(name, value)
}

// DelSession セッションを削除する
func (c *BaseController) DelSession(name string) {
	if !isSession() {
		delete(session, name)
		return
	}

	if c.CruSession == nil {
		c.StartSession()
	}
	c.CruSession.Delete(name)
}
