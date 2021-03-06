package user

import (
	"github.com/wheatandcat/dotstamp_server/tests"

	. "gopkg.in/check.v1"
)

type TestProfile struct {
	test.TSuite
}

func init() {
	var t test.Accessor = &TestProfile{}
	t.SetTableNameList([]string{
		"user_profile_images",
	})

	var _ = Suite(t)
}

func (t *TestProfile) TestGetProfileImageListByUserID(c *C) {
	r, _ := GetProfileImageListByUserID(1)

	c.Check(r[0].ID, Equals, uint(1))
}

func (t *TestProfile) TestGetIDAndAddProfileImage(c *C) {
	r, _ := GetIDAndAddProfileImage(1)

	c.Check(r, Equals, uint(3))
}
