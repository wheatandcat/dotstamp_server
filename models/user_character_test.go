package models

import (
	"dotstamp_server/tests"

	. "gopkg.in/check.v1"
)

type TestUserCharacter struct {
	test.TSuite
}

func init() {
	var t test.Accessor = &TestUserCharacter{}
	t.SetTableNameList([]string{
		"user_characters",
	})

	var _ = Suite(t)
}

func (t *TestUserCharacter) TestAdd(c *C) {
	u := &UserCharacter{
		UserID:   1,
		Name:     "test",
		Info:     "{}",
		Priority: 0,
	}

	u.Add()

	r, _, _ := u.GetListByUserID(u.UserID)

	c.Check(r[0].Name, Equals, "abc")
	c.Check(r[1].Name, Equals, "test")
}

func (t *TestUserCharacter) TestGetListByUserID(c *C) {
	u := &UserCharacter{}
	r, _, _ := u.GetListByUserID(2)

	c.Check(r[0].Name, Equals, "def")
}

func (t *TestUserCharacter) TestGetListByIDList(c *C) {
	e := []int{
		1,
		2,
	}

	u := &UserCharacter{}
	r, _, _ := u.GetListByIDList(e)

	c.Check(r[0].Name, Equals, "abc")
	c.Check(r[1].Name, Equals, "def")
}
