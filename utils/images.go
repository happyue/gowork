package utils

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// ImageSize returns image URL of Qiniu image processing style with the specified width and height.
func ImageSize(imageURL string, width, height int) string {
	if strings.Contains(imageURL, "imageView") || !strings.Contains(imageURL, "img.hacpai.com") {
		return imageURL
	}

	return imageURL + "?imageView2/1/w/" + strconv.Itoa(width) + "/h/" + strconv.Itoa(height) + "/interlace/1/q/100"
}

// RandImage returns an image URL randomly for article thumbnail.
// https://github.com/b3log/bing
func RandImage() string {
	min := time.Date(2017, 11, 04, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Now().Unix()
	delta := max - min
	sec := rand.Int63n(delta) + min

	return time.Unix(sec, 0).Format("https://img.hacpai.com/bing/20060102.jpg")
}

// RandImages returns random image URLs.
func RandImages(n int) []string {
	var ret []string

	i := 0
	for {
		if i >= n*5 {
			break
		}

		url := RandImage()
		if !contains(url, ret) {
			ret = append(ret, url)
		}

		if len(ret) >= n {
			return ret
		}

		i++
	}

	return ret
}

func contains(str string, slice []string) bool {
	for _, s := range slice {
		if str == s {
			return true
		}
	}

	return false
}
