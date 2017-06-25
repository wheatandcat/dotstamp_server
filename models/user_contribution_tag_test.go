package models

import (
	"github.com/wheatandcat/dotstamp_server/tests"

	. "gopkg.in/check.v1"
)

type TestUserContributionTag struct {
	test.TSuite
}

func init() {
	var t test.Accessor = &TestUserContributionTag{}
	t.SetTableNameList([]string{
		"user_contribution_tags",
	})

	var _ = Suite(t)
}

func (t *TestUserContributionTag) TestAddList(c *C) {
	userTag := []UserContributionTag{
		{
			UserContributionID: 3,
			Name:               "aaa",
		},
		{
			UserContributionID: 4,
			Name:               "bbb",
		},
	}

	u := UserContributionTag{}
	u.AddList(userTag)

	r, _, _ := u.GetListByUserContributionIDList([]int{3, 4})

	c.Check(r[0].UserContributionID, Equals, 3)
	c.Check(r[1].UserContributionID, Equals, 4)
}

func (t *TestUserContributionTag) TestGetListByUserContributionID(c *C) {
	u := UserContributionTag{}

	r, _, _ := u.GetListByUserContributionID(1)

	c.Check(r[0].UserContributionID, Equals, 1)
	c.Check(r[0].Name, Equals, "abc")
}

func (t *TestUserContributionTag) TestGetListByUserContributionIDList(c *C) {
	u := UserContributionTag{}

	r, _, _ := u.GetListByUserContributionIDList([]int{1, 2})

	c.Check(r[0].UserContributionID, Equals, 1)
	c.Check(r[0].Name, Equals, "abc")
}

func (t *TestUserContributionTag) TestGetByID(c *C) {
	u := UserContributionTag{}

	r, _, _ := u.GetByID(1)

	c.Check(r.ID, Equals, uint(1))
	c.Check(r.Name, Equals, "abc")
}

func (t *TestUserContributionTag) TestSave(c *C) {
	u := UserContributionTag{}
	u, _, _ = u.GetByID(1)
	u.Name = "ddd"

	u.Save()

	r, _, _ := u.GetByID(1)
	c.Check(r.ID, Equals, uint(1))
	c.Check(r.Name, Equals, "ddd")
}
