package controllersMovie

import (
	"errors"
	"strconv"

	"github.com/wheatandcat/dotstamp_server/controllers"
	"github.com/wheatandcat/dotstamp_server/utils/contribution"
	"github.com/wheatandcat/dotstamp_server/utils/movie"
)

// PostResponse 作成レスポンス
type PostResponse struct {
	Warning bool   `json:"warning"`
	Message string `json:"message"`
}

// Post 作成する
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

	u, err := contributions.GetByUserContributionID(id)
	if err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	if userID != u.UserID {
		c.ServerError(errors.New("diff UserID"), controllers.ErrCodeCommon, userID)
		return
	}

	if err = movie.ExecMakeMovie(id); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	c.Data["json"] = PostResponse{
		Warning: false,
		Message: "",
	}

	c.ServeJSON()
}
