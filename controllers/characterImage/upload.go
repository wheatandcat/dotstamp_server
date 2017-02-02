package controllersCharacterImage

import (
	"dotstamp_server/controllers"
	"dotstamp_server/utils/character"
	"dotstamp_server/utils/image"
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
func (t *UploadController) Post() {
	userID := t.GetUserID()
	if !t.IsNoLogin(userID) {
		t.ServerLoginNotFound()
		return
	}

	id, err := characters.AddImage(userID, 0, 0)
	if err != nil {
		t.ServerError(err, controllers.ErrCodeCommon)
		return
	}

	filePath := characters.GetImageName(id)

	tmpImagePath := "./static/files/tmp/character/_tmp_" + filePath
	t.SaveToFile("file", tmpImagePath)

	tmpRootImagePath := "./static/files/tmp/character/" + filePath

	if err := images.PngToJpeg(tmpImagePath, tmpRootImagePath); err != nil {
		t.ServerError(err, controllers.ErrImageConversion)
		return
	}

	outputImagePath := "./static/files/character/" + filePath

	if err := images.Resize(tmpRootImagePath, outputImagePath, 180, 180); err != nil {
		t.ServerError(err, controllers.ErrImageResize)
		return
	}

	image := characters.GetImageListByUserID(userID)

	t.Data["json"] = UploadResponse{
		Image: image,
	}

	t.ServeJSON()
}
