package main

import (
	"image"
	"os"
	"fmt"
	"image/jpeg"
)

func ImageEncodeJPEG(ph string, img image.Image, option int) error {
	// 确保文件父目录存在
	// 打开文件等待写入
	f ,err:= os.Create(ph)
	if err!=nil {
		fmt.Println(err.Error())
	}
	// 保证文件正常关闭
	defer f.Close()
	// 写入文件
	return jpeg.Encode(f, img, &jpeg.Options{option})
}
