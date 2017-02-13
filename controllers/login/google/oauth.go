package controllersLoginGoogle

import (
	"dotstamp_server/controllers"
	"log"

	"golang.org/x/oauth2"
)

// OauthController 新規登録コントローラ
type OauthController struct {
	controllers.BaseController
}

// Get 新規ログイン
func (new *OauthController) Get() {
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

	authURL := conf.AuthCodeURL("state-token")
	log.Println(authURL)

	new.Redirect(authURL, 302)
}
