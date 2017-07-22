package controllersSound

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/wheatandcat/dotstamp_server/controllers"
	"github.com/wheatandcat/dotstamp_server/models"
	"github.com/wheatandcat/dotstamp_server/utils/contribution"

	validator "gopkg.in/go-playground/validator.v9"
)

// PutRequest リクエスト
type PutRequest struct {
	SoundStatus int `form:"soundStatus" validate:"min=1,max=2"`
}

// PutResponse レスポンス
type PutResponse struct {
	Warning bool   `json:"warning"`
	Message string `json:"message"`
}

// Put 保存する
func (c *MainController) Put() {
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

	request := PutRequest{}
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &request); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	validate := validator.New()
	if err = validate.Struct(request); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	u, err := contributions.GetByUserContributionID(id)
	if err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	if userID != u.UserID {
		c.ServerError(errors.New("diff UserID"), controllers.ErrCodeCommon, userID)
		return
	}

	s, err := contributions.GetSoundByUserContributionID(id)
	if err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	if s.ID == uint(0) {
		c.ServerError(errors.New("is added sound"), controllers.ErrCodeCommon, userID)
		return
	}

	if !contributions.ExistsSound(id) {
		c.ServerError(errors.New("not exists file"), controllers.ErrCodeCommon, userID)
		return
	}

	tx := models.Begin()

	s.SoundStatus = request.SoundStatus
	if err := s.Save(); err != nil {
		models.Rollback(tx)
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	models.Commit(tx)

	c.Data["json"] = PutResponse{
		Warning: false,
		Message: "",
	}

	c.ServeJSON()
}
