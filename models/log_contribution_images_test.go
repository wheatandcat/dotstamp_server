package models

import (
	"github.com/wheatandcat/dotstamp_server/tests"

	. "gopkg.in/check.v1"
)

type TestLogContributionImage struct {
	test.TSuite
}

func init() {
	var t test.Accessor = &TestLogContributionImage{}
	t.SetTableNameList([]string{
		"log_contribution_images",
	})

	var _ = Suite(t)
}

func (t *TestLogContributionImage) TestGetIDAndAdd(c *C) {
	u := &LogContributionImage{
		UserContributionID: 1,
	}

	r, _ := u.GetIDAndAdd()
	c.Check(r, Equals, uint(2))
}
