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

func setUpSaveVoiceList() {
	test.Setup()
	test.SetupFixture([]string{
		"user_masters",
		"user_contributions",
		"user_contribution_sound_details",
	})
}

func TestSaveVoiceListPost(t *testing.T) {
	setUpSaveVoiceList()

	json := `{
		"voiceType":2
	}`

	r, err := http.NewRequest(
		"PUT",
		"/api/sounds/1/voice/all/",
		bytes.NewBuffer([]byte(json)),
	)

	if err != nil {
		panic(err)
	}

	r.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	Convey("/sounds/1/voice/all/\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}
