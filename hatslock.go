package main

import (
	"syscall"
	"unsafe"

	"github.com/FFFEGO/hatslock/assets"
	"github.com/getlantern/systray"
)

var keyboardHook uintptr

func setUpTray() {
	systray.SetIcon(assets.IconData)
	systray.SetTooltip("hatslock")
	mQuit := systray.AddMenuItem("Quit", "Quit hatslock")
	go func() {
		<-mQuit.ClickedCh
		systray.Quit()
	}()
}

func changeLanguage() {
	ret, _, _ := procGetForegroundWindow.Call()
	procPostMessage.Call(
		ret, uintptr(WM_INPUTLANGCHANGEREQUEST), uintptr(2), uintptr(0))
}

func callback(nCode int, wParam uintptr, kbdStruct *KBDLLHOOKSTRUCT) uintptr {
	if nCode == HC_ACTION && wParam == WM_KEYDOWN {
		if kbdStruct.vkCode == VK_CAPITAL {
			changeLanguage()
			return 1
		}
	}
	ret, _, _ := procCallNextHookEx.Call(
		keyboardHook,
		uintptr(nCode),
		uintptr(wParam),
		uintptr(unsafe.Pointer(kbdStruct)))
	return ret
}

func setUpHook() {
	ret, _, _ := procSetWindowsHookEx.Call(
		uintptr(WH_KEYBOARD_LL),
		uintptr(syscall.NewCallback(callback)),
		uintptr(0),
		uintptr(0),
	)
	keyboardHook = ret
}

func getMessage(msg *MSG) int {
	ret, _, _ := procGetMessage.Call(
		uintptr(unsafe.Pointer(msg)), uintptr(0), uintptr(0), uintptr(0))
	return int(ret)
}

func runLoop() {
	var msg MSG
	for getMessage(&msg) != 0 {
	}
}

func start() {
	setUpTray()
	setUpHook()
	runLoop()
}

func stop() {
	procUnhookWindowsHookEx.Call(keyboardHook)
}

func main() {
	systray.Run(start, stop)
}
