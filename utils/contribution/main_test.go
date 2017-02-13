package contributions

import (
	test "dotstamp_server/tests"

	. "gopkg.in/check.v1"
)

type TestMain struct {
	test.TSuite
}

func init() {
	var t test.Accessor = &TestMain{}
	t.SetTableNameList([]string{
		"user_masters",
		"user_contributions",
		"user_contribution_details",
		"user_contribution_tags",
	})

	var _ = Suite(t)
}

func (t *TestMain) TestAdd(c *C) {
	r, _ := Add(100, "abc", "{}")

	c.Check(r, Equals, uint(3))
}

func (t *TestMain) TestSave(c *C) {
	Save(1, 1, "abcdef")

	r := GetByUserContributionID(1)

	c.Check(r.Title, Equals, "abcdef")
}

func (t *TestMain) TestDeleteByID(c *C) {
	DeleteByID(1, 100)
	uc := GetByUserContributionID(1)
	ucd := GetDetailByUserContributionID(1)

	c.Check(uc.ID, Equals, uint(1))
	c.Check(ucd.ID, Equals, uint(1))

	DeleteByID(1, 1)
	uc = GetByUserContributionID(1)
	ucd = GetDetailByUserContributionID(1)

	c.Check(uc.ID, Equals, uint(0))
	c.Check(ucd.ID, Equals, uint(0))
}

func (t *TestMain) TestGetListByUserID(c *C) {
	r := GetListByUserID(1)

	c.Check(r[0].UserID, Equals, 1)
}

func (t *TestMain) TestGetByUserContributionID(c *C) {
	r := GetByUserContributionID(1)

	c.Check(r.Title, Equals, "test001")
}

func (t *TestMain) TestGetContributionByUserContributionID(c *C) {
	r, _ := GetContributionByUserContributionID(1)

	c.Check(r.ID, Equals, uint(1))
}

func (t *TestMain) TestGetByTop(c *C) {
	r, _ := GetByTop(0, 10)

	c.Check(r[0].ID, Equals, uint(2))
	c.Check(r[0].Tag[0].Name, Equals, "def")
}
