package controllersTag

import (
	"errors"

	"github.com/wheatandcat/dotstamp_server/controllers"
	"github.com/wheatandcat/dotstamp_server/models"
	"github.com/wheatandcat/dotstamp_server/utils/contribution"
	"github.com/wheatandcat/dotstamp_server/utils/tag"

	validator "gopkg.in/go-playground/validator.v9"
)

// AddController 追加コントローラ
type AddController struct {
	controllers.BaseController
}

// AddRequest 追加リクエスト
type AddRequest struct {
	UserContributionID int    `form:"userContributionId"`
	Name               string `form:"name" validate:"min=1,max=20"`
}

// AddResponse 追加レスポンス
type AddResponse struct {
	Warning bool       `json:"warning"`
	Message string     `json:"message"`
	Tag     []tags.Tag `json:"tag"`
}

// Post 追加する
func (c *AddController) Post() {
	userID := c.GetUserID()
	if !c.IsNoLogin(userID) {
		c.ServerLoginNotFound()
		return
	}

	request := AddRequest{}
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
	if err := models.Lock("user_masters", userID); err != nil {
		models.Rollback(tx)
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	contribution, err := contributions.GetByUserContributionID(request.UserContributionID)
	if err != nil {
		models.Rollback(tx)
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	if contribution.ID == uint(0) {
		models.Rollback(tx)
		c.ServerError(errors.New("not found UserContributionID"), controllers.ErrContributionNotFound, userID)
		return
	}

	if contribution.UserID != userID {
		models.Rollback(tx)
		c.ServerError(errors.New("difference UserID"), controllers.ErrContributionNoUser, userID)
		return
	}

	tagList, err := tags.GetListByUserContributionID(request.UserContributionID)
	if err != nil {
		models.Rollback(tx)
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	if len(tagList) > tags.TagMaxNumber {
		models.Rollback(tx)
		c.ServerError(errors.New("max number over tag"), controllers.ErrTagMaxNumberOver, userID)
		return
	}

	for _, tag := range tagList {
		if tag.Name == request.Name {
			models.Rollback(tx)
			c.ServerError(errors.New("tag name overlap"), controllers.ErrTagNameOverlap, userID)
			return
		}
	}

	if err = tags.Add(request.UserContributionID, request.Name); err != nil {
		models.Rollback(tx)
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	tagList, err = tags.GetListByUserContributionID(request.UserContributionID)
	if err != nil {
		models.Rollback(tx)
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	models.Commit(tx)

	if contribution.ViewStatus == models.ViewStatusPublic {
		t, err := tags.GetTagNameJoin(request.UserContributionID)
		if err != nil {
			c.ServerError(err, controllers.ErrContributionSave, userID)
			return
		}

		detail, err := contributions.GetDetailByUserContributionID(request.UserContributionID)
		if err != nil {
			c.ServerError(err, controllers.ErrContributionSave, userID)
			return
		}

		b, err := contributions.GetSearchWordBody(detail.Body)
		if err != nil {
			c.ServerError(err, controllers.ErrContributionNew, userID)
			return
		}

		searchWord := contributions.SearchWord{
			Title: contribution.Title,
			Body:  b,
			Tag:   t,
		}

		s := contributions.JoinSearchWord(searchWord)
		if err := contributions.AddOrSaveSearch(request.UserContributionID, s); err != nil {
			c.ServerError(err, controllers.ErrContributionSave, userID)
			return
		}
	}

	c.Data["json"] = AddResponse{
		Warning: false,
		Message: "",
		Tag:     tagList,
	}

	c.ServeJSON()
}
