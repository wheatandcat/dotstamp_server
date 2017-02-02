package controllersCharacter

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
	Character []characters.Character
}

// Post 一覧を取得する
func (t *ListController) Post() {
	userID := t.GetUserID()
	if !t.IsNoLogin(userID) {
		t.ServerLoginNotFound()
		return
	}

	character := characters.GetListByUserID(userID)

	t.Data["json"] = ListResponse{
		Character: character,
	}

	t.ServeJSON()
}
