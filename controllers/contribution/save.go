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
func (c *SaveController) Post() {
	userID := c.GetUserID()
	if !c.IsNoLogin(userID) {
		c.ServerLoginNotFound()
		return
	}

	request := SaveRequest{}
	if err := c.ParseForm(&request); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon)
		return
	}

	if err := contributions.Save(request.UserContributionID, userID, request.Title); err != nil {
		c.ServerError(err, controllers.ErrContributionSave)
	}

	if err := contributions.SaveDetail(request.UserContributionID, request.Body); err != nil {
		c.ServerError(err, controllers.ErrContributionSave)
	}

	c.Data["json"] = request.UserContributionID
	c.ServeJSON()
}
