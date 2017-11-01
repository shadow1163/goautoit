// +build windows
// +build amd64

package goautoit

import (
	"log"
	"syscall"
	"unicode/utf16"
	"unsafe"
)

//properties available in AutoItX.
const (
	SWHide            = 0
	SWMaximize        = 3
	SWMinimize        = 6
	SWRestore         = 9
	SWShow            = 5
	SWShowDefault     = 10
	SWShowMaximized   = 3
	SWShowMinimized   = 2
	SWShowminNoActive = 7
	SWShowMa          = 8
	SWShowNoActive    = 4
	SWShowNormal      = 1
)

var (
	dll64              *syscall.LazyDLL
	winMinimizeAll     *syscall.LazyProc
	winMinimizeAllundo *syscall.LazyProc
	winGetTitle        *syscall.LazyProc
	winGetText         *syscall.LazyProc
	send               *syscall.LazyProc
	run                *syscall.LazyProc
)

func init() {
	// dll64, err = syscall.LoadDLL("D:\\Program Files (x86)\\AutoIt3\\AutoItX\\AutoItX3_x64.dll")
	// defer dll64.Release()
	dll64 = syscall.NewLazyDLL("D:\\Program Files (x86)\\AutoIt3\\AutoItX\\AutoItX3_x64.dll")
	winMinimizeAll = dll64.NewProc("AU3_WinMinimizeAll")
	winMinimizeAllundo = dll64.NewProc("AU3_WinMinimizeAllUndo")
	winGetTitle = dll64.NewProc("AU3_WinGetTitle")
	winGetText = dll64.NewProc("AU3_WinGetText")
	send = dll64.NewProc("AU3_Send")
	run = dll64.NewProc("AU3_Run")
}

// WinMinimizeAll -- all windows should be minimize
func WinMinimizeAll() {
	winMinimizeAll.Call()
}

//WinMinimizeAllUndo -- undo minimize all windows
func WinMinimizeAllUndo() {
	winMinimizeAllundo.Call()
}

//WinGetTitle -- get windows title
func WinGetTitle(szTitle, szText string, bufSize int) string {
	// szTitle := "[active]"
	// szText := ""
	// bufSize := 256
	buff := make([]uint16, int(bufSize))
	ret, _, lastErr := winGetTitle.Call(strPtr(szTitle), strPtr(szText), uintptr(unsafe.Pointer(&buff[0])), intPtr(bufSize))
	log.Println(ret)
	log.Println(lastErr)
	pos := findTermChr(buff)
	log.Println(string(utf16.Decode(buff[0:pos])))
	// log.Println(pos)
	return (string(utf16.Decode(buff[0:pos])))
}

//WinGetText -- get text in window
func WinGetText(szTitle, szText string, bufSize int) string {
	buff := make([]uint16, int(bufSize))
	winGetText.Call(strPtr(szTitle), strPtr(szText), uintptr(unsafe.Pointer(&buff[0])), intPtr(bufSize))
	pos := findTermChr(buff)
	log.Println(string(utf16.Decode(buff[0:pos])))
	return (string(utf16.Decode(buff[0:pos])))
}

// Run -- Run a windows program
// flag 3(max) 6(min) 9(normal) 0(hide)
func Run(szProgram string, args ...interface{}) {
	var szDir string
	var flag int
	var ok bool
	if len(args) == 0 {
		szDir = ""
		flag = SWShowNormal
	} else if len(args) == 1 {
		if szDir, ok = args[0].(string); !ok {
			panic("szDir must be a string")
		}
		flag = SWShowNormal
	} else if len(args) == 2 {
		if szDir, ok = args[0].(string); !ok {
			panic("szDir must be a string")
		}
		if flag, ok = args[1].(int); !ok {
			panic("flag must be a int")
		}
	} else {
		panic("Too more parameter")
	}
	run.Call(strPtr(szProgram), strPtr(szDir), intPtr(flag))
}

//Send -- Send simulates input on the keyboard
// flag: 0: normal, 1: raw
func Send(key string, args ...interface{}) {
	var nMode int
	var ok bool
	if len(args) == 0 {
		nMode = 0
	} else if len(args) == 1 {
		if nMode, ok = args[0].(int); !ok {
			panic("nMode must be a int")
		}
	} else {
		panic("Too more parameter")
	}
	send.Call(strPtr(key), intPtr(nMode))
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
