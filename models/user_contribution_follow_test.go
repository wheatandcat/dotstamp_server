package models

import (
	"dotstamp_server/tests"

	. "gopkg.in/check.v1"
)

type TestUserContributionFollow struct {
	test.TSuite
}

func init() {
	var t test.Accessor = &TestUserContributionFollow{}
	t.SetTableNameList([]string{
		"user_contribution_follow",
	})

	var _ = Suite(t)
}
