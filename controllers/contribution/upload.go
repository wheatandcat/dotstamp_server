package controllersContribution

import (
	"dotstamp_server/controllers"
	"dotstamp_server/utils"
	"fmt"
)

// UploadController Uploadコントローラ
type UploadController struct {
	controllers.BaseController
}

// Post 画像アップロード
func (c *UploadController) Post() {
	userID := c.GetUserID()
	if !c.IsNoLogin(userID) {
		c.ServerLoginNotFound()
		return
	}

	_, header, _ := c.GetFile("file")
	displayUserID := 1
	userContributionID := 1

	fileName := utils.SrringToEncryption(header.Filename)
	filePath := fmt.Sprintf("%08d", displayUserID) + "_" + fmt.Sprintf("%08d", userContributionID) + "_" + fileName + ".jpg"

	c.SaveToFile("file", "./static/files/talk/"+filePath)

	c.Data["json"] = filePath
	c.ServeJSON()
}
