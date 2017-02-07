package controllersLogin

import (
	"dotstamp_server/controllers"
	"dotstamp_server/utils/user"
)

// NewController 新規登録コントローラ
type NewController struct {
	controllers.BaseController
}

// NewRequest 新規リクエスト
type NewRequest struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

// Post 新規ログイン
func (t *NewController) Post() {
	request := NewRequest{}

	if err := t.ParseForm(&request); err != nil {
		t.ServerError(err, controllers.ErrCodeCommon)
		return
	}

	userID, err := user.Add(request.Email, request.Email, request.Password)
	if err != nil {
		t.ServerError(err, controllers.ErrCreateUser)
		return
	}

	t.SetSession("user_id", userID)

	t.Data["json"] = userID

	t.ServeJSON()
}
