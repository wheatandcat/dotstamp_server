package follows

import (
	"dotstamp_server/tests"

	. "gopkg.in/check.v1"
)

type TestTotal struct {
	test.TSuite
}

func init() {
	var t test.Accessor = &TestTotal{}
	t.SetTableNameList([]string{
		"user_contribution_follows",
		"contribution_total_follows",
	})

	var _ = Suite(t)
}

func (t *TestTotal) TestAddTotal(c *C) {
	AddTotal(2, 10)
}

func (t *TestTotal) TestGetTotalListByUserContributionIDList(c *C) {
	r, _ := GetTotalListByUserContributionIDList([]int{1})

	c.Check(r[0].UserContributionID, Equals, 1)
}

func (t *TestTotal) TestGetTotalMapByUserContributionIDList(c *C) {
	r, _ := GetTotalMapByUserContributionIDList([]int{1})

	c.Check(r[1], Equals, 1)
}
