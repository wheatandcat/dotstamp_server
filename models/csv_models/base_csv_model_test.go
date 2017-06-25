package csvModels

import (
	"github.com/wheatandcat/dotstamp_server/tests"

	. "gopkg.in/check.v1"
)

type TestBaseCsvModel struct {
	test.TSuite
}

func init() {
	var t test.Accessor = &TestBaseCsvModel{}
	t.SetTableNameList([]string{})

	var _ = Suite(t)
}

func (t *TestBaseCsvModel) TestGetMampAll(c *C) {
	r, _ := GetMampAll("contribution_sound_body_replace.csv")

	c.Check(r[0]["id"], Equals, "1")
}

func (t *TestBaseCsvModel) TestGetAll(c *C) {
	r := []ContributionSoundBodyReplace{}

	GetAll("contribution_sound_body_replace.csv", &r)

	c.Check(r[0].ID, Equals, "1")
}
