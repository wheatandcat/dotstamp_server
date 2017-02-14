package routers

import (
	"dotstamp_server/controllers"
	"dotstamp_server/controllers/character"
	"dotstamp_server/controllers/characterImage"
	"dotstamp_server/controllers/contribution"
	"dotstamp_server/controllers/login"
	"dotstamp_server/controllers/user"
	"dotstamp_server/controllers/user/forget_password"
	"dotstamp_server/controllers/user/profile"
	"dotstamp_server/controllers/work"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/contribution/new/", &controllersContribution.NewController{})
	beego.Router("/contribution/list/", &controllersContribution.ListController{})
	beego.Router("/contribution/upload/", &controllersContribution.UploadController{})
	beego.Router("/contribution/save/", &controllersContribution.SaveController{})
	beego.Router("/contribution/show/:id([0-9]+)", &controllersContribution.ShowController{})
	beego.Router("/contribution/edit/:id([0-9]+)", &controllersContribution.EditController{})
	beego.Router("/contribution/delete/:id([0-9]+)", &controllersContribution.DeleteController{})

	beego.Router("/work/history", &controllersWork.WorkHistoryController{})

	beego.Router("/login/auth/", &controllersLogin.AuthController{})
	beego.Router("/login/new/", &controllersLogin.NewController{})
	beego.Router("/login/check/", &controllersLogin.CheckController{})
	beego.Router("/login/logout/", &controllersLogin.LogoutController{})
	beego.Router("/login/callback/", &controllersLogin.CallbackController{})

	beego.Router("/user/contributionList/", &controllersUser.ContributionListController{})
	beego.Router("/user/save/", &controllersUser.ShowController{})
	beego.Router("/user/show/", &controllersUser.ShowController{})
	beego.Router("/user/profile/upload/", &controllersUserProfile.UploadController{})
	beego.Router("/user/forget_password/add/", &controllersForgetPassword.AddController{})
	beego.Router("/user/forget_password/check/", &controllersForgetPassword.CheckController{})

	beego.Router("/character/list/", &controllersCharacter.ListController{})

	beego.Router("/characterImage/list/", &controllersCharacterImage.ListController{})
	beego.Router("/characterImage/upload/", &controllersCharacterImage.UploadController{})
	beego.Router("/characterImage/delete/:id([0-9]+)", &controllersCharacterImage.DeleteController{})
}
