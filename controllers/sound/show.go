package controllersSound

import (
	"dotstamp_server/controllers"
	"dotstamp_server/models"
	"dotstamp_server/utils/contribution"
	"errors"

	validator "gopkg.in/go-playground/validator.v9"
)

// ShowController 確認コントローラ
type ShowController struct {
	controllers.BaseController
}

// ShowRequest 確認リクエスト
type ShowRequest struct {
	UserContributionID int `form:"userContributionId" validate:"min=1"`
}

// ShowResponse 確認レスポンス
type ShowResponse struct {
	List        []models.UserContributionSoundDetail
	SoundStatus int
	SoundFile   bool
}

// Post 確認する
func (c *ShowController) Post() {
	userID := c.GetUserID()
	if !c.IsNoLogin(userID) {
		c.ServerLoginNotFound()
		return
	}

	request := ShowRequest{}
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
		c.ServerError(errors.New("not dound ID"), controllers.ErrCodeCommon)
		return
	}

	list, err := contributions.GetSoundDetailListByUserContributionID(request.UserContributionID)
	if err != nil {
		c.ServerError(err, controllers.ErrCodeCommon)
		return
	}

	c.Data["json"] = ShowResponse{
		List:        list,
		SoundFile:   contributions.ExistsSound(request.UserContributionID),
		SoundStatus: s.SoundStatus,
	}

	c.ServeJSON()
}
