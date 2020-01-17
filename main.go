package main

import (
	_ "image/jpeg"
	_ "image/png"
	"io"
	"log"
	"strings"

	"github.com/masterZSH/cxk/pkg/config"
	"github.com/masterZSH/cxk/pkg/gif"
	"github.com/masterZSH/cxk/pkg/request"
)

func main() {
	// 0 1 and space
	sysConfig, colors := config.GetConfig()
	var file io.ReadCloser
	var err error
	file, err = request.GetGifDataByURL(sysConfig.GifURL)
	if err != nil {
		log.Fatal(err)
	}

	//chars := []string{"M", "8", "0", "V", "1", "i", ":", "*", "|", ".", " "}
	chars := strings.Split(sysConfig.Characters, "")
	bgColor, penColor := colors[sysConfig.BgColorType], colors[sysConfig.PenColorType]
	gif.Convert(file, chars, sysConfig.SubWidth, sysConfig.SubHeight, sysConfig.ImageOut, bgColor, penColor)
}
