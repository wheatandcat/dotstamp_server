package controllersUser

import (
	"dotstamp_server/controllers"
	"dotstamp_server/utils/contribution"
)

// ContributionListController 投稿リスト
type ContributionListController struct {
	controllers.BaseController
}

// Post ユーザー投稿一覧を取得する
func (c *ContributionListController) Post() {
	userID := c.GetUserID()
	if !c.IsNoLogin(userID) {
		c.ServerLoginNotFound()
		return
	}

	userContributionlist, err := contributions.GetListByUserID(userID)
	if err != nil {
		c.ServerError(err, controllers.ErrCodeCommon)
		return
	}

	c.Data["json"] = userContributionlist
	c.ServeJSON()
}
