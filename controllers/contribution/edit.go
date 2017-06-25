package controllersContribution

import (
	"github.com/wheatandcat/dotstamp_server/controllers"
	"github.com/wheatandcat/dotstamp_server/models"
	"github.com/wheatandcat/dotstamp_server/utils/contribution"
	"strconv"
)

// EditController 編集コントローラ
type EditController struct {
	controllers.BaseController
}

// EditResponse 編集レスポンス
type EditResponse struct {
	contributions.Contribution
	Sound     bool
	SoundFile bool
}

// Post 編集する
func (t *EditController) Post() {
	userID := t.GetUserID()
	if !t.IsNoLogin(userID) {
		t.ServerLoginNotFound()
		return
	}

	id, err := strconv.Atoi(t.Ctx.Input.Param(":id"))
	if err != nil {
		t.ServerError(err, controllers.ErrParameter, userID)
		return
	}

	c, err := contributions.GetContributionByUserContributionID(id)
	if err != nil {
		t.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	if int(c.User.ID) != userID {
		t.ServerError(err, controllers.ErrCodeUserNotFound, userID)
		return
	}

	s, err := contributions.GetSoundByUserContributionID(id)
	if err != nil {
		t.ServerError(err, controllers.ErrParameter, userID)
		return
	}

	soundFile := false
	if s.SoundStatus == models.SoundStatusPublic {
		soundFile = contributions.ExistsSound(id)
	}

	t.Data["json"] = EditResponse{
		Contribution: c,
		Sound:        (s.ID != uint(0)),
		SoundFile:    soundFile,
	}

	t.ServeJSON()
}
