package user

import (
	"dotstamp_server/tests"
	"dotstamp_server/utils"
	"time"

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
	AddForgetPassword("test@abcdef.com", "aaaa")

	r, _ := GetForgetPasswordByEmail("test@abcdef.com")

	c.Check(r.Email, Equals, "test@abcdef.com")
}

func (t *TestPassword) TestIsUpdatePassword(c *C) {
	loc, _ := time.LoadLocation("Asia/Tokyo")
	utils.SetNow(time.Date(2015, 1, 1, 11, 30, 0, 0, loc))
	r, _ := IsUpdatePassword("test@tedt.com", "abcdef")

	c.Check(r, Equals, false)

	utils.SetNow(time.Date(2015, 1, 1, 10, 30, 0, 0, loc))

	r, _ = IsUpdatePassword("test@tedt.com", "abcdef")

	c.Check(r, Equals, true)

	r, _ = IsUpdatePassword("test@tedt.com", "abcdefe")

	c.Check(r, Equals, false)
}

func (t *TestPassword) TestDeleteByEmail(c *C) {
	err := DeleteByEmail("test@tedt.com")
	c.Check(err, Equals, nil)

	r, _ := GetForgetPasswordByEmail("test@abcdef.com")
	c.Check(r.ID, Equals, uint(0))

	err = DeleteByEmail("test@tedt.com")
	c.Check(err, Equals, nil)
}
