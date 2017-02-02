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
func (w *WorkHistoryController) Post() {
	uID := 1

	userContribution := &models.UserContribution{}

	userContributionlist := userContribution.GetListByUserID(uID)

	w.Data["json"] = userContributionlist
	w.ServeJSON()
}
