package problem

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
		"log_problem_contribution_reports",
	})

	var _ = Suite(t)
}

func (t *TestMain) TestAdd(c *C) {
	r := Add(1, 1, 1)

	c.Check(r, Equals, nil)
}
