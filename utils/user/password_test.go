package user

import (
	"dotstamp_server/tests"

	. "gopkg.in/check.v1"
)

type TestPassword struct {
	test.TSuite
}

func init() {
	var t test.Accessor = &TestPassword{}
	t.SetTableNameList([]string{
		"user_forget_passwords",
	})

	var _ = Suite(t)
}

func (t *TestPassword) TestAddForgetPassword(c *C) {
	AddForgetPassword("test@abcdef.com")

	r, _ := GetForgetPasswordByEmail("test@abcdef.com")

	c.Check(r.Email, Equals, "test@abcdef.com")
}

func (t *TestPassword) TestIsUpdatePassword(c *C) {
	r, _ := IsUpdatePassword("test@tedt.com", "abcdef")

	c.Check(r, Equals, true)

	r, _ = IsUpdatePassword("test@tedt.com", "abcdefe")

	c.Check(r, Equals, false)
}
