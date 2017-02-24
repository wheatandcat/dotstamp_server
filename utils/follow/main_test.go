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

func (t *TestMain) TestGetCountByUserContributionID(c *C) {
	r, _ := GetCountByUserContributionID(1)

	c.Check(r, Equals, 2)
}

func (t *TestMain) TestGetByUserIDAndUserContributionID(c *C) {
	r, _ := GetByUserIDAndUserContributionID(1000, 1)

	c.Check(r.ID, Equals, uint(1))
}

func (t *TestMain) TestGetCountByUserIDAndUserContributionID(c *C) {
	r, _ := GetCountByUserIDAndUserContributionID(1000, 1)

	c.Check(r, Equals, 1)
}

func (t *TestMain) TestGetListByUserID(c *C) {
	r, _ := GetListByUserID(1000, "ID desc", 10, 0)

	c.Check(r[0].ID, Equals, uint(1))
}

func (t *TestMain) TestGetOrderValueListByUserID(c *C) {
	r, _ := GetOrderValueListByUserID(1000, "ID desc", 10, 0)

	c.Check(r[0].Order, Equals, 0)

	r, _ = GetOrderValueListByUserID(1, "ID desc", 10, 0)

	c.Check(len(r), Equals, 0)
}

func (t *TestMain) TestGetListByUserContributionIDList(c *C) {
	r, _ := GetListByUserContributionIDList([]int{1, 2})

	c.Check(r[0].UserContributionID, Equals, 1)
	c.Check(r[1].UserContributionID, Equals, 1)
}

func (t *TestMain) TestGetFollowCountMap(c *C) {
	u, _ := GetListByUserContributionIDList([]int{1, 2})

	r := GetFollowCountMap(u)
	c.Check(r[2], Equals, 1)
}

func (t *TestMain) TestGetCountByUserID(c *C) {
	r, _ := GetCountByUserID(1000, "ID desc", 10, 0)

	c.Check(r, Equals, 1)
}
