package user

import (
	"errors"

	"github.com/wheatandcat/dotstamp_server/models"
	"github.com/wheatandcat/dotstamp_server/utils"

	"github.com/astaxie/beego"
	"gopkg.in/go-playground/validator.v9"
)

// User ユーザー情報
type User struct {
	ID             uint   `json:"id"`
	Name           string `json:"name"`
	ProfileImageID int    `json:"profileImageID"`
}

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
func GetByEmail(email string) (models.UserMaster, error) {
	u := &models.UserMaster{}

	r, _, err := u.GetByEmail(email)

	return r, err
}

// GetByEmailAndPassword メールアドレスとパスワードから取得する
func GetByEmailAndPassword(email string, password string) (u models.UserMaster, err error) {
	u, err = GetByEmail(email)
	if err != nil {
		return models.UserMaster{}, err
	}

	if u.Password != GetPassword(password) {
		return models.UserMaster{}, errors.New("password diffrent")
	}

	return u, nil
}

// GetByUserID ユーザIDから取得する
func GetByUserID(userID int) (User, error) {
	u := &models.UserMaster{}
	user := User{}

	err := u.GetScanByID(userID, &user)
	if err != nil {
		return user, err
	}

	return user, err
}

// UpadateToProfileImageID プロフィール画像IDを更新する
func UpadateToProfileImageID(uID int, pID int) error {
	u := &models.UserMaster{}
	userMaster, _, err := u.GetByID(uID)
	if err != nil {
		return err
	}

	userMaster.ProfileImageID = pID

	return userMaster.Save()
}

// Upadate 更新する
func Upadate(uID int, n string) error {
	u := &models.UserMaster{}
	userMaster, _, err := u.GetByID(uID)
	if err != nil {
		return err
	}

	userMaster.Name = n

	return userMaster.Save()
}

// GetMaptByUserIDList ユーザIDリストからマップを取得する
func GetMaptByUserIDList(userIDList []int) (map[int]User, error) {
	userMap := map[int]User{}

	u := &models.UserMaster{}
	userList := []User{}
	err := u.GetScanByIDList(userIDList, &userList)
	if err != nil {
		return userMap, err
	}

	for _, user := range userList {
		userMap[int(user.ID)] = user
	}

	return userMap, nil
}

// UpadateToPassword パスワードを更新する
func UpadateToPassword(email string, password string) error {
	u, err := GetByEmail(email)
	if err != nil {
		return err
	}

	u.Password = GetPassword(password)

	return u.Save()
}
