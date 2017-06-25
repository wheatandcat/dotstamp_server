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
		"user_contribution_sound_lengths",
	})
}

func TestContributionShow(t *testing.T) {
	Convey("tasks/moveTime/main.go\n", t, func() {
		Convey("ContributionShow", func() {
			r := contribution()
			So(r, ShouldEqual, nil)
		})
	})
}
