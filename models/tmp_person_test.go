package models

import (
	"dotstamp_server/tests"

	. "gopkg.in/check.v1"
)

type TestTmpPerson struct {
	test.TSuite
}

func init() {
	var t test.Accessor = &TestTmpPerson{}
	t.SetTableNameList([]string{
		"tmp_person",
	})

	var _ = Suite(t)
}

func (t *TestTmpPerson) TestSave(c *C) {
	tmpPerson := &TmpPerson{
		UserID: 1000,
		Name:   "test",
	}

	tmpPerson.Add()

	r := tmpPerson.GetFindAll()

	c.Check(r[0].ID, Equals, 1)
	c.Check(r[1].ID, Equals, 2)
}
