package controllersCharacterImage

import (
	"strconv"
	"dotstamp_server/controllers"
	"dotstamp_server/utils/character"
)

// DeleteController Deleteコントローラ
type DeleteController struct {
	controllers.BaseController
}

// DeleteResponse 削除レスポンス
type DeleteResponse struct {
	Image []characters.Image
}

// Post 画像を削除する
func (t *DeleteController) Post() {
	userID := t.GetUserID()
	if !t.IsNoLogin(userID) {
		t.ServerLoginNotFound()
		return
	}

	id, err := strconv.Atoi(t.Ctx.Input.Param(":id"))
	if err != nil {
		t.ServerError(err, controllers.ErrParameter)
		return
	}

	if err := characters.DeleteByID(id, userID); err != nil {
		t.ServerError(err, controllers.ErrCodeUserNotFound)
		return
	}

	image := characters.GetImageListByUserID(userID)

	t.Data["json"] = UploadResponse{
		Image: image,
	}

	t.ServeJSON()
}
