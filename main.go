package main

import (
	_ "image/jpeg"
	_ "image/png"
	"io"
	"log"
	"strings"

	"github.com/masterZSH/cxk/pkg/request"
	"github.com/masterZSH/cxk/pkg/config"

)



func main() {
	// 0 1 and space

	var file io.ReadCloser
	var err error
	file, err = request.GetGifDataByURL()
	if err != nil {
		log.Fatal(err)
	}

	//chars := []string{"M", "8", "0", "V", "1", "i", ":", "*", "|", ".", " "}
	chars := strings.Split(config.characters, "")
	bgColor, penColor := colors[bgColorType], colors[penColorType]
	convert(file, chars, subWidth, subHeight, imageOut, bgColor, penColor)
}
