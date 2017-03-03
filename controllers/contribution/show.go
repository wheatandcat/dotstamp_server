package controllersContribution

import (
	"dotstamp_server/controllers"
	"dotstamp_server/models"
	"dotstamp_server/utils/contribution"
	"dotstamp_server/utils/follow"
	"strconv"
)

// ShowController showコントローラ
type ShowController struct {
	controllers.BaseController
}

// ShowResponse 確認レスポンス
type ShowResponse struct {
	contributions.Contribution
	FollowCount int
	Following   bool
	SoundFile   bool
}

// Post 投稿詳細を取得する
func (c *ShowController) Post() {
	id, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if err != nil {
		c.ServerError(err, controllers.ErrParameter)
		return
	}

	contribution, err := contributions.GetContributionByUserContributionID(id)
	if err != nil {
		c.ServerError(err, controllers.ErrCodeCommon)
		return
	}

	followCount, err := follows.GetCountByUserContributionID(id)
	if err != nil {
		c.ServerError(err, controllers.ErrCodeCommon)
		return
	}

	following := false

	userID := c.GetUserID()
	if c.IsNoLogin(userID) {
		var count int
		count, err = follows.GetCountByUserIDAndUserContributionID(userID, id)
		if err != nil {
			c.ServerError(err, controllers.ErrCodeCommon)
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
		c.ServerError(err, controllers.ErrCodeCommon)
		return
	}

	soundFile := false
	if s.SoundStatus == models.SoundStatusPublic {
		soundFile = contributions.ExistsSound(id)
	}

	c.Data["json"] = ShowResponse{
		Contribution: contribution,
		FollowCount:  followCount,
		Following:    following,
		SoundFile:    soundFile,
	}

	c.ServeJSON()
}
