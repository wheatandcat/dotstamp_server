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
		"user_contribution_sound_details",
		"user_contribution_movies",
	})
}

func TestExec(t *testing.T) {
	r := MakeMovie(1)

	Convey("tasks/makeMovie/main.go\n", t, func() {
		Convey("MakeMovie", func() {
			So(r, ShouldEqual, nil)
		})
	})
}
