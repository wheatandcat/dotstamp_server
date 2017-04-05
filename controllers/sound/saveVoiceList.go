package controllersSound

import (
	"dotstamp_server/controllers"
	"dotstamp_server/models"
	"dotstamp_server/utils/contribution"
	"errors"

	validator "gopkg.in/go-playground/validator.v9"
)

// SaveVoiceListController ボイスリスト更新コントローラ
type SaveVoiceListController struct {
	controllers.BaseController
}

// SaveVoiceListRequest ボイスリスト更新リクエスト
type SaveVoiceListRequest struct {
	UserContributionID int `form:"userContributionId" validate:"min=1"`
	VoiceType          int `form:"voice_type" validate:"min=1"`
}

// SaveVoiceListResponse ボイスリスト更新レスポンス
type SaveVoiceListResponse struct {
	Warning bool
	Message string
}

// Post ボイスリストを更新する
func (c *SaveVoiceListController) Post() {
	userID := c.GetUserID()
	if !c.IsNoLogin(userID) {
		c.ServerLoginNotFound()
		return
	}

	request := SaveVoiceListRequest{}
	if err := c.ParseForm(&request); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	u, err := contributions.GetByUserContributionID(request.UserContributionID)
	if err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	if userID != u.UserID {
		c.ServerError(errors.New("diff UserID"), controllers.ErrCodeCommon, userID)
		return
	}

	if err := contributions.UpdatesSoundToMakeStatusAndVoiceTypeByUserContributionID(request.UserContributionID, models.MakeStatusUncreated, request.VoiceType); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	c.Data["json"] = SaveVoiceListResponse{
		Warning: false,
		Message: "",
	}

	c.ServeJSON()
}
