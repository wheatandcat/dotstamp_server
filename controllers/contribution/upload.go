package controllersContribution

import (
	"encoding/json"
	"strconv"

	"github.com/wheatandcat/dotstamp_server/controllers"
	"github.com/wheatandcat/dotstamp_server/utils/contribution"
)

// UploadController アップロードコントローラ
type UploadController struct {
	controllers.BaseController
}

// UploadRequest アップロードリクエスト
type UploadRequest struct {
	ID int `form:"id"`
}

// UploadResponse アップロードレスポンス
type UploadResponse struct {
	Warning bool   `json:"warning"`
	Message string `json:"message"`
	Path    string `json:"path"`
}

// Post 画像アップロード
func (c *UploadController) Post() {
	userID := c.GetUserID()
	if !c.IsNoLogin(userID) {
		c.ServerLoginNotFound()
		return
	}

	request := UploadRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &request); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, 0)
		return
	}

	id, err := contributions.GetImageIDAndAdd(request.ID)
	if err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, 0)
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
