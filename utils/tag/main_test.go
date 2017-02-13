package tags

import (
	"dotstamp_server/models"
	"dotstamp_server/tests"

	. "gopkg.in/check.v1"
)

type TestMain struct {
	test.TSuite
}

func init() {
	var t test.Accessor = &TestMain{}
	t.SetTableNameList([]string{
		"user_contribution_tags",
	})

	var _ = Suite(t)
}

func (t *TestMain) TestSave(c *C) {
	id := 1
	n := "bbb"

	Save(id, n)

	u := models.UserContributionTag{}
	r := u.GetByID(id)

	c.Check(r.Name, Equals, "bbb")
}

func (t *TestMain) TestDeleteByID(c *C) {
	DeleteByID(1)
	u := models.UserContributionTag{}
	r := u.GetByID(1)

	c.Check(r.ID, Equals, uint(0))
}

func (t *TestMain) TestAddList(c *C) {
	uID := 3
	n := "abc defg　hijkl"
	AddList(uID, n)

	u := models.UserContributionTag{}

	r := u.GetListByUserContributionID(uID)

	c.Check(r[0].Name, Equals, "abc")
	c.Check(r[1].Name, Equals, "defg")
	c.Check(r[2].Name, Equals, "hijkl")
}

func (t *TestMain) TestGetListByUserContributionID(c *C) {
	r, _ := GetListByUserContributionID(1)

	c.Check(r[0].Name, Equals, "abc")
}

func (t *TestMain) TestGetMapByUserContributionIDList(c *C) {
	r, _ := GetMapByUserContributionIDList([]int{1})

	c.Check(r[1][0].Name, Equals, "abc")
}
