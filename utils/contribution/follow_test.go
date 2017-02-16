package contributions

import (
	"dotstamp_server/models"
	"dotstamp_server/tests"

	. "gopkg.in/check.v1"
)

type TestContributionFollow struct {
	test.TSuite
}

func init() {
	var t test.Accessor = &TestContributionFollow{}
	t.SetTableNameList([]string{
		"user_contribution_follows",
	})

	var _ = Suite(t)
}

func (t *TestContributionFollow) TestAddFollow(c *C) {
	AddFollow(1, 100)

	r, _ := GetFollowListByUserContributionID(100)

	c.Check(r[0].UserContributionID, Equals, 100)
}

func (t *TestContributionFollow) TestDeleteFollow(c *C) {
	DeleteFollow(1)

	u := models.UserContributionFollow{}
	r, _, _ := u.GetByID(1)

	c.Check(r.ID, Equals, uint(0))
}
