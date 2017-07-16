package controllersSound

import (
	"errors"
	"strconv"

	"github.com/wheatandcat/dotstamp_server/controllers"
	"github.com/wheatandcat/dotstamp_server/models"
	"github.com/wheatandcat/dotstamp_server/utils/contribution"
)

// GetResponse 確認レスポンス
type GetResponse struct {
	List        []models.UserContributionSoundDetail `json:"list"`
	SoundStatus int                                  `json:"soundStatus"`
	SoundFile   bool                                 `json:"soundFile"`
	MovieFile   bool                                 `json:"movieFile"`
	Movie       models.UserContributionMovie         `json:"movie"`
}

// Get 確認する
func (c *MainController) Get() {
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

	u, err := contributions.GetByUserContributionID(id)
	if err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	if userID != u.UserID {
		c.ServerError(errors.New("diff UserID"), controllers.ErrCodeCommon, userID)
		return
	}

	s, err := contributions.GetSoundByUserContributionID(id)
	if err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	if s.ID == uint(0) {
		c.ServerError(errors.New("not dound ID"), controllers.ErrCodeCommon, userID)
		return
	}

	list, err := contributions.GetSoundDetailListByUserContributionID(id)
	if err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	movie, err := contributions.GetMovie(id, models.MovieTypeYoutube)
	if err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	c.Data["json"] = GetResponse{
		List:        list,
		SoundFile:   contributions.ExistsSound(id),
		SoundStatus: s.SoundStatus,
		MovieFile:   contributions.ExistsMovie(id),
		Movie:       movie,
	}

	c.ServeJSON()
}
