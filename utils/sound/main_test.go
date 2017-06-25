package sound

import (
	"github.com/wheatandcat/dotstamp_server/models/csv_models"
	"github.com/wheatandcat/dotstamp_server/tests"

	. "gopkg.in/check.v1"
)

type TestMain struct {
	test.TSuite
}

func init() {
	var t test.Accessor = &TestMain{}
	t.SetTableNameList([]string{})

	var _ = Suite(t)
}

func (t *TestMain) TestAddSound(c *C) {
	AddTmpSound("こんにちは", "0_1", csvModels.VoiceTypeMeiNormal)
	AddTmpSound("おはようございます", "0_2", csvModels.VoiceTypeMeiAngry)
	AddTmpSound("今日はいい天気ですね", "0_3", csvModels.VoiceTypeMeiBashful)
	AddTmpSound("散歩に行きましょう", "0_4", csvModels.VoiceTypeMeiHappy)

	AddTmpSound("テスト", "0_5", csvModels.VoiceTypeYukkuri)
}

func (t *TestMain) TestJoin(c *C) {
	list := []string{
		"0_1",
		"0_2",
		"0_3",
		"0_4",
	}

	AddTmpSound("こんにちは", list[0], csvModels.VoiceTypeMeiNormal)
	AddTmpSound("おはようございます", list[1], csvModels.VoiceTypeMeiAngry)
	AddTmpSound("今日はいい天気ですね", list[2], csvModels.VoiceTypeMeiBashful)
	AddTmpSound("散歩に行きましょう", list[3], csvModels.VoiceTypeMeiHappy)

	Join(list, "0")
}

func (t *TestMain) TestRemoveDetailFile(c *C) {
	list := []string{
		"0_1",
	}

	AddTmpSound("こんにちは", list[0], csvModels.VoiceTypeMeiNormal)

	r := RemoveDetailFile(list[0])

	c.Check(r, Equals, nil)

	r = RemoveDetailFile(list[0])

	c.Check(r, Not(Equals), nil)
}

func (t *TestMain) TestRemoveJoinFile(c *C) {
	list := []string{
		"0_1",
		"0_2",
	}

	AddTmpSound("こんにちは", list[0], csvModels.VoiceTypeMeiNormal)
	AddTmpSound("おはようございます", list[1], csvModels.VoiceTypeMeiAngry)
	Join(list, "0")

	r := ToM4a("0")
	c.Check(r, Equals, nil)

	r = RemoveJoinFile("0")

	c.Check(r, Equals, nil)
}

func (t *TestMain) TestGetLength(c *C) {
	_, r := GetLength("-1")
	c.Check(r, Equals, nil)

	_, r = GetLength("0")
	c.Check(r, Equals, nil)
}
