package controllersFollow

import (
	"github.com/wheatandcat/dotstamp_server/controllers"
	"github.com/wheatandcat/dotstamp_server/utils/contribution"
	"github.com/wheatandcat/dotstamp_server/utils/follow"
)

// ListController リストコントローラ
type ListController struct {
	controllers.BaseController
}

// ListRequest リストリクエスト
type ListRequest struct {
	Order int `form:"order" validate:"min=1,max=2"`
	Page  int `form:"page"`
	Limit int `form:"limit" validate:"min=1,max=50"`
}

// ListResponse リストレスポンス
type ListResponse struct {
	List  []contributions.Contribution
	Count int
}

// Post 追加する
func (c *ListController) Post() {
	userID := c.GetUserID()
	if !c.IsNoLogin(userID) {
		c.ServerLoginNotFound()
		return
	}

	request := ListRequest{}
	if err := c.ParseForm(&request); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	orderMap := map[int]string{
		1: "ID desc",
		2: "ID asc",
	}

	offset := (request.Page - 1) * request.Limit

	orderList, err := follows.GetOrderValueListByUserID(userID, orderMap[request.Order], request.Limit, offset)
	if err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	list, err := contributions.GetListByFollowOrderValue(orderList)
	if err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	list = contributions.ContributionListToPublic(list)

	count, err := follows.GetCountByUserID(userID, orderMap[request.Order])
	if err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	c.Data["json"] = ListResponse{
		List:  list,
		Count: count,
	}
	c.ServeJSON()
}
