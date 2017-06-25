package models

import (
	"github.com/wheatandcat/dotstamp_server/tests"

	. "gopkg.in/check.v1"
)

type TestLogBugReport struct {
	test.TSuite
}

func init() {
	var t test.Accessor = &TestLogBugReport{}
	t.SetTableNameList([]string{
		"log_bug_reports",
	})

	var _ = Suite(t)
}

func (t *TestLogBugReport) TestAdd(c *C) {
	u := &LogBugReport{
		UserID: 1,
		Body:   "abc",
	}

	r := u.Add()
	c.Check(r, Equals, nil)
}
