package contributions

import (
	"dotstamp_server/tests"

	. "gopkg.in/check.v1"
)

type TestUserContributionSearch struct {
	test.TSuite
}

func init() {
	var t test.Accessor = &TestUserContributionSearch{}
	t.SetTableNameList([]string{
		"user_contribution_searches",
	})

	var _ = Suite(t)
}

func (t *TestUserContributionSearch) TestAddSearch(c *C) {
	AddSearch(10, "aaaaaa")
	r, _ := GetSearchByUserContributionID(10)

	c.Check(r.UserContributionID, Equals, 10)
}

func (t *TestUserContributionSearch) TestGetSearchByUserContributionID(c *C) {
	r, _ := GetSearchByUserContributionID(1)

	c.Check(r.UserContributionID, Equals, 1)
}
