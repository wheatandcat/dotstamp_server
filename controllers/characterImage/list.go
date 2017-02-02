package controllersCharacterImage

import (
	"dotstamp_server/controllers"
	"dotstamp_server/utils/character"
)

// ListController listコントローラ
type ListController struct {
	controllers.BaseController
}

// ListResponse リストレスポンス
type ListResponse struct {
	Image []characters.Image
}

// Post 一覧を取得する
func (t *ListController) Post() {
	userID := t.GetUserID()
	if !t.IsNoLogin(userID) {
		t.ServerLoginNotFound()
		return
	}

	image := characters.GetImageListByUserID(userID)

	t.Data["json"] = ListResponse{
		Image: image,
	}

	t.ServeJSON()
}
