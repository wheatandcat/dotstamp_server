package controllersSound

import (
	"encoding/json"
	"log"

	"github.com/wheatandcat/dotstamp_server/controllers"
	"github.com/wheatandcat/dotstamp_server/models"
	"github.com/wheatandcat/dotstamp_server/utils/contribution"

	validator "gopkg.in/go-playground/validator.v9"
)

// SaveVoiceController ボイス保存コントローラ
type SaveVoiceController struct {
	controllers.BaseController
}

// SaveVoicedRequest ボイス保存リクエスト
type SaveVoicedRequest struct {
	ID        uint `form:"id" validate:"min=1"`
	VoiceType int  `form:"voiceType" validate:"min=1"`
}

// SaveVoiceResponse ボイス保存レスポンス
type SaveVoiceResponse struct {
	Warning bool   `json:"warning"`
	Message string `json:"message"`
	ID      uint   `json:"id"`
}

// Put ボイス保存する
func (c *SaveVoiceController) Put() {
	userID := c.GetUserID()
	if !c.IsNoLogin(userID) {
		c.ServerLoginNotFound()
		return
	}

	request := SaveVoicedRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &request); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}
	log.Println(request)
	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	tx := models.Begin()

	if err := contributions.SaveSoundDetailTVoiceType(request.ID, request.VoiceType, userID); err != nil {
		models.Rollback(tx)
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	models.Commit(tx)

	u := models.UserContributionSoundDetail{}
	r, _, err := u.GetByID(request.ID)
	if err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	if err := contributions.AddTmpSound(r); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	c.Data["json"] = SaveVoiceResponse{
		Warning: false,
		Message: "",
		ID:      request.ID,
	}

	c.ServeJSON()
}
