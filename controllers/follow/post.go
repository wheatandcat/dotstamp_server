package controllersFollow

import (
	"strconv"

	"github.com/wheatandcat/dotstamp_server/controllers"
	"github.com/wheatandcat/dotstamp_server/models"
	"github.com/wheatandcat/dotstamp_server/utils/contribution"
	"github.com/wheatandcat/dotstamp_server/utils/follow"
)

// PostResponse 追加レスポンス
type PostResponse struct {
	Warning     bool   `json:"warning"`
	Message     string `json:"message"`
	FollowCount int    `json:"followCount"`
}

// Post 追加する
func (c *MainController) Post() {
	userID := c.GetUserID()
	if !c.IsNoLogin(userID) {
		c.ServerLoginNotFound()
		return
	}

	id, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if err != nil {
		c.ServerError(err, controllers.ErrParameter, userID)
		return
	}

	tx := models.Begin()

	if err = models.Lock("user_masters", userID); err != nil {
		models.Rollback(tx)
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	userContribution, err := contributions.GetByUserContributionID(id)
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

	check, err := follows.GetCountByUserIDAndUserContributionID(userID, id)
	if err != nil {
		models.Rollback(tx)
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	if check > 0 {
		models.Rollback(tx)
		c.ServerError(err, controllers.ErrFollowed, userID)
		return
	}

	if err = follows.Add(userID, id); err != nil {
		models.Rollback(tx)
		c.ServerError(err, controllers.ErrAddFollow, userID)
		return
	}

	count, err := follows.GetCountByUserContributionID(id)
	if err != nil {
		models.Rollback(tx)
		c.ServerError(err, controllers.ErrAddFollow, userID)
		return
	}

	models.Commit(tx)

	c.Data["json"] = PostResponse{
		Warning:     false,
		Message:     "",
		FollowCount: count,
	}

	c.ServeJSON()
}
