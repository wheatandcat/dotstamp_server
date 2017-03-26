package controllersMovie

import (
	"context"
	"dotstamp_server/controllers"
	"dotstamp_server/models"
	"dotstamp_server/utils/contribution"
	"dotstamp_server/utils/movie"
	"errors"
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
)

// CallbackController コールバックコントローラ
type CallbackController struct {
	controllers.BaseController
}

// CallbackRequest コールバッククエスト
type CallbackRequest struct {
	Code               string `form:"code"`
	UserContributionID int    `form:"state"`
}

// Get コールバックする
func (c *CallbackController) Get() {
	request := CallbackRequest{}
	if err := c.ParseForm(&request); err != nil {
		c.RedirectError(err, 0)
		return
	}

	fmt.Println("中・・・")

	context := context.Background()

	config := movie.GetConnect()

	tok, err := config.Exchange(context, request.Code)
	if err != nil {
		c.RedirectError(err, 0)
		return
	}

	if tok.Valid() == false {
		c.RedirectError(errors.New("vaild token"), 0)
		return
	}

	client := config.Client(context, tok)

	u, err := contributions.GetByUserContributionID(request.UserContributionID)
	if err != nil {
		c.RedirectError(err, 0)
		return
	}

	if !contributions.ExistsMovie(request.UserContributionID) {
		c.RedirectError(errors.New("not found movie"), 0)
		return
	}

	userMovie, err := contributions.GetMovie(int(u.ID), models.MovieTypeYoutube)
	if err != nil {
		c.RedirectError(err, 0)
		return
	}

	if userMovie.MovieStatus == models.StatusRunning {
		c.RedirectError(errors.New("making movie"), 0)
		return
	}

	contributions.AddOrSaveMovie(int(u.ID), "", models.MovieTypeYoutube, models.StatusRunning)

	videoStatus := "unlisted"
	if beego.AppConfig.String("runmode") == "prod" {
		videoStatus = "public"

	}

	m := movie.Upload{
		UserContributionID: strconv.Itoa(request.UserContributionID),
		Title:              u.Title,
		Description:        beego.AppConfig.String("topurl") + "contribution/show/" + strconv.Itoa(request.UserContributionID),
		CategoryID:         "22",
		VideoStatus:        videoStatus,
	}

	id, err := movie.UploadToYoutube(client, m)
	if err != nil {
		contributions.AddOrSaveMovie(int(u.ID), "", models.MovieTypeYoutube, models.StatusError)
		c.RedirectError(errors.New("diff user_id"), 0)
		return
	}

	contributions.AddOrSaveMovie(int(u.ID), id, models.MovieTypeYoutube, models.StatusPublic)

	c.Redirect(beego.AppConfig.String("topurl")+"static/html/success.html", 302)
}
