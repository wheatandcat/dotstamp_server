package models

import (
	"dotstamp_server/tests"

	. "gopkg.in/check.v1"
)

type TestUserContributionSoundDital struct {
	test.TSuite
}

func init() {
	var t test.Accessor = &TestUserContributionSoundDital{}
	t.SetTableNameList([]string{
		"user_contribution_sound_details",
	})

	var _ = Suite(t)
}

func (t *TestUserContributionSoundDital) TestAdd(c *C) {
	u := &UserContributionSoundDital{
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
