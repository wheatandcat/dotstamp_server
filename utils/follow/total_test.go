package follows

import (
	"github.com/wheatandcat/dotstamp_server/tests"

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

func (t *TestTotal) TestTruncateTotal(c *C) {
	TruncateTotal()
}

func (t *TestTotal) TestAddTotal(c *C) {
	AddTotal(2, 10)
}

func (t *TestTotal) TestAddTotalMap(c *C) {
	TruncateTotal()

	m := map[int]int{}
	m[1] = 10
	m[2] = 3

	AddTotalMap(m)

	r, _ := GetTotalListByUserContributionIDList([]int{1, 2})

	c.Check(r[1].UserContributionID, Equals, 1)
	c.Check(r[1].Count, Equals, 10)

	c.Check(r[0].UserContributionID, Equals, 2)
	c.Check(r[0].Count, Equals, 3)
}

func (t *TestTotal) TestGetTotalListByUserContributionIDList(c *C) {
	r, _ := GetTotalListByUserContributionIDList([]int{1})

	c.Check(r[0].UserContributionID, Equals, 1)
}

func (t *TestTotal) TestGetTotalMapByUserContributionIDList(c *C) {
	r, _ := GetTotalMapByUserContributionIDList([]int{1})

	c.Check(r[1], Equals, 1)
}
