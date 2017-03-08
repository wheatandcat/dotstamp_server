package controllersSound

import (
	"dotstamp_server/controllers"
	"dotstamp_server/models"
	"dotstamp_server/utils/contribution"

	validator "gopkg.in/go-playground/validator.v9"
)

// SaveBodyController 本文保存コントローラ
type SaveBodyController struct {
	controllers.BaseController
}

// SaveBodydRequest 本文保存リクエスト
type SaveBodydRequest struct {
	ID   uint   `form:"id" validate:"min=1"`
	Body string `form:"body" validate:"max=256"`
}

// SaveBodyResponse 本文保存レスポンス
type SaveBodyResponse struct {
	Warning bool
	Message string
}

// Post 本文保存する
func (c *SaveBodyController) Post() {
	userID := c.GetUserID()
	if !c.IsNoLogin(userID) {
		c.ServerLoginNotFound()
		return
	}

	request := SaveBodydRequest{}
	if err := c.ParseForm(&request); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	tx := models.Begin()

	if err := contributions.SaveSoundDetailToBodySound(request.ID, request.Body, userID); err != nil {
		models.Rollback(tx)
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	models.Commit(tx)

	c.Data["json"] = SaveBodyResponse{
		Warning: false,
		Message: "",
	}

	c.ServeJSON()
}
