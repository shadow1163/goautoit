// +build windows
// +build amd64

package goautoit

import (
	"log"
	"syscall"
)

var dll64 *syscall.DLL
var err error

func init() {
	dll64, err = syscall.LoadDLL("D:\\Program Files (x86)\\AutoIt3\\AutoItX\\AutoItX3_x64.dll")
	defer dll64.Release()
	// dll64 = syscall.NewLazyDLL("D:\\Program Files (x86)\\AutoIt3\\AutoItX\\AutoItX3_x64.dll")
}

// WinMinimizeAll -- all windows should be minimize
func WinMinimizeAll() {
	// WMA := dll64.NewProc("AU3_WinMinimizeAll")
	// WMA.Call()
}

//WinMinimizeAllUndo -- undo minimize all windows
func WinMinimizeAllUndo() {
	// WMA := dll64.NewProc("AU3_WinMinimizeAllUndo")
	// WMA.Call()
}

//WinGetTitle -- get windows title
func WinGetTitle() {
	WGT, err := dll64.FindProc("AU3_WinGetTitle")
	if err != nil {
		log.Fatal(err)
	}
	_, _, lastErr := WGT.Call()
	log.Println(lastErr)
}
