// +build windows
// +build amd64

package goautoit

import "log"

//MouseClickDrag -- Perform a mouse click and drag operation.
func MouseClickDrag(button string, x1, y1, x2, y2 int, args ...interface{}) int {
	var nSpeed int
	var ok bool

	if len(args) == 0 {
		nSpeed = 10
	} else if len(args) == 1 {
		if nSpeed, ok = args[0].(int); !ok {
			panic("nSpeed must be a int")
		}
	} else {
		panic("Error parameters")
	}
	ret, _, lastErr := mouseClickDrag.Call(strPtr(button), intPtr(x1), intPtr(x2), intPtr(y2), intPtr(nSpeed))
	if int(ret) != 1 {
		log.Print("failure!!!")
		log.Println(lastErr)
	}
	return int(ret)
}

// MouseDown -- Perform a mouse down event at the current mouse position.
func MouseDown(args ...interface{}) int {
	var button string
	var ok bool

	if len(args) == 0 {
		button = DefaultMouseButton
	} else if len(args) == 1 {
		if button, ok = args[0].(string); !ok {
			panic("nSpeed must be a int")
		}
	} else {
		panic("Error parameters")
	}
	ret, _, lastErr := mouseDown.Call(strPtr(button))
	if int(ret) != 1 {
		log.Print("failure!!!")
		log.Println(lastErr)
	}
	return int(ret)
}
