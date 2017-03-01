package images

import (
	"image"
	"image/jpeg"
	"os"

	"github.com/nfnt/resize"
)

// Resize 画像リサイズ
func Resize(src string, dest string, width uint, height uint) error {
	file, err := os.Open(src)
	if err != nil {
		return err
	}

	img, err := jpeg.Decode(file)
	if err != nil {
		return err
	}

	widthMax, heightMax, _ := getWidthAndHeight(src, width, height)
	if widthMax == 0 && heightMax == 0 {
		widthMax = width
	}

	defer file.Close()

	m := resize.Resize(widthMax, heightMax, img, resize.Lanczos3)

	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()

	jpeg.Encode(out, m, nil)

	return nil
}

// getWidthAndHeight 幅と高さを取得する
func getWidthAndHeight(src string, width uint, height uint) (uint, uint, error) {
	file, err := os.Open(src)
	if err != nil {
		return 0, 0, err
	}

	conf, _, err := image.DecodeConfig(file)
	if err != nil {
		return 0, 0, err
	}

	if conf.Width > conf.Height {
		height = 0
	} else {
		width = 0
	}

	return width, height, nil
}
