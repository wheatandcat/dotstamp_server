package controllersContribution

import (
	"dotstamp_server/controllers"
	"dotstamp_server/utils/contribution"
	"strconv"
)

// DeleteController Deleteコントローラ
type DeleteController struct {
	controllers.BaseController
}

// Post 画像を削除する
func (t *DeleteController) Post() {
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

	if err = contributions.DeleteByID(id, userID); err != nil {
		t.ServerError(err, controllers.ErrCodeUserNotFound)
		return
	}

	userContributionlist, err := contributions.GetListByUserID(userID)
	if err != nil {
		t.ServerError(err, controllers.ErrCodeCommon)
		return
	}

	t.Data["json"] = userContributionlist

	t.ServeJSON()
}
