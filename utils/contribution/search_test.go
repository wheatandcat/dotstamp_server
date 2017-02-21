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

func (t *TestUserContributionSearch) TestGetSearchWordBody(c *C) {
	s := `[{"priority":0,"body":"あああ","character":{"Id":128,"Character_id":0,"Priority":0,"imageType":4},"directionType":1,"talkType":1,"edit":false},{"priority":1,"body":"あああ","character":{"Id":125,"Character_id":0,"Priority":0,"imageType":4},"directionType":1,"talkType":1,"edit":false},{"priority":2,"body":"いいいい","character":{"Id":126,"Character_id":0,"Priority":0,"imageType":4},"directionType":1,"talkType":1,"edit":false}]`
	r, _ := GetSearchWordBody(s)

	c.Check(r, Equals, "ああああああいいいい")
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

func (t *TestUserContributionSearch) TestAddOrSaveSearch(c *C) {
	AddOrSaveSearch(1, "ああああああいいいい")
	r, _ := GetSearchByUserContributionID(1)

	c.Check(r.Search, Equals, "ああああああいいいい")

	AddOrSaveSearch(10, "ああああああいいいい")
	r, _ = GetSearchByUserContributionID(10)

	c.Check(r.Search, Equals, "ああああああいいいい")
}

func (t *TestUserContributionSearch) TestDeleteSearchByUserContributionID(c *C) {
	DeleteSearchByUserContributionID(1)

	r, _ := GetSearchByUserContributionID(1)
	c.Check(r.ID, Equals, uint(0))

	DeleteSearchByUserContributionID(10)

	r, _ = GetSearchByUserContributionID(10)
	c.Check(r.ID, Equals, uint(0))
}

func (t *TestUserContributionSearch) TestGetSearchValueListBySearch(c *C) {
	r, _ := GetSearchValueListBySearch("a", "ID desc", 10, 0)

	c.Check(r[0].UserContributionID, Equals, 3)
	c.Check(r[0].Order, Equals, 0)

	c.Check(r[1].UserContributionID, Equals, 1)
	c.Check(r[1].Order, Equals, 1)

	r, _ = GetSearchValueListBySearch("a", "ID desc", 10, 10)

	c.Check(len(r), Equals, 0)
}