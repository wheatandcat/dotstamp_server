package contributions

import (
	"dotstamp_server/tests"

	. "gopkg.in/check.v1"
)

type TestLog struct {
	test.TSuite
}

func init() {
	var t test.Accessor = &TestLog{}
	t.SetTableNameList([]string{
		"log_user_contributions",
	})

	var _ = Suite(t)
}

func (t *TestLog) TestAdd(c *C) {
	r := AddLog(1, 1)

	c.Check(r, Equals, nil)
}
