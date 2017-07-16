package controllersLogin

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	_ "github.com/wheatandcat/dotstamp_server/routers"
	"github.com/wheatandcat/dotstamp_server/tests"

	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
)

func init() {
	test.Setup()

	test.SetupFixture([]string{
		"user_masters",
	})
}

func TestNewPost(t *testing.T) {
	values := url.Values{}
	values.Set("email", "test_xyz@test.com")
	values.Set("password", "testtest")

	r, err := http.NewRequest(
		"POST",
		"/api/users/new/",
		strings.NewReader(values.Encode()),
	)

	if err != nil {
		panic(err)
	}

	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	Convey("/users/new/\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}
