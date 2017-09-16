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
	"github.com/wheatandcat/dotstamp_server/controllers/native"
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

	beego.Router("/api/characters/", &controllersCharacter.MainController{})
	beego.Router("/api/characters/:id([0-9]+)", &controllersCharacter.DeleteController{})

	beego.Router("/api/facebook/oauth/", &controllersFacebook.OauthController{})
	beego.Router("/api/facebook/callback/", &controllersFacebook.CallbackController{})

	beego.Router("/api/follows/:id([0-9]+)", &controllersFollow.MainController{})
	beego.Router("/api/follows/list/", &controllersFollow.ListController{})

	beego.Router("/api/forget_password/", &controllersForgetPassword.MainController{})
	beego.Router("/api/forget_password/check/:email/:keyword", &controllersForgetPassword.CheckController{})

	beego.Router("/api/login/auth/", &controllersLogin.AuthController{})
	beego.Router("/api/login/check/", &controllersLogin.CheckController{})
	beego.Router("/api/login/callback/", &controllersLogin.CallbackController{})

	beego.Router("/api/logout/", &controllersLogin.LogoutController{})

	beego.Router("/api/google/oauth/", &controllersGoogle.OauthController{})
	beego.Router("/api/google/callback/", &controllersGoogle.CallbackController{})

	beego.Router("/api/movies/:id([0-9]+)", &controllersMovie.MainController{})
	beego.Router("/api/movies/:id([0-9]+)/upload/", &controllersMovie.UploadController{})
	beego.Router("/api/movies/connect/:id([0-9]+)", &controllersMovie.ConnectController{})
	beego.Router("/api/movies/callback/", &controllersMovie.CallbackController{})

	beego.Router("/api/native/callback/", &controllersNative.CallbackController{})
	beego.Router("/api/native/dev-callback/", &controllersNative.DevCallbackController{})
	beego.Router("/api/native/redirect/", &controllersNative.RedirectController{})
	beego.Router("/api/native/dev-redirect/", &controllersNative.DevRedirectController{})

	beego.Router("/api/problem/", &controllersProblem.AddController{})

	beego.Router("/api/profile/", &controllersUserProfile.UploadController{})

	beego.Router("/api/question/", &controllersQuestion.AddController{})

	beego.Router("/api/sounds/:id([0-9]+)/", &controllersSound.MainController{})
	beego.Router("/api/sounds/:id([0-9]+)/make/", &controllersSound.MakeController{})
	beego.Router("/api/sounds/:id([0-9]+)/reflect/", &controllersSound.ReflectController{})
	beego.Router("/api/sounds/:id([0-9]+)/voice/all/", &controllersSound.SaveVoiceListController{})
	beego.Router("/api/sounds/body/", &controllersSound.SaveBodyController{})
	beego.Router("/api/sounds/voice/", &controllersSound.SaveVoiceController{})
	beego.Router("/api/sounds/length/", &controllersSound.LengthController{})

	beego.Router("/api/tags/", &controllersTag.MainController{})

	beego.Router("/api/twitter/oauth/", &controllersTwitter.OauthController{})
	beego.Router("/api/twitter/callback/", &controllersTwitter.CallbackController{})

	beego.Router("/api/me/", &controllersUser.MainController{})
	beego.Router("/api/users/new/", &controllersLogin.NewController{})
	beego.Router("/api/users/contribution/list/", &controllersUser.ContributionListController{})

	beego.Router("/*", &controllers.MainController{})
}
