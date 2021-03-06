package controllersUser

import (
	"encoding/json"

	"github.com/wheatandcat/dotstamp_server/controllers"
	"github.com/wheatandcat/dotstamp_server/models"
	"github.com/wheatandcat/dotstamp_server/utils/contribution"
)

// ContributionListController 投稿リストコントローラー
type ContributionListController struct {
	controllers.BaseController
}

// ContributionListRequest 投稿リストリクエスト
type ContributionListRequest struct {
	Order int `form:"order"`
	Page  int `form:"page"`
	Limit int `form:"limit"`
}

// ContributionListResponse 投稿リストレスポンス
type ContributionListResponse struct {
	PrivteList []models.UserContribution `json:"privtes"`
	List       []models.UserContribution `json:"list"`
	TitleList  []string                  `json:"titles"`
	Count      int                       `json:"count"`
}

// Post ユーザー投稿一覧を取得する
func (c *ContributionListController) Post() {
	userID := c.GetUserID()
	if !c.IsNoLogin(userID) {
		c.ServerLoginNotFound()
		return
	}

	request := ContributionListRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &request); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	orderMap := map[int]string{
		1: "ID desc",
		2: "ID asc",
	}

	limit := 1000
	offset := 0

	userContributionlist, err := contributions.GetListByUserID(userID, orderMap[request.Order], limit, offset)
	if err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	title := []string{}

	for _, v := range userContributionlist {
		title = append(title, v.Title)
	}

	privte := []models.UserContribution{}

	for _, v := range userContributionlist {
		if v.ViewStatus == models.ViewStatusPrivate {
			privte = append(privte, v)
		}
	}

	count, err := contributions.GetCountByUserID(userID, orderMap[request.Order])
	if err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	c.Data["json"] = ContributionListResponse{
		PrivteList: privte,
		List:       userContributionlist,
		Count:      count,
		TitleList:  title,
	}
	c.ServeJSON()
}
