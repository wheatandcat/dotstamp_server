package models

import (
	"dotstamp_server/tests"

	. "gopkg.in/check.v1"
)

type TestUserContribution struct {
	test.TSuite
}

func init() {
	var t test.Accessor = &TestUserContribution{}
	t.SetTableNameList([]string{
		"user_contribution",
	})

	var _ = Suite(t)
}

func (t *TestUserContribution) TestGetIDAndAdd(c *C) {
	u := &UserContribution{
		UserID: 100,
		Title:  "ああああ",
	}

	r, _ := u.GetIDAndAdd()

	c.Check(r, Equals, 3)
}

func (t *TestUserContribution) TestDelete(c *C) {
	u := &UserContribution{}
	uc := u.GetByID(1)
	uc.Delete()

	r := u.GetByID(1)

	c.Check(r.DeleteFlag, Equals, 0)
}

func (t *TestUserContribution) TestGetByID(c *C) {
	u := &UserContribution{}

	r := u.GetByID(1)

	c.Check(r.ID, Equals, 1)
}

func (t *TestUserContribution) TestGetByTop(c *C) {
	u := &UserContribution{}

	r := u.GetByTop(0, 1)

	c.Check(r[0].ID, Equals, 2)
}
