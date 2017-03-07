package models

import (
	"dotstamp_server/tests"

	. "gopkg.in/check.v1"
)

type TestLogQuestion struct {
	test.TSuite
}

func init() {
	var t test.Accessor = &TestLogQuestion{}
	t.SetTableNameList([]string{
		"log_questions",
	})

	var _ = Suite(t)
}

func (t *TestLogQuestion) TestAdd(c *C) {
	l := &LogQuestion{
		UserID: 1,
		Email:  "abc@test.com",
		Body:   "aaaa",
	}

	r := l.Add()
	c.Check(r, Equals, nil)
}
