package main

import (
	"syscall"
)

var (
	user32                  = syscall.NewLazyDLL("user32.dll")
	procSetWindowsHookEx    = user32.NewProc("SetWindowsHookExW")
	procCallNextHookEx      = user32.NewProc("CallNextHookEx")
	procUnhookWindowsHookEx = user32.NewProc("UnhookWindowsHookEx")
	procGetMessage          = user32.NewProc("GetMessageW")
	procPostMessage         = user32.NewProc("PostMessageW")
	procGetForegroundWindow = user32.NewProc("GetForegroundWindow")
)

const (
	// WH_KEYBOARD_LL is a procedure id to monitor low-level keyboard events
	WH_KEYBOARD_LL = 13
	// HC_ACTION code that passed to the hook callback when wParam and lParam
	// parameters contain information about a keyboard message
	HC_ACTION = 0
	// WM_KEYDOWN message is posted when a nonsystem key is pressed
	WM_KEYDOWN = 0x0100
	// WM_INPUTLANGCHANGEREQUEST message is posted when the user chooses a new
	// input language
	WM_INPUTLANGCHANGEREQUEST = 0x0050
	// VK_CAPITAL is a CAPS LOCK key code
	VK_CAPITAL = 0x14
)

// KBDLLHOOKSTRUCT contains information about a low-level keyboard input event
type KBDLLHOOKSTRUCT struct {
	vkCode      uint32
	scanCode    uint32
	flags       uint32
	time        uint32
	dwExtraInfo uintptr
}

// MSG contains message information from a thread's message queue
type MSG struct {
	hwnd     uintptr
	message  uint32
	wParam   uintptr
	lParam   uintptr
	time     uint32
	pt       uintptr
	lPrivate uint32
}
