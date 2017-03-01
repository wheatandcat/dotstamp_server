package controllersContribution

import (
	"dotstamp_server/controllers"
	"dotstamp_server/models"
	"dotstamp_server/utils/contribution"
	"dotstamp_server/utils/tag"

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
	if err := c.ParseForm(&request); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon)
		return
	}

	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon)
		return
	}

	tx := models.Begin()

	userContributionID, err := contributions.Add(userID, request.Title, request.Body, request.ViewStatus)
	if err != nil {
		models.Rollback(tx)
		c.ServerError(err, controllers.ErrContributionNew)
		return
	}

	tag := request.Tag
	if tag != "" {
		if err := tags.AddList(int(userContributionID), tag); err != nil {
			models.Rollback(tx)
			c.ServerError(err, controllers.ErrContributionNew)
			return
		}
	}

	if request.ViewStatus == models.ViewStatusPublic {
		b, err := contributions.GetSearchWordBody(request.Body)
		if err != nil {
			models.Rollback(tx)
			c.ServerError(err, controllers.ErrContributionNew)
			return
		}

		searchWord := contributions.SearchWord{
			Title: request.Title,
			Body:  b,
			Tag:   request.Tag,
		}

		s := contributions.JoinSearchWord(searchWord)
		if err := contributions.AddSearch(int(userContributionID), s); err != nil {
			models.Rollback(tx)
			c.ServerError(err, controllers.ErrContributionNew)
			return
		}
	}

	models.Commit(tx)

	c.Data["json"] = userContributionID
	c.ServeJSON()
}
