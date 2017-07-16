package controllersCharacter

import (
	"strconv"

	"github.com/wheatandcat/dotstamp_server/controllers"
	"github.com/wheatandcat/dotstamp_server/models"
	"github.com/wheatandcat/dotstamp_server/utils/character"
)

// DeleteController Deleteコントローラ
type DeleteController struct {
	controllers.BaseController
}

// DeleteResponse 削除レスポンス
type DeleteResponse struct {
	Image []characters.Image `json:"images"`
}

// Delete 画像を削除する
func (c *DeleteController) Delete() {
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

	if err = characters.DeleteByID(id, userID); err != nil {
		models.Rollback(tx)
		c.ServerError(err, controllers.ErrCodeUserNotFound, userID)
		return
	}

	models.Commit(tx)

	image, err := characters.GetImageListByUserID(userID)
	if err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	c.Data["json"] = DeleteResponse{
		Image: image,
	}

	c.ServeJSON()
}
