package controllersSound

import (
	"dotstamp_server/controllers"
	"dotstamp_server/utils/contribution"
	"errors"
)

// MakeController 作成コントローラ
type MakeController struct {
	controllers.BaseController
}

// MakeRequest 作成リクエスト
type MakeRequest struct {
	UserContributionID int `form:"userContributionId" validate:"min=1"`
}

// MakeResponse 作成レスポンス
type MakeResponse struct {
	Warning bool
	Message string
}

// Post 作成する
func (c *MakeController) Post() {
	userID := c.GetUserID()
	if !c.IsNoLogin(userID) {
		c.ServerLoginNotFound()
		return
	}

	request := MakeRequest{}
	if err := c.ParseForm(&request); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon)
		return
	}

	u, err := contributions.GetByUserContributionID(request.UserContributionID)
	if err != nil {
		c.ServerError(err, controllers.ErrCodeCommon)
		return
	}

	if userID != u.UserID {
		c.ServerError(errors.New("diff UserID"), controllers.ErrCodeCommon)
		return
	}

	list, err := contributions.GetSoundDetailListByUserContributionID(request.UserContributionID)
	if err != nil {
		c.ServerError(err, controllers.ErrCodeCommon)
		return
	}

	if err := contributions.MakeSoundFile(request.UserContributionID, list); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon)
		return
	}

	c.Data["json"] = MakeResponse{
		Warning: false,
		Message: "",
	}

	c.ServeJSON()

}
