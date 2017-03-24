package movie

import (
	"dotstamp_server/tests"

	. "gopkg.in/check.v1"
)

type TestYoutube struct {
	test.TSuite
}

func init() {
	var t test.Accessor = &TestYoutube{}
	t.SetTableNameList([]string{})

	var _ = Suite(t)
}

func (t *TestYoutube) TestUploadToYoutube(c *C) {

	GetConnect()
}
