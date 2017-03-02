package controllersSound

import (
	"dotstamp_server/controllers"
	"dotstamp_server/models"
	"dotstamp_server/utils/contribution"

	validator "gopkg.in/go-playground/validator.v9"
)

// SaveVoiceController ボイス保存コントローラ
type SaveVoiceController struct {
	controllers.BaseController
}

// SaveVoicedRequest ボイス保存リクエスト
type SaveVoicedRequest struct {
	ID        uint `form:"id" validate:"min=1"`
	VoiceType int  `form:"voice_type" validate:"min=1"`
}

// SaveVoiceResponse ボイス保存レスポンス
type SaveVoiceResponse struct {
	Warning bool
	Message string
}

// Post ボイス保存する
func (c *SaveVoiceController) Post() {
	userID := c.GetUserID()
	if !c.IsNoLogin(userID) {
		c.ServerLoginNotFound()
		return
	}

	request := SaveVoicedRequest{}
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
	models.Lock("user_masters", userID)

	if err := contributions.SaveSoundDetailTVoiceType(request.ID, request.VoiceType, userID); err != nil {
		models.Rollback(tx)
		c.ServerError(err, controllers.ErrCodeCommon)
		return
	}

	models.Commit(tx)

	c.Data["json"] = SaveVoiceResponse{
		Warning: false,
		Message: "",
	}

	c.ServeJSON()
}
