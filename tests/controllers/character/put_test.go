package controllersCharacterImage

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

func setupPut() {
	test.Setup()
	test.SetupFixture([]string{
		"user_masters",
		"user_character_images",
	})
}

func TestSavePost(t *testing.T) {
	setupPut()

	values := url.Values{}
	values.Set("id", "1")
	values.Set("voiceType", "2")

	r, err := http.NewRequest(
		"PUT",
		"/api/characters/",
		strings.NewReader(values.Encode()),
	)

	if err != nil {
		panic(err)
	}

	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	Convey("PUT /characters/\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}
