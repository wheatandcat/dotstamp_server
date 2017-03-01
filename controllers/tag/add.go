package controllersTag

import (
	"dotstamp_server/controllers"
	"dotstamp_server/models"
	"dotstamp_server/utils/contribution"
	"dotstamp_server/utils/tag"
	"errors"

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
	Warning bool
	Message string
	Tag     []tags.Tag
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
		c.ServerError(err, controllers.ErrCodeCommon)
		return
	}

	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon)
		return
	}

	tx := models.Begin()
	models.Lock("user_masters", userID)

	contribution, err := contributions.GetByUserContributionID(request.UserContributionID)
	if err != nil {
		c.ServerError(err, controllers.ErrCodeCommon)
		return
	}

	if contribution.ID == uint(0) {
		c.ServerError(errors.New("not found UserContributionID"), controllers.ErrContributionNotFound)
		return
	}

	if contribution.UserID != userID {
		c.ServerError(errors.New("difference UserID"), controllers.ErrContributionNoUser)
		return
	}

	tagList, err := tags.GetListByUserContributionID(request.UserContributionID)
	if err != nil {
		c.ServerError(err, controllers.ErrCodeCommon)
		return
	}

	if len(tagList) > tags.TagMaxNumber {
		c.ServerError(errors.New("max number over tag"), controllers.ErrTagMaxNumberOver)
		return
	}

	for _, tag := range tagList {
		if tag.Name == request.Name {
			c.ServerError(errors.New("tag name overlap"), controllers.ErrTagNameOverlap)
			return
		}
	}

	if err = tags.Add(request.UserContributionID, request.Name); err != nil {
		models.Rollback(tx)
		c.ServerError(err, controllers.ErrCodeCommon)
		return
	}

	tagList, err = tags.GetListByUserContributionID(request.UserContributionID)
	if err != nil {
		models.Rollback(tx)
		c.ServerError(err, controllers.ErrCodeCommon)
		return
	}

	if contribution.ViewStatus == models.ViewStatusPublic {
		t, err := tags.GetTagNameJoin(request.UserContributionID)
		if err != nil {
			models.Rollback(tx)
			c.ServerError(err, controllers.ErrContributionSave)
			return
		}

		detail, err := contributions.GetDetailByUserContributionID(request.UserContributionID)
		if err != nil {
			models.Rollback(tx)
			c.ServerError(err, controllers.ErrContributionSave)
			return
		}

		b, err := contributions.GetSearchWordBody(detail.Body)
		if err != nil {
			models.Rollback(tx)
			c.ServerError(err, controllers.ErrContributionNew)
			return
		}

		searchWord := contributions.SearchWord{
			Title: contribution.Title,
			Body:  b,
			Tag:   t,
		}

		s := contributions.JoinSearchWord(searchWord)
		if err := contributions.AddOrSaveSearch(request.UserContributionID, s); err != nil {
			models.Rollback(tx)
			c.ServerError(err, controllers.ErrContributionSave)
			return
		}
	}

	models.Commit(tx)

	c.Data["json"] = AddResponse{
		Warning: false,
		Message: "",
		Tag:     tagList,
	}

	c.ServeJSON()
}
