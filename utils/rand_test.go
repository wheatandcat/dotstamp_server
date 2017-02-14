package utils

import (
	"dotstamp_server/tests"

	"github.com/astaxie/beego"

	. "gopkg.in/check.v1"
)

type TestRand struct {
	test.TSuite
}

func init() {
	var t test.Accessor = &TestCommon{}
	t.SetTableNameList([]string{})

	var _ = Suite(t)
}

func (t *TestCommon) TestGetRandNum(c *C) {
	SetTestRandNum(10)
	r := GetRandNum(15)

	c.Check(r, Equals, 10)

	beego.BConfig.RunMode = "dev"

	SetTestRandNum(15)
	r = GetRandNum(10)

	c.Check(r, Not(Equals), 15)
	beego.BConfig.RunMode = "test"
}

func (t *TestCommon) TestGetRandString(c *C) {
	SetTestRandString("abc")
	r := GetRandString(15)

	c.Check(r, Equals, "abc")

	beego.BConfig.RunMode = "dev"

	SetTestRandString("abc")
	r = GetRandString(15)

	c.Check(r, Not(Equals), "abc")
	beego.BConfig.RunMode = "test"
}
