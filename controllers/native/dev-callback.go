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

// DevResponse レスポンス
type DevResponse struct {
	Login bool   `json:"login"`
	Email string `json:"email"`
}

// Get コールバックする
func (c *DevCallbackController) Get() {
	request := DevCallbackRequest{}
	if err := c.ParseForm(&request); err != nil {
		c.RedirectError(err, 0)
		return
	}

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

	login := false

	if u.ID != 0 {
		c.SetSession("user_id", u.ID)
		login = true
	}

	c.Data["json"] = DevResponse{
		Email: res.Email,
		Login: login,
	}

	c.ServeJSON()
}
