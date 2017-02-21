package controllersFollow

import (
	"dotstamp_server/controllers"
	"dotstamp_server/utils/contribution"
	"dotstamp_server/utils/follow"
)

// AddController 追加コントローラ
type AddController struct {
	controllers.BaseController
}

// AddRequest 追加リクエスト
type AddRequest struct {
	UserContributionID int `form:"userContributionId"`
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
		c.ServerLoginNotFound()
		return
	}

	request := AddRequest{}
	if err := c.ParseForm(&request); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon)
		return
	}

	userContribution, err := contributions.GetByUserContributionID(request.UserContributionID)
	if err != nil {
		c.ServerError(err, controllers.ErrCodeCommon)
		return
	}

	if userContribution.ID == uint(0) {
		c.ServerError(err, controllers.ErrContributionNotFound)
		return
	}

	count, err := follows.GetCountByUserIDAndUserContributionID(userID, request.UserContributionID)
	if err != nil {
		c.ServerError(err, controllers.ErrCodeCommon)
		return
	}

	if count > 0 {
		c.ServerError(err, controllers.ErrFollowed)
		return
	}

	if err := follows.Add(userID, request.UserContributionID); err != nil {
		c.ServerError(err, controllers.ErrAddFollow)
		return
	}

	c.Data["json"] = AddResponse{
		Warning: false,
		Message: "",
	}

	c.ServeJSON()
}
