package controllersMovie

import (
	"context"
	"errors"
	"strconv"

	"github.com/wheatandcat/dotstamp_server/controllers"
	"github.com/wheatandcat/dotstamp_server/models"
	"github.com/wheatandcat/dotstamp_server/utils"
	"github.com/wheatandcat/dotstamp_server/utils/contribution"
	"github.com/wheatandcat/dotstamp_server/utils/movie"

	"github.com/astaxie/beego"
)

// UploadController アップロードコントローラ
type UploadController struct {
	controllers.BaseController
}

// UploadResponse アップロードレスポンス
type UploadResponse struct {
	Warning bool   `json:"warning"`
	Message string `json:"message"`
	MovieID string `json:"movieID"`
}

// Post アップロードする
func (c *UploadController) Post() {
	userID := c.GetUserID()
	if !c.IsNoLogin(userID) {
		c.ServerLoginNotFound()
		return
	}

	cid, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if err != nil {
		c.ServerError(err, controllers.ErrParameter, userID)
		return
	}

	u, err := contributions.GetByUserContributionID(cid)
	if err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	if !contributions.ExistsMovie(cid) {
		c.ServerError(errors.New("not found movie"), controllers.ErrCodeCommon, userID)
		return
	}

	userMovie, err := contributions.GetMovie(int(u.ID), models.MovieTypeYoutube)
	if err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	if userMovie.MovieStatus == models.StatusRunning {
		c.ServerError(errors.New("making movie"), controllers.ErrCodeCommon, userID)
		return
	}

	if userMovie.MovieStatus == models.StatusUploading {
		c.ServerError(errors.New("uploading"), controllers.ErrCodeCommon, userID)
		return
	}

	userMovie.MovieStatus = models.StatusUploading
	userMovie.Save()

	upload, err := contributions.GetUploadByUserContributionID(cid)
	if err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	id, err := uploadYoutube(u, upload.Token)
	if err != nil {
		contributions.AddOrSaveMovie(int(u.ID), "", models.MovieTypeYoutube, models.StatusError)
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	if err = contributions.AddOrSaveMovie(int(u.ID), id, models.MovieTypeYoutube, models.StatusPublic); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	c.Data["json"] = UploadResponse{
		Warning: false,
		Message: "",
		MovieID: id,
	}

	c.ServeJSON()
}

func uploadYoutube(u models.UserContribution, token string) (string, error) {
	if utils.IsTest() {
		return "test", nil
	}

	context := context.Background()

	config := movie.GetConnect()

	tok, err := config.Exchange(context, token)
	if err != nil {
		return "", err
	}

	if tok.Valid() == false {
		return "", errors.New("vaild token")
	}

	client := config.Client(context, tok)

	if err = contributions.AddOrSaveMovie(int(u.ID), "", models.MovieTypeYoutube, models.StatusRunning); err != nil {
		return "", err
	}

	videoStatus := "unlisted"
	if beego.AppConfig.String("runmode") == "prod" {
		videoStatus = "public"
	}

	m := movie.Upload{
		UserContributionID: strconv.Itoa(int(u.ID)),
		Title:              u.Title,
		Description:        "元記事はこちら " + beego.AppConfig.String("topurl") + "#/contribution/show/" + strconv.Itoa(int(u.ID)),
		CategoryID:         "22",
		VideoStatus:        videoStatus,
	}

	return movie.UploadToYoutube(client, m)
}
