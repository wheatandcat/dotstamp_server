package user

import (
	"github.com/wheatandcat/dotstamp_server/tests"

	. "gopkg.in/check.v1"
)

type TestMain struct {
	test.TSuite
}

func init() {
	var t test.Accessor = &TestMain{}
	t.SetTableNameList([]string{
		"user_masters",
	})

	var _ = Suite(t)
}

func (t *TestMain) TestAdd(c *C) {
	uID, _ := Add("test@gmail.com", "abc", "abcdef")

	c.Check(uID, Equals, uint(3))

	r, _ := GetByEmail("test@gmail.com")

	c.Check(r.ID, Equals, uint(3))
	c.Check(r.Email, Equals, "test@gmail.com")
}

func (t *TestMain) TestGetByEmailAndPassword(c *C) {
	r, _ := GetByEmailAndPassword("test@tedt.com", "abc")
	c.Check(r.ID, Equals, uint(1))

	r, _ = GetByEmailAndPassword("test@tedt.com", "kbk")
	c.Check(r.ID, Equals, uint(0))
}

func (t *TestMain) TestGetByUserID(c *C) {
	r, _ := GetByUserID(1)

	c.Check(r.ID, Equals, uint(1))
}

func (t *TestMain) TestUpadateToProfileImageID(c *C) {
	UpadateToProfileImageID(1, 3)
	r, _ := GetByUserID(1)

	c.Check(r.ProfileImageID, Equals, 3)
}

func (t *TestMain) TestUpadate(c *C) {
	Upadate(1, "abcdef")
	r, _ := GetByUserID(1)

	c.Check(r.Name, Equals, "abcdef")
}

func (t *TestMain) TestGetMaptByUserIDList(c *C) {
	r, _ := GetMaptByUserIDList([]int{1})

	c.Check(r[1].ID, Equals, uint(1))
}

func (t *TestMain) TestUpadateToPassword(c *C) {
	UpadateToPassword("test@tedt.com", "bbb")
	r, _ := GetByEmail("test@tedt.com")

	c.Check(r.Password, Equals, GetPassword("bbb"))
}
