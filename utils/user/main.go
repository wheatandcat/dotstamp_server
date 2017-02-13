package user

import (
	"dotstamp_server/models"
	"dotstamp_server/utils"
	"errors"

	"github.com/astaxie/beego"
	"github.com/mitchellh/mapstructure"
	"gopkg.in/go-playground/validator.v9"
)

// GetPassword パスワードを取得する
func GetPassword(pass string) string {
	key := beego.AppConfig.String("loginKey")

	return utils.SrringToEncryption(pass + key)
}

// Add ユーザー登録する
func Add(email string, name string, pass string) (uint, error) {
	u := models.UserMaster{
		Email:    email,
		Name:     name,
		Password: GetPassword(pass),
	}

	validate := validator.New()
	if err := validate.Struct(u); err != nil {
		return 0, err
	}

	return u.GetIDAndAdd()
}

// GetByEmail メールアドレスから取得する
func GetByEmail(email string) models.UserMaster {
	u := &models.UserMaster{}

	return u.GetByEmail(email)
}

// GetByEmailAndPassword メールアドレスとパスワードから取得する
func GetByEmailAndPassword(email string, password string) (u models.UserMaster, err error) {
	u = GetByEmail(email)

	if u.Password != GetPassword(password) {
		return models.UserMaster{}, errors.New("password diffrent")
	}

	return u, nil
}

// GetByUserID ユーザIDから取得する
func GetByUserID(userID int) (User, error) {
	u := &models.UserMaster{}
	userMaster := u.GetByID(userID)

	return userMaster
}

// UpadateToProfileImageID プロフィール画像IDを更新する
func UpadateToProfileImageID(uID int, pID int) error {
	u := &models.UserMaster{}
	userMaster := u.GetByID(uID)

	userMaster.ProfileImageID = pID

	return userMaster.Save()
}

// Upadate 更新する
func Upadate(uID int, n string) error {
	u := &models.UserMaster{}
	userMaster := u.GetByID(uID)

	userMaster.Name = n

	return userMaster.Save()
}

// GetMaptByUserIDList ユーザIDリストからマップを取得する
func GetMaptByUserIDList(userIDList []int) (userMap map[int]User, err error) {
	u := &models.UserMaster{}
	userMaster := u.GetListByIDList(userIDList)

	userList := []User{}
	if err = mapstructure.Decode(utils.StructListToMapList(userMaster), &userList); err != nil {
		return map[int]User{}, err
	}

	userMap = map[int]User{}
	for _, user := range userList {
		userMap[int(user.ID)] = user
	}

	return userMap, nil
}
