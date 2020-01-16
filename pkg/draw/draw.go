package drw

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
)

// Draw 绘画主方法
func Draw(m *image.Paletted, chars []string, subWidth, subHeight int, imageSwitch bool, bgColor, penColor color.RGBA) string {
	imageWidth, imageHeight := m.Bounds().Max.X, m.Bounds().Max.Y
	var img *image.NRGBA
	if imageSwitch {
		img = initGif(imageWidth, imageHeight, bgColor)
	}
	piecesX, piecesY := imageWidth/subWidth, imageHeight/subHeight
	var buff bytes.Buffer
	for y := 0; y < piecesY; y++ {
		offsetY := y * subHeight
		for x := 0; x < piecesX; x++ {
			offsetX := x * subWidth
			averageBrightness := calcBrightness(m, image.Rect(offsetX, offsetY, offsetX+subWidth, offsetY+subHeight))
			char := getCharByBrightness(chars, averageBrightness)
			buff.WriteString(char)
			if img != nil {
				addCharToImage(img, char, x*subWidth, y*subHeight, penColor)
			}
		}
		buff.WriteString("\n")
	}
	fmt.Printf("%s", buff.String())
	return buff.String()
}
