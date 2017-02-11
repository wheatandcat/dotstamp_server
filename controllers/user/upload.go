package controllersUser

import (
	"dotstamp_server/controllers"
	"dotstamp_server/utils/user"
	"strconv"
)

// UploadController アップロード
type UploadController struct {
	controllers.BaseController
}

// Post アップロード
func (c *UploadController) Post() {
	userID := c.GetUserID()
	if !c.IsNoLogin(userID) {
		c.ServerLoginNotFound()
		return
	}

	id, err := user.GetIDAndAddProfileImage(userID)
	if err != nil {
		c.ServerError(err, controllers.ErrCodeCommon)
	}

	fileName := strconv.Itoa(id) + ".jpg"

	var code int
	if code, err = c.SetImageFileResize(fileName, "icon", 60, 60); err != nil {
		c.ServerError(err, code)
		return
	}

	c.Data["json"] = fileName
	c.ServeJSON()
}
