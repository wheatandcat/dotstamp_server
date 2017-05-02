package controllersForgetPassword

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
		"user_masters",
		"user_forget_passwords",
	})
}

func TestCheckPost(t *testing.T) {

	r, err := http.NewRequest(
		"POST",
		"/api/user/forget_password/check/rH.Zw7xSMXghDIT_/uXiJ05lV/",
		nil,
	)

	if err != nil {
		panic(err)
	}

	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	Convey("/user/forget_password/check/\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}
