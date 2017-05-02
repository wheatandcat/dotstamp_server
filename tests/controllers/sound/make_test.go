package controllersSound

import (
	_ "dotstamp_server/routers"
	"dotstamp_server/tests"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

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

	values := url.Values{}
	values.Set("userContributionId", "1")

	r, err := http.NewRequest(
		"POST",
		"/api/sound/make/",
		strings.NewReader(values.Encode()),
	)

	if err != nil {
		panic(err)
	}

	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	Convey("/sound/make/\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}
