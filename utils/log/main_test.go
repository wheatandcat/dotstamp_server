package logs

import (
	"github.com/wheatandcat/dotstamp_server/tests"

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

func (t *TestMain) TestErr(c *C) {
	r := Err("test", 1)

	c.Check(r, Equals, nil)
}

func (t *TestMain) TestBatch(c *C) {
	r := Batch("abc", "calcFollow")

	c.Check(r, Equals, nil)
}

func (t *TestMain) TestLogFile(c *C) {
	logfile, r := LogFile("test")
	defer logfile.Close()

	c.Check(r, Equals, nil)
}

func (t *TestMain) TestRemoveLogFile(c *C) {
	RemoveLogFile("test")
}
