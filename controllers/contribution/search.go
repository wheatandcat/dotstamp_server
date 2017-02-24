package controllersContribution

import (
	"dotstamp_server/controllers"
	"dotstamp_server/utils/contribution"
)

// SearchController 検索コントローラ
type SearchController struct {
	controllers.BaseController
}

// SearchRequest 検索リクエスト
type SearchRequest struct {
	Search string `form:"search"`
	Order  int    `form:"order"`
	Page   int    `form:"page"`
	Limit  int    `form:"limit"`
}

// SearchResponse 検索レスポンス
type SearchResponse struct {
	List  []contributions.Contribution
	Count int
}

// Post 検索を取得を取得する
func (c *SearchController) Post() {
	request := SearchRequest{}
	if err := c.ParseForm(&request); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon)
		return
	}

	orderMap := map[int]string{
		1: "ID desc",
		2: "ID asc",
	}

	offset := (request.Page - 1) * request.Limit

	searchValue, err := contributions.GetSearchValueListBySearch(request.Search, orderMap[request.Order], request.Limit, offset)
	if err != nil {
		c.ServerError(err, controllers.ErrContributionSearch)
		return
	}

	if len(searchValue) == 0 {
		c.Data["json"] = []contributions.Contribution{}
		c.ServeJSON()
		return
	}

	contributionlist, err := contributions.GetListBySearchValue(searchValue)
	if err != nil {
		c.ServerError(err, controllers.ErrContributionSearch)
		return
	}

	c.Data["json"] = SearchResponse{
		List:  contributionlist,
		Count: 0,
	}

	c.ServeJSON()
}
