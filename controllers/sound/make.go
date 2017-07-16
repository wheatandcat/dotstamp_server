package controllersSound

import (
	"errors"
	"strconv"

	"github.com/wheatandcat/dotstamp_server/controllers"
	"github.com/wheatandcat/dotstamp_server/models"
	"github.com/wheatandcat/dotstamp_server/utils/contribution"
)

// MakeController 作成コントローラ
type MakeController struct {
	controllers.BaseController
}

// MakeResponse 作成レスポンス
type MakeResponse struct {
	Warning bool   `json:"warning"`
	Message string `json:"message"`
}

// Post 作成する
func (c *MakeController) Post() {
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

	u, err := contributions.GetByUserContributionID(id)
	if err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	if userID != u.UserID {
		c.ServerError(errors.New("diff UserID"), controllers.ErrCodeCommon, userID)
		return
	}

	list, err := contributions.GetSoundDetailListByUserContributionID(id)
	if err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	if err := contributions.MakeSoundFile(id, list); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	if err := contributions.UpdateSoundToMakeStatus(id, models.MakeStatusMade); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	c.Data["json"] = MakeResponse{
		Warning: false,
		Message: "",
	}

	c.ServeJSON()

}
