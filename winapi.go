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
	procPostMessage         = user32.NewProc("PostMessageA")
	procGetForegroundWindow = user32.NewProc("GetForegroundWindow")
)

const (
	WH_KEYBOARD_LL            = 13
	HC_ACTION                 = 0
	WM_KEYDOWN                = 0x0100
	KFL_SETFORPROCESS         = 0x00000100
	WM_INPUTLANGCHANGEREQUEST = 0x0050
	VK_CAPITAL                = 0x14
)

type KBDLLHOOKSTRUCT struct {
	VkCode      uint32
	ScanCode    uint32
	Flags       uint32
	Time        uint32
	DwExtraInfo uintptr
}

type MSG struct {
	Hwnd    uintptr
	Message uint32
	WParam  uintptr
	LParam  uintptr
	Time    uint32
	Pt      uintptr
}
