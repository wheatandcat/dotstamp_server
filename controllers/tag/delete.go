package controllersTag

import (
	"errors"

	"github.com/wheatandcat/dotstamp_server/controllers"
	"github.com/wheatandcat/dotstamp_server/models"
	"github.com/wheatandcat/dotstamp_server/utils/contribution"
	"github.com/wheatandcat/dotstamp_server/utils/tag"
)

// DeleteController 削除コントローラ
type DeleteController struct {
	controllers.BaseController
}

// DeleteRequest 削除リクエスト
type DeleteRequest struct {
	UserContributionID int `form:"userContributionId"`
	ID                 int `form:"id"`
}

// DeleteResponse 削除レスポンス
type DeleteResponse struct {
	Warning bool       `json:"warning"`
	Message string     `json:"message"`
	Tag     []tags.Tag `json:"tag"`
}

// Post 削除する
func (c *DeleteController) Post() {
	userID := c.GetUserID()
	if !c.IsNoLogin(userID) {
		c.ServerLoginNotFound()
		return
	}

	request := DeleteRequest{}
	if err := c.ParseForm(&request); err != nil {
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

	if err = tags.DeleteByIDAndUserContributionID(request.ID, request.UserContributionID); err != nil {
		models.Rollback(tx)
		c.ServerError(err, controllers.ErrContributionNoUser, userID)
		return
	}

	tagList, err := tags.GetListByUserContributionID(request.UserContributionID)
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

	c.Data["json"] = DeleteResponse{
		Warning: false,
		Message: "",
		Tag:     tagList,
	}

	c.ServeJSON()
}
