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
func (c *SaveController) Post() {
	id, err := c.GetInt("tagId")
	if err != nil {
		c.ServerError(err, controllers.ErrParameter)
		return
	}

	if err := tags.Save(id, c.GetString("tagName")); err != nil {
		c.ServerError(err, controllers.ErrContributionTagSave)
		return
	}

	c.Data["json"] = SaveResponse{
		Message: "",
	}

	c.ServeJSON()
}
