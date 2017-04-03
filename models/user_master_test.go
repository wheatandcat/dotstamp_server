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
		"user_masters",
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

	c.Check(r, Equals, uint(3))
}

func (t *TestUserMaster) TestSave(c *C) {
	u := &UserMaster{}
	user, _, _ := u.GetByID(1)

	user.Name = "test"

	user.Save()

	r, _, _ := u.GetByID(1)

	c.Check(r.Name, Equals, "test")
}

func (t *TestUserMaster) TestGetByID(c *C) {
	u := &UserMaster{}
	r, _, _ := u.GetByID(1)

	c.Check(r.ID, Equals, uint(1))
}

func (t *TestUserMaster) TestGetByEmail(c *C) {
	u := &UserMaster{}
	r, _, _ := u.GetByEmail("test@tedt.com")

	c.Check(r.ID, Equals, uint(1))
}

// User ユーザー情報
type User struct {
	ID             uint
	Name           string
	ProfileImageID int
}

func (t *TestUserMaster) TestGetScanByID(c *C) {
	u := &UserMaster{}
	scan := User{}

	u.GetScanByID(1, &scan)

	c.Check(scan.ID, Equals, uint(1))
}

func (t *TestUserMaster) TestGetListByIDList(c *C) {
	u := &UserMaster{}
	r, _, _ := u.GetListByIDList([]int{1, 2})

	c.Check(r[0].ID, Equals, uint(1))
	c.Check(r[1].ID, Equals, uint(2))
}

func (t *TestUserMaster) TestGetScanByIDList(c *C) {
	u := &UserMaster{}
	scan := []User{}

	u.GetScanByIDList([]int{1}, &scan)

	c.Check(scan[0].ID, Equals, uint(1))
}
