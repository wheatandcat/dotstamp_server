package controllersContribution

import (
	"encoding/json"

	"github.com/wheatandcat/dotstamp_server/controllers"
	"github.com/wheatandcat/dotstamp_server/models"
	"github.com/wheatandcat/dotstamp_server/utils/contribution"
	"github.com/wheatandcat/dotstamp_server/utils/tag"

	validator "gopkg.in/go-playground/validator.v9"
)

// NewController Newコントローラ
type NewController struct {
	controllers.BaseController
}

// NewRequest 新規リクエスト
type NewRequest struct {
	Title      string `form:"title" validate:"min=1,max=100"`
	Body       string `form:"body" validate:"min=1"`
	ViewStatus int    `form:"viewStatus"`
	Tag        string `form:"tag"`
}

// Post 新規登録する
func (c *NewController) Post() {
	userID := c.GetUserID()
	if !c.IsNoLogin(userID) {
		c.ServerLoginNotFound()
		return
	}

	request := NewRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &request); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	tx := models.Begin()

	userContributionID, err := contributions.Add(userID, request.Title, request.Body, request.ViewStatus)
	if err != nil {
		models.Rollback(tx)
		c.ServerError(err, controllers.ErrContributionNew, userID)
		return
	}

	tag := request.Tag
	if tag != "" {
		if err := tags.AddList(int(userContributionID), tag); err != nil {
			models.Rollback(tx)
			c.ServerError(err, controllers.ErrContributionNew, userID)
			return
		}
	}

	models.Commit(tx)

	if request.ViewStatus == models.ViewStatusPublic {
		b, err := contributions.GetSearchWordBody(request.Body)
		if err != nil {
			c.ServerError(err, controllers.ErrContributionNew, userID)
			return
		}

		searchWord := contributions.SearchWord{
			Title: request.Title,
			Body:  b,
			Tag:   request.Tag,
		}

		s := contributions.JoinSearchWord(searchWord)
		if err := contributions.AddSearch(int(userContributionID), s); err != nil {
			c.ServerError(err, controllers.ErrContributionNew, userID)
			return
		}
	}

	c.Data["json"] = userContributionID
	c.ServeJSON()
}
