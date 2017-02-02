package models

import (
	"dotstamp_server/tests"
	"dotstamp_server/utils"

	. "gopkg.in/check.v1"
)

type TestTmpWork struct {
	test.TSuite
}

func init() {
	var t test.Accessor = &TestTmpWork{}
	t.SetTableNameList([]string{
		"tmp_work",
	})

	var _ = Suite(t)
}

func (t *TestTmpWork) TestSave(c *C) {
	Released, _ := utils.StringToDate("2016-11-01")

	TmpWork := &TmpWork{
		UserID:     1000,
		CategoryID: 1,
		Name:       "test",
		AuthorID:   1,
		CountryID:  1,
		Released:   Released,
	}

	TmpWork.Add()

	r := TmpWork.GetFindAll()

	c.Check(r[0].ID, Equals, 1)
	c.Check(r[1].ID, Equals, 2)
}
