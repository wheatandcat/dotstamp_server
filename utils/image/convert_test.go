package images

import (
	"github.com/wheatandcat/dotstamp_server/tests"

	. "gopkg.in/check.v1"
)

type TestConvert struct {
	test.TSuite
}

func init() {
	var t test.Accessor = &TestConvert{}
	t.SetTableNameList([]string{})

	var _ = Suite(t)
}

func (t *TestResize) TestPngToJpeg(c *C) {
	imageRoot := "../../tests/files/"
	inputPath := imageRoot + "input/def.png"
	outputPath := imageRoot + "output/def.jpg"
	r := PngToJpeg(inputPath, outputPath)

	c.Check(r, Equals, nil)

	inputPath = imageRoot + "input/abc.jpg"
	r = PngToJpeg(inputPath, outputPath)

	c.Check(r, Equals, nil)

}
