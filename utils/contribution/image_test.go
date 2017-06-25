package contributions

import (
	"github.com/wheatandcat/dotstamp_server/tests"

	. "gopkg.in/check.v1"
)

type TestImage struct {
	test.TSuite
}

func init() {
	var t test.Accessor = &TestImage{}
	t.SetTableNameList([]string{
		"log_contribution_images",
	})

	var _ = Suite(t)
}

func (t *TestImage) TestGetImageIDAndAdd(c *C) {
	r, _ := GetImageIDAndAdd(1)

	c.Check(r, Equals, uint(2))
}
