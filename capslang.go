package main

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/fffego/capslang/assets"
	"github.com/getlantern/systray"
)

var (
	user32                     = syscall.NewLazyDLL("user32.dll")
	procSetWindowsHookEx       = user32.NewProc("SetWindowsHookExW")
	procCallNextHookEx         = user32.NewProc("CallNextHookEx")
	procUnhookWindowsHookEx    = user32.NewProc("UnhookWindowsHookEx")
	procGetKeyboardLayout      = user32.NewProc("GetKeyboardLayout")
	procActivateKeyboardLayout = user32.NewProc("ActivateKeyboardLayout")
	procGetMessage             = user32.NewProc("GetMessageW")
	procPostMessage            = user32.NewProc("PostMessageA")
	procGetForegroundWindow    = user32.NewProc("GetForegroundWindow")
	keyboardHook               uintptr
)

const (
	wh_keyboard_ll            = 13
	wm_keydown                = 256
	KFL_SETFORPROCESS         = 0x00000100
	WM_INPUTLANGCHANGEREQUEST = 0x0050
)

type KBDLLHOOKSTRUCT struct {
	VkCode      uint32
	ScanCode    uint32
	Flags       uint32
	Time        uint32
	DwExtraInfo uintptr
}

func getMessage(msg uintptr, hwnd uintptr, msgFilterMin uint32, msgFilterMax uint32) int {
	ret, _, _ := procGetMessage.Call(
		uintptr(unsafe.Pointer(msg)),
		uintptr(hwnd),
		uintptr(msgFilterMin),
		uintptr(msgFilterMax))
	return int(ret)
}

func activateKeyboardLayout(hkl uintptr, flag uint) uintptr {
	ret, _, err := procActivateKeyboardLayout.Call(
		uintptr(hkl),
		uintptr(flag))
	fmt.Println(err)
	fmt.Println(ret)
	return ret
}

func callback(nCode int, wparam uintptr, lparam uintptr) uintptr {
	if nCode == 0 && wparam == wm_keydown {
		kbdstruct := (*KBDLLHOOKSTRUCT)(unsafe.Pointer(lparam))
		code := byte(kbdstruct.VkCode)
		if code == 20 {
			fg, _, _ := procGetForegroundWindow.Call()
			procPostMessage.Call(fg, uintptr(WM_INPUTLANGCHANGEREQUEST), uintptr(2), uintptr(0))
			fmt.Println("CAPS")
			return 1
		}
	}
	ret, _, _ := procCallNextHookEx.Call(
		uintptr(keyboardHook), uintptr(nCode), uintptr(wparam), uintptr(lparam),
	)
	return uintptr(ret)
}

func start() {
	systray.SetIcon(assets.IconData)
	systray.SetTooltip("capslang")
	mQuit := systray.AddMenuItem("Quit", "Quit the app")
	go func() {
		<-mQuit.ClickedCh
		systray.Quit()
	}()

	ret, _, _ := procSetWindowsHookEx.Call(
		uintptr(wh_keyboard_ll),
		uintptr(syscall.NewCallback(callback)),
		uintptr(0),
		uintptr(0),
	)

	keyboardHook = ret

	var msg uintptr
	for getMessage(msg, 0, 0, 0) != 0 {
	}
}

func stop() {
	procUnhookWindowsHookEx.Call(uintptr(keyboardHook))
}

func main() {
	systray.Run(start, stop)
}
