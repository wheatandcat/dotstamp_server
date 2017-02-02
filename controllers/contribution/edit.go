package controllersContribution

import (
	"strconv"
	"dotstamp_server/controllers"
	"dotstamp_server/utils/contribution"
)

// EditController Editコントローラ
type EditController struct {
	controllers.BaseController
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
		t.ServerError(err, controllers.ErrParameter)
		return
	}

	c, err := contributions.GetContributionByUserContributionID(id)
	if err != nil {
		t.ServerError(err, controllers.ErrCodeCommon)
		return
	}

	if c.User.ID != userID {
		t.ServerError(err, controllers.ErrCodeUserNotFound)
		return
	}

	t.Data["json"] = c

	t.ServeJSON()
}
