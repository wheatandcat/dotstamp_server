package controllersSound

import (
	"github.com/wheatandcat/dotstamp_server/controllers"
	"github.com/wheatandcat/dotstamp_server/models"
)

// LengthController コントローラ
type LengthController struct {
	controllers.BaseController
}

// LengthResponse レスポンス
type LengthResponse struct {
	Character int
}

// Get 1文字あたりの長さを取得する
func (c *LengthController) Get() {
	u := models.UserContributionSoundLength{}
	list, _, err := u.GetByTop(0, 1000)
	if err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, 0)
		return
	}

	s := 0
	l := 0
	for _, v := range list {
		if v.Second == 0 {
			continue
		}
		s += v.Second
		l += v.Length
	}

	c.Data["json"] = LengthResponse{
		Character: int(s / l),
	}
	c.ServeJSON()
}
