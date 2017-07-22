package controllersTags

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/wheatandcat/dotstamp_server/routers"
	"github.com/wheatandcat/dotstamp_server/tests"

	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
)

func setUpPost() {
	test.Setup()
	test.SetupFixture([]string{
		"user_contributions",
		"user_contribution_details",
		"user_contribution_tags",
	})
}

func TestAddPost(t *testing.T) {
	setUpPost()

	json := `{
		"userContributionId":1,
		"name":"foo"
	}`

	r, err := http.NewRequest(
		"POST",
		"/api/tags/",
		bytes.NewBuffer([]byte(json)),
	)

	if err != nil {
		panic(err)
	}

	r.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	Convey("POST /tags/\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}
