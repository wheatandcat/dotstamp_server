package controllersCharacterImage

import (
	"dotstamp_server/controllers"
	"dotstamp_server/utils/character"
)

// UploadController アップロードコントローラ
type UploadController struct {
	controllers.BaseController
}

// UploadResponse アップロードレスポンス
type UploadResponse struct {
	Image []characters.Image
}

// Post 画像アップロード
func (c *UploadController) Post() {
	userID := c.GetUserID()
	if !c.IsNoLogin(userID) {
		c.ServerLoginNotFound()
		return
	}

	id, err := characters.AddImage(userID, 0, 0)
	if err != nil {
		c.ServerError(err, controllers.ErrCodeCommon)
		return
	}

	filePath := characters.GetImageName(id)

	var code int
	if code, err = c.SetImageFileResize(filePath, "character", 180, 180); err != nil {
		c.ServerError(err, code)
		return
	}

	image, err := characters.GetImageListByUserID(userID)
	if err != nil {
		c.ServerError(err, controllers.ErrCodeCommon)
		return
	}

	c.Data["json"] = UploadResponse{
		Image: image,
	}

	c.ServeJSON()
}
