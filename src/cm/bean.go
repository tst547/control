package cm

type Base struct {
	err int
	msg interface{}
}
type IpMsg struct {
	addr string
}
type WinMsg struct {
	isReady bool
	port uint8
}

type FileMsg struct {
	files []File
}
type File struct {
	name string
	size uint32
}
