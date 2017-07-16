package controllersContribution

import (
	"strconv"

	"github.com/wheatandcat/dotstamp_server/controllers"
	"github.com/wheatandcat/dotstamp_server/utils/contribution"
)

// ListController listコントローラ
type ListController struct {
	controllers.BaseController
}

// Get 一覧を取得する
func (c *ListController) Get() {
	order, err := strconv.Atoi(c.Ctx.Input.Param(":order"))
	if err != nil {
		c.ServerError(err, controllers.ErrParameter, 0)
		return
	}

	contributionlist, err := contributions.GetListByTop(0, (order+1)*10)

	if err != nil {
		c.ServerError(err, controllers.ErrParameter, 0)
		return
	}

	c.Data["json"] = contributionlist
	c.ServeJSON()
}
