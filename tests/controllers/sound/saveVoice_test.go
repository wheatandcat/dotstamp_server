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

func setUpSaveVoice() {
	test.Setup()
	test.SetupFixture([]string{
		"user_masters",
		"user_contributions",
		"user_contribution_sound_details",
	})
}

func TestSaveVoicePost(t *testing.T) {
	setUpSaveVoice()

	json := `{
		"id":1,
		"voiceType":2
	}`

	r, err := http.NewRequest(
		"PUT",
		"/api/sounds/voice/",
		bytes.NewBuffer([]byte(json)),
	)

	if err != nil {
		panic(err)
	}

	r.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	Convey("/sounds/voice/\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}
