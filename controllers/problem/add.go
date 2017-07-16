package controllersProblem

import (
	"github.com/wheatandcat/dotstamp_server/controllers"
	"github.com/wheatandcat/dotstamp_server/utils/problem"
)

// AddController 追加コントローラ
type AddController struct {
	controllers.BaseController
}

// AddRequest 追加リクエスト
type AddRequest struct {
	ID   int `form:"id"`
	Type int `form:"type"`
}

// AddResponse 追加レスポンス
type AddResponse struct {
	Warning bool   `json:"warning"`
	Message string `json:"message"`
}

// Post 追加する
func (c *AddController) Post() {
	userID := c.GetUserID()
	if !c.IsNoLogin(userID) {
		userID = 0
	}

	request := AddRequest{}
	if err := c.ParseForm(&request); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	if err := problem.Add(userID, request.ID, request.Type); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	c.Data["json"] = AddResponse{
		Warning: false,
		Message: "",
	}

	c.ServeJSON()
}
