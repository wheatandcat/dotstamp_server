package models

import (
	"dotstamp_server/tests"

	. "gopkg.in/check.v1"
)

type TestUserContributionSoundDetail struct {
	test.TSuite
}

func init() {
	var t test.Accessor = &TestUserContributionSoundDetail{}
	t.SetTableNameList([]string{
		"user_contribution_sound_details",
	})

	var _ = Suite(t)
}

func (t *TestUserContributionSoundDetail) TestAdd(c *C) {
	u := &UserContributionSoundDetail{
		UserContributionID: 100,
		Priority:           1,
		TalkType:           1,
		Body:               "abc",
		BodySound:          "def",
		VoiceType:          1,
	}

	u.Add()

	r, _, _ := u.GetListByUserContributionID(100)

	c.Check(r[0].UserContributionID, Equals, 100)
}

func (t *TestUserContributionSoundDetail) TestSave(c *C) {
	u := &UserContributionSoundDetail{}

	uc, _, _ := u.GetByID(1)
	uc.Body = "aaaaa"

	r := uc.Save()

	c.Check(r, Equals, nil)
}

func (t *TestUserContributionSoundDetail) TestGetByID(c *C) {
	u := &UserContributionSoundDetail{}

	r, _, _ := u.GetByID(1)

	c.Check(r.ID, Equals, uint(1))
}

func (t *TestUserContributionSoundDetail) TestUpdateToMakeStatusByUserContributionID(c *C) {
	u := &UserContributionSoundDetail{}

	u.UpdateToMakeStatusByUserContributionID(1, 1)

	r, _, _ := u.GetListByUserContributionID(1)

	c.Check(r[0].MakeStatus, Equals, 1)
}

func (t *TestUserContributionSoundDetail) TestGetListByMakeStatusMade(c *C) {
	u := &UserContributionSoundDetail{}

	r, _, _ := u.GetListByMakeStatusMade()

	c.Check(r[0].MakeStatus, Equals, 1)
}
