package models

import (
	"time"
	"dotstamp_server/tests"

	. "gopkg.in/check.v1"
)

type TestUserContributionTag struct {
	test.TSuite
}

func init() {
	var t test.Accessor = &TestUserContributionTag{}
	t.SetTableNameList([]string{
		"user_contribution_tag",
	})

	var _ = Suite(t)
}

func (t *TestUserContributionTag) TestAddList(c *C) {
	u := []UserContributionTag{
		{
			UserContributionID: 2,
			Name:               "test_edf",
			Updated:            time.Now(),
			Created:            time.Now(),
		},
	}

	userContributionTag := UserContributionTag{}
	userContributionTag.AddList(u)

	r := userContributionTag.GetFindAll()

	c.Check(r[0].UserContributionID, Equals, 1)
	c.Check(r[1].UserContributionID, Equals, 2)
}

func (t *TestUserContributionTag) TestGetListByUserContributionID(c *C) {
	u := UserContributionTag{}

	r := u.GetListByUserContributionID(1)

	c.Check(r[0].UserContributionID, Equals, 1)
	c.Check(r[0].Name, Equals, "abc")
}

func (t *TestUserContributionTag) TestGetListByUserContributionIDList(c *C) {
	u := UserContributionTag{}

	r := u.GetListByUserContributionIDList([]int{1, 2})

	c.Check(r[0].UserContributionID, Equals, 1)
	c.Check(r[0].Name, Equals, "abc")

	c.Check(r[1].UserContributionID, Equals, 2)
	c.Check(r[1].Name, Equals, "def")
}

func (t *TestUserContributionTag) TestGetByID(c *C) {
	u := UserContributionTag{}

	r := u.GetByID(1)

	c.Check(r.ID, Equals, 1)
	c.Check(r.Name, Equals, "abc")
}

func (t *TestUserContributionTag) TestSave(c *C) {
	u := UserContributionTag{}
	u = u.GetByID(1)
	u.Name = "ddd"

	u.Save()
	r := u.GetByID(1)
	c.Check(r.ID, Equals, 1)
	c.Check(r.Name, Equals, "ddd")
}
