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
func (t *TopController) Post() {
	contributionList, err := contributions.GetByTop(0, 10)
	if err != nil {
		t.ServerError(err, controllers.ErrCodeCommon)
	}

	t.Data["json"] = contributionList
	t.ServeJSON()
}
