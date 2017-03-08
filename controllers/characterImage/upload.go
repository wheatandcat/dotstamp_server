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

	c.Data["json"] = UploadResponse{
		Image: image,
	}

	c.ServeJSON()
}
