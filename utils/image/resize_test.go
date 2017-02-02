package images

import (
	"dotstamp_server/tests"

	. "gopkg.in/check.v1"
)

type TestResize struct {
	test.TSuite
}

func init() {
	var t test.Accessor = &TestResize{}
	t.SetTableNameList([]string{})

	var _ = Suite(t)
}

func (t *TestResize) TestResize(c *C) {
	imageRoot := "../../tests/files/"
	inputPath := imageRoot + "input/abc.jpg"
	outputPath := imageRoot + "output/abc.jpg"
	err := Resize(inputPath, outputPath, 100, 100)
	if err != nil {
		panic(err)
	}
}
