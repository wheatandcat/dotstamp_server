package controllersContribution

import (
	"dotstamp_server/controllers"
	"dotstamp_server/utils/contribution"
	"strconv"
)

// UploadController アップロードコントローラ
type UploadController struct {
	controllers.BaseController
}

// UploadRequest アップロードリクエスト
type UploadRequest struct {
	UserContributionID int `form:"userContributionId"`
}

// UploadResponse アップロードレスポンス
type UploadResponse struct {
	Warning bool
	Message string
	Path    string
}

// Post 画像アップロード
func (c *UploadController) Post() {
	userID := c.GetUserID()
	if !c.IsNoLogin(userID) {
		c.ServerLoginNotFound()
		return
	}

	request := UploadRequest{}
	if err := c.ParseForm(&request); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon)
		return
	}

	id, err := contributions.GetImageIDAndAdd(request.UserContributionID)
	if err != nil {
		c.ServerError(err, controllers.ErrCodeCommon)
		return
	}

	filePath := strconv.Itoa(int(id)) + ".jpg"

	c.ToFile("./static/files/talk/" + filePath)

	c.Data["json"] = UploadResponse{
		Warning: false,
		Message: "",
		Path:    filePath,
	}
	c.ServeJSON()
}
