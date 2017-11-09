package screen

import (
	"image"
	"github.com/antonlahti/go-winapi"
	"syscall"
	"unsafe"
)

var(
	rect = new(image.Rectangle)
)

func ChooseWin() (*image.Rectangle,error) {
	ret,err := rectInfo()
	rect.Min.X=int(ret.Left)
	rect.Min.Y=int(ret.Top)
	rect.Max.X=int(ret.Right-ret.Left)
	rect.Max.Y=int(ret.Bottom-ret.Top)
	return rect,err
}
/**
	获取前端窗体
 */
func rectInfo() (*winapi.RECT,error) {
	var rect winapi.RECT
	moduser32 := syscall.NewLazyDLL("user32.dll")
	procGetFWin := moduser32.NewProc("GetForegroundWindow")
	procGetWinR := moduser32.NewProc("GetWindowRect")
	ret,_,err := procGetFWin.Call()
	procGetWinR.Call(
		ret, uintptr(unsafe.Pointer(&rect)))
	return &rect,err
}