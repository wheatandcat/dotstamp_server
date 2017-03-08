package controllersContribution

import (
	"dotstamp_server/controllers"
	"dotstamp_server/utils/contribution"
)

// ListController listコントローラ
type ListController struct {
	controllers.BaseController
}

// ListRequest リストリクエスト
type ListRequest struct {
	Order int `form:"order"`
}

// Post 一覧を取得する
func (t *ListController) Post() {
	request := ListRequest{}
	if err := t.ParseForm(&request); err != nil {
		t.ServerError(err, controllers.ErrCodeCommon, 0)
		return
	}

	contributionlist, err := contributions.GetListByTop(0, (request.Order+1)*10)

	if err != nil {
		t.ServerError(err, controllers.ErrParameter, 0)
		return
	}

	t.Data["json"] = contributionlist
	t.ServeJSON()
}
