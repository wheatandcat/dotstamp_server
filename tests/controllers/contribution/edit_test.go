package controllersContribution

import (
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/wheatandcat/dotstamp_server/routers"
	"github.com/wheatandcat/dotstamp_server/tests"

	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
)

func setupEdit() {
	test.Setup()
	test.SetupFixture([]string{
		"user_masters",
		"user_contributions",
		"user_contribution_details",
		"user_contribution_tags",
		"log_user_contributions",
	})
}

func TestEdit(t *testing.T) {
	setupEdit()

	r, err := http.NewRequest(
		"GET",
		"/api/contributions/edit/1",
		nil,
	)
	if err != nil {
		panic(err)
	}

	r.Header.Set("Content-Type", " application/json")

	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	Convey("GET /contribution/edit/1\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}
