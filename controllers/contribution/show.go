package controllersContribution

import (
	"dotstamp_server/controllers"
	"dotstamp_server/utils/contribution"
	"strconv"
)

// ShowController showコントローラ
type ShowController struct {
	controllers.BaseController
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

	c.Data["json"] = contribution

	c.ServeJSON()
}
