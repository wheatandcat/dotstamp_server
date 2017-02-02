package controllersContribution

import (
	"dotstamp_server/controllers"
	"dotstamp_server/utils/contribution"
)

// SaveController Saveコントローラ
type SaveController struct {
	controllers.BaseController
}

// SaveRequest 保存リクエスト
type SaveRequest struct {
	UserContributionID int    `form:"userContributionId"`
	Title              string `form:"title"`
	Body               string `form:"body"`
}

// Post 保存する
func (t *SaveController) Post() {
	userID := t.GetUserID()
	if !t.IsNoLogin(userID) {
		t.ServerLoginNotFound()
		return
	}

	request := SaveRequest{}
	if err := t.ParseForm(&request); err != nil {
		t.ServerError(err, controllers.ErrCodeCommon)
		return
	}

	if err := contributions.Save(request.UserContributionID, userID, request.Title); err != nil {
		t.ServerError(err, controllers.ErrContributionSave)
	}

	if err := contributions.SaveDetail(request.UserContributionID, request.Body); err != nil {
		t.ServerError(err, controllers.ErrContributionSave)
	}

	t.Data["json"] = request.UserContributionID
	t.ServeJSON()
}
