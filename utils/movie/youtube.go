package movie

import (
	"net/http"

	"github.com/astaxie/beego"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	demoFunc  = make(map[string]func(*http.Client, []string))
	demoScope = make(map[string]string)
)

// GetConnect 接続を取得する
func GetConnect() *oauth2.Config {
	config := &oauth2.Config{
		ClientID:     beego.AppConfig.String("youtubeClientID"),
		ClientSecret: beego.AppConfig.String("youtubeClientSecret"),
		Endpoint:     google.Endpoint,
		Scopes:       []string{"https://www.googleapis.com/auth/youtube.upload"},
		RedirectURL:  beego.AppConfig.String("callBackUrl") + "api/movie/callback",
	}

	return config
}
