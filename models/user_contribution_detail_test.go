package models

import (
	"dotstamp_server/tests"

	. "gopkg.in/check.v1"
)

type TestUserContributionDetail struct {
	test.TSuite
}

func init() {
	var t test.Accessor = &TestUserContributionDetail{}
	t.SetTableNameList([]string{
		"user_contribution_details",
	})

	var _ = Suite(t)
}

func (t *TestUserContributionDetail) TestAdd(c *C) {
	u := &UserContributionDetail{
		UserContributionID: 100,
		Body:               "{}",
	}

	u.Add()

	r := u.GetByUserContributionID(100)

	c.Check(r.UserContributionID, Equals, 100)
}

func (t *TestUserContributionDetail) TestSave(c *C) {
	u := &UserContributionDetail{}
	ucd := u.GetByUserContributionID(1)
	ucd.Body = "abc"
	ucd.Save()

	r := u.GetByUserContributionID(1)

	c.Check(r.Body, Equals, "abc")
}

func (t *TestUserContributionDetail) TestDelete(c *C) {
	u := &UserContributionDetail{}
	ucd := u.GetByUserContributionID(1)
	ucd.Delete()

	//r := u.GetByUserContributionID(1)
}

func (t *TestUserContributionDetail) TestGetByUserContributionID(c *C) {
	u := &UserContributionDetail{}

	r := u.GetByUserContributionID(1)

	c.Check(r.UserContributionID, Equals, 1)
}
