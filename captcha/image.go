package captcha

import (
	"bytes"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

func createCaptchaImage(key string) io.Reader {
	file, err := os.Open("./assets/captcha/base.png")
	if err != nil {
		log.Fatalf("CAPTCHA | Failed to open file: %s", err.Error())
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		log.Fatalf("CAPTCHA | Failed to decode image: %s", err.Error())
	}
	dst := image.NewRGBA(img.Bounds())
	draw.Draw(dst, dst.Bounds(), img, image.Point{}, draw.Src)

	fontBin, err := ioutil.ReadFile("./assets/captcha/font.ttf")
	if err != nil {
		log.Fatalf("CAPTCHA | Failed to load font: %s", err.Error())
	}
	ft, err := truetype.Parse(fontBin)
	if err != nil {
		log.Fatalf("CAPTCHA | Failed to parse font: %s", err.Error())
	}

	x, y := 600, 325
	dot := fixed.Point26_6{X: fixed.Int26_6(x * 64), Y: fixed.Int26_6(y * 64)}

	d := &font.Drawer{
		Dst: dst,
		Src: image.NewUniform(color.RGBA{0, 0, 0, 255}),
		Face: truetype.NewFace(ft, &truetype.Options{
			Size: 200,
		}),
		Dot: dot,
	}
	d.DrawString(key)

	buf := new(bytes.Buffer)
	png.Encode(buf, dst)
	return buf
}
