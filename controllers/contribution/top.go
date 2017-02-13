package controllersContribution

import (
	"dotstamp_server/controllers"
	"dotstamp_server/utils/contribution"
)

// TopController Topコントローラ
type TopController struct {
	controllers.BaseController
}

// Post 新着情報を取得する
func (c *TopController) Post() {
	contributionList, err := contributions.GetByTop(0, 10)
	if err != nil {
		c.ServerError(err, controllers.ErrCodeCommon)
	}

	c.Data["json"] = contributionList
	c.ServeJSON()
}
