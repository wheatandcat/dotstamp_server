package controllersUser

import (
	"dotstamp_server/controllers"
	"dotstamp_server/models"
	"dotstamp_server/utils/user"
)

// SaveController 保存
type SaveController struct {
	controllers.BaseController
}

// SaveRequest 保存リクエスト
type SaveRequest struct {
	Name string `form:"name"`
}

// SaveResponse 保存レスポンス
type SaveResponse struct {
	Success bool
}

// Post ユーザー情報
func (c *SaveController) Post() {
	userID := c.GetUserID()
	if !c.IsNoLogin(userID) {
		c.ServerLoginNotFound()
		return
	}

	request := SaveRequest{}

	if err := c.ParseForm(&request); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon)
		return
	}

	tx := models.Begin()

	if err := user.Upadate(userID, request.Name); err != nil {
		models.Rollback(tx)
		c.ServerError(err, controllers.ErrUserSave)
		return
	}

	models.Commit(tx)

	c.Data["json"] = SaveResponse{
		Success: true,
	}

	c.ServeJSON()
}
