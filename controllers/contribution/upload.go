package controllersContribution

import (
	"fmt"
	"dotstamp_server/controllers"
	"dotstamp_server/utils"
)

// UploadController Uploadコントローラ
type UploadController struct {
	controllers.BaseController
}

// Post 画像アップロード
func (upload *UploadController) Post() {
	_, header, _ := upload.GetFile("file")
	displayUserID := 1
	userContributionID := 1

	fileName := utils.SrringToEncryption(header.Filename)
	filePath := fmt.Sprintf("%08d", displayUserID) + "_" + fmt.Sprintf("%08d", userContributionID) + "_" + fileName + ".jpg"

	upload.SaveToFile("file", "./static/files/talk/"+filePath)

	upload.Data["json"] = filePath
	upload.ServeJSON()
}
