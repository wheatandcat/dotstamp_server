package models

import (
	"dotstamp_server/tests"

	. "gopkg.in/check.v1"
)

type TestUserContributionSound struct {
	test.TSuite
}

func init() {
	var t test.Accessor = &TestUserContributionSound{}
	t.SetTableNameList([]string{
		"user_contribution_sounds",
	})

	var _ = Suite(t)
}

func (t *TestUserContributionSound) TestAdd(c *C) {
	u := &UserContributionSound{
		UserContributionID: 100,
		SoundStatus:        1,
	}

	u.Add()

	r, _, _ := u.GetByUserContributionID(100)

	c.Check(r.UserContributionID, Equals, 100)
}
