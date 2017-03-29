package routers

import (
	"dotstamp_server/controllers"
	"dotstamp_server/controllers/bug"
	"dotstamp_server/controllers/characterImage"
	"dotstamp_server/controllers/contribution"
	"dotstamp_server/controllers/follow"
	"dotstamp_server/controllers/login"
	"dotstamp_server/controllers/movie"
	"dotstamp_server/controllers/problem"
	"dotstamp_server/controllers/question"
	"dotstamp_server/controllers/sound"
	"dotstamp_server/controllers/tag"
	"dotstamp_server/controllers/user"
	"dotstamp_server/controllers/user/forget_password"
	"dotstamp_server/controllers/user/profile"

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
	beego.Router("/contribution/search/", &controllersContribution.SearchController{})

	beego.Router("/login/auth/", &controllersLogin.AuthController{})
	beego.Router("/login/new/", &controllersLogin.NewController{})
	beego.Router("/login/check/", &controllersLogin.CheckController{})
	beego.Router("/login/logout/", &controllersLogin.LogoutController{})
	beego.Router("/login/callback/", &controllersLogin.CallbackController{})

	beego.Router("/user/contributionList/", &controllersUser.ContributionListController{})
	beego.Router("/user/save/", &controllersUser.SaveController{})
	beego.Router("/user/show/", &controllersUser.ShowController{})
	beego.Router("/user/profile/upload/", &controllersUserProfile.UploadController{})
	beego.Router("/user/forget_password/add/", &controllersForgetPassword.AddController{})
	beego.Router("/user/forget_password/check/:email/:keyword", &controllersForgetPassword.CheckController{})
	beego.Router("/user/forget_password/save/", &controllersForgetPassword.SaveController{})

	beego.Router("/characterImage/list/", &controllersCharacterImage.ListController{})
	beego.Router("/characterImage/upload/", &controllersCharacterImage.UploadController{})
	beego.Router("/characterImage/delete/:id([0-9]+)", &controllersCharacterImage.DeleteController{})
	beego.Router("/characterImage/save/", &controllersCharacterImage.SaveController{})

	beego.Router("/follow/add/", &controllersFollow.AddController{})
	beego.Router("/follow/delete/", &controllersFollow.DeleteController{})
	beego.Router("/follow/list/", &controllersFollow.ListController{})

	beego.Router("/tag/add/", &controllersTag.AddController{})
	beego.Router("/tag/delete/", &controllersTag.DeleteController{})

	beego.Router("/bug/add/", &controllersBug.AddController{})

	beego.Router("/problem/add/", &controllersProblem.AddController{})

	beego.Router("/sound/add/", &controllersSound.AddController{})
	beego.Router("/sound/make/", &controllersSound.MakeController{})
	beego.Router("/sound/show/", &controllersSound.ShowController{})
	beego.Router("/sound/save/", &controllersSound.SaveController{})
	beego.Router("/sound/saveBody/", &controllersSound.SaveBodyController{})
	beego.Router("/sound/saveVoice/", &controllersSound.SaveVoiceController{})
	beego.Router("/sound/reflect/", &controllersSound.ReflectController{})

	beego.Router("/movie/make/", &controllersMovie.MakeController{})
	beego.Router("/movie/connect/:id([0-9]+)", &controllersMovie.ConnectController{})
	beego.Router("/movie/callback/", &controllersMovie.CallbackController{})
	beego.Router("/movie/upload/", &controllersMovie.UploadController{})
	beego.Router("/movie/check/", &controllersMovie.CheckController{})

	beego.Router("/question/add/", &controllersQuestion.AddController{})
}
