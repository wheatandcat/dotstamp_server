package characters

import (
	"dotstamp_server/tests"

	. "gopkg.in/check.v1"
)

type TestMain struct {
	test.TSuite
}

func init() {
	var t test.Accessor = &TestMain{}
	t.SetTableNameList([]string{
		"user_character",
	})

	var _ = Suite(t)
}

func (t *TestMain) TestAdd(c *C) {
	uID := 10

	Add(uID, "test_abc", "{}", 1)

	r := GetListByUserID(uID)

	c.Check(r[0].Name, Equals, "test_abc")
}

func (t *TestMain) TestGetListByUserID(c *C) {
	r := GetListByUserID(1)

	c.Check(r[0].Name, Equals, "abc")
}
