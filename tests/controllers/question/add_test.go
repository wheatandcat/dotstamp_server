package controllersQuestion

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

func init() {
	test.Setup()
	test.SetupFixture([]string{
		"log_questions",
	})
}

func TestAddPost(t *testing.T) {
	values := url.Values{}
	values.Set("body", "abc")
	values.Set("email", "abc@test.com")

	r, err := http.NewRequest(
		"POST",
		"/question/add/",
		strings.NewReader(values.Encode()),
	)

	if err != nil {
		panic(err)
	}

	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	Convey("/question/add/\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}
