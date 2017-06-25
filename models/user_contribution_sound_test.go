package models

import (
	"github.com/wheatandcat/dotstamp_server/tests"

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

func (t *TestUserContributionSound) TestSave(c *C) {
	u := &UserContributionSound{}

	uc, _, _ := u.GetByUserContributionID(1)
	uc.SoundStatus = 2

	r := uc.Save()

	c.Check(r, Equals, nil)
}

func (t *TestUserContributionSound) TestGetListByUserContributionIDList(c *C) {
	u := &UserContributionSound{}

	r, _, _ := u.GetListByUserContributionIDList([]int{1})

	c.Check(r[0].ID, Equals, uint(1))
}
