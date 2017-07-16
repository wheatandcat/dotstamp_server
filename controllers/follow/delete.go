package controllersFollow

import (
	"github.com/wheatandcat/dotstamp_server/controllers"
	"github.com/wheatandcat/dotstamp_server/models"
	"github.com/wheatandcat/dotstamp_server/utils/contribution"
	"github.com/wheatandcat/dotstamp_server/utils/follow"
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
	Warning     bool   `json:"warning"`
	Message     string `json:"message"`
	FollowCount int    `json:"followCount"`
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
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	tx := models.Begin()

	if err := models.Lock("user_masters", userID); err != nil {
		models.Rollback(tx)
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	userContribution, err := contributions.GetByUserContributionID(request.UserContributionID)
	if err != nil {
		models.Rollback(tx)
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	if userContribution.ID == uint(0) {
		models.Rollback(tx)
		c.ServerError(err, controllers.ErrContributionNotFound, userID)
		return
	}

	userfollow, err := follows.GetByUserIDAndUserContributionID(userID, request.UserContributionID)
	if err != nil {
		models.Rollback(tx)
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	if userfollow.ID == uint(0) {
		models.Rollback(tx)
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	if err = follows.Delete(userfollow.ID); err != nil {
		models.Rollback(tx)
		c.ServerError(err, controllers.ErrAddFollow, userID)
		return
	}

	count, err := follows.GetCountByUserContributionID(request.UserContributionID)
	if err != nil {
		models.Rollback(tx)
		c.ServerError(err, controllers.ErrAddFollow, userID)
		return
	}

	models.Commit(tx)

	c.Data["json"] = DeleteResponse{
		Warning:     false,
		Message:     "",
		FollowCount: count,
	}

	c.ServeJSON()
}
