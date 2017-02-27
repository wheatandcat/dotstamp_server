package images

import (
	"errors"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"os"
)

// FileSizeMax 最大アップロードファイルサイズ
const FileSizeMax = 600000

// PngToJpeg pngをjpegに変換する
func PngToJpeg(src string, dest string) error {
	var inFile *os.File
	var outFile *os.File
	var img image.Image
	var err error

	if inFile, err = os.Open(src); err != nil {
		return err
	}

	defer inFile.Close()

	if getFormat(inFile) == "jpg" {
		return nil
	}

	fi, err := inFile.Stat()
	if err != nil {
		return err
	}

	if fi.Size() > FileSizeMax {
		return errors.New("max file size over")
	}

	if img, err = png.Decode(inFile); err != nil {
		return err
	}

	img = alphaToWhite(img)

	if outFile, err = os.Create(dest); err != nil {
		return err
	}

	option := &jpeg.Options{Quality: 100}

	if err = jpeg.Encode(outFile, img, option); err != nil {
		return err
	}

	defer outFile.Close()

	return nil
}

func getFormat(file *os.File) string {
	bytes := make([]byte, 4)
	n, _ := file.ReadAt(bytes, 0)
	if n < 4 {
		return ""
	}
	if bytes[0] == 0x89 && bytes[1] == 0x50 && bytes[2] == 0x4E && bytes[3] == 0x47 {
		return "png"
	}
	if bytes[0] == 0xFF && bytes[1] == 0xD8 {
		return "jpg"
	}
	if bytes[0] == 0x47 && bytes[1] == 0x49 && bytes[2] == 0x46 && bytes[3] == 0x38 {
		return "gif"
	}
	if bytes[0] == 0x42 && bytes[1] == 0x4D {
		return "bmp"
	}
	return ""
}

// alphaToWhite 透明色を白色にする
func alphaToWhite(inputImage image.Image) image.Image {
	rect := inputImage.Bounds()
	width := rect.Size().X
	height := rect.Size().Y
	rgba := image.NewRGBA(rect)

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			var col color.RGBA
			// 座標(x,y)のR, G, B, α の値を取得
			r, g, b, a := inputImage.At(x, y).RGBA()

			if a == 0 {
				// 透過色は白色に変換する
				r = 65535
				g = 65535
				b = 65535
				a = 65535
			}

			col.R = uint8(r)
			col.G = uint8(g)
			col.B = uint8(b)
			col.A = uint8(a)
			rgba.Set(x, y, col)
		}
	}

	return rgba.SubImage(rect)
}

// ToPng pngに変換する
func ToPng(src string, dest string) error {
	var inFile *os.File
	var outFile *os.File
	var img image.Image
	var err error

	if inFile, err = os.Open(src); err != nil {
		return err
	}

	defer inFile.Close()

	if img, _, err = image.Decode(inFile); err != nil {
		return err
	}

	if outFile, err = os.Create(dest); err != nil {
		return err
	}

	if err = png.Encode(outFile, img); err != nil {
		return err
	}

	defer outFile.Close()

	return nil
}
