// +build windows
// +build amd64

package goautoit

import (
	"syscall"
)

var dll64 *syscall.LazyDLL

func init() {
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
	// WGT.Call(LPCWSTRPointer(""), LPCWSTRPointer(""), 256, 256)
}
