package controllersFollow

import (
	"dotstamp_server/controllers"
	"dotstamp_server/utils/contribution"
	"dotstamp_server/utils/follow"
)

// DeleteController 削除コントローラ
type DeleteController struct {
	controllers.BaseController
}

// DeleteRequest 削除リクエスト
type DeleteRequest struct {
	UserContributionID int `form:"userContributionId"`
}

// DeleteResponse 削除レスポンス
type DeleteResponse struct {
	Warning     bool
	Message     string
	FollowCount int
}

// Post 削除する
func (c *DeleteController) Post() {
	userID := c.GetUserID()
	if !c.IsNoLogin(userID) {
		c.ServerLoginNotFound()
		return
	}

	request := DeleteRequest{}
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

	userfollow, err := follows.GetByUserIDAndUserContributionID(userID, request.UserContributionID)
	if err != nil {
		c.ServerError(err, controllers.ErrCodeCommon)
		return
	}

	if userfollow.ID == uint(0) {
		c.ServerError(err, controllers.ErrCodeCommon)
		return
	}

	if err = follows.Delete(userfollow.ID); err != nil {
		c.ServerError(err, controllers.ErrAddFollow)
		return
	}

	count, err := follows.GetCountByUserContributionID(request.UserContributionID)
	if err != nil {
		c.ServerError(err, controllers.ErrAddFollow)
		return
	}

	c.Data["json"] = DeleteResponse{
		Warning:     false,
		Message:     "",
		FollowCount: count,
	}

	c.ServeJSON()
}
