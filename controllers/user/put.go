package controllersUser

import (
	"encoding/json"

	"github.com/wheatandcat/dotstamp_server/controllers"
	"github.com/wheatandcat/dotstamp_server/models"
	"github.com/wheatandcat/dotstamp_server/utils/user"

	validator "gopkg.in/go-playground/validator.v9"
)

// PutRequest 保存リクエスト
type PutRequest struct {
	Name string `form:"name" validate:"min=1,max=100"`
}

// PutResponse 保存レスポンス
type PutResponse struct {
	Success bool `json:"success"`
}

// Put ユーザー情報
func (c *MainController) Put() {
	userID := c.GetUserID()
	if !c.IsNoLogin(userID) {
		c.ServerLoginNotFound()
		return
	}

	request := PutRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &request); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	tx := models.Begin()

	if err := user.Upadate(userID, request.Name); err != nil {
		models.Rollback(tx)
		c.ServerError(err, controllers.ErrUserSave, userID)
		return
	}

	models.Commit(tx)

	c.Data["json"] = PutResponse{
		Success: true,
	}

	c.ServeJSON()
}
