package csvModels

import (
	"dotstamp_server/tests"

	. "gopkg.in/check.v1"
)

type TestWork struct {
	test.TSuite
}

func init() {
	var t test.Accessor = &TestWork{}
	t.SetTableNameList([]string{})

	var _ = Suite(t)
}
