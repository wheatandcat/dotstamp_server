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
	UserContributionID int  `form:"userContributionId" validate:"min=1"`
	Overwrite          bool `form:"overwrite"`
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

	request := ReflectRequest{}
	if err := c.ParseForm(&request); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	tx := models.Begin()
	models.Lock("user_masters", userID)

	u, err := contributions.GetByUserContributionID(request.UserContributionID)
	if err != nil {
		models.Rollback(tx)
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	if userID != u.UserID {
		models.Rollback(tx)
		c.ServerError(errors.New("diff UserID"), controllers.ErrCodeCommon, userID)
		return
	}

	image, err := characters.GetImageListByUserID(userID)
	if err != nil {
		models.Rollback(tx)
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	charVoiceMap := map[int]int{}

	for _, v := range image {
		charVoiceMap[int(v.ID)] = v.VoiceType
	}

	body, err := contributions.GetBodyByUserContributionID(request.UserContributionID)
	if err != nil {
		models.Rollback(tx)
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	for k, v := range body {
		body[k].Character.VoiceType = charVoiceMap[v.Character.ID]
	}

	list, err := contributions.GetSoundDetailListByUserContributionID(request.UserContributionID)
	if err != nil {
		models.Rollback(tx)
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	// 既存データ更新
	bodyMap := map[int]contributions.GetBody{}

	for _, v := range body {
		bodyMap[v.Priority] = v
	}

	var bodySound string

	for _, v := range list {
		if bodyMap[v.Priority].Body == v.Body && !request.Overwrite {
			continue
		}

		v.Body = bodyMap[v.Priority].Body
		if request.Overwrite {
			bodySound, err = contributions.ReplaceBodeySound(bodyMap[v.Priority].Body)
			if err != nil {
				models.Rollback(tx)
				c.ServerError(err, controllers.ErrCodeCommon, userID)
				return
			}

			// 元のデータと一致する場合は更新しない
			if bodySound == v.BodySound {
				continue
			}

			v.BodySound = bodySound
			v.MakeStatus = models.MakeStatusUncreated
		}

		if err = v.Save(); err != nil {
			models.Rollback(tx)
			c.ServerError(err, controllers.ErrCodeCommon, userID)
			return
		}
	}

	// 新規データ追加
	if len(body) > len(list) {
		addBody := body[len(list):]

		err = contributions.AddSoundDetailList(request.UserContributionID, addBody)
		if err != nil {
			models.Rollback(tx)
			c.ServerError(err, controllers.ErrCodeCommon, userID)
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
