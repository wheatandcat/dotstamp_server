package controllersGoogle

import (
	"context"
	"dotstamp_server/controllers"
	"dotstamp_server/utils"
	"dotstamp_server/utils/oauth/google"
	"errors"

	"github.com/astaxie/beego"

	v2 "google.golang.org/api/oauth2/v2"
)

// CallbackController コールバックコントローラ
type CallbackController struct {
	controllers.BaseController
}

// CallbackRequest コールバックリクエスト
type CallbackRequest struct {
	Code  string `form:"code"`
	State int    `form:"state"`
}

// Get コールバックする
func (c *CallbackController) Get() {
	request := CallbackRequest{}
	if err := c.ParseForm(&request); err != nil {
		c.RedirectError(err, 0)
		return
	}

	config := google.GetConnect()
	context := context.Background()

	t, err := config.Exchange(context, request.Code)
	if err != nil {
		c.RedirectError(err, 0)
		return
	}

	if c.GetSession("googleOauthState") != request.State {
		c.RedirectError(errors.New("vaild state"), 0)
		return
	}

	if t.Valid() == false {
		c.RedirectError(errors.New("vaild token"), 0)
		return
	}

	s, err := v2.New(config.Client(context, t))
	if err != nil {
		c.RedirectError(err, 0)
		return
	}

	info, err := s.Tokeninfo().AccessToken(t.AccessToken).Context(context).Do()
	if err != nil {
		c.RedirectError(err, 0)
		return
	}

	e, err := utils.Encrypter([]byte(info.Email))
	if err != nil {
		c.RedirectError(err, 0)
		return
	}

	url := beego.AppConfig.String("topurl") + "/oauth/" + utils.Urlencode(e)
	c.Redirect(url, 302)
}
