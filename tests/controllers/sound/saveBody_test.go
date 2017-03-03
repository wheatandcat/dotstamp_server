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

func setUpSaveBody() {
	test.Setup()
	test.SetupFixture([]string{
		"user_masters",
		"user_contributions",
		"user_contribution_sound_details",
	})
}

func TestSaveBodyPost(t *testing.T) {
	setUpSaveBody()

	values := url.Values{}
	values.Set("id", "1")
	values.Set("body", "あいうえお")

	r, err := http.NewRequest(
		"POST",
		"/sound/saveBody/",
		strings.NewReader(values.Encode()),
	)

	if err != nil {
		panic(err)
	}

	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	Convey("/sound/SaveBody/\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}