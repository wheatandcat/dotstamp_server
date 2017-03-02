package controllersCharacterImage

import (
	"dotstamp_server/controllers"
	"dotstamp_server/models"
	"dotstamp_server/utils/character"

	validator "gopkg.in/go-playground/validator.v9"
)

// SaveController 保存コントローラ
type SaveController struct {
	controllers.BaseController
}

// SaveRequest 保存リクエスト
type SaveRequest struct {
	ID        int `form:"id" validate:"min=1"`
	VoiceType int `form:"voiceType" validate:"min=1"`
}

// SaveResponse 保存レスポンス
type SaveResponse struct {
	Warning bool
	Message string
}

// Post 保存する
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

	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon)
		return
	}

	tx := models.Begin()

	if err := characters.SaveToVoiceType(request.ID, request.VoiceType, int(userID)); err != nil {
		models.Rollback(tx)
		c.ServerError(err, controllers.ErrCodeUserNotFound)
		return
	}

	models.Commit(tx)

	c.Data["json"] = SaveResponse{
		Warning: false,
		Message: "",
	}

	c.ServeJSON()
}
