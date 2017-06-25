package controllersMovie

import (
	"github.com/wheatandcat/dotstamp_server/controllers"
	"github.com/wheatandcat/dotstamp_server/utils/contribution"
	"github.com/wheatandcat/dotstamp_server/utils/movie"
	"errors"

	validator "gopkg.in/go-playground/validator.v9"
)

// MakeController 作成コントローラ
type MakeController struct {
	controllers.BaseController
}

// MakeRequest 作成リクエスト
type MakeRequest struct {
	UserContributionID int `form:"userContributionId" validate:"min=1"`
}

// MakeResponse 作成レスポンス
type MakeResponse struct {
	Warning bool
	Message string
}

// Post 作成する
func (c *MakeController) Post() {
	userID := c.GetUserID()
	if !c.IsNoLogin(userID) {
		c.ServerLoginNotFound()
		return
	}

	request := MakeRequest{}
	if err := c.ParseForm(&request); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	u, err := contributions.GetByUserContributionID(request.UserContributionID)
	if err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	if userID != u.UserID {
		c.ServerError(errors.New("diff UserID"), controllers.ErrCodeCommon, userID)
		return
	}

	if err = movie.ExecMakeMovie(request.UserContributionID); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	c.Data["json"] = MakeResponse{
		Warning: false,
		Message: "",
	}

	c.ServeJSON()
}
