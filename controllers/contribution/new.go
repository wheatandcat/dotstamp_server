package controllersContribution

import (
	"dotstamp_server/controllers"
	"dotstamp_server/utils/contribution"
	"dotstamp_server/utils/tag"
)

// NewController Newコントローラ
type NewController struct {
	controllers.BaseController
}

// Post 新規登録する
func (t *NewController) Post() {
	userID := t.GetUserID()
	if !t.IsNoLogin(userID) {
		t.ServerLoginNotFound()
		return
	}

	userContributionID, err := contributions.Add(userID, t.GetString("title"), t.GetString("body"))
	if err != nil {
		t.ServerError(err, controllers.ErrContributionNew)
		return
	}

	tag := t.GetString("tag")
	if tag != "" {
		if err := tags.AddList(userContributionID, tag); err != nil {
			t.ServerError(err, controllers.ErrContributionNew)
			return
		}
	}

	t.Data["json"] = userContributionID
	t.ServeJSON()
}
