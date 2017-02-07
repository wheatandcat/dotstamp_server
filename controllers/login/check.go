package controllersLogin

import (
	"dotstamp_server/controllers"
	"dotstamp_server/utils/user"
)

// CheckController 登録確認コントローラ
type CheckController struct {
	controllers.BaseController
}

// CheckRequest 確認リクエスト
type CheckRequest struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

// Post ログイン
func (t *CheckController) Post() {
	request := CheckRequest{}

	if err := t.ParseForm(&request); err != nil {
		t.ServerError(err, controllers.ErrCodeCommon)
		return
	}

	u, err := user.GetByEmailAndPassword(request.Email, request.Password)
	if err != nil {
		t.ServerError(err, controllers.ErrCreateUser)
		return
	}

	t.SetSession("user_id", u.ID)

	t.Data["json"] = true

	t.ServeJSON()
}
