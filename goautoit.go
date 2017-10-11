// +build windows
// +build amd64

package goautoit

import (
	"log"
	"syscall"
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
	szTitle := ""
	szText := ""
	szRetText := make([]byte, 256)
	ret, _, lastErr := WGT.Call(uintptr(unsafe.Pointer(&szTitle)), uintptr(unsafe.Pointer(&szText)), uintptr(unsafe.Pointer(&szRetText)), 256)
	log.Println(ret)
	log.Println(lastErr)
	log.Println(string(szRetText))
}
