package models

import (
	"dotstamp_server/tests"

	. "gopkg.in/check.v1"
)

type TestUserCharacterImage struct {
	test.TSuite
}

func init() {
	var t test.Accessor = &TestUserCharacterImage{}
	t.SetTableNameList([]string{
		"user_character_images",
	})

	var _ = Suite(t)
}

func (t *TestUserCharacterImage) TestAdd(c *C) {
	u := &UserCharacterImage{
		UserID:      1,
		CharacterID: 0,
		Priority:    0,
	}

	u.Add()

	r := u.GetListByUserID(u.UserID)

	c.Check(r[0].CharacterID, Equals, 1)
	c.Check(r[1].CharacterID, Equals, 0)

	c.Check(u.ID, Equals, uint(3))
}

func (t *TestUserCharacterImage) TestGetListByUserID(c *C) {
	u := &UserCharacterImage{}
	r := u.GetListByUserID(2)

	c.Check(r[0].CharacterID, Equals, 2)
}

func (t *TestUserCharacterImage) TestGetByID(c *C) {
	id := 1
	u := &UserCharacterImage{}
	r := u.GetByID(id)

	c.Check(r.ID, Equals, uint(1))
}

func (t *TestUserCharacterImage) TestDelete(c *C) {
	u := &UserCharacterImage{}

	userCharacterImage := u.GetByID(1)
	userCharacterImage.Delete()

	r := u.GetByID(1)

	c.Check(r.ID, Not(Equals), 1)
}
