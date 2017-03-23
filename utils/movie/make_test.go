package movie

import (
	"dotstamp_server/tests"

	. "gopkg.in/check.v1"
)

type TestMain struct {
	test.TSuite
}

func init() {
	var t test.Accessor = &TestMain{}
	t.SetTableNameList([]string{})

	var _ = Suite(t)
}

func (t *TestMain) TestMake(c *C) {
	r := Make("0")

	c.Check(r, Equals, nil)
}

func (t *TestMain) TestUploadYoutube(c *C) {
	u := Upload{
		UserContributionID: "1",
		Title:              "abc",
		Description:        "abc",
		CategoryID:         "22",
		VideoStatus:        "public",
	}

	r := UploadYoutube(u)

	c.Check(r, Equals, nil)
}
