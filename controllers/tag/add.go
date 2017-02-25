package controllersTag

import (
	"dotstamp_server/controllers"
	"dotstamp_server/models"
	"dotstamp_server/utils/contribution"
	"dotstamp_server/utils/tag"
)

// AddController 追加コントローラ
type AddController struct {
	controllers.BaseController
}

// AddRequest 追加リクエスト
type AddRequest struct {
	UserContributionID int    `form:"userContributionId"`
	Name               string `form:"name"`
}

// AddResponse 追加レスポンス
type AddResponse struct {
	Warning bool
	Message string
	list    []tags.Tag
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

	contribution, err := contributions.GetByUserContributionID(request.UserContributionID)
	if err != nil {
		c.ServerError(err, controllers.ErrCodeCommon)
		return
	}

	if contribution.ID == uint(0) {
		c.ServerError(err, controllers.ErrContributionNotFound)
		return
	}

	if contribution.UserID != userID {
		c.ServerError(err, controllers.ErrContributionNoUser)
		return
	}

	tagList, err := tags.GetListByUserContributionID(request.UserContributionID)
	if err != nil {
		c.ServerError(err, controllers.ErrCodeCommon)
		return
	}

	if len(tagList) > tags.TagMaxNumber {
		c.ServerError(err, controllers.ErrTagMaxNumberOver)
		return
	}

	for _, tag := range tagList {
		if tag.Name == request.Name {
			c.ServerError(err, controllers.ErrTagNameOverlap)
			return
		}
	}

	if err = tags.Add(request.UserContributionID, request.Name); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon)
		return
	}

	tagList, err = tags.GetListByUserContributionID(request.UserContributionID)
	if err != nil {
		c.ServerError(err, controllers.ErrCodeCommon)
		return
	}

	if contribution.ViewStatus == models.ViewStatusPublic {
		t, err := tags.GetTagNameJoin(request.UserContributionID)
		if err != nil {
			c.ServerError(err, controllers.ErrContributionSave)
			return
		}

		detail, err := contributions.GetDetailByUserContributionID(request.UserContributionID)
		if err != nil {
			c.ServerError(err, controllers.ErrContributionSave)
			return
		}

		b, err := contributions.GetSearchWordBody(detail.Body)
		if err != nil {
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
			c.ServerError(err, controllers.ErrContributionSave)
			return
		}
	}

	c.Data["json"] = AddResponse{
		Warning: false,
		Message: "",
		list:    tagList,
	}

	c.ServeJSON()
}
