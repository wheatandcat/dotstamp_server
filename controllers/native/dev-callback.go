package controllersNative

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/wheatandcat/dotstamp_server/controllers"
	"github.com/wheatandcat/dotstamp_server/utils/user"
)

// DevCallbackController コールバックコントローラ
type DevCallbackController struct {
	controllers.BaseController
}

// DevCallbackRequest コールバックリクエスト
type DevCallbackRequest struct {
	AccessToken string `form:"access_token"`
}

// Get コールバックする
func (c *DevCallbackController) Get() {
	request := CallbackRequest{}

	url := "https://graph.facebook.com/me?access_token=" + request.AccessToken + "&fields=email"
	r, _ := http.Get(url)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		c.RedirectError(err, 0)
		return
	}

	res := Response{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		c.RedirectError(err, 0)
		return
	}

	u, err := user.GetByEmail(res.Email)
	if err != nil {
		c.RedirectError(err, 0)
		return
	}

	if u.ID != 0 {
		c.SetSession("user_id", u.ID)
		c.Data["login"] = true
	} else {
		c.Data["login"] = false
	}

	c.Data["email"] = res.Email
	c.TplName = "dev_oauth.tpl"
}
