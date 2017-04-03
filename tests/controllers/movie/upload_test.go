package controllersMovie

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

	values := url.Values{}
	values.Set("userContributionId", "1")

	r, err := http.NewRequest(
		"POST",
		"/movie/upload/",
		strings.NewReader(values.Encode()),
	)

	if err != nil {
		panic(err)
	}

	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	Convey("/movie/upload/\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}
