package controllersQuestion

import (
	"encoding/json"

	"github.com/wheatandcat/dotstamp_server/controllers"
	"github.com/wheatandcat/dotstamp_server/utils/question"

	validator "gopkg.in/go-playground/validator.v9"
)

// AddController 追加コントローラ
type AddController struct {
	controllers.BaseController
}

// AddRequest 追加リクエスト
type AddRequest struct {
	Body  string `form:"body" validate:"min=1"`
	Email string `form:"email" validate:"required,email"`
}

// AddResponse 追加レスポンス
type AddResponse struct {
	Warning bool   `json:"warning"`
	Message string `json:"message"`
}

// Post 追加する
func (c *AddController) Post() {
	userID := c.GetUserID()
	if !c.IsNoLogin(userID) {
		userID = 0
	}

	request := AddRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &request); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	if err := question.Add(userID, request.Body, request.Email); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	c.Data["json"] = AddResponse{
		Warning: false,
		Message: "",
	}

	c.ServeJSON()
}
