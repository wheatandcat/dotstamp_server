package bug

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
		"log_bug_reports",
	})

	var _ = Suite(t)
}

func (t *TestMain) TestAdd(c *C) {
	r := Add(1, "abc")

	c.Check(r, Equals, nil)
}
