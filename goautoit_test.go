package goautoit

import (
	"testing"
)

func TestWindows(t *testing.T) {
	// WinMinimizeAll()
	// WinMinimizeAllUndo()

	// WinGetText()
	// Run("notepad.exe", "", SWMaximize)
	Run("notepad.exe", 1)
	// time.Sleep(2 * time.Second)
	// Send("yes PPG", 0)
	// WinGetTitle("", "", 256)
}
