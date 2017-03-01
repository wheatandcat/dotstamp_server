package controllersLoginGoogle

import (
	"dotstamp_server/controllers"
	"fmt"
	"log"
	"reflect"

	"golang.org/x/oauth2"
	v2 "google.golang.org/api/oauth2/v2"
	"google.golang.org/appengine"
)

// CallbackController 新規登録コントローラ
type CallbackController struct {
	controllers.BaseController
}

// Get 認証の受け口
func (c *CallbackController) Get() {
	client_id := `951607719169-cuoivge74lju8p5atm9jn77m6822hvo9.apps.googleusercontent.com`
	client_secret := `L-fHGOut1vzmB_EWfBgecM0y`

	conf := &oauth2.Config{
		ClientID:     client_id,
		ClientSecret: client_secret,
		RedirectURL:  `http://local.org:8080/login/callback`,
		Scopes:       []string{" https://www.googleapis.com/auth/drive.readonly"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://accounts.google.com/o/oauth2/auth",
			TokenURL: "https://accounts.google.com/o/oauth2/token",
		},
	}

	log.Println("001-----")
	test := c.Ctx.Request
	log.Println(test)
	fmt.Println(reflect.TypeOf(test))
	context := appengine.NewContext(test)
	log.Println("002-----")
	code := c.GetString("code")
	log.Println(code)

	// 認証コードからtokenを取得します
	tok, err := conf.Exchange(context, code)
	if err != nil {
		panic(err)
	}

	// tokenが正しいことを確認します
	if tok.Valid() == false {
		log.Println("NG")
	}

	// oauth2 clinet serviceを取得します
	// 特にuserの情報が必要ない場合はスルーです
	service, err := v2.New(conf.Client(context, tok))
	if err != nil {
		panic(err)
	}
	log.Println(service)

	// token情報を取得します
	// ここにEmailやUser IDなどが入っています
	// 特にuserの情報が必要ない場合はスルーです
	tokenInfo, err := service.Tokeninfo().AccessToken(tok.AccessToken).Context(context).Do()
	if err != nil {
		panic(err)
	}

	log.Println(tokenInfo)

	c.Data["json"] = "aaa"
	c.ServeJSON()
}
