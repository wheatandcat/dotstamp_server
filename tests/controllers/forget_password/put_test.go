package controllersForgetPassword

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

func setUpPut() {
	test.Setup()

	test.SetupFixture([]string{
		"user_masters",
		"user_forget_passwords",
	})
}

func TestSavePost(t *testing.T) {
	setUpPut()

	json := `{
		"email":"vHWexIhSOGxjAIjz.t.3o8DN2_cv4ozt3TOb",
		"keyword":"gEyG9YZUN31mLKbA18GFpxVc_h8fGFdtn2dNU9SwqG7uakosOKeNU0we4Ahpvishbf4-",
		"password":"testtest"
	}`

	r, err := http.NewRequest(
		"PUT",
		"/api/forget_password/",
		bytes.NewBuffer([]byte(json)),
	)

	if err != nil {
		panic(err)
	}

	r.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	Convey("PUT /user/forget_password/\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}
