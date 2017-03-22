package sound

import (
	"dotstamp_server/tests"

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
	AddTmpSound("こんにちは", "0_1", "mei/mei_normal.htsvoice")
	AddTmpSound("おはようございます", "0_2", "mei/mei_normal.htsvoice")
	AddTmpSound("今日はいい天気ですね", "0_3", "mei/mei_normal.htsvoice")
	AddTmpSound("散歩に行きましょう", "0_4", "mei/mei_normal.htsvoice")
}

func (t *TestMain) TestJoin(c *C) {
	list := []string{
		"0_1",
		"0_2",
		"0_3",
		"0_4",
	}

	AddTmpSound("こんにちは", list[0], "mei/mei_normal.htsvoice")
	AddTmpSound("おはようございます", list[1], "mei/mei_normal.htsvoice")
	AddTmpSound("今日はいい天気ですね", list[2], "mei/mei_normal.htsvoice")
	AddTmpSound("散歩に行きましょう", list[3], "mei/mei_normal.htsvoice")

	Join(list, "0")
}

func (t *TestMain) TestToM4a(c *C) {
	r := ToM4a("0")

	c.Check(r, Equals, nil)
}
