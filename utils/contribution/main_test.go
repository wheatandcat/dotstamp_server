package contributions

import (
	test "dotstamp_server/tests"
	"dotstamp_server/utils/follow"

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
	r, _ := Add(100, "abc", "{}", 1)

	c.Check(r, Equals, uint(3))
}

func (t *TestMain) TestSave(c *C) {
	Save(1, 1, "abcdef", 1)

	r, _ := GetByUserContributionID(1)

	c.Check(r.Title, Equals, "abcdef")
}

func (t *TestMain) TestDeleteByID(c *C) {
	DeleteByID(1, 100)
	uc, _ := GetByUserContributionID(1)
	ucd, _ := GetDetailByUserContributionID(1)

	c.Check(uc.ID, Equals, uint(1))
	c.Check(ucd.ID, Equals, uint(1))

	DeleteByID(1, 1)
	uc, _ = GetByUserContributionID(1)
	ucd, _ = GetDetailByUserContributionID(1)

	c.Check(uc.ID, Equals, uint(0))
	c.Check(ucd.ID, Equals, uint(0))
}

func (t *TestMain) TestGetListByUserID(c *C) {
	r, _ := GetListByUserID(1)

	c.Check(r[0].UserID, Equals, 1)
}

func (t *TestMain) TestGetByUserContributionID(c *C) {
	r, _ := GetByUserContributionID(1)

	c.Check(r.Title, Equals, "test001")
}

func (t *TestMain) TestGetContributionByUserContributionID(c *C) {
	r, _ := GetContributionByUserContributionID(1)

	c.Check(r.ID, Equals, uint(1))
}

func (t *TestMain) TestGetListByTop(c *C) {
	r, _ := GetListByTop(0, 10)

	c.Check(r[0].ID, Equals, uint(2))
	c.Check(len(r[0].Tag), Equals, 0)
}

func (t *TestMain) TestGetListBySearchValue(c *C) {
	s := []SearchValue{
		{
			UserContributionID: 1,
			Search:             "aaaaabbbbcccc",
			Order:              1,
		},
		{
			UserContributionID: 2,
			Search:             "xxxyyyzzz",
			Order:              0,
		},
	}

	r, _ := GetListBySearchValue(s)

	c.Check(r[1].ID, Equals, uint(1))
	c.Check(r[1].Tag[0].Name, Equals, "abc")
	c.Check(r[1].Search, Equals, "aaaaabbbbcccc")

	c.Check(r[0].ID, Equals, uint(2))
	c.Check(len(r[0].Tag), Equals, 0)
	c.Check(r[0].Search, Equals, "xxxyyyzzz")
}

func (t *TestMain) TestGetListByFollowOrderValue(c *C) {
	f := []follows.OrderValue{
		{
			UserContributionID: 1,
			Order:              1,
		},
		{
			UserContributionID: 2,
			Order:              0,
		},
	}

	r, _ := GetListByFollowOrderValue(f)

	c.Check(r[1].ID, Equals, uint(1))
	c.Check(r[1].Tag[0].Name, Equals, "abc")
}
