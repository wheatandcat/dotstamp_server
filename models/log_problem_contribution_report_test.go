package models

import (
	"github.com/wheatandcat/dotstamp_server/tests"

	. "gopkg.in/check.v1"
)

type TestLogProblemContributionReport struct {
	test.TSuite
}

func init() {
	var t test.Accessor = &TestLogProblemContributionReport{}
	t.SetTableNameList([]string{
		"log_problem_contribution_reports",
	})

	var _ = Suite(t)
}

func (t *TestLogProblemContributionReport) TestAdd(c *C) {
	l := &LogProblemContributionReport{
		UserID:             1,
		UserContributionID: 1,
		Type:               1,
	}

	r := l.Add()
	c.Check(r, Equals, nil)
}
