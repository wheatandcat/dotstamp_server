package controllersUser

import (
	"dotstamp_server/controllers"
	"dotstamp_server/utils/user"
)

// ShowController 詳細確認
type ShowController struct {
	controllers.BaseController
}

// Post ユーザー情報
func (t *ShowController) Post() {
	userID := t.GetUserID()
	if !t.IsNoLogin(userID) {
		t.ServerLoginNotFound()
		return
	}

	u, err := user.GetByUserID(userID)
	if err != nil {
		t.ServerError(err, controllers.ErrCodeUserNotFound)
	}

	t.Data["json"] = u
	t.ServeJSON()
}
