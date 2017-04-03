package models

import (
	"dotstamp_server/tests"

	. "gopkg.in/check.v1"
)

type TestUserContributionMovie struct {
	test.TSuite
}

func init() {
	var t test.Accessor = &TestUserContributionMovie{}
	t.SetTableNameList([]string{
		"user_contribution_movies",
	})

	var _ = Suite(t)
}

func (t *TestUserContributionMovie) TestAdd(c *C) {
	u := &UserContributionMovie{
		UserContributionID: 100,
		MovieStatus:        1,
		MovieType:          1,
		MovieID:            "abc",
	}

	r := u.Add()

	c.Check(r, Equals, nil)
}

func (t *TestUserContributionMovie) TestSave(c *C) {
	u := &UserContributionMovie{}
	uc, _, _ := u.GetByUserContributionID(1, 1)

	uc.MovieStatus = 2

	r := uc.Save()

	c.Check(r, Equals, nil)
}

func (t *TestUserContributionMovie) TestGetListByUserContributionIDList(c *C) {
	u := &UserContributionMovie{}
	r, _, _ := u.GetListByUserContributionIDList([]int{1}, 1)

	c.Check(r[0].ID, Equals, uint(1))
}

func (t *TestUserContributionMovie) TestGetListByMovieStatusPublic(c *C) {
	u := &UserContributionMovie{}
	r, _, _ := u.GetListByMovieStatusPublic()

	c.Check(r[0].MovieStatus, Equals, StatusPublic)
}
