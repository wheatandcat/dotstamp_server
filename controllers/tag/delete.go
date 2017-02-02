package controllersTag

import (
	"dotstamp_server/controllers"
	"dotstamp_server/utils/tag"
)

// DeleteController 削除コントローラ
type DeleteController struct {
	controllers.BaseController
}

// DeleteResponse 削除レスポンス
type DeleteResponse struct {
	Tag []tags.Tag
}

// Post 削除する
func (delte *DeleteController) Post() {
	//id, _ := this.GetInt("tagId")

	delte.Data["json"] = DeleteResponse{}

	delte.ServeJSON()
}
