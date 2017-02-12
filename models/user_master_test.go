package models

import (
	"dotstamp_server/tests"

	. "gopkg.in/check.v1"
)

type TestUserMaster struct {
	test.TSuite
}

func init() {
	var t test.Accessor = &TestUserMaster{}
	t.SetTableNameList([]string{
		"user_master",
	})

	var _ = Suite(t)
}

func (t *TestUserMaster) TestGetIDAndAdd(c *C) {
	u := &UserMaster{
		Name:     "abc",
		Email:    "def@text.com",
		Password: "abcdef",
	}

	r, _ := u.GetIDAndAdd()

	c.Check(r, Equals, 3)
}

func (t *TestUserMaster) TestSave(c *C) {
	u := &UserMaster{}
	user := u.GetByID(1)

	user.Name = "test"

	user.Save()

	r := u.GetByID(1)

	c.Check(r.Name, Equals, "test")
}

func (t *TestUserMaster) TestGetByID(c *C) {
	u := &UserMaster{}
	r := u.GetByID(1)

	c.Check(r.ID, Equals, 1)
}

func (t *TestUserMaster) TestGetByEmail(c *C) {
	u := &UserMaster{}
	r := u.GetByEmail("test@tedt.com")

	c.Check(r.ID, Equals, 1)
}

func (t *TestUserMaster) TestGetListByIDList(c *C) {
	u := &UserMaster{}
	r := u.GetListByIDList([]int{1, 2})

	c.Check(r[0].ID, Equals, 1)
	c.Check(r[1].ID, Equals, 2)
}
