package utils

import (
	"dotstamp_server/tests"

	. "gopkg.in/check.v1"
)

type TestLog struct {
	test.TSuite
}

func init() {
	var t test.Accessor = &TestLog{}
	t.SetTableNameList([]string{})

	var _ = Suite(t)
}

func (t *TestLog) LogFile(c *C) {
	logfile, r := LogFile("test")
	defer logfile.Close()

	c.Check(r, Equals, nil)
}
