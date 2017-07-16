package routers

import (
	"github.com/wheatandcat/dotstamp_server/controllers"
	"github.com/wheatandcat/dotstamp_server/controllers/bug"
	"github.com/wheatandcat/dotstamp_server/controllers/character"
	"github.com/wheatandcat/dotstamp_server/controllers/contribution"
	"github.com/wheatandcat/dotstamp_server/controllers/facebook"
	"github.com/wheatandcat/dotstamp_server/controllers/follow"
	"github.com/wheatandcat/dotstamp_server/controllers/forget_password"
	"github.com/wheatandcat/dotstamp_server/controllers/google"
	"github.com/wheatandcat/dotstamp_server/controllers/login"
	"github.com/wheatandcat/dotstamp_server/controllers/movie"
	"github.com/wheatandcat/dotstamp_server/controllers/problem"
	"github.com/wheatandcat/dotstamp_server/controllers/profile"
	"github.com/wheatandcat/dotstamp_server/controllers/question"
	"github.com/wheatandcat/dotstamp_server/controllers/sound"
	"github.com/wheatandcat/dotstamp_server/controllers/tag"
	"github.com/wheatandcat/dotstamp_server/controllers/twitter"
	"github.com/wheatandcat/dotstamp_server/controllers/user"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/api/bug/", &controllersBug.AddController{})

	beego.Router("/api/contributions/list/:order([0-9]+)", &controllersContribution.ListController{})
	beego.Router("/api/contributions/new/", &controllersContribution.NewController{})
	beego.Router("/api/contributions/:id([0-9]+)", &controllersContribution.MainController{})
	beego.Router("/api/contributions/upload/", &controllersContribution.UploadController{})
	beego.Router("/api/contributions/edit/:id([0-9]+)", &controllersContribution.EditController{})
	beego.Router("/api/contributions/search/", &controllersContribution.SearchController{})

	beego.Router("/api/character/", &controllersCharacter.MainController{})
	beego.Router("/api/character/:id([0-9]+)", &controllersCharacter.DeleteController{})

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

	beego.Router("/api/follow/add/", &controllersFollow.AddController{})
	beego.Router("/api/follow/delete/", &controllersFollow.DeleteController{})
	beego.Router("/api/follow/list/", &controllersFollow.ListController{})

	beego.Router("/api/tag/add/", &controllersTag.AddController{})
	beego.Router("/api/tag/delete/", &controllersTag.DeleteController{})

	beego.Router("/api/problem/add/", &controllersProblem.AddController{})

	beego.Router("/api/sound/add/", &controllersSound.AddController{})
	beego.Router("/api/sound/make/", &controllersSound.MakeController{})
	beego.Router("/api/sound/show/", &controllersSound.ShowController{})
	beego.Router("/api/sound/save/", &controllersSound.SaveController{})
	beego.Router("/api/sound/saveBody/", &controllersSound.SaveBodyController{})
	beego.Router("/api/sound/saveVoice/", &controllersSound.SaveVoiceController{})
	beego.Router("/api/sound/saveVoiceList/", &controllersSound.SaveVoiceListController{})
	beego.Router("/api/sound/reflect/", &controllersSound.ReflectController{})
	beego.Router("/api/sound/length/", &controllersSound.LengthController{})

	beego.Router("/api/movie/make/", &controllersMovie.MakeController{})
	beego.Router("/api/movie/connect/:id([0-9]+)", &controllersMovie.ConnectController{})
	beego.Router("/api/movie/callback/", &controllersMovie.CallbackController{})
	beego.Router("/api/movie/upload/", &controllersMovie.UploadController{})
	beego.Router("/api/movie/check/", &controllersMovie.CheckController{})

	beego.Router("/api/question/add/", &controllersQuestion.AddController{})

	beego.Router("/*", &controllers.MainController{})
}
