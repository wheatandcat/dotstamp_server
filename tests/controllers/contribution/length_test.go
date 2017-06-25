package controllersContribution

import (
	_ "dotstamp_server/routers"
	"dotstamp_server/tests"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
)

func init() {
	test.Setup()
	test.SetupFixture([]string{
		"user_contribution_sound_lengths",
	})
}

func TestLengthGet(t *testing.T) {
	r, err := http.NewRequest(
		"GET",
		"/api/sound/length/",
		nil,
	)
	if err != nil {
		panic(err)
	}

	r.Header.Set("Content-Type", " application/json")

	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	Convey("/contribution/list/\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}
