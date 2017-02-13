package characters

import (
	"dotstamp_server/models"
	"dotstamp_server/tests"

	. "gopkg.in/check.v1"
)

type TestImage struct {
	test.TSuite
}

func init() {
	var t test.Accessor = &TestImage{}
	t.SetTableNameList([]string{
		"user_character_images",
	})

	var _ = Suite(t)
}

func (t *TestImage) TestAddImage(c *C) {
	uID := 10

	AddImage(uID, 1, 1)

	r := GetImageListByUserID(uID)

	c.Check(r[0].CharacterID, Equals, 1)
}

func (t *TestImage) TestGetImageListByUserID(c *C) {

	r := GetImageListByUserID(1)

	c.Check(r[0].CharacterID, Equals, 1)
}

func (t *TestImage) TestDeleteByID(c *C) {
	u := models.UserCharacterImage{}

	e := DeleteByID(1, 2)
	c.Check(e, Equals, e)

	r := u.GetByID(1)
	c.Check(r.ID, Equals, uint(1))

	DeleteByID(1, 1)
	r = u.GetByID(1)

	c.Check(r.ID, Not(Equals), uint(1))
}

func (t *TestImage) TestGetImageName(c *C) {

	r := GetImageName(1)

	c.Check(r, Equals, "6b86b273ff34fce19d6b804eff5a3f5747ada4eaa22f1d49c01e52ddb7875b4b.jpg")
}
