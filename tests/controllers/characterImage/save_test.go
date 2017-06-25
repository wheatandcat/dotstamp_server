package controllersCharacterImage

import (
	_ "github.com/wheatandcat/dotstamp_server/routers"
	"github.com/wheatandcat/dotstamp_server/tests"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
)

func setupSave() {
	test.Setup()
	test.SetupFixture([]string{
		"user_masters",
		"user_character_images",
	})
}

func TestSavePost(t *testing.T) {
	setupSave()

	values := url.Values{}
	values.Set("id", "1")
	values.Set("voiceType", "2")

	r, err := http.NewRequest(
		"POST",
		"/api/characterImage/save/",
		strings.NewReader(values.Encode()),
	)

	if err != nil {
		panic(err)
	}

	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	Convey("/characterImage/save/\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}
