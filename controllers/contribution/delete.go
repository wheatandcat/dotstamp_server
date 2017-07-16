package controllersContribution

import (
	"strconv"

	"github.com/wheatandcat/dotstamp_server/controllers"
	"github.com/wheatandcat/dotstamp_server/models"
	"github.com/wheatandcat/dotstamp_server/utils/contribution"
)

// Delete 削除する
func (c *MainController) Delete() {
	userID := c.GetUserID()
	if !c.IsNoLogin(userID) {
		c.ServerLoginNotFound()
		return
	}

	id, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if err != nil {
		c.ServerError(err, controllers.ErrParameter, userID)
		return
	}

	tx := models.Begin()

	if err = contributions.DeleteByID(id, userID); err != nil {
		models.Rollback(tx)
		c.ServerError(err, controllers.ErrContributionNotFound, userID)
		return
	}

	models.Commit(tx)

	if err = contributions.DeleteSearchByUserContributionID(id); err != nil {
		c.ServerError(err, controllers.ErrContributionSave, userID)
		return
	}

	c.Data["json"] = true

	c.ServeJSON()
}
