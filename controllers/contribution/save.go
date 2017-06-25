package controllersContribution

import (
	"github.com/wheatandcat/dotstamp_server/controllers"
	"github.com/wheatandcat/dotstamp_server/models"
	"github.com/wheatandcat/dotstamp_server/utils/contribution"
	"github.com/wheatandcat/dotstamp_server/utils/tag"

	validator "gopkg.in/go-playground/validator.v9"
)

// SaveController Saveコントローラ
type SaveController struct {
	controllers.BaseController
}

// SaveRequest 保存リクエスト
type SaveRequest struct {
	UserContributionID int    `form:"userContributionId"`
	Title              string `form:"title" validate:"min=1,max=100"`
	Body               string `form:"body" validate:"min=1"`
	ViewStatus         int    `form:"viewStatus"`
}

// Post 保存する
func (c *SaveController) Post() {
	userID := c.GetUserID()
	if !c.IsNoLogin(userID) {
		c.ServerLoginNotFound()
		return
	}

	request := SaveRequest{}
	if err := c.ParseForm(&request); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	tx := models.Begin()

	if err := contributions.Save(request.UserContributionID, userID, request.Title, request.ViewStatus); err != nil {
		models.Rollback(tx)
		c.ServerError(err, controllers.ErrContributionSave, userID)
		return
	}

	if err := contributions.SaveDetail(request.UserContributionID, request.Body); err != nil {
		models.Rollback(tx)
		c.ServerError(err, controllers.ErrContributionSave, userID)
		return
	}

	models.Commit(tx)

	if request.ViewStatus == models.ViewStatusPublic {
		t, err := tags.GetTagNameJoin(request.UserContributionID)
		if err != nil {
			c.ServerError(err, controllers.ErrContributionSave, userID)
			return
		}

		b, err := contributions.GetSearchWordBody(request.Body)
		if err != nil {
			c.ServerError(err, controllers.ErrContributionNew, userID)
			return
		}

		searchWord := contributions.SearchWord{
			Title: request.Title,
			Body:  b,
			Tag:   t,
		}

		s := contributions.JoinSearchWord(searchWord)
		if err := contributions.AddOrSaveSearch(request.UserContributionID, s); err != nil {
			c.ServerError(err, controllers.ErrContributionSave, userID)
			return
		}

	} else {
		if err := contributions.DeleteSearchByUserContributionID(request.UserContributionID); err != nil {
			c.ServerError(err, controllers.ErrContributionSave, userID)
			return
		}
	}

	c.Data["json"] = request.UserContributionID
	c.ServeJSON()
}
