package controllers

import (
	"errors"

	"github.com/astaxie/beego"
)

// noUserID ユーザ無しID
const noUserID = 0

// BaseController ベースコントローラ
type BaseController struct {
	beego.Controller
}

// ErrorResponse エラー発生レスポンス
type ErrorResponse struct {
	Message string
	ErrCode int
}

// Accessor ベースアクセサー
type Accessor interface {
	GetUserID() int
	ServerError()
}

const (
	// ErrCodeCommon 汎用エラー
	ErrCodeCommon = 1
	// ErrCodeUserNotFound ユーザ情報が取得できない or 不一致
	ErrCodeUserNotFound = 2
	// ErrCodeLoginNotFound ログインしていない
	ErrCodeLoginNotFound = 3
	// ErrCreateUser ユーザ登録に失敗
	ErrCreateUser = 4
	// ErrParameter パラメータエラー
	ErrParameter = 5
	// ErrImageConversion 画像変換エラー
	ErrImageConversion = 6
	// ErrImageResize 画像縮小エラー
	ErrImageResize = 7
	// ErrContributionNew 投稿失敗
	ErrContributionNew = 8
	// ErrContributionSave 投稿保存失敗
	ErrContributionSave = 9
	// ErrContributionTagSave 投稿タグ保存失敗
	ErrContributionTagSave = 10
	// ErrUserSave ユーザ保存失敗
	ErrUserSave = 11
)

// errResponseMap エラーレスポンスマップ
var errResponseMap = map[int]ErrorResponse{
	ErrCodeCommon: {
		Message: "エラーが発生しました。",
	},
	ErrCodeUserNotFound: {
		Message: "ユーザ情報が取得できませんでした。もう一度、ログインして下さい。",
	},
	ErrCodeLoginNotFound: {
		Message: "この画面は、ログインしていないユーザーに使用できません",
	},
	ErrCreateUser: {
		Message: "ユーザ作成に失敗しました。もう一度登録お願い致します。",
	},
	ErrParameter: {
		Message: "不正なパラメータが送信されました。",
	},
	ErrImageConversion: {
		Message: "画像の変換に失敗しました。",
	},
	ErrImageResize: {
		Message: "画像のリサイズに失敗しました。",
	},
	ErrContributionNew: {
		Message: "投稿失敗しました。",
	},
	ErrContributionSave: {
		Message: "保存に失敗しました。",
	},
	ErrContributionTagSave: {
		Message: "タグ保存に失敗しました。",
	},
	ErrUserSave: {
		Message: "ユーザ保存に失敗しました。",
	},
}

// getErroResponse エラーレスポンスを取得する
func getErroResponse(errCode int) ErrorResponse {

	err := errResponseMap[errCode]
	err.ErrCode = errCode

	return err
}

// IsNoLogin ログインしているか判定する
func (c *BaseController) IsNoLogin(userID int) bool {
	if userID == noUserID {
		return false
	}

	return true
}

// ServerLoginNotFound ログイン無しで観覧できない
func (c *BaseController) ServerLoginNotFound() {
	c.ServerError(errors.New("login not found"), ErrCodeLoginNotFound)
}

// ServerError サーバーエラーにする
func (c *BaseController) ServerError(err error, errCode int) {
	beego.Error("Error :", err.Error())

	c.Ctx.ResponseWriter.WriteHeader(500)
	c.Data["json"] = getErroResponse(errCode)

	c.ServeJSON()
}

// isTest テスト環境か判定する
func isTest() bool {
	if beego.AppConfig.String("runmode") == "test" {
		return true
	}

	return false
}
