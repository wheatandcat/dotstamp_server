package controllersSound

import (
	"dotstamp_server/controllers"
	"dotstamp_server/models"
	"dotstamp_server/utils/character"
	"dotstamp_server/utils/contribution"
	"errors"

	validator "gopkg.in/go-playground/validator.v9"
)

// ReflectController 反映コントローラ
type ReflectController struct {
	controllers.BaseController
}

// ReflectRequest 反映リクエスト
type ReflectRequest struct {
	UserContributionID int `form:"userContributionId" validate:"min=1"`
}

// ReflectResponse 反映レスポンス
type ReflectResponse struct {
	Warning bool
	Message string
}

// Post 反映する
func (c *ReflectController) Post() {
	userID := c.GetUserID()
	if !c.IsNoLogin(userID) {
		c.ServerLoginNotFound()
		return
	}

	request := AddRequest{}
	if err := c.ParseForm(&request); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon)
		return
	}

	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon)
		return
	}

	tx := models.Begin()
	models.Lock("user_masters", userID)

	u, err := contributions.GetByUserContributionID(request.UserContributionID)
	if err != nil {
		c.ServerError(err, controllers.ErrCodeCommon)
		return
	}

	if userID != u.UserID {
		c.ServerError(errors.New("diff UserID"), controllers.ErrCodeCommon)
		return
	}

	image, err := characters.GetImageListByUserID(userID)
	if err != nil {
		c.ServerError(err, controllers.ErrCodeCommon)
		return
	}

	charVoiceMap := map[int]int{}

	for _, v := range image {
		charVoiceMap[int(v.ID)] = v.VoiceType
	}

	body, err := contributions.GetBodyByUserContributionID(request.UserContributionID)
	if err != nil {
		c.ServerError(err, controllers.ErrCodeCommon)
		return
	}

	for k, v := range body {
		body[k].Character.VoiceType = charVoiceMap[v.Character.ID]
	}

	list, err := contributions.GetSoundDetailListByUserContributionID(request.UserContributionID)
	if err != nil {
		c.ServerError(err, controllers.ErrCodeCommon)
		return
	}

	// 既存データ更新
	bodyMap := map[int]contributions.GetBody{}

	for _, v := range body {
		bodyMap[v.Priority] = v
	}

	for _, v := range list {
		if bodyMap[v.Priority].Body == v.Body {
			continue
		}

		v.Body = bodyMap[v.Priority].Body
		if err = v.Save(); err != nil {
			models.Rollback(tx)
			c.ServerError(err, controllers.ErrCodeCommon)
			return
		}
	}

	// 新規データ追加
	if len(body) > len(list) {
		addBody := body[len(list):]

		err = contributions.AddSoundDetailList(request.UserContributionID, addBody)
		if err != nil {
			models.Rollback(tx)
			c.ServerError(err, controllers.ErrCodeCommon)
			return
		}
	}

	models.Commit(tx)

	c.Data["json"] = ReflectResponse{
		Warning: false,
		Message: "",
	}

	c.ServeJSON()

}
