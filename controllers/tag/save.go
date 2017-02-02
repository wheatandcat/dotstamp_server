package controllersTag

import (
	"dotstamp_server/controllers"
	"dotstamp_server/utils/tag"
)

// SaveController 保存コントローラ
type SaveController struct {
	controllers.BaseController
}

// SaveResponse 保存レスポンス
type SaveResponse struct {
	Message string
}

// Post 保存する
func (t *SaveController) Post() {
	id, err := t.GetInt("tagId")
	if err != nil {
		t.ServerError(err, controllers.ErrParameter)
		return
	}

	if err := tags.Save(id, t.GetString("tagName")); err != nil {
		t.ServerError(err, controllers.ErrContributionTagSave)
		return
	}

	t.Data["json"] = SaveResponse{
		Message: "",
	}

	t.ServeJSON()
}
