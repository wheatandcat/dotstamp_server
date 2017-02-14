package maill

import (
	"dotstamp_server/tests"

	. "gopkg.in/check.v1"
)

type TestMain struct {
	test.TSuite
}

func init() {
	var t test.Accessor = &TestMain{}
	t.SetTableNameList([]string{
		"user_contribution_tags",
	})

	var _ = Suite(t)
}

func (t *TestMain) TestSend(c *C) {
	b := Body{
		From:    "dotstamplocaltest@gmail.com",
		To:      "dotstamplocaltest2@gmail.com",
		Subject: "タイトルです",
		Message: "本文です",
	}

	body := GetBody(b)
	r := Send("dotstamplocaltest@gmail.com", body)

	c.Check(r, Equals, nil)
}

func (t *TestMain) TestGetBody(c *C) {
	b := Body{
		From:    "dotstamplocaltest@gmail.com",
		To:      "dotstamplocaltest2@gmail.com",
		Subject: "タイトルです",
		Message: "本文です",
	}

	GetBody(b)
}
