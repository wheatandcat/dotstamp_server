package controllersGoogle

import (
	"context"
	"github.com/wheatandcat/dotstamp_server/controllers"
	"github.com/wheatandcat/dotstamp_server/utils/oauth/google"
	"github.com/wheatandcat/dotstamp_server/utils/user"
	"errors"
	"net/url"

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
	State string `form:"state"`
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

	u, err := user.GetByEmail(info.Email)
	if err != nil {
		c.RedirectError(err, 0)
		return
	}

	if u.ID != 0 {
		c.SetSession("user_id", u.ID)
		c.Redirect(beego.AppConfig.String("topurl"), 302)
		return
	}

	url := beego.AppConfig.String("topurl") + "oauth/?email=" + url.QueryEscape(info.Email)
	c.Redirect(url, 302)
}
