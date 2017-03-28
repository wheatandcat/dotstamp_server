package controllersMovie

import (
	"dotstamp_server/controllers"
	"dotstamp_server/models"
	"dotstamp_server/utils/contribution"
	"dotstamp_server/utils/movie"
	"dotstamp_server/utils/sound"
	"errors"
	"strconv"

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

	list, err := contributions.GetSoundDetailListByUserContributionID(request.UserContributionID)
	if err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	// 音声ファイル作成
	if err = contributions.MakeSoundFile(request.UserContributionID, list); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	if err = contributions.UpdateSoundToMakeStatus(request.UserContributionID, models.MakeStatusMade); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	if err = sound.ToM4a(strconv.Itoa(request.UserContributionID)); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	// 動画ファイル作成
	if err = movie.Make(strconv.Itoa(request.UserContributionID)); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	if err = movie.ToFilter(strconv.Itoa(request.UserContributionID)); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	userMovie, err := contributions.GetMovie(request.UserContributionID, models.MovieTypeYoutube)
	if err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	if userMovie.MovieStatus != 0 && userMovie.MovieStatus != models.StatusReMeake {
		userMovie.MovieStatus = models.StatusReMeake
		userMovie.MovieID = ""

		if err = userMovie.Save(); err != nil {
			c.ServerError(err, controllers.ErrCodeCommon, userID)
			return
		}
	}

	c.Data["json"] = MakeResponse{
		Warning: false,
		Message: "",
	}

	c.ServeJSON()

}
