package controllersSound

import (
	"dotstamp_server/controllers"
	"dotstamp_server/models"
	"dotstamp_server/utils/contribution"
	"errors"

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
	Warning     bool
	Message     string
	FollowCount int
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

	u, err := contributions.GetByUserContributionID(request.UserContributionID)
	if err != nil {
		c.ServerError(err, controllers.ErrCodeCommon)
		return
	}

	if userID != u.UserID {
		c.ServerError(errors.New("diff UserID"), controllers.ErrCodeCommon)
		return
	}

	s, err := contributions.GetSoundByUserContributionID(request.UserContributionID)
	if err != nil {
		c.ServerError(err, controllers.ErrCodeCommon)
		return
	}

	if s.ID == uint(0) {
		c.ServerError(errors.New("is added sound"), controllers.ErrCodeCommon)
		return
	}

	if !contributions.ExistsSound(request.UserContributionID) {
		c.ServerError(errors.New("not exists file"), controllers.ErrCodeCommon)
		return
	}

	tx := models.Begin()

	s.SoundStatus = request.SoundStatus
	if err := s.Save(); err != nil {
		models.Rollback(tx)
		c.ServerError(err, controllers.ErrCodeCommon)
		return
	}

	models.Commit(tx)

	c.Data["json"] = AddResponse{
		Warning: false,
		Message: "",
	}

	c.ServeJSON()
}