package controllersSound

import (
	"errors"
	"strconv"

	"github.com/wheatandcat/dotstamp_server/controllers"
	"github.com/wheatandcat/dotstamp_server/models"
	"github.com/wheatandcat/dotstamp_server/utils/character"
	"github.com/wheatandcat/dotstamp_server/utils/contribution"
)

// PostResponse 追加レスポンス
type PostResponse struct {
	Warning bool   `json:"warning"`
	Message string `json:"message"`
}

// Post 追加する
func (c *MainController) Post() {
	userID := c.GetUserID()
	if !c.IsNoLogin(userID) {
		c.ServerLoginNotFound()
		return
	}

	id, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if err != nil {
		c.ServerError(err, controllers.ErrParameter, userID)
		return
	}

	tx := models.Begin()
	models.Lock("user_masters", userID)

	u, err := contributions.GetByUserContributionID(id)
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

	s, err := contributions.GetSoundByUserContributionID(id)
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

	err = contributions.AddSound(id, models.SoundStatusPrivate)
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

	body, err := contributions.GetBodyByUserContributionID(id)
	if err != nil {
		models.Rollback(tx)
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	for k, v := range body {
		body[k].Character.VoiceType = charVoiceMap[v.Character.ID]
	}

	err = contributions.AddSoundDetailList(id, body)
	if err != nil {
		models.Rollback(tx)
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	models.Commit(tx)

	c.Data["json"] = PostResponse{
		Warning: false,
		Message: "",
	}

	c.ServeJSON()

}
