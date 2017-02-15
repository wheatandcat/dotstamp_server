package utils

import (
	"dotstamp_server/tests"
	"time"

	"github.com/jinzhu/gorm"

	. "gopkg.in/check.v1"
)

type TestCommon struct {
	test.TSuite
}

func init() {
	var t test.Accessor = &TestCommon{}
	t.SetTableNameList([]string{})

	var _ = Suite(t)
}

func (t *TestCommon) TestStringToDate(c *C) {
	r, _ := StringToDate("2014-12-31")

	c.Check(r, Equals, time.Date(2014, 12, 31, 0, 0, 0, 0, time.UTC))
}

func (t *TestCommon) TestGetAppPath(c *C) {
	r := GetAppPath()
	c.Check(r, Not(IsNil))
}

func (t *TestCommon) TestGetArrayCombile(c *C) {
	r, _ := GetArrayCombile([]string{"id", "data"}, []string{"1", "100"})

	c.Check(r["id"], Equals, "1")
	c.Check(r["data"], Equals, "100")

	_, e := GetArrayCombile([]string{"id", "data"}, []string{"1", "100", "3"})

	c.Check(e.Error(), Equals, "Both parameters should have an equal number of elements")
}

type TDB struct {
	Title      string
	UserID     int `json:"user_id"`
	gorm.Model `model:"true"`
}

func (t *TestCommon) TestDbStructToMap(c *C) {
	s := &TDB{
		Title:  "テスト",
		UserID: 1,
	}

	r := DbStructToMap(s)

	c.Check(r["user_id"], Equals, 1)
}

func (t *TestCommon) TestDbStructListToMapList(c *C) {
	s := []TDB{
		{
			Title:  "テスト",
			UserID: 1,
		},
	}

	r := DbStructListToMapList(s)

	c.Check(r[0]["user_id"], Equals, 1)
}

type T struct {
	ID    int
	Title string
}

func (t *TestCommon) TestStructToMap(c *C) {
	s := &T{
		ID:    1,
		Title: "テスト",
	}

	r := StructToMap(s)

	c.Check(r["ID"], Equals, 1)
	c.Check(r["Title"], Equals, "テスト")
}

func (t *TestCommon) TestStructListToMapList(c *C) {
	s := []T{
		{
			ID:    1,
			Title: "テスト",
		},
	}

	r := StructListToMapList(s)

	e := []map[string]interface{}{
		{
			"ID":    1,
			"Title": "テスト",
		},
	}

	c.Check(r[0]["ID"], Equals, e[0]["ID"])
	c.Check(r[0]["Title"], Equals, e[0]["Title"])
}

func (t *TestCommon) TestUrldecode(c *C) {
	e := Urlencode("test@add.com")
	r, _ := Urldecode(e)

	c.Check(r, Equals, "test@add.com")
}
