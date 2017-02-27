package controllers

import "dotstamp_server/utils/image"

// SetImageFileResize リサイズするファイルを指定する
func (c *BaseController) SetImageFileResize(f string, p string, w uint, h uint) (int, error) {
	if isTest() {
		return 0, nil
	}

	tmpPath := "./static/files/tmp/" + p + "/_tmp_" + f
	c.ToFile(tmpPath)

	tmpRootPath := "./static/files/tmp/" + p + "/" + f

	if err := images.PngToJpeg(tmpPath, tmpRootPath); err != nil {
		return ErrImageConversion, err
	}

	outputPath := "./static/files/" + p + "/" + f

	if err := images.Resize(tmpRootPath, outputPath, w, h); err != nil {
		return ErrImageResize, err
	}

	return 0, nil
}

// ToFile ファイルを保存する
func (c *BaseController) ToFile(path string) {
	if isTest() {
		return
	}

	c.SaveToFile("file", path)
}
