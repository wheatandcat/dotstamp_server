package csvModels

import (
	"dotstamp_server/tests"

	. "gopkg.in/check.v1"
)

type TestVoiceType struct {
	test.TSuite
}

func init() {
	var t test.Accessor = &TestVoiceType{}
	t.SetTableNameList([]string{})

	var _ = Suite(t)
}

func (t *TestVoiceType) TestGetStructAll(c *C) {

	m := &VoiceType{}

	r, _ := m.GetStructAll()

	c.Check(r[0].ID, Equals, "1")
}

func (t *TestVoiceType) TestGetStruct(c *C) {

	m := &VoiceType{}

	r, _ := m.GetStruct(1)

	c.Check(r.ID, Equals, "1")

	r, _ = m.GetStruct(0)

	c.Check(r.ID, Equals, "")
}
