package models

import (
	"dotstamp_server/tests"

	. "gopkg.in/check.v1"
)

type TestUserWorkHistory struct {
	test.TSuite
}

func init() {
	var t test.Accessor = &TestUserWorkHistory{}
	t.SetTableNameList([]string{
		"user_work_histories",
	})

	var _ = Suite(t)
}

func (t *TestUserWorkHistory) TestGetListByUserID(c *C) {
	userWorkHistory := &UserWorkHistory{}

	r, _, _ := userWorkHistory.GetListByUserID(1000)
	c.Check(r[0].ID, Equals, uint(1))
	c.Check(r[0].UserID, Equals, 1000)
}
