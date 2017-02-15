package mail

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

	r := GetBody(b)
	c.Log(string(r))
}

func (t *TestMain) TestGetForgetpasswordBody(c *C) {
	f := ForgetpasswordTemplate{
		URL:   "http://abc.com/?p=aaaaaa&e=bbbbb",
		Host:  "http:/abc.com",
		Email: "test@abc.com",
	}

	m := GetForgetpasswordBody(f)
	b := Body{
		From:    "dotstamplocaltest@gmail.com",
		To:      "dotstamplocaltest2@gmail.com",
		Subject: "[dotstamp]パスワード再設定",
		Message: string(m),
	}

	r := GetBody(b)
	c.Log(string(r))
}

func (t *TestMain) TestGetForgetpasswordURL(c *C) {
	r, _ := GetForgetpasswordURL("test@tedt.com", "abcdef")

	c.Check(r, Not(Equals), "test@tedt.com/abcdef")
}
