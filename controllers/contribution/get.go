package controllersContribution

import (
	"strconv"

	"github.com/wheatandcat/dotstamp_server/controllers"
	"github.com/wheatandcat/dotstamp_server/models"
	"github.com/wheatandcat/dotstamp_server/utils/contribution"
	"github.com/wheatandcat/dotstamp_server/utils/follow"
)

// GetResponse レスポンス
type GetResponse struct {
	contributions.Contribution
	FollowCount int                          `json:"followCount"`
	Following   bool                         `json:"following"`
	SoundFile   bool                         `json:"soundFile"`
	Movie       models.UserContributionMovie `json:"movie"`
}

// Get 投稿詳細を取得する
func (c *MainController) Get() {
	id, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if err != nil {
		c.ServerError(err, controllers.ErrParameter, 0)
		return
	}

	contribution, err := contributions.GetContributionByUserContributionID(id)
	if err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, 0)
		return
	}

	followCount, err := follows.GetCountByUserContributionID(id)
	if err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, 0)
		return
	}

	following := false

	userID := c.GetUserID()
	if c.IsNoLogin(userID) {
		var count int
		count, err = follows.GetCountByUserIDAndUserContributionID(userID, id)
		if err != nil {
			c.ServerError(err, controllers.ErrCodeCommon, 0)
			return
		}

		if count > 0 {
			following = true
		}
	}

	if contribution.User.ID != uint(userID) {
		contribution = contributions.ContributionToPublic(contribution)
	}

	s, err := contributions.GetSoundByUserContributionID(id)
	if err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, 0)
		return
	}

	soundFile := false
	if s.SoundStatus == models.SoundStatusPublic {
		soundFile = contributions.ExistsSound(id)
	}

	movie, err := contributions.GetMovie(id, models.MovieTypeYoutube)
	if err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}
	contributions.AddLog(userID, id)

	c.Data["json"] = GetResponse{
		Contribution: contribution,
		FollowCount:  followCount,
		Following:    following,
		SoundFile:    soundFile,
		Movie:        movie,
	}

	c.ServeJSON()
}
