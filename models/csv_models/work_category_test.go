package csvModels

import (
	"dotstamp_server/tests"

	. "gopkg.in/check.v1"
)

type TestWorkCategory struct {
	test.TSuite
}

func init() {
	var t test.Accessor = &TestWorkCategory{}
	t.SetTableNameList([]string{})

	var _ = Suite(t)
}
