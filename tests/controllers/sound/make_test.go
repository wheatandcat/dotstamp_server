package controllersSound

import (
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/wheatandcat/dotstamp_server/routers"
	"github.com/wheatandcat/dotstamp_server/tests"

	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
)

func setUpMake() {
	test.Setup()
	test.SetupFixture([]string{
		"user_contributions",
		"user_contribution_details",
		"user_contribution_sounds",
		"user_contribution_sound_details",
		"user_contribution_movies",
	})
}

func TestMakePost(t *testing.T) {
	setUpMake()

	test.CopyTestFile(1)

	r, err := http.NewRequest(
		"POST",
		"/api/sounds/1/make/",
		nil,
	)

	if err != nil {
		panic(err)
	}

	r.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	Convey("/sounds/1/make/\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}
