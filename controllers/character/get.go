package controllersCharacter

import (
	"github.com/wheatandcat/dotstamp_server/controllers"
	"github.com/wheatandcat/dotstamp_server/utils/character"
)

// GettResponse レスポンス
type GettResponse struct {
	Image []characters.Image `json:"images"`
}

// Get 一覧を取得する
func (c *MainController) Get() {
	userID := c.GetUserID()
	if !c.IsNoLogin(userID) {
		c.ServerLoginNotFound()
		return
	}

	image, err := characters.GetImageListByUserID(userID)
	if err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	c.Data["json"] = GettResponse{
		Image: image,
	}

	c.ServeJSON()
}
