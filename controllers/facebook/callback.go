package controllersFacebook

import (
	"dotstamp_server/controllers"
	"dotstamp_server/utils/oauth/facebook"
	"dotstamp_server/utils/user"
	"errors"
	"net/url"

	"github.com/astaxie/beego"
	fb "github.com/huandu/facebook"
	"golang.org/x/oauth2"
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

	config := facebook.GetConnect()

	t, err := config.Exchange(oauth2.NoContext, request.Code)
	if err != nil {
		c.RedirectError(err, 0)
		return
	}

	if c.GetSession("facebookOauthState") != request.State {
		c.RedirectError(errors.New("vaild state"), 0)
		return
	}

	if t.Valid() == false {
		c.RedirectError(errors.New("vaild token"), 0)
		return
	}

	client := config.Client(oauth2.NoContext, t)
	session := &fb.Session{
		Version:    "v2.8",
		HttpClient: client,
	}

	res, err := session.Get("/me?fields=id,name,email", nil)
	if err != nil {
		c.RedirectError(err, 0)
		return
	}

	u, err := user.GetByEmail(res["email"].(string))
	if err != nil {
		c.RedirectError(err, 0)
		return
	}

	if u.ID != 0 {
		c.SetSession("user_id", u.ID)
		c.Redirect(beego.AppConfig.String("topurl"), 302)
		return
	}

	url := beego.AppConfig.String("topurl") + "oauth/?email=" + url.QueryEscape(res["email"].(string))
	c.Redirect(url, 302)
}
