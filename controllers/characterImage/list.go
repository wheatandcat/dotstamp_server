package controllersCharacterImage

import (
	"github.com/wheatandcat/dotstamp_server/controllers"
	"github.com/wheatandcat/dotstamp_server/utils/character"
)

// ListController listコントローラ
type ListController struct {
	controllers.BaseController
}

// ListResponse リストレスポンス
type ListResponse struct {
	Image []characters.Image
}

// Post 一覧を取得する
func (c *ListController) Post() {
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

	c.Data["json"] = ListResponse{
		Image: image,
	}

	c.ServeJSON()
}
