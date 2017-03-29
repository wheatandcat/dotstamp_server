package controllersMovie

import (
	"dotstamp_server/controllers"
	"dotstamp_server/models"
	"dotstamp_server/utils/contribution"

	validator "gopkg.in/go-playground/validator.v9"
)

// CheckController 確認コントローラ
type CheckController struct {
	controllers.BaseController
}

// CheckRequest 確認リクエスト
type CheckRequest struct {
	UserContributionID int `form:"userContributionId" validate:"min=1"`
}

// CheckResponse 確認レスポンス
type CheckResponse struct {
	Warning     bool
	Message     string
	MovieStatus int
}

// Post 確認する
func (c *CheckController) Post() {
	request := CheckRequest{}
	if err := c.ParseForm(&request); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, 0)
		return
	}

	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, 0)
		return
	}

	userMovie, err := contributions.GetMovie(request.UserContributionID, models.MovieTypeYoutube)
	if err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, 0)
		return
	}

	c.Data["json"] = CheckResponse{
		Warning:     false,
		Message:     "",
		MovieStatus: userMovie.MovieStatus,
	}

	c.ServeJSON()
}
