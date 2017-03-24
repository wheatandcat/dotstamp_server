package controllersMovie

import (
	"context"
	"dotstamp_server/controllers"
	"dotstamp_server/models"
	"dotstamp_server/utils/contribution"
	"dotstamp_server/utils/movie"
	"log"
)

// MakeController 作成コントローラ
type MakeController struct {
	controllers.BaseController
}

// MakeRequest 作成リクエスト
type MakeRequest struct {
	Code  string `form:"code"`
	State string `form:"state"`
}

// MakeResponse 作成レスポンス
type MakeResponse struct {
	Warning bool
	Message string
}

// Get 作成する
func (c *MakeController) Get() {
	request := MakeRequest{}
	if err := c.ParseForm(&request); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, 0)
		return
	}

	log.Println(request)
	context := context.Background()

	config := movie.GetConnect()
	config.RedirectURL = "http://192.168.33.10.xip.io:8080/movie/make"

	tok, err := config.Exchange(context, request.Code)
	if err != nil {
		log.Println("aaaa")
		c.ServerError(err, controllers.ErrCodeCommon, 0)
		return
	}

	if tok.Valid() == false {
		log.Println("aaaa")
		c.ServerError(err, controllers.ErrCodeCommon, 0)
		return
	}

	client := config.Client(context, tok)

	m := movie.Upload{
		UserContributionID: "21",
		Title:              "test",
		Description:        "test",
		CategoryID:         "22",
		VideoStatus:        "unlisted",
	}

	id, err := movie.UploadToYoutube(client, m)
	if err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, 0)
		return
	}

	contributions.AddOrSaveMovie(1, id, models.MovieTypeYoutube, models.StatusPublic)

	c.Data["json"] = MakeResponse{
		Warning: false,
		Message: "",
	}

	c.ServeJSON()

}
