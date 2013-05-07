package win32

import (
	"errors"
	"fmt"
	"syscall"
	"unsafe"
)

var (
	dwmapi                     = syscall.NewLazyDLL("dwmapi.dll")
	pDwmEnableBlurBehindWindow = dwmapi.NewProc("DwmEnableBlurBehindWindow")
)

type DWM_BLURBEHIND struct {
	DwFlags                uint32
	FEnable                int32
	HRgbBlur               uintptr
	FTransitionOnMaximized int32
}

const (
	DWM_BB_ENABLE                = 0x00000001 //     A value for the fEnable member has been specified.
	DWM_BB_BLURREGION            = 0x00000002 //     A value for the hRgnBlur member has been specified.
	DWM_BB_TRANSITIONONMAXIMIZED = 0x00000004 //     A value for the fTransitionOnMaximized member has been specified.
)

func DwmEnableBlurBehindWindow(hwnd HWND, pBlurBehind *DWM_BLURBEHIND) error {
	err := pDwmEnableBlurBehindWindow.Find()
	if err != nil {
		return errors.New("No window blur support: " + err.Error())
	}

	cRet, _, _ := pDwmEnableBlurBehindWindow.Call(uintptr(unsafe.Pointer(hwnd)), uintptr(unsafe.Pointer(pBlurBehind)))
	ret := int32(cRet)

	if ret >= 0 {
		return nil
	} else {
		return errors.New(fmt.Sprintf("No window blur support: DwmEnableBlurBehindWindow(): HRESULT = %d", ret))
	}
}
