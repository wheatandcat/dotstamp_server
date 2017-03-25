package contributions

import (
	"dotstamp_server/tests"

	. "gopkg.in/check.v1"
)

type TestMovie struct {
	test.TSuite
}

func init() {
	var t test.Accessor = &TestMovie{}
	t.SetTableNameList([]string{
		"user_contributions",
		"user_contribution_movie",
	})

	var _ = Suite(t)
}

func (t *TestSound) TestAddMovie(c *C) {
	r := AddMovie(100, "abc", 1, 1)

	c.Check(r, Equals, nil)
}

func (t *TestSound) TestGetMovie(c *C) {
	r, _ := GetMovie(1, 1)

	c.Check(r.UserContributionID, Equals, 1)
}

func (t *TestSound) TestAddOrSaveMovie(c *C) {
	r := AddOrSaveMovie(100, "abc", 1, 1)

	c.Check(r, Equals, nil)

	r = AddOrSaveMovie(1, "abc", 1, 1)

	c.Check(r, Equals, nil)
}

func (t *TestSound) TestExistsMovie(c *C) {
	r := ExistsMovie(1)

	c.Check(r, Equals, true)
}
