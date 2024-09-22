package utils

import (
	"fmt"
	"github.com/laojianzi/mdavatar"
	"image"
	"image/png"
	"os"
)

// #2453de rgb(36, 83, 222)
func GenDefaultAvatar(name string, path string) (string, error) {
	//取最后两个
	if len(name) > 6 {
		name = name[len(name)-6:]
	}
	filename := fmt.Sprintf("%s.png", name)
	option := mdavatar.WithAvatarTextHandle(func(s string, enableAsianFontChar bool) string {
		return name
	})
	//判断path是否存在
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// 创建文件夹
		err = os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return "", err
		}
	}
	//path + filename判断文件是否存在
	if _, err := os.Stat(path + filename); err == nil {
		return path + filename, nil
	}

	rgba := image.NewRGBA(image.Rect(0, 0, 512, 512))
	toRGBA, err2 := hexToRGBA("#2453de")
	if err2 != nil {
		return "", err2
	}
	rgba.Set(0, 0, toRGBA)
	for y := 0; y < 512; y++ {
		for x := 0; x < 512; x++ {
			rgba.Set(x, y, toRGBA) // 在每个像素上设置蓝色
		}
	}
	size := mdavatar.WithAvatarSize(512)
	background := mdavatar.WithBackground(rgba)
	avatar, err := mdavatar.New(name, background, option, size).Build()
	if err != nil {
		return "", err
	}
	file, err := os.Create(path + filename)
	if err != nil {
		return "", err
	}
	if err := png.Encode(file, avatar); err != nil {
		return "", err
	}
	defer file.Close()
	return file.Name(), nil
}
