package controllersContribution

import (
	"strconv"
	"dotstamp_server/controllers"
	"dotstamp_server/utils/contribution"
)

// ShowController showコントローラ
type ShowController struct {
	controllers.BaseController
}

// Post 投稿詳細を取得する
func (t *ShowController) Post() {
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

	t.Data["json"] = c

	t.ServeJSON()
}
