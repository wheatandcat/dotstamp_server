package models

import (
	"dotstamp_server/tests"

	. "gopkg.in/check.v1"
)

type TestUserContributionFollow struct {
	test.TSuite
}

func init() {
	var t test.Accessor = &TestUserContributionFollow{}
	t.SetTableNameList([]string{
		"user_contribution_follows",
	})

	var _ = Suite(t)
}

func (t *TestUserContributionFollow) TestAdd(c *C) {
	u := &UserContributionFollow{
		UserID:             1,
		UserContributionID: 100,
	}

	u.Add()

	r, _, _ := u.GetListByUserContributionID(100)

	c.Check(r[0].UserContributionID, Equals, 100)
}

func (t *TestUserContributionFollow) TestDelete(c *C) {
	u := &UserContributionFollow{}
	f, _, _ := u.GetByID(1)

	f.Delete()

	r, _, _ := u.GetByID(1)

	c.Check(r.ID, Equals, uint(0))
}

func (t *TestUserContributionFollow) TestByID(c *C) {
	u := &UserContributionFollow{}
	r, _, _ := u.GetByID(1)

	c.Check(r.ID, Equals, uint(1))
}

func (t *TestUserContributionFollow) TestGetListByUserContributionID(c *C) {
	u := &UserContributionFollow{}
	r, _, _ := u.GetListByUserContributionID(1)

	c.Check(r[0].ID, Equals, uint(1))
}

func (t *TestUserContributionFollow) TestGetListByUserContributionIDList(c *C) {
	u := &UserContributionFollow{}
	r, _, _ := u.GetListByUserContributionIDList([]int{
		1,
	})

	c.Check(r[0].ID, Equals, uint(1))
}
