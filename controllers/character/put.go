package controllersCharacter

import (
	"encoding/json"

	"github.com/wheatandcat/dotstamp_server/controllers"
	"github.com/wheatandcat/dotstamp_server/models"
	"github.com/wheatandcat/dotstamp_server/utils/character"

	validator "gopkg.in/go-playground/validator.v9"
)

// PutRequest リクエスト
type PutRequest struct {
	ID        int `form:"id" validate:"min=1"`
	VoiceType int `form:"voiceType" validate:"min=1"`
}

// PutResponse 保存レスポンス
type PutResponse struct {
	Warning bool   `json:"warning"`
	Message string `json:"message"`
}

// Put 保存する
func (c *MainController) Put() {
	userID := c.GetUserID()
	if !c.IsNoLogin(userID) {
		c.ServerLoginNotFound()
		return
	}

	request := PutRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &request); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	tx := models.Begin()

	if err := characters.SaveToVoiceType(request.ID, request.VoiceType, int(userID)); err != nil {
		models.Rollback(tx)
		c.ServerError(err, controllers.ErrCodeUserNotFound, userID)
		return
	}

	models.Commit(tx)

	c.Data["json"] = PutResponse{
		Warning: false,
		Message: "",
	}

	c.ServeJSON()
}
