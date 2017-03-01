package contributions

import (
	"dotstamp_server/tests"

	. "gopkg.in/check.v1"
)

type TestSound struct {
	test.TSuite
}

func init() {
	var t test.Accessor = &TestSound{}
	t.SetTableNameList([]string{
		"user_contribution_sounds",
	})

	var _ = Suite(t)
}

func (t *TestSound) TestGetByUserContributionID(c *C) {
	r, _ := GetSoundByUserContributionID(1)

	c.Check(r.UserContributionID, Equals, 1)
}

func (t *TestSound) TestAddSound(c *C) {
	AddSound(100, 1)

	r, _ := GetSoundByUserContributionID(100)

	c.Check(r.UserContributionID, Equals, 100)
}
