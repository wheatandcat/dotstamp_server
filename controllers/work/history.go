package controllersWork

import (
	"dotstamp_server/controllers"
	"dotstamp_server/models"
)

// WorkHistoryController 作品履歴コントローラ
type WorkHistoryController struct {
	controllers.BaseController
}

// Post ユーザー歴史一覧取得
func (c *WorkHistoryController) Post() {
	uID := 1

	userContribution := &models.UserContribution{}

	userContributionlist, _, err := userContribution.GetListByUserID(uID)
	if err != nil {
		c.ServerError(err, controllers.ErrCodeCommon)
		return
	}
	c.Data["json"] = userContributionlist
	c.ServeJSON()
}
