package controllersForgetPassword

import (
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/wheatandcat/dotstamp_server/routers"
	"github.com/wheatandcat/dotstamp_server/tests"

	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
)

func setUpCheck() {
	test.Setup()

	test.SetupFixture([]string{
		"user_masters",
		"user_forget_passwords",
	})
}

func TestCheck(t *testing.T) {
	setUpCheck()

	r, err := http.NewRequest(
		"GET",
		"/api/forget_password/check/rH.Zw7xSMXghDIT_/uXiJ05lV/",
		nil,
	)

	if err != nil {
		panic(err)
	}

	r.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	Convey("/forget_password/check/\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}
