package contributionSearch

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
		"user_contribution_details",
		"user_contribution_follows",
		"contribution_total_follows",
		"user_contribution_searches",
	})
}

func TestExec(t *testing.T) {
	r := ResetSearch()

	Convey("tasks/contributionSearch/main.go\n", t, func() {
		Convey("ResetSearch", func() {
			So(r, ShouldEqual, nil)
		})
	})

}
