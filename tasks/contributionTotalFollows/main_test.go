package main

import (
	_ "dotstamp_server/routers"
	"dotstamp_server/tests"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func init() {
	test.Setup()
	test.SetupFixture([]string{
		"user_contributions",
		"user_contribution_follows",
		"contribution_total_follows",
	})
}

func TestExec(t *testing.T) {
	r := Exec()

	Convey("tasks/contributionTotalFollows/main.go\n", t, func() {
		Convey("Check Error", func() {
			So(r, ShouldEqual, nil)
		})
	})
}
