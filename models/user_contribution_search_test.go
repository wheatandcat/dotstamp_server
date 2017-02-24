package models

import (
	"dotstamp_server/tests"

	. "gopkg.in/check.v1"
)

type TestUserContributionSearch struct {
	test.TSuite
}

func init() {
	var t test.Accessor = &TestUserContributionSearch{}
	t.SetTableNameList([]string{
		"user_contribution_searches",
	})

	var _ = Suite(t)
}

func (t *TestUserContributionSearch) TestAdd(c *C) {
	u := &UserContributionSearch{
		UserContributionID: 2,
		Search:             "abcdef",
	}

	u.Add()

	r, _, _ := u.GetByUserContributionID(2)

	c.Check(r.UserContributionID, Equals, 2)
}

func (t *TestUserContributionSearch) TestSave(c *C) {
	u := &UserContributionSearch{}

	ucs, _, _ := u.GetByUserContributionID(1)
	ucs.Search = "aaaaa"
	ucs.Save()

	r, _, _ := u.GetByUserContributionID(1)

	c.Check(r.Search, Equals, "aaaaa")
}

func (t *TestUserContributionSearch) TestDelete(c *C) {
	u := &UserContributionSearch{}

	ucs, _, _ := u.GetByUserContributionID(1)
	ucs.Delete()

	r, _, _ := u.GetByUserContributionID(1)

	c.Check(r.ID, Equals, uint(0))
}

func (t *TestUserContributionSearch) TestGetByUserContributionID(c *C) {
	u := &UserContributionSearch{}

	r, _, _ := u.GetByUserContributionID(1)

	c.Check(r.UserContributionID, Equals, 1)
}

func (t *TestUserContributionSearch) TestGetListByUserContributionIDList(c *C) {
	u := &UserContributionSearch{}

	r, _, _ := u.GetListByUserContributionIDList([]int{1})

	c.Check(r[0].UserContributionID, Equals, 1)
}

func (t *TestUserContributionSearch) TestGetListBySearch(c *C) {
	u := &UserContributionSearch{}

	r, _, _ := u.GetListBySearch("a", "ID desc", 10, 0)

	c.Check(r[0].Search, Equals, "abcdef")
	c.Check(r[1].Search, Equals, "abc")
}

func (t *TestUserContributionSearch) TestGetCountBySearch(c *C) {
	u := &UserContributionSearch{}

	r, _ := u.GetCountBySearch("a", "ID desc")

	c.Check(r, Equals, 2)
}
