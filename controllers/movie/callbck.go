package controllersMovie

import (
	"context"
	"dotstamp_server/controllers"
	"dotstamp_server/models"
	"dotstamp_server/utils/contribution"
	"dotstamp_server/utils/movie"
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
	userID := c.GetUserID()
	if !c.IsNoLogin(userID) {
		c.Redirect(beego.AppConfig.String("errorUrl"), 302)
		return
	}

	request := CallbackRequest{}
	if err := c.ParseForm(&request); err != nil {
		c.Redirect(beego.AppConfig.String("errorUrl"), 302)
		return
	}

	context := context.Background()

	config := movie.GetConnect()

	tok, err := config.Exchange(context, request.Code)
	if err != nil {
		c.Redirect(beego.AppConfig.String("errorUrl"), 302)
		return
	}

	if tok.Valid() == false {
		c.Redirect(beego.AppConfig.String("errorUrl"), 302)
		return
	}

	client := config.Client(context, tok)

	u, err := contributions.GetByUserContributionID(request.UserContributionID)
	if err != nil {
		c.Redirect(beego.AppConfig.String("errorUrl"), 302)
		return
	}

	if userID != u.UserID {
		c.Redirect(beego.AppConfig.String("errorUrl"), 302)
		return
	}

	if !contributions.ExistsMovie(request.UserContributionID) {
		c.Redirect(beego.AppConfig.String("errorUrl"), 302)
		return
	}

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
		c.Redirect(beego.AppConfig.String("errorUrl"), 302)
		return
	}

	contributions.AddOrSaveMovie(int(u.ID), id, models.MovieTypeYoutube, models.StatusPublic)

	c.Redirect(beego.AppConfig.String("topurl")+"static/html/success.html", 302)
}
