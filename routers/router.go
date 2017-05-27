package routers

import (
	"dotstamp_server/controllers"
	"dotstamp_server/controllers/bug"
	"dotstamp_server/controllers/characterImage"
	"dotstamp_server/controllers/contribution"
	"dotstamp_server/controllers/facebook"
	"dotstamp_server/controllers/follow"
	"dotstamp_server/controllers/forget_password"
	"dotstamp_server/controllers/google"
	"dotstamp_server/controllers/login"
	"dotstamp_server/controllers/movie"
	"dotstamp_server/controllers/problem"
	"dotstamp_server/controllers/profile"
	"dotstamp_server/controllers/question"
	"dotstamp_server/controllers/sound"
	"dotstamp_server/controllers/tag"
	"dotstamp_server/controllers/twitter"
	"dotstamp_server/controllers/user"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/api/contribution/new/", &controllersContribution.NewController{})
	beego.Router("/api/contribution/list/", &controllersContribution.ListController{})
	beego.Router("/api/contribution/upload/", &controllersContribution.UploadController{})
	beego.Router("/api/contribution/save/", &controllersContribution.SaveController{})
	beego.Router("/api/contribution/show/:id([0-9]+)", &controllersContribution.ShowController{})
	beego.Router("/api/contribution/edit/:id([0-9]+)", &controllersContribution.EditController{})
	beego.Router("/api/contribution/delete/:id([0-9]+)", &controllersContribution.DeleteController{})
	beego.Router("/api/contribution/search/", &controllersContribution.SearchController{})

	beego.Router("/api/login/auth/", &controllersLogin.AuthController{})
	beego.Router("/api/login/new/", &controllersLogin.NewController{})
	beego.Router("/api/login/check/", &controllersLogin.CheckController{})
	beego.Router("/api/login/logout/", &controllersLogin.LogoutController{})
	beego.Router("/api/login/callback/", &controllersLogin.CallbackController{})
	beego.Router("/api/google/oauth/", &controllersGoogle.OauthController{})
	beego.Router("/api/google/callback/", &controllersGoogle.CallbackController{})

	beego.Router("/api/twitter/oauth/", &controllersTwitter.OauthController{})
	beego.Router("/api/twitter/callback/", &controllersTwitter.CallbackController{})

	beego.Router("/api/facebook/oauth/", &controllersFacebook.OauthController{})
	beego.Router("/api/facebook/callback/", &controllersFacebook.CallbackController{})

	beego.Router("/api/user/contributionList/", &controllersUser.ContributionListController{})
	beego.Router("/api/user/save/", &controllersUser.SaveController{})
	beego.Router("/api/user/show/", &controllersUser.ShowController{})
	beego.Router("/api/user/profile/upload/", &controllersUserProfile.UploadController{})
	beego.Router("/api/user/forget_password/add/", &controllersForgetPassword.AddController{})
	beego.Router("/api/user/forget_password/check/:email/:keyword", &controllersForgetPassword.CheckController{})
	beego.Router("/api/user/forget_password/save/", &controllersForgetPassword.SaveController{})

	beego.Router("/api/characterImage/list/", &controllersCharacterImage.ListController{})
	beego.Router("/api/characterImage/upload/", &controllersCharacterImage.UploadController{})
	beego.Router("/api/characterImage/delete/:id([0-9]+)", &controllersCharacterImage.DeleteController{})
	beego.Router("/api/characterImage/save/", &controllersCharacterImage.SaveController{})

	beego.Router("/api/follow/add/", &controllersFollow.AddController{})
	beego.Router("/api/follow/delete/", &controllersFollow.DeleteController{})
	beego.Router("/api/follow/list/", &controllersFollow.ListController{})

	beego.Router("/api/tag/add/", &controllersTag.AddController{})
	beego.Router("/api/tag/delete/", &controllersTag.DeleteController{})

	beego.Router("/api/bug/add/", &controllersBug.AddController{})

	beego.Router("/api/problem/add/", &controllersProblem.AddController{})

	beego.Router("/api/sound/add/", &controllersSound.AddController{})
	beego.Router("/api/sound/make/", &controllersSound.MakeController{})
	beego.Router("/api/sound/show/", &controllersSound.ShowController{})
	beego.Router("/api/sound/save/", &controllersSound.SaveController{})
	beego.Router("/api/sound/saveBody/", &controllersSound.SaveBodyController{})
	beego.Router("/api/sound/saveVoice/", &controllersSound.SaveVoiceController{})
	beego.Router("/api/sound/saveVoiceList/", &controllersSound.SaveVoiceListController{})
	beego.Router("/api/sound/reflect/", &controllersSound.ReflectController{})

	beego.Router("/api/movie/make/", &controllersMovie.MakeController{})
	beego.Router("/api/movie/connect/:id([0-9]+)", &controllersMovie.ConnectController{})
	beego.Router("/api/movie/callback/", &controllersMovie.CallbackController{})
	beego.Router("/api/movie/upload/", &controllersMovie.UploadController{})
	beego.Router("/api/movie/check/", &controllersMovie.CheckController{})

	beego.Router("/api/question/add/", &controllersQuestion.AddController{})

	beego.Router("/*", &controllers.MainController{})
}
