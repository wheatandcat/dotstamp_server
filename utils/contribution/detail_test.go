package contributions

import (
	"dotstamp_server/tests"

	. "gopkg.in/check.v1"
)

type TestDetail struct {
	test.TSuite
}

func init() {
	var t test.Accessor = &TestDetail{}
	t.SetTableNameList([]string{
		"user_contribution_details",
	})

	var _ = Suite(t)
}

func (t *TestDetail) TestSaveDetail(c *C) {
	b := `[{"priority":0,"body":"あああ","character":{"Id":128,"Character_id":0,"Priority":0,"FileName":"2747b7c718564ba5f066f0523b03e17f6a496b06851333d2d59ab6d863225848.jpg","imageType":4},"directionType":1,"talkType":1,"edit":false},{"priority":1,"body":"あああ","character":{"Id":125,"Character_id":0,"Priority":0,"FileName":"0f8ef3377b30fc47f96b48247f463a726a802f62f3faa03d56403751d2f66c67.jpg","imageType":4},"directionType":1,"talkType":1,"edit":false},{"priority":2,"body":"あああ","character":{"Id":126,"Character_id":0,"Priority":0,"FileName":"65a699905c02619370bcf9207f5a477c3d67130ca71ec6f750e07fe8d510b084.jpg","imageType":4},"directionType":1,"talkType":1,"edit":false}]`
	uID := 1

	SaveDetail(uID, b)

	r, _ := GetBodyByUserContributionID(uID)

	c.Check(r[0].Character.ID, Equals, 128)
	c.Check(r[0].Character.FileName, Equals, "128.jpg")
}

func (t *TestDetail) TestStirngToSaveBody(c *C) {
	b := `[{"priority":0,"body":"あああ","character":{"Id":128,"Character_id":0,"Priority":0,"FileName":"2747b7c718564ba5f066f0523b03e17f6a496b06851333d2d59ab6d863225848.jpg","imageType":4},"directionType":1,"talkType":1,"edit":false},{"priority":1,"body":"あああ","character":{"Id":125,"Character_id":0,"Priority":0,"FileName":"0f8ef3377b30fc47f96b48247f463a726a802f62f3faa03d56403751d2f66c67.jpg","imageType":4},"directionType":1,"talkType":1,"edit":false},{"priority":2,"body":"あああ","character":{"Id":126,"Character_id":0,"Priority":0,"FileName":"65a699905c02619370bcf9207f5a477c3d67130ca71ec6f750e07fe8d510b084.jpg","imageType":4},"directionType":1,"talkType":1,"edit":false}]`

	r, _ := StirngToSaveBody(b)

	c.Check(r[0].Body, Equals, "あああ")
	c.Check(r[0].Character.ID, Equals, 128)
}

func (t *TestDetail) TestStirngToGetBody(c *C) {
	b := `[{"priority":0,"body":"あああ","character":{"Id":128,"Character_id":0,"Priority":0,"FileName":"default1.png","imageType":4},"directionType":1,"talkType":1,"edit":false},{"priority":1,"body":"あああ","character":{"Id":125,"Character_id":0,"Priority":0,"imageType":4},"directionType":1,"talkType":1,"edit":false},{"priority":2,"body":"あああ","character":{"Id":126,"Character_id":0,"Priority":0,"imageType":4},"directionType":1,"talkType":1,"edit":false}]`

	r, _ := StirngToGetBody(b)

	c.Check(r[0].Body, Equals, "あああ")
	c.Check(r[0].Character.ID, Equals, 128)
	c.Check(r[0].Character.FileName, Equals, "128.jpg")
}
