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

func (t *TestMain) TestAdd(c *C) {
	Add(100, "abc")

	r, _ := GetListByUserContributionID(100)

	c.Check(r[0].Name, Equals, "abc")
}

func (t *TestMain) TestSave(c *C) {
	id := 1
	n := "bbb"

	Save(id, n)

	u := models.UserContributionTag{}
	r, _, _ := u.GetByID(id)

	c.Check(r.Name, Equals, "bbb")
}

func (t *TestMain) TestDeleteByIDAndUserContributionID(c *C) {
	DeleteByIDAndUserContributionID(1, 10)
	u := models.UserContributionTag{}
	r, _, _ := u.GetByID(1)

	c.Check(r.ID, Not(Equals), uint(0))

	DeleteByIDAndUserContributionID(1, 1)
	r, _, _ = u.GetByID(1)

	c.Check(r.ID, Equals, uint(0))
}

func (t *TestMain) TestAddList(c *C) {
	uID := 3
	n := "abc defg　hijkl"
	AddList(uID, n)

	u := models.UserContributionTag{}

	r, _, _ := u.GetListByUserContributionID(uID)

	c.Check(r[0].Name, Equals, "abc")
	c.Check(r[1].Name, Equals, "defg")
	c.Check(r[2].Name, Equals, "hijkl")

	err := AddList(uID, "aa aa bb cc")
	c.Check(err, Equals, nil)

	err = AddList(uID, "zz xx yy ww vv uu tt ss mm nn")
	c.Check(err, Equals, nil)

	err = AddList(uID, "zz xx yy ww vv uu tt ss mm nn ll")
	c.Check(err, Not(Equals), nil)
}

func (t *TestMain) TestGetListByUserContributionID(c *C) {
	r, _ := GetListByUserContributionID(1)

	c.Check(r[0].Name, Equals, "abc")

	r, _ = GetListByUserContributionID(100)
	c.Check(len(r), Equals, 0)
}

func (t *TestMain) TestGetMapByUserContributionIDList(c *C) {
	r, _ := GetMapByUserContributionIDList([]int{1, 10})

	c.Check(r[1][0].Name, Equals, "abc")

}

func (t *TestMain) TestGetTagNameJoin(c *C) {
	r, _ := GetTagNameJoin(1)

	c.Check(r, Equals, "abc")
}
