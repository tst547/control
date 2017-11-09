package stream

import (
	"screen"
	"time"
	"github.com/vova616/screenshot"
	"fmt"
	"image"
)

func OutPutImage() error {
	rect,err := screen.ChooseWin()
	if err!=nil {
		return err
	}
	for{
		begin := time.Now()
		img, _ := screenshot.CaptureRect(rect)
		fmt.Println(time.Now().Sub(begin).Seconds())
		myImg := image.Image(img)
		ImageEncodeJPEG("c:\\t.jpg",myImg,100)
	}
}