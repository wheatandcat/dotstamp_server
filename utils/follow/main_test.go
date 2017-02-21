package follows

import (
	"dotstamp_server/models"
	"dotstamp_server/tests"

	. "gopkg.in/check.v1"
)

type TestMain struct {
	test.TSuite
}

func init() {
	var t test.Accessor = &TestMain{}
	t.SetTableNameList([]string{
		"user_contribution_follows",
	})

	var _ = Suite(t)
}

func (t *TestMain) TestAdd(c *C) {
	Add(1, 100)

	r, _ := GetListByUserContributionID(100)

	c.Check(r[0].UserContributionID, Equals, 100)
}

func (t *TestMain) TestDelete(c *C) {
	Delete(1)

	u := models.UserContributionFollow{}
	r, _, _ := u.GetByID(1)

	c.Check(r.ID, Equals, uint(0))
}

func (t *TestMain) TestGetByUserIDAndUserContributionID(c *C) {
	r, _ := GetByUserIDAndUserContributionID(1000, 1)

	c.Check(r.ID, Equals, uint(1))
}

func (t *TestMain) TestGetCountByUserIDAndUserContributionID(c *C) {
	r, _ := GetCountByUserIDAndUserContributionID(1000, 1)

	c.Check(r, Equals, 1)
}
