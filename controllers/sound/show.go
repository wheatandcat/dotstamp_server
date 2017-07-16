package controllersSound

import (
	"errors"

	"github.com/wheatandcat/dotstamp_server/controllers"
	"github.com/wheatandcat/dotstamp_server/models"
	"github.com/wheatandcat/dotstamp_server/utils/contribution"

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
	List        []models.UserContributionSoundDetail `json:"list"`
	SoundStatus int                                  `json:"soundStatus"`
	SoundFile   bool                                 `json:"soundFile"`
	MovieFile   bool                                 `json:"movieFile"`
	Movie       models.UserContributionMovie         `json:"movie"`
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
		c.ServerError(errors.New("not dound ID"), controllers.ErrCodeCommon, userID)
		return
	}

	list, err := contributions.GetSoundDetailListByUserContributionID(request.UserContributionID)
	if err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	movie, err := contributions.GetMovie(request.UserContributionID, models.MovieTypeYoutube)
	if err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	c.Data["json"] = ShowResponse{
		List:        list,
		SoundFile:   contributions.ExistsSound(request.UserContributionID),
		SoundStatus: s.SoundStatus,
		MovieFile:   contributions.ExistsMovie(request.UserContributionID),
		Movie:       movie,
	}

	c.ServeJSON()
}
