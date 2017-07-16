package controllersSound

import (
	"errors"

	"github.com/wheatandcat/dotstamp_server/controllers"
	"github.com/wheatandcat/dotstamp_server/models"
	"github.com/wheatandcat/dotstamp_server/utils/contribution"

	validator "gopkg.in/go-playground/validator.v9"
)

// SaveController 追加コントローラ
type SaveController struct {
	controllers.BaseController
}

// SaveRequest 追加リクエスト
type SaveRequest struct {
	UserContributionID int `form:"userContributionId" validate:"min=1"`
	SoundStatus        int `form:"soundStatus" validate:"min=1,max=2"`
}

// SaveResponse 追加レスポンス
type SaveResponse struct {
	Warning bool   `json:"warning"`
	Message string `json:"message"`
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

	s, err := contributions.GetSoundByUserContributionID(request.UserContributionID)
	if err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	if s.ID == uint(0) {
		c.ServerError(errors.New("is added sound"), controllers.ErrCodeCommon, userID)
		return
	}

	if !contributions.ExistsSound(request.UserContributionID) {
		c.ServerError(errors.New("not exists file"), controllers.ErrCodeCommon, userID)
		return
	}

	tx := models.Begin()

	s.SoundStatus = request.SoundStatus
	if err := s.Save(); err != nil {
		models.Rollback(tx)
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	models.Commit(tx)

	c.Data["json"] = AddResponse{
		Warning: false,
		Message: "",
	}

	c.ServeJSON()
}
