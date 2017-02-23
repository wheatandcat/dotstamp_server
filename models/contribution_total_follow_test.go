package models

import (
	"dotstamp_server/tests"

	. "gopkg.in/check.v1"
)

type TestContributionTotalFollows struct {
	test.TSuite
}

func init() {
	var t test.Accessor = &TestContributionTotalFollows{}
	t.SetTableNameList([]string{
		"contribution_total_follows",
	})

	var _ = Suite(t)
}

func (t *TestContributionTotalFollows) TestAdd(c *C) {
	u := &ContributionTotalFollows{
		UserContributionID: 100,
		Count:              1,
	}

	u.Add()

	r, _, _ := u.GetListByUserContributionID([]int{100})

	c.Check(r[0].UserContributionID, Equals, 100)
	c.Check(r[0].Count, Equals, 1)
}

func (t *TestContributionTotalFollows) TestSave(c *C) {
	u := &ContributionTotalFollows{}
	user, _, _ := u.GetListByUserContributionID([]int{1})

	user[0].Count = 10
	user[0].Save()

	r, _, _ := u.GetListByUserContributionID([]int{1})

	c.Check(r[0].Count, Equals, 10)
}

func (t *TestContributionTotalFollows) TestGetListByUserContributionID(c *C) {
	u := &ContributionTotalFollows{}

	r, _, _ := u.GetListByUserContributionID([]int{1, 10})

	c.Check(r[0].UserContributionID, Equals, 1)
}

func (t *TestContributionTotalFollows) TestTruncate(c *C) {
	u := &ContributionTotalFollows{}

	u.Truncate()

	r, _, _ := u.GetListByUserContributionID([]int{1})

	c.Check(len(r), Equals, 0)
}
