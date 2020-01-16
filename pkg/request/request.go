package request

import (
	"io"
	"net/http"
)

// GetGifDataByURL 获取图片数据
func GetGifDataByURL(gifURL string) (io.ReadCloser, error) {
	var respImgData *http.Response
	respImgData, err := http.Get(gifURL)
	if err != nil {
		return nil, err
	}
	return respImgData.Body, nil
}
