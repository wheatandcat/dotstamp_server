package controllersProblem

import (
	"dotstamp_server/controllers"
	"dotstamp_server/utils/problem"
)

// AddController 追加コントローラ
type AddController struct {
	controllers.BaseController
}

// AddRequest 追加リクエスト
type AddRequest struct {
	UserContributionID int `form:"userContributionId"`
	Type               int `form:"type"`
}

// AddResponse 追加レスポンス
type AddResponse struct {
	Warning bool
	Message string
}

// Post 追加する
func (c *AddController) Post() {
	userID := c.GetUserID()
	if !c.IsNoLogin(userID) {
		userID = 0
	}

	request := AddRequest{}
	if err := c.ParseForm(&request); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon)
		return
	}

	if err := problem.Add(userID, request.UserContributionID, request.Type); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon)
		return
	}

	c.Data["json"] = AddResponse{
		Warning: false,
		Message: "",
	}

	c.ServeJSON()
}