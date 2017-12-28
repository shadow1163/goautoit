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

	INTDEFAULT = -2147483647
)

// HWND -- window handle
type HWND uintptr

var (
	dll64              *syscall.LazyDLL
	winMinimizeAll     *syscall.LazyProc
	winMinimizeAllundo *syscall.LazyProc
	winGetTitle        *syscall.LazyProc
	winGetText         *syscall.LazyProc
	send               *syscall.LazyProc
	run                *syscall.LazyProc
	winWait            *syscall.LazyProc
	controlClick       *syscall.LazyProc
	mouseClick         *syscall.LazyProc
	clipGet            *syscall.LazyProc
	clipPut            *syscall.LazyProc
	winGetHandle       *syscall.LazyProc
	winCloseByHandle   *syscall.LazyProc
	controlSend        *syscall.LazyProc
	controlSetText     *syscall.LazyProc
)

func init() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
	// dll64, err = syscall.LoadDLL("D:\\Program Files (x86)\\AutoIt3\\AutoItX\\AutoItX3_x64.dll")
	// defer dll64.Release()
	dll64 = syscall.NewLazyDLL("lib\\AutoItX3_x64.dll")
	winMinimizeAll = dll64.NewProc("AU3_WinMinimizeAll")
	winMinimizeAllundo = dll64.NewProc("AU3_WinMinimizeAllUndo")
	winGetTitle = dll64.NewProc("AU3_WinGetTitle")
	winGetText = dll64.NewProc("AU3_WinGetText")
	send = dll64.NewProc("AU3_Send")
	run = dll64.NewProc("AU3_Run")
	winWait = dll64.NewProc("AU3_WinWait")
	controlClick = dll64.NewProc("AU3_ControlClick")
	mouseClick = dll64.NewProc("AU3_MouseClick")
	clipGet = dll64.NewProc("AU3_ClipGet")
	clipPut = dll64.NewProc("AU3_ClipPut")
	winGetHandle = dll64.NewProc("AU3_WinGetHandle")
	winCloseByHandle = dll64.NewProc("AU3_WinCloseByHandle")
	controlSend = dll64.NewProc("AU3_ControlSend")
	controlSetText = dll64.NewProc("AU3_ControlSetText")
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
	return (goWString(buff))
}

//WinGetText -- get text in window
func WinGetText(szTitle, szText string, bufSize int) string {
	buff := make([]uint16, int(bufSize))
	winGetText.Call(strPtr(szTitle), strPtr(szText), uintptr(unsafe.Pointer(&buff[0])), intPtr(bufSize))
	return (goWString(buff))
}

// Run -- Run a windows program
// flag 3(max) 6(min) 9(normal) 0(hide)
func Run(szProgram string, args ...interface{}) int {
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
	pid, _, lastErr := run.Call(strPtr(szProgram), strPtr(szDir), intPtr(flag))
	// log.Println(pid)
	if int(pid) == 0 {
		log.Println(lastErr)
	}
	return int(pid)
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

//WinWait -- wait window to active
//
func WinWait(szTitle string, args ...interface{}) int {
	var szText string
	var nTimeout int
	var ok bool
	if len(args) == 0 {
		szText = ""
		nTimeout = 0
	} else if len(args) == 1 {
		if szText, ok = args[0].(string); !ok {
			panic("szText must be a string")
		}
		nTimeout = 0
	} else if len(args) == 2 {
		if szText, ok = args[0].(string); !ok {
			panic("szText must be a string")
		}
		if nTimeout, ok = args[1].(int); !ok {
			panic("nTimeout must be a int")
		}
	} else {
		panic("Too more parameter")
	}

	handle, _, lastErr := winWait.Call(strPtr(szTitle), strPtr(szText), intPtr(nTimeout))
	if int(handle) == 0 {
		log.Print("timeout or failure!!!")
		log.Println(lastErr)
	}
	return int(handle)
}

//MouseClick --
func MouseClick(button string, args ...interface{}) int {
	var x, y, nClicks, nSpeed int
	var ok bool

	if len(args) == 0 {
		x = INTDEFAULT
		y = INTDEFAULT
		nClicks = 1
		nSpeed = 10
	} else if len(args) == 2 {
		if x, ok = args[0].(int); !ok {
			panic("x must be a int")
		}
		if y, ok = args[1].(int); !ok {
			panic("y must be a int")
		}
		nClicks = 1
		nSpeed = 10
	} else if len(args) == 3 {
		if x, ok = args[0].(int); !ok {
			panic("x must be a int")
		}
		if y, ok = args[1].(int); !ok {
			panic("y must be a int")
		}
		if nClicks, ok = args[2].(int); !ok {
			panic("nClicks must be a int")
		}
		nSpeed = 10
	} else if len(args) == 4 {
		if x, ok = args[0].(int); !ok {
			panic("x must be a int")
		}
		if y, ok = args[1].(int); !ok {
			panic("y must be a int")
		}
		if nClicks, ok = args[2].(int); !ok {
			panic("nClicks must be a int")
		}
		if nSpeed, ok = args[3].(int); !ok {
			panic("nSpeed must be a int")
		}
	} else {
		panic("Error parameters")
	}
	ret, _, lastErr := mouseClick.Call(strPtr(button), intPtr(x), intPtr(y), intPtr(nClicks), intPtr(nSpeed))
	if int(ret) != 1 {
		log.Print("failure!!!")
		log.Println(lastErr)
	}
	return int(ret)
}

//ControlClick --
func ControlClick(title, text, control string, args ...interface{}) int {
	var button string
	var x, y, nClicks int
	var ok bool

	if len(args) == 0 {
		button = "left"
		nClicks = 1
		x = INTDEFAULT
		y = INTDEFAULT
	} else if len(args) == 1 {
		if button, ok = args[0].(string); !ok {
			panic("button must be a string")
		}
		nClicks = 1
		x = INTDEFAULT
		y = INTDEFAULT
	} else if len(args) == 2 {
		if button, ok = args[0].(string); !ok {
			panic("button must be a string")
		}
		if nClicks, ok = args[1].(int); !ok {
			panic("nClicks must be a int")
		}
		x = INTDEFAULT
		y = INTDEFAULT
	} else if len(args) == 4 {
		if button, ok = args[0].(string); !ok {
			panic("button must be a string")
		}
		if nClicks, ok = args[1].(int); !ok {
			panic("nClicks must be a int")
		}
		if x, ok = args[2].(int); !ok {
			panic("x must be a int")
		}
		if y, ok = args[3].(int); !ok {
			panic("y must be a int")
		}
	} else {
		panic("Error parameters")
	}
	ret, _, lastErr := controlClick.Call(strPtr(title), strPtr(text), strPtr(control), strPtr(button), intPtr(nClicks), intPtr(x), intPtr(y))
	if int(ret) == 0 {
		log.Print("failure!!!")
		log.Println(lastErr)
	}
	return int(ret)
}

//ClipGet -- get a string from clip
func ClipGet(args ...interface{}) string {
	var nBufSize int
	var ok bool
	if len(args) == 0 {
		nBufSize = 256
	} else if len(args) == 1 {
		if nBufSize, ok = args[0].(int); !ok {
			panic("nBufSize must be a int")
		}
	} else {
		panic("Error parameters")
	}
	clip := make([]uint16, int(nBufSize))
	clipGet.Call(uintptr(unsafe.Pointer(&clip[0])), intPtr(nBufSize))
	return (goWString(clip))
}

// ClipPut -- put a string to clip
func ClipPut(szClip string) int {
	ret, _, lastErr := clipPut.Call(strPtr(szClip))
	if int(ret) == 0 {
		log.Println(lastErr)
	}
	return int(ret)
}

//WinGetHandle -- get window handle
func WinGetHandle(title string, args ...interface{}) HWND {
	var text string
	var ok bool
	if len(args) == 0 {
		text = ""
	} else if len(args) == 1 {
		if text, ok = args[0].(string); !ok {
			panic("text must be a string")
		}
	} else {
		panic("Error parameters")
	}
	ret, _, lastErr := winGetHandle.Call(strPtr(title), strPtr(text))
	if int(ret) == 0 {
		log.Print("failure!!!")
		log.Println(lastErr)
	}
	return HWND(ret)
}

// WinCloseByHandle --
func WinCloseByHandle(hwnd HWND) int {
	ret, _, lastErr := winCloseByHandle.Call(uintptr(hwnd))
	if int(ret) == 0 {
		log.Print("failure!!!")
		log.Println(lastErr)
	}
	return int(ret)
}

//ControlSend --
func ControlSend(title, text, control, sendText string, args ...interface{}) int {
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
	ret, _, lastErr := controlSend.Call(strPtr(title), strPtr(text), strPtr(control), strPtr(sendText), intPtr(nMode))
	if int(ret) == 0 {
		log.Println(lastErr)
	}
	return int(ret)
}

//ControlSetText --
func ControlSetText(title, text, control, newText string) int {
	ret, _, lastErr := controlSetText.Call(strPtr(title), strPtr(text), strPtr(control), strPtr(newText))
	if int(ret) == 0 {
		log.Println(lastErr)
	}
	return int(ret)
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

// GoWString -- Convert a uint16 arrry C string to a Go String
func goWString(s []uint16) string {
	pos := findTermChr(s)
	// log.Println(string(utf16.Decode(s[0:pos])))
	return (string(utf16.Decode(s[0:pos])))
}
