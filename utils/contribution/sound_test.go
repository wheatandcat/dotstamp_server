package contributions

import (
	"dotstamp_server/models"
	"dotstamp_server/tests"

	. "gopkg.in/check.v1"
)

type TestSound struct {
	test.TSuite
}

func init() {
	var t test.Accessor = &TestSound{}
	t.SetTableNameList([]string{
		"user_contributions",
		"user_contribution_sounds",
		"user_contribution_sound_details",
	})

	var _ = Suite(t)
}

func (t *TestSound) TestGetByUserContributionID(c *C) {
	r, _ := GetSoundByUserContributionID(1)

	c.Check(r.UserContributionID, Equals, 1)
}

func (t *TestSound) TestAddSound(c *C) {
	AddSound(100, 1)

	r, _ := GetSoundByUserContributionID(100)

	c.Check(r.UserContributionID, Equals, 100)
}

func (t *TestSound) TestAddSoundDetailList(c *C) {
	char := GetCharacter{
		VoiceType: 1,
	}

	list := []GetBody{
		{
			Body:      "abc",
			TalkType:  models.TalkTypeImage,
			Priority:  1,
			Character: char,
		},
		{
			Body:      "def'a!-jhg",
			TalkType:  models.TalkTypeText,
			Priority:  2,
			Character: char,
		},
	}

	AddSoundDetailList(100, list)
	r, _ := GetSoundDetailListByUserContributionID(100)

	c.Check(r[1].UserContributionID, Equals, 100)
	c.Check(r[1].BodySound, Equals, "defajhg")
}

func (t *TestSound) TestSaveSoundDetailToBodySound(c *C) {
	err := SaveSoundDetailToBodySound(uint(1), "abcdef", 1)
	c.Check(err, Equals, nil)

	err = SaveSoundDetailToBodySound(uint(1), "abcdef", 2)
	c.Check(err, Not(Equals), nil)
}

func (t *TestSound) TestSaveSoundDetailTVoiceType(c *C) {
	err := SaveSoundDetailTVoiceType(uint(1), 1, 1)
	c.Check(err, Equals, nil)

	err = SaveSoundDetailTVoiceType(uint(1), 1, 2)
	c.Check(err, Not(Equals), nil)
}

func (t *TestSound) TestMakeSoundFile(c *C) {
	list := []models.UserContributionSoundDetail{
		{
			UserContributionID: 0,
			Priority:           1,
			VoiceType:          VoiceTypeMeiNormal,
			BodySound:          "今日は雨だ",
		},
		{
			UserContributionID: 0,
			Priority:           2,
			VoiceType:          VoiceTypeMeiAngry,
			BodySound:          "",
		},
		{
			UserContributionID: 0,
			Priority:           3,
			VoiceType:          0,
			BodySound:          "明日は晴れだ",
		},
		{
			UserContributionID: 0,
			Priority:           4,
			VoiceType:          VoiceTypeMeiAngry,
			BodySound:          "でも、そのあと雨だ",
		},
		{
			UserContributionID: 0,
			Priority:           5,
			VoiceType:          VoiceTypeMeiBashful,
			BodySound:          "明後日は曇りだ",
		},
		{
			UserContributionID: 0,
			Priority:           6,
			VoiceType:          VoiceTypeMeiHappy,
			BodySound:          "3日後は晴れる",
		},
		{
			UserContributionID: 0,
			Priority:           7,
			VoiceType:          VoiceTypeMeiSad,
			BodySound:          "来週も晴れるといいな",
		},
		{
			UserContributionID: 0,
			Priority:           8,
			VoiceType:          VoiceTypeM100,
			BodySound:          "来月も晴れるといいな",
		},
	}

	r := MakeSoundFile(0, list)
	c.Check(r, Equals, nil)
}

func (t *TestSound) TestExistsSound(c *C) {
	r := ExistsSound(0)

	c.Check(r, Equals, true)
}

func (t *TestSound) TestUpdateSoundToMakeStatus(c *C) {
	r := UpdateSoundToMakeStatus(1, 1)

	c.Check(r, Equals, nil)
}

func (t *TestSound) TestReplaceBodeySound(c *C) {
	r, _ := ReplaceBodeySound("映画通とガンダム00")

	c.Check(r, Equals, "えいがつうとガンダムダブルオー")
}
