package models

import (
	"time"
)

// UserMaster ユーザー情報
type UserMaster struct {
	ID             int `beedb:"PK"`
	Name           string
	Email          string `validate:"required,email"`
	Password       string
	ProfileImageID int `sql:"profile_image_id"`
	DeleteFlag     int `sql:"delete_flag"`
	Created        time.Time
	Updated        time.Time
}

// GetIDAndAdd 投稿してIDを取得する
func (u *UserMaster) GetIDAndAdd() (int, error) {
	u.DeleteFlag = DeleteFlagOff
	u.Created = time.Now()
	u.Updated = time.Now()

	if err := Save(u); err != nil {
		return 0, err
	}

	return u.ID, nil
}

// GetByEmail メールアドレスから取得する
func (u *UserMaster) GetByEmail(email string) (userMaster UserMaster) {
	whereList := []map[string]interface{}{
		{"Email": email},
		{"DeleteFlag": DeleteFlagOff},
	}
	option := make(map[string]interface{})

	GetWhere(&userMaster, "Email LIKE :Email AND Delete_flag = :DeleteFlag", whereList, option)

	return
}

// GetByID ユーザ IDから取得する
func (u *UserMaster) GetByID(id int) (userMaster UserMaster) {
	whereList := []map[string]interface{}{
		{"ID": id},
		{"DeleteFlag": DeleteFlagOff},
	}
	option := make(map[string]interface{})

	GetWhere(&userMaster, "ID = :ID AND Delete_flag = :DeleteFlag", whereList, option)

	return
}

// GetListByIDList ユーザ IDリストからリストを取得する
func (u *UserMaster) GetListByIDList(idList []int) (userMaster []UserMaster) {
	whereList := []map[string]interface{}{
		{"ID": idList},
		{"DeleteFlag": DeleteFlagOff},
	}
	option := make(map[string]interface{})

	GetListWhere(&userMaster, "ID IN :ID AND Delete_flag = :DeleteFlag", whereList, option)

	return
}
