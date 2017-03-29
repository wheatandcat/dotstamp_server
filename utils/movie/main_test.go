package movie

import (
	"dotstamp_server/tests"

	. "gopkg.in/check.v1"
)

type TestMain struct {
	test.TSuite
}

func init() {
	var t test.Accessor = &TestMain{}
	t.SetTableNameList([]string{})

	var _ = Suite(t)
}

func (t *TestMain) TestMake(c *C) {
	r := Make("0")

	c.Check(r, Equals, nil)
}

func (t *TestMain) TestToFilter(c *C) {
	r := ToFilter("0")

	c.Check(r, Equals, nil)
}

func (t *TestMain) TestExecMakeMovie(c *C) {
	r := ExecMakeMovie(0)

	c.Check(r, Equals, nil)
}
