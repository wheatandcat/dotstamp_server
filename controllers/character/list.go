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
func (c *ListController) Post() {
	userID := c.GetUserID()
	if !c.IsNoLogin(userID) {
		c.ServerLoginNotFound()
		return
	}

	character, err := characters.GetListByUserID(userID)
	if err != nil {
		c.ServerError(err, controllers.ErrCodeCommon)
		return
	}

	c.Data["json"] = ListResponse{
		Character: character,
	}

	c.ServeJSON()
}
