package models

import (
	"dotstamp_server/tests"

	. "gopkg.in/check.v1"
)

type TestUserProfileImage struct {
	test.TSuite
}

func init() {
	var t test.Accessor = &TestUserProfileImage{}
	t.SetTableNameList([]string{
		"user_profile_images",
	})

	var _ = Suite(t)
}

func (t *TestUserProfileImage) TestAdd(c *C) {
	u := &UserProfileImage{
		UserID: 1,
	}

	u.GetIDAndAdd()

	r, _, _ := u.GetListByUserID(u.UserID)

	c.Check(r[0].ID, Equals, uint(1))
	c.Check(r[1].ID, Equals, uint(3))

	c.Check(r[1].UserID, Equals, 1)
}

func (t *TestUserProfileImage) TestGetListByUserID(c *C) {
	u := &UserProfileImage{}

	r, _, _ := u.GetListByUserID(1)

	c.Check(r[0].ID, Equals, uint(1))
}
