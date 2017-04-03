package contributions

import (
	"dotstamp_server/models"
	"dotstamp_server/tests"
	"dotstamp_server/utils"
	"time"

	. "gopkg.in/check.v1"
)

type TestMovie struct {
	test.TSuite
}

func init() {
	var t test.Accessor = &TestMovie{}
	t.SetTableNameList([]string{
		"user_contributions",
		"user_contribution_movies",
	})

	var _ = Suite(t)
}

func (t *TestMovie) TestAddMovie(c *C) {
	r := AddMovie(100, "abc", 1, 1)

	c.Check(r, Equals, nil)
}

func (t *TestMovie) TestGetMovie(c *C) {
	r, _ := GetMovie(1, 1)

	c.Check(r.UserContributionID, Equals, 1)
}

func (t *TestMovie) TestAddOrSaveMovie(c *C) {
	r := AddOrSaveMovie(101, "abc", 1, 1)

	c.Check(r, Equals, nil)

	r = AddOrSaveMovie(1, "abc", 1, 1)

	c.Check(r, Equals, nil)
}

func (t *TestMovie) TestExistsMovie(c *C) {
	r := ExistsMovie(1)

	c.Check(r, Equals, true)
}

func (t *TestMovie) TestGetMovieListByMovieStatusPublic(c *C) {
	r, _ := GetMovieListByMovieStatusPublic()

	c.Check(r[0].MovieStatus, Equals, models.StatusPublic)
}

func (t *TestMovie) TestGetMovieListBySpecifiedDays(c *C) {
	loc, _ := time.LoadLocation("Asia/Tokyo")
	utils.SetNow(time.Date(2015, 1, 3, 11, 00, 0, 0, loc))

	u, _ := GetMovieListByMovieStatusPublic()

	r := GetMovieListBySpecifiedDays(u, 2)

	c.Check(r[0].MovieStatus, Equals, models.StatusPublic)
}
