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
func (t *ContributionListController) Post() {
	userID := t.GetUserID()
	if !t.IsNoLogin(userID) {
		t.ServerLoginNotFound()
		return
	}
	

	userContributionlist := contributions.GetListByUserID(userID)

	t.Data["json"] = userContributionlist
	t.ServeJSON()
}
