package models

import (
	"github.com/wheatandcat/dotstamp_server/tests"

	. "gopkg.in/check.v1"
)

type TestUserForgetPassword struct {
	test.TSuite
}

func init() {
	var t test.Accessor = &TestUserForgetPassword{}
	t.SetTableNameList([]string{
		"user_forget_passwords",
	})

	var _ = Suite(t)
}

func (t *TestUserForgetPassword) TestAdd(c *C) {
	u := &UserForgetPassword{
		Email:   "abc@test.com",
		Keyword: "abcdef",
	}
	u.Add()

	r, _, _ := u.GetByEmail("abc@test.com")

	c.Check(r.ID, Equals, uint(3))
}

func (t *TestUserForgetPassword) TestDelete(c *C) {
	u := &UserForgetPassword{}
	user, _, _ := u.GetByEmail("test@tedt.com")

	user.Delete()

	r, _, _ := u.GetByEmail("test@tedt.com")

	c.Check(r.ID, Equals, uint(0))
}

func (t *TestUserForgetPassword) TestDeleteList(c *C) {
	u := &UserForgetPassword{}
	user, _, _ := u.GetListByEmail("test@tedt.com")

	u.DeleteList(user)

	r, _, _ := u.GetByEmail("test@tedt.com")

	c.Check(r.ID, Equals, uint(0))
}

func (t *TestUserForgetPassword) TestGetByEmail(c *C) {
	u := &UserForgetPassword{}
	r, _, _ := u.GetByEmail("test@tedt.com")

	c.Check(r.ID, Equals, uint(1))
}

func (t *TestUserForgetPassword) TestGetListByEmail(c *C) {
	u := &UserForgetPassword{}
	r, _, _ := u.GetListByEmail("test@tedt.com")

	c.Check(r[0].ID, Equals, uint(1))
}
