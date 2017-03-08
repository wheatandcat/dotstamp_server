package controllersSound

import (
	"dotstamp_server/controllers"
	"dotstamp_server/models"
	"dotstamp_server/utils/character"
	"dotstamp_server/utils/contribution"
	"errors"

	validator "gopkg.in/go-playground/validator.v9"
)

// AddController 追加コントローラ
type AddController struct {
	controllers.BaseController
}

// AddRequest 追加リクエスト
type AddRequest struct {
	UserContributionID int `form:"userContributionId" validate:"min=1"`
}

// AddResponse 追加レスポンス
type AddResponse struct {
	Warning     bool
	Message     string
	FollowCount int
}

// Post 追加する
func (c *AddController) Post() {
	userID := c.GetUserID()
	if !c.IsNoLogin(userID) {
		c.ServerLoginNotFound()
		return
	}

	request := AddRequest{}
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
	models.Lock("user_masters", userID)

	u, err := contributions.GetByUserContributionID(request.UserContributionID)
	if err != nil {
		models.Rollback(tx)
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	if userID != u.UserID {
		models.Rollback(tx)
		c.ServerError(errors.New("diff UserID"), controllers.ErrCodeCommon, userID)
		return
	}

	s, err := contributions.GetSoundByUserContributionID(request.UserContributionID)
	if err != nil {
		models.Rollback(tx)
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	if s.ID != uint(0) {
		models.Rollback(tx)
		c.ServerError(errors.New("is added sound"), controllers.ErrCodeCommon, userID)
		return
	}

	err = contributions.AddSound(request.UserContributionID, models.SoundStatusPrivate)
	if err != nil {
		models.Rollback(tx)
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	image, err := characters.GetImageListByUserID(userID)
	if err != nil {
		models.Rollback(tx)
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	charVoiceMap := map[int]int{}

	for _, v := range image {
		charVoiceMap[int(v.ID)] = v.VoiceType
	}

	body, err := contributions.GetBodyByUserContributionID(request.UserContributionID)
	if err != nil {
		models.Rollback(tx)
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	for k, v := range body {
		body[k].Character.VoiceType = charVoiceMap[v.Character.ID]
	}

	err = contributions.AddSoundDetailList(request.UserContributionID, body)
	if err != nil {
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
