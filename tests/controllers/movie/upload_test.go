package controllersMovie

import (
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/wheatandcat/dotstamp_server/routers"
	"github.com/wheatandcat/dotstamp_server/tests"

	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
)

func setUpUpload() {
	test.Setup()
	test.SetupFixture([]string{
		"user_contributions",
		"user_contribution_movies",
		"user_contribution_uploads",
	})
}

func TestUploadPost(t *testing.T) {
	setUpUpload()

	test.CopyTestFile(1)

	r, err := http.NewRequest(
		"POST",
		"/api/movies/1/upload/",
		nil,
	)

	if err != nil {
		panic(err)
	}

	r.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	Convey("/movies/1/upload/\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}
