package utils

import (
	"time"

	"github.com/wheatandcat/dotstamp_server/tests"

	"github.com/astaxie/beego"

	. "gopkg.in/check.v1"
)

type TestTime struct {
	test.TSuite
}

func init() {
	var t test.Accessor = &TestTime{}
	t.SetTableNameList([]string{})

	var _ = Suite(t)
}

func (t *TestCommon) TestNow(c *C) {
	beego.BConfig.RunMode = "dev"

	r1 := Now()

	beego.BConfig.RunMode = "test"

	loc, _ := time.LoadLocation("Asia/Tokyo")
	SetNow(time.Date(2015, 1, 1, 0, 0, 0, 0, loc))

	r2 := Now()

	c.Check(r1, Not(Equals), r2)
}
