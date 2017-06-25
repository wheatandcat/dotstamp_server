package contributions

import (
	"github.com/wheatandcat/dotstamp_server/tests"

	. "gopkg.in/check.v1"
)

type TestUpload struct {
	test.TSuite
}

func init() {
	var t test.Accessor = &TestUpload{}
	t.SetTableNameList([]string{
		"user_contribution_uploads",
	})

	var _ = Suite(t)
}

func (t *TestUpload) TestGetByUserContributionID(c *C) {
	r, _ := GetUploadByUserContributionID(1)

	c.Check(r.UserContributionID, Equals, 1)
}
