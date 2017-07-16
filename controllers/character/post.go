package controllersCharacter

import (
	"github.com/wheatandcat/dotstamp_server/controllers"
	"github.com/wheatandcat/dotstamp_server/utils/character"
)

// PostResponse アップロードレスポンス
type PostResponse struct {
	Image []characters.Image `json:"images"`
}

// Post 画像アップロード
func (c *MainController) Post() {
	userID := c.GetUserID()
	if !c.IsNoLogin(userID) {
		c.ServerLoginNotFound()
		return
	}

	id, err := characters.AddImage(userID, 0, 0)
	if err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	filePath := characters.GetImageName(id)

	var code int
	if code, err = c.SetImageFileResize(filePath, "character", 180, 180); err != nil {
		characters.DeleteByID(int(id), userID)
		c.ServerError(err, code, userID)
		return
	}

	image, err := characters.GetImageListByUserID(userID)
	if err != nil {
		characters.DeleteByID(int(id), userID)
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	c.Data["json"] = PostResponse{
		Image: image,
	}

	c.ServeJSON()
}
