package controllersSound

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

	json := `{
			"id":1,
			"body":"あいうえお"
	}`

	r, err := http.NewRequest(
		"PUT",
		"/api/sounds/body/",
		bytes.NewBuffer([]byte(json)),
	)

	if err != nil {
		panic(err)
	}

	r.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	Convey("/sounds/body/\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}
