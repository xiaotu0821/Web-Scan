package option

import (
	"fmt"
	"syscall"
	Configs "web_scan/config"
)

var (
	kernel32    *syscall.LazyDLL  = syscall.NewLazyDLL(`kernel32.dll`)
	proc        *syscall.LazyProc = kernel32.NewProc(`SetConsoleTextAttribute`)
	CloseHandle *syscall.LazyProc = kernel32.NewProc(`CloseHandle`)
)

type Color struct {
	Black        int // 黑色
	Blue         int // 蓝色
	Green        int // 绿色
	Cyan         int // 青色
	Red          int // 红色
	Purple       int // 紫色
	Yellow       int // 黄色
	Light_gray   int // 淡灰色（系统默认值）
	Gray         int // 灰色
	Light_blue   int // 亮蓝色
	Light_green  int // 亮绿色
	Light_cyan   int // 亮青色
	Light_red    int // 亮红色
	Light_purple int // 亮紫色
	Light_yellow int // 亮黄色
	White        int // 白色
}

var FontColor Color = Color{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

func InfoPrint(in string) {
	handle, _, _ := proc.Call(uintptr(syscall.Stdout), uintptr(FontColor.Light_cyan))
	Configs.ColorInfo.Println(in)
	CloseHandle.Call(handle)
}

func SuccessPrint(in string) {
	handle, _, _ := proc.Call(uintptr(syscall.Stdout), uintptr(FontColor.Light_green))
	Configs.ColorSuccess.Println(in)
	CloseHandle.Call(handle)
}

func FailPrint(in string) {
	handle, _, _ := proc.Call(uintptr(syscall.Stdout), uintptr(FontColor.Light_yellow))
	Configs.ColorFail.Println(in)
	CloseHandle.Call(handle)
}

func MistakPrint(in string) {
	handle, _, _ := proc.Call(uintptr(syscall.Stdout), uintptr(FontColor.Light_red))
	Configs.ColorMistake.Println(in)
	CloseHandle.Call(handle)
}

func SendPrint(in string) {
	handle, _, _ := proc.Call(uintptr(syscall.Stdout), uintptr(FontColor.White))
	Configs.ColorSend.Println(in)
	CloseHandle.Call(handle)
}

func OrgPrint(in string, i int) {
	handle, _, _ := proc.Call(uintptr(syscall.Stdout), uintptr(i))
	fmt.Println(in)
	CloseHandle.Call(handle)
}
