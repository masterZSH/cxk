package gif

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"time"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

// InitGif 初始化gif图片
func InitGif(width, height int, bgColor color.RGBA) *image.NRGBA {
	img := image.NewNRGBA(image.Rect(0, 0, width, height))
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, bgColor)
		}
	}
	return img
}

// CalcBrightness 计算灰度值
func CalcBrightness(img image.Image, rect image.Rectangle) float64 {
	var averageBrightness float64
	width, height := rect.Max.X-rect.Min.X, rect.Max.Y-rect.Min.Y
	var brightness float64
	for x := rect.Min.X; x < rect.Max.X; x++ {
		for y := rect.Min.Y; y < rect.Max.Y; y++ {
			r, g, b, _ := img.At(x, y).RGBA()
			brightness = float64(r>>8+g>>8+b>>8) / 3
			averageBrightness += brightness
		}
	}
	averageBrightness /= float64(width * height)
	return averageBrightness
}

// GetCharByBrightness 获取字符
func GetCharByBrightness(chars []string, brightness float64) string {
	index := int(brightness*float64(len(chars))) >> 8
	if index == len(chars) {
		index--
	}
	return chars[len(chars)-index-1]
}

// AddCharToImage 添加字符到图片
func AddCharToImage(img *image.NRGBA, char string, x, y int, penColor color.RGBA) {
	face := basicfont.Face7x13
	point := fixed.Point26_6{X: fixed.Int26_6(x * 64), Y: fixed.Int26_6(y * 64)}
	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(penColor),
		Face: face,
		Dot:  point,
	}
	d.DrawString(char)
}

// Convert 转换
func Convert(f io.ReadCloser, chars []string, subWidth, subHeight int, imageSwitch bool, bgColor, penColor color.RGBA) int {
	defer f.Close()
	var charsLength int = len(chars)
	if charsLength == 0 {
		return 0
	}
	if subWidth == 0 || subHeight == 0 {
		return 0
	}
	tgif, err := gif.DecodeAll(f)
	if err != nil {
		return 0
	}
	for i, m := range tgif.Image {
		delay := 10 * tgif.Delay[i]
		formatStr := ""
		formatStr += string(delay)
		formatStr += "ms"
		dur, _ := time.ParseDuration(formatStr)
		time.Sleep(dur)
		Draw(m, chars, subWidth, subHeight, imageSwitch, bgColor, penColor)
	}
	return 1
}
