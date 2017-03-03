package models

import (
	"dotstamp_server/tests"

	. "gopkg.in/check.v1"
)

type TestLogUserContribution struct {
	test.TSuite
}

func init() {
	var t test.Accessor = &TestLogUserContribution{}
	t.SetTableNameList([]string{
		"log_user_contributions",
	})

	var _ = Suite(t)
}

func (t *TestLogUserContribution) TestAdd(c *C) {
	u := &LogUserContribution{
		UserID:             1,
		UserContributionID: 1,
	}

	r := u.Add()
	c.Check(r, Equals, nil)
}
