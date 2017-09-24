package controllersNative

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/wheatandcat/dotstamp_server/controllers"
	"github.com/wheatandcat/dotstamp_server/utils/user"
)

// CallbackController コールバックコントローラ
type CallbackController struct {
	controllers.BaseController
}

// CallbackRequest コールバックリクエスト
type CallbackRequest struct {
	AccessToken string `form:"access_token"`
}

// Response レスポンス
type Response struct {
	Email string `json:"email"`
}

// Get コールバックする
func (c *CallbackController) Get() {
	request := CallbackRequest{}

	url := "https://graph.facebook.com/me?access_token=" + request.AccessToken
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
	c.TplName = "oauth.tpl"
}
