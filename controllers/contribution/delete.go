package controllersContribution

import (
	"dotstamp_server/controllers"
	"dotstamp_server/models"
	"dotstamp_server/utils/contribution"
	"strconv"
)

// DeleteController Deleteコントローラ
type DeleteController struct {
	controllers.BaseController
}

// Post 画像を削除する
func (c *DeleteController) Post() {
	userID := c.GetUserID()
	if !c.IsNoLogin(userID) {
		c.ServerLoginNotFound()
		return
	}

	id, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if err != nil {
		c.ServerError(err, controllers.ErrParameter)
		return
	}

	tx := models.Begin()

	if err = contributions.DeleteByID(id, userID); err != nil {
		models.Rollback(tx)
		c.ServerError(err, controllers.ErrContributionNotFound)
		return
	}

	if err = contributions.DeleteSearchByUserContributionID(id); err != nil {
		models.Rollback(tx)
		c.ServerError(err, controllers.ErrContributionSave)
		return
	}

	models.Commit(tx)

	c.Data["json"] = true

	c.ServeJSON()
}
