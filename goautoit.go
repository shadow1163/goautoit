// +build windows
// +build amd64

package goautoit

import (
	"log"
	"syscall"
	"unicode/utf16"
	"unsafe"
)

var dll64 *syscall.LazyDLL
var err error

func init() {
	// dll64, err = syscall.LoadDLL("D:\\Program Files (x86)\\AutoIt3\\AutoItX\\AutoItX3_x64.dll")
	// defer dll64.Release()
	dll64 = syscall.NewLazyDLL("D:\\Program Files (x86)\\AutoIt3\\AutoItX\\AutoItX3_x64.dll")
}

// WinMinimizeAll -- all windows should be minimize
func WinMinimizeAll() {
	WMA := dll64.NewProc("AU3_WinMinimizeAll")
	WMA.Call()
}

//WinMinimizeAllUndo -- undo minimize all windows
func WinMinimizeAllUndo() {
	WMA := dll64.NewProc("AU3_WinMinimizeAllUndo")
	WMA.Call()
}

//WinGetTitle -- get windows title
func WinGetTitle() {
	WGT := dll64.NewProc("AU3_WinGetTitle")
	szTitle := "[active]"
	szText := ""
	// szRetText := ""
	bufSize := 256
	buff := make([]uint16, bufSize)
	// var buff *uint16
	ret, _, lastErr := WGT.Call(strPtr(szTitle), strPtr(szText), uintptr(unsafe.Pointer(&buff[0])), intPtr(bufSize))
	log.Println(ret)
	log.Println(lastErr)
	log.Println(len(buff))
	log.Println(buff[0])
	// log.Println(GoWString(buff))
	pos := findTermChr(buff)
	log.Println(string(utf16.Decode(buff[0:pos])))
	// log.Println(pos)
	// log.Println(len(szRetText))
	// log.Println(GoWString(buff))
}

//WinGetText -- get text in window
func WinGetText() {
	// WGT := dll64.NewProc("AU3_WinGetText")
	// bufSize := 256
	// buff := make([]uint16, bufSize)
	// ret, _, lastErr := WGT.Call()
	// pos := findTermChr(buff)
	// log.Println(ret)
	// log.Println(lastErr)
	// log.Println(string(utf16.Decode(buff[0:pos])))
}

// Run -- Run a windows program
// flag 3(max) 6(min) 9(normal) 0(hide)
func Run() {
	run := dll64.NewProc("AU3_Run")
	szProgram := "notepad.exe"
	ret, _, lastErr := run.Call(strPtr(szProgram), strPtr("."), intPtr(3))
	log.Println(ret)
	log.Println(lastErr)
}

func findTermChr(buff []uint16) int {
	for i, char := range buff {
		if char == 0x0 {
			return i
		}
	}
	panic("not supposed to happen")
}

func intPtr(n int) uintptr {
	return uintptr(n)
}

func strPtr(s string) uintptr {
	return uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(s)))
}

// GoWString -- Convert a *uint16 C string to a Go String
func GoWString(s *uint16) string {
	if s == nil {
		return ""
	}

	p := (*[1<<30 - 1]uint16)(unsafe.Pointer(s))

	// find the string length
	sz := 0
	for p[sz] != 0 {
		sz++
	}

	return string(utf16.Decode(p[:sz:sz]))
}
