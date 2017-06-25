package models

import (
	"github.com/wheatandcat/dotstamp_server/tests"

	. "gopkg.in/check.v1"
)

type TestUserContributionSoundLength struct {
	test.TSuite
}

func init() {
	var t test.Accessor = &TestUserContributionSoundLength{}
	t.SetTableNameList([]string{
		"user_contribution_sound_lengths",
	})

	var _ = Suite(t)
}

func (t *TestUserContributionSoundLength) TestAdd(c *C) {
	u := &UserContributionSoundLength{
		UserContributionID: 100,
		Second:             1,
	}

	u.Add()

	r, _, _ := u.GetByUserContributionID(100)

	c.Check(r.UserContributionID, Equals, 100)
}

func (t *TestUserContributionSoundLength) TestSave(c *C) {
	u := &UserContributionSoundLength{}

	uc, _, _ := u.GetByUserContributionID(1)
	uc.Second = 2

	r := uc.Save()

	c.Check(r, Equals, nil)
}

func (t *TestUserContributionSoundLength) TestGetByUserContributionID(c *C) {
	u := &UserContributionSoundLength{}

	r, _, _ := u.GetByUserContributionID(1)

	c.Check(r.ID, Equals, uint(1))
}

func (t *TestUserContributionSoundLength) TestGetByTop(c *C) {
	u := &UserContributionSoundLength{}

	r, _, _ := u.GetByTop(0, 1)

	c.Check(r[0].ID, Equals, uint(1))
}
