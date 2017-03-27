package models

import (
	"dotstamp_server/tests"

	. "gopkg.in/check.v1"
)

type TestUserContributionUpload struct {
	test.TSuite
}

func init() {
	var t test.Accessor = &TestUserContributionUpload{}
	t.SetTableNameList([]string{
		"user_contribution_uploads",
	})

	var _ = Suite(t)
}

func (t *TestUserContributionUpload) TestAdd(c *C) {
	u := &UserContributionUpload{
		UserContributionID: 100,
		Token:              "abc",
	}

	r := u.Add()

	c.Check(r, Equals, nil)
}

func (t *TestUserContributionUpload) TestSave(c *C) {
	u := &UserContributionUpload{}
	uc, _, _ := u.GetByUserContributionID(1)

	uc.Token = "def"

	r := uc.Save()

	c.Check(r, Equals, nil)
}
