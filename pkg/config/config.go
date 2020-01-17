package config

import (
	"image/color"
)

// Config 配置结构体
type Config struct {
	// GifURL gif 图片地址
	GifURL string
	// 画的字符
	Characters string
	// 宽度
	SubWidth int
	// 高度
	SubHeight int
	// 输出图片
	ImageOut bool
	// 背景颜色类型
	BgColorType string
	// 画笔颜色
	PenColorType string
}

// Colors 颜色map
var Colors map[string]color.RGBA = map[string]color.RGBA{"black": {0, 0, 0, 255},
	"gray":  {140, 140, 140, 255},
	"red":   {255, 0, 0, 255},
	"green": {0, 128, 0, 255},
	"blue":  {0, 0, 255, 255}}

func GetConfig() (Config, map[string]color.RGBA) {
	var config = Config{
		"https://mynovelsave.oss-cn-beijing.aliyuncs.com/pic/cxk.gif",
		"01 ",
		10,
		10,
		false,
		"black",
		"gray",
	}
	return config, Colors
}
