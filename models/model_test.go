package models

import (
	"dotstamp_server/tests"

	. "gopkg.in/check.v1"
)

type TestModel struct {
	test.TSuite
}

func init() {
	var t test.Accessor = &TestModel{}
	t.SetTableNameList([]string{
		"user_masters",
	})

	var _ = Suite(t)
}

// UserMasterTest ユーザーテスト情報
type UserMasterTest struct {
	ID             uint
	Name           string
	Email          string `validate:"required,email"`
	ProfileImageID int    `json:"profile_image_id"`
}

func (t *TestModel) TestGetWhere(c *C) {
	u := UserMaster{}
	whereList := []map[string]interface{}{
		{"ID": 1},
	}
	option := map[string]interface{}{
		"order":  "ID desc",
		"offset": 0,
		"limit":  1,
	}

	db, _ := GetWhere(&u, "ID = :ID", whereList, option)

	r := u
	c.Check(r.ID, Equals, uint(1))
	c.Check(r.Name, Equals, "abc")
	c.Check(r.Email, Equals, "test@tedt.com")

	ut := UserMasterTest{}
	db.Table("user_masters").Scan(&ut)

	c.Check(ut.ID, Equals, uint(1))
}

func (t *TestModel) TestGetWhereRecordNotFound(c *C) {
	u := UserMaster{}
	whereList := []map[string]interface{}{
		{"ID": 999},
	}
	option := map[string]interface{}{
		"order":  "ID desc",
		"offset": 0,
		"limit":  1,
	}

	_, err := GetWhere(&u, "ID = :ID", whereList, option)
	c.Check(err, IsNil)
	r := u
	c.Check(r.ID, Equals, uint(0))

}

func (t *TestModel) TestGetLisWhere(c *C) {
	u := []UserMaster{}
	whereList := []map[string]interface{}{
		{"ID": 1},
	}
	option := map[string]interface{}{
		"order":  "ID desc",
		"offset": 0,
		"limit":  1,
	}

	db, _ := GetListWhere(&u, "ID = :ID", whereList, option)

	r := u
	c.Check(r[0].ID, Equals, uint(1))
	c.Check(r[0].Name, Equals, "abc")
	c.Check(r[0].Email, Equals, "test@tedt.com")

	ut := []UserMasterTest{}
	db.Table("user_masters").Scan(&ut)

	c.Check(ut[0].ID, Equals, uint(1))
}

func (t *TestModel) TestGetLisWhereRecordNotFound(c *C) {
	u := []UserMaster{}
	whereList := []map[string]interface{}{
		{"ID": 999},
	}
	option := map[string]interface{}{
		"order":  "ID desc",
		"offset": 0,
		"limit":  1,
	}

	_, err := GetListWhere(&u, "ID = :ID", whereList, option)
	c.Check(err, IsNil)
}

func (t *TestModel) TestCreate(c *C) {
	u := UserMaster{
		Name:           "abcdef",
		Email:          "abc@com",
		Password:       "xxxx",
		ProfileImageID: 1,
	}

	Create(&u)

	whereList := []map[string]interface{}{
		{"ID": 3},
	}
	option := make(map[string]interface{})

	GetWhere(&u, "ID = :ID", whereList, option)

	r := u
	c.Check(r.ID, Equals, uint(3))
	c.Check(r.Name, Equals, "abcdef")
}

func (t *TestModel) TestSave(c *C) {
	whereList := []map[string]interface{}{
		{"ID": 1},
	}
	option := make(map[string]interface{})
	u := UserMaster{}

	GetWhere(&u, "ID = :ID", whereList, option)

	u.Name = "xyz"

	Save(&u)

	GetWhere(&u, "ID = :ID", whereList, option)

	r := u
	c.Check(r.ID, Equals, uint(1))
	c.Check(r.Name, Equals, "xyz")
}
