package controllersContribution

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
}

func TestNewPost(t *testing.T) {
	values := url.Values{}
	values.Set("title", "テスト")
	values.Set("body", `[{"priority":0,"body":"あああ","character":{"Id":128,"Character_id":0,"Priority":0,"FileName":"2747b7c718564ba5f066f0523b03e17f6a496b06851333d2d59ab6d863225848.jpg","imageType":4},"directionType":1,"talkType":1,"edit":false},{"priority":1,"body":"あああ","character":{"Id":125,"Character_id":0,"Priority":0,"FileName":"0f8ef3377b30fc47f96b48247f463a726a802f62f3faa03d56403751d2f66c67.jpg","imageType":4},"directionType":1,"talkType":1,"edit":false},{"priority":2,"body":"あああ","character":{"Id":126,"Character_id":0,"Priority":0,"FileName":"65a699905c02619370bcf9207f5a477c3d67130ca71ec6f750e07fe8d510b084.jpg","imageType":4},"directionType":1,"talkType":1,"edit":false}]`)
	// models.ViewStatusPublic
	values.Set("viewStatus", "1")
	values.Set("tag", "abc defg　hijkl")

	r, err := http.NewRequest(
		"POST",
		"/api/contributions/new/",
		strings.NewReader(values.Encode()),
	)

	if err != nil {
		panic(err)
	}

	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	Convey("/contribution/new/\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}
