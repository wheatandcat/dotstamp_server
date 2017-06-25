package main

import (
	_ "github.com/wheatandcat/dotstamp_server/routers"
	"github.com/wheatandcat/dotstamp_server/tests"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func init() {
	test.Setup()
	test.SetupFixture([]string{
		"user_contributions",
		"user_contribution_sound_details",
		"user_contribution_movies",
	})
}

func TestExec(t *testing.T) {
	test.CopyTestFile(1)

	r := MakeMovie(1)

	Convey("tasks/makeMovie/main.go\n", t, func() {
		Convey("MakeMovie", func() {
			So(r, ShouldEqual, nil)
		})
	})
}
