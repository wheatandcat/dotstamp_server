package controllersMovie

import (
	"strconv"

	"github.com/wheatandcat/dotstamp_server/controllers"
	"github.com/wheatandcat/dotstamp_server/models"
	"github.com/wheatandcat/dotstamp_server/utils/contribution"
)

// GetResponse 確認レスポンス
type GetResponse struct {
	Warning     bool   `json:"warning"`
	Message     string `json:"message"`
	MovieStatus int    `json:"movieStatus"`
}

// Get 確認する
func (c *MainController) Get() {
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

	userMovie, err := contributions.GetMovie(id, models.MovieTypeYoutube)
	if err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, 0)
		return
	}

	c.Data["json"] = GetResponse{
		Warning:     false,
		Message:     "",
		MovieStatus: userMovie.MovieStatus,
	}

	c.ServeJSON()
}
