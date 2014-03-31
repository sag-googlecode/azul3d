// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// +build windows

package win32

// Windows Type Information:
//  http://en.wikibooks.org/wiki/Windows_Programming/Handles_and_Data_Types
//
// Most windows types are in:
//  C:\mingw\x86_64-w64-mingw32\include\windef.h
//
// Also look at windows.h in the same folder, and wingdi.h for the gdi32 headers.

/*
#define UNICODE
#include <windows.h>

#cgo LDFLAGS: -luser32 -lgdi32 -lkernel32 -lmsimg32

WORD win32_MAKELANGID(USHORT usPrimaryLanguage, USHORT usSubLanguage);

LRESULT CALLBACK win32_WndProcWrapper(HWND hwnd, UINT msg, WPARAM wParam, LPARAM lParam);

WNDPROC win32_WndProcWrapperHandle;

DWORD win32_DEVMODE_dmDisplayFixedOutput(DEVMODE* dm);
POINTL win32_DEVMODE_dmPosition(DEVMODE* dm);

MONITORENUMPROC win32_MonitorEnumProcCallbackHandle;
HOOKPROC win32_LowLevelKeyboardHookCallbackHandle;

LPTSTR macro_MAKEINTRESOURCE(WORD wInteger);

*/
import "C"

import (
	"reflect"
	"sync"
	"unicode/utf16"
	"unsafe"
)

type LPTSTR C.LPTSTR

// Decodes a UTF-16 encoded C.LPTSTR/C.LPWSTR to a UTF-8 encoded Go string.
//
// if cstr == nil: "" is returned
//
// This function does not touch/free the memory held by the cstr parameter.
func LPTSTRToString(cstr C.LPTSTR) string {
	if cstr == nil {
		return ""
	}
	strlen := int(C.wcslen((*C.wchar_t)(unsafe.Pointer(cstr))))

	var wstr []uint16
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&wstr))
	sliceHeader.Cap = strlen
	sliceHeader.Len = strlen
	sliceHeader.Data = uintptr(unsafe.Pointer(cstr))

	return string(utf16.Decode(wstr))
}

// Encodes a UTF-8 encoded Go string to a UTF-16 encoded C.LPTSTR/C.LPWSTR.
//
// if len(g) == 0: nil is returned.
//
// Note: The returned C.LPTSTR should be free'd at some point; it is malloc'd
func StringToLPTSTR(g string) C.LPTSTR {
	if len(g) == 0 {
		return nil
	}

	u16 := utf16.Encode([]rune(g))

	// u16 is uint16 type
	nBytes := C.size_t(len(u16) * 2)

	// Allocate a buffer
	cstr := (C.LPTSTR)(C.calloc(1, nBytes+2)) // +2 for uint16 NULL terminator

	// Memcpy the UTF-16 encoded string into the buffer
	C.memcpy(unsafe.Pointer(cstr), unsafe.Pointer(&u16[0]), nBytes)

	return cstr
}

type HDC C.HDC

// http://msdn.microsoft.com/en-us/library/windows/desktop/dd183569(v=vs.85).aspx
type DISPLAY_DEVICE C.DISPLAY_DEVICE

/*
  typedef struct _DISPLAY_DEVICEA {
    DWORD StateFlags;
    CHAR DeviceID[128];
    CHAR DeviceKey[128];
  } DISPLAY_DEVICEA,*PDISPLAY_DEVICEA,*LPDISPLAY_DEVICEA;

  typedef struct _DISPLAY_DEVICEW {
    WCHAR DeviceString[128];
    DWORD StateFlags;
    WCHAR DeviceID[128];
    WCHAR DeviceKey[128];
  } DISPLAY_DEVICEW,*PDISPLAY_DEVICEW,*LPDISPLAY_DEVICEW;
*/

func (c *DISPLAY_DEVICE) GetDeviceName() string {
	return LPTSTRToString((C.LPTSTR)(unsafe.Pointer(&c.DeviceName)))
}

func (c *DISPLAY_DEVICE) GetDeviceString() string {
	return LPTSTRToString((C.LPTSTR)(unsafe.Pointer(&c.DeviceString)))
}

const (
	DISPLAY_DEVICE_ACTIVE           = C.DISPLAY_DEVICE_ACTIVE
	DISPLAY_DEVICE_ATTACHED         = C.DISPLAY_DEVICE_ATTACHED
	DISPLAY_DEVICE_MIRRORING_DRIVER = C.DISPLAY_DEVICE_MIRRORING_DRIVER
	DISPLAY_DEVICE_MODESPRUNED      = C.DISPLAY_DEVICE_MODESPRUNED
	DISPLAY_DEVICE_PRIMARY_DEVICE   = C.DISPLAY_DEVICE_PRIMARY_DEVICE
	DISPLAY_DEVICE_REMOVABLE        = C.DISPLAY_DEVICE_REMOVABLE
	DISPLAY_DEVICE_VGA_COMPATIBLE   = C.DISPLAY_DEVICE_VGA_COMPATIBLE
)

func (c *DISPLAY_DEVICE) GetStateFlags() DWORD {
	return DWORD(c.StateFlags)
}

func (c *DISPLAY_DEVICE) GetDeviceID() string {
	return LPTSTRToString((C.LPTSTR)(unsafe.Pointer(&c.DeviceID)))
}

func (c *DISPLAY_DEVICE) GetDeviceKey() string {
	return LPTSTRToString((C.LPTSTR)(unsafe.Pointer(&c.DeviceKey)))
}

// http://msdn.microsoft.com/en-us/library/windows/desktop/dd162609(v=vs.85).aspx
func EnumDisplayDevices(lpDevice string, iDevNum, dwFlags DWORD) (ret bool, displayDevice *DISPLAY_DEVICE) {
	dd := new(C.DISPLAY_DEVICE)
	dd.cb = C.DWORD(unsafe.Sizeof(*dd))

	cDevice := StringToLPTSTR(lpDevice)
	defer C.free(unsafe.Pointer(cDevice))

	ret = C.EnumDisplayDevices(cDevice, C.DWORD(iDevNum), (C.PDISPLAY_DEVICE)(unsafe.Pointer(dd)), C.DWORD(dwFlags)) != 0
	displayDevice = (*DISPLAY_DEVICE)(dd)
	return
}

type DEVMODE C.DEVMODE

func NewDEVMODE() *DEVMODE {
	m := DEVMODE{}
	m.dmSize = C.WORD(unsafe.Sizeof(m))
	return &m
}

func (m *DEVMODE) DmDeviceName() string {
	return LPTSTRToString((C.LPTSTR)(unsafe.Pointer(&m.dmDeviceName)))
}

func (m *DEVMODE) DmBitsPerPel() DWORD {
	return DWORD(m.dmBitsPerPel)
}

func (m *DEVMODE) DmPelsWidth() DWORD {
	return DWORD(m.dmPelsWidth)
}

func (m *DEVMODE) DmPelsHeight() DWORD {
	return DWORD(m.dmPelsHeight)
}

func (m *DEVMODE) DmDisplayFrequency() DWORD {
	return DWORD(m.dmDisplayFrequency)
}

const (
	DMDFO_DEFAULT = C.DMDFO_DEFAULT
	DMDFO_CENTER  = C.DMDFO_CENTER
	DMDFO_STRETCH = C.DMDFO_STRETCH
)

func (m *DEVMODE) DmDisplayFixedOutput() DWORD {
	return DWORD(C.win32_DEVMODE_dmDisplayFixedOutput((*C.DEVMODE)(m)))
}

type POINTL C.POINTL

func (c POINTL) X() LONG {
	return LONG(c.x)
}
func (c POINTL) Y() LONG {
	return LONG(c.y)
}

func (m *DEVMODE) DmPosition() POINTL {
	return POINTL(C.win32_DEVMODE_dmPosition((*C.DEVMODE)(m)))
}

const (
	DM_BITSPERPEL       = C.DM_BITSPERPEL
	DM_PELSWIDTH        = C.DM_PELSWIDTH
	DM_PELSHEIGHT       = C.DM_PELSHEIGHT
	DM_DISPLAYFLAGS     = C.DM_DISPLAYFLAGS
	DM_DISPLAYFREQUENCY = C.DM_DISPLAYFREQUENCY
	DM_POSITION         = C.DM_POSITION
)

func (m *DEVMODE) SetDmFields(fields DWORD) {
	m.dmFields = C.DWORD(fields)
}

func (m *DEVMODE) SetDmPelsWidth(value DWORD) {
	m.dmPelsWidth = C.DWORD(value)
}

func (m *DEVMODE) SetDmPelsHeight(value DWORD) {
	m.dmPelsHeight = C.DWORD(value)
}

func (m *DEVMODE) SetDmBitsPerPel(value DWORD) {
	m.dmBitsPerPel = C.DWORD(value)
}

func (m *DEVMODE) SetDmDisplayFrequency(value DWORD) {
	m.dmDisplayFrequency = C.DWORD(value)
}

/*
typedef struct _devicemode {
  TCHAR dmDeviceName[CCHDEVICENAME];
  WORD  dmSpecVersion;
  WORD  dmDriverVersion;
  WORD  dmSize;
  WORD  dmDriverExtra;
  DWORD dmFields;
  union {
    struct {
      short dmOrientation;
      short dmPaperSize;
      short dmPaperLength;
      short dmPaperWidth;
      short dmScale;
      short dmCopies;
      short dmDefaultSource;
      short dmPrintQuality;
    };
    struct {
      POINTL dmPosition;
      DWORD  dmDisplayOrientation;
      DWORD  dmDisplayFixedOutput;
    };
  };
  short dmColor;
  short dmDuplex;
  short dmYResolution;
  short dmTTOption;
  short dmCollate;
  TCHAR dmFormName[CCHFORMNAME];
  WORD  dmLogPixels;
  //DWORD dmBitsPerPel;
  //DWORD dmPelsWidth;
  //DWORD dmPelsHeight;
  union {
    DWORD dmDisplayFlags;
    DWORD dmNup;
  };
  //DWORD dmDisplayFrequency;
#if (WINVER >= 0x0400)
  DWORD dmICMMethod;
  DWORD dmICMIntent;
  DWORD dmMediaType;
  DWORD dmDitherType;
  DWORD dmReserved1;
  DWORD dmReserved2;
#if (WINVER >= 0x0500) || (_WIN32_WINNT >= 0x0400)
  DWORD dmPanningWidth;
  DWORD dmPanningHeight;
#endif
#endif
} DEVMODE, *PDEVMODE, *LPDEVMODE;
*/

const ENUM_CURRENT_SETTINGS = C.ENUM_CURRENT_SETTINGS

func EnumDisplaySettings(lpszDeviceName string, iModeNum DWORD) (ret bool, devMode *DEVMODE) {
	var mode C.DEVMODE

	cDeviceName := StringToLPTSTR(lpszDeviceName)
	defer C.free(unsafe.Pointer(cDeviceName))

	ret = C.EnumDisplaySettings(cDeviceName, C.DWORD(iModeNum), (C.LPDEVMODE)(unsafe.Pointer(&mode))) != 0
	devMode = (*DEVMODE)(&mode)
	return
}

type VIDEOPARAMETERS C.VIDEOPARAMETERS

/*
typedef struct _VIDEOPARAMETERS {
  GUID  guid;
  ULONG dwOffset;
  ULONG dwCommand;
  ULONG dwFlags;
  ULONG dwMode;
  ULONG dwTVStandard;
  ULONG dwAvailableModes;
  ULONG dwAvailableTVStandard;
  ULONG dwFlickerFilter;
  ULONG dwOverScanX;
  ULONG dwOverScanY;
  ULONG dwMaxUnscaledX;
  ULONG dwMaxUnscaledY;
  ULONG dwPositionX;
  ULONG dwPositionY;
  ULONG dwBrightness;
  ULONG dwContrast;

  ULONG dwCPType;
  ULONG dwCPCommand;
  ULONG dwCPStandard;
  ULONG dwCPKey;
  ULONG bCP_APSTriggerBits;
  UCHAR bOEMCopyProtection[256];
} VIDOEPARAMETERS, *PVIDEOPARAMETERS;
*/

const (
	DISP_CHANGE_SUCCESSFUL  = C.DISP_CHANGE_SUCCESSFUL
	DISP_CHANGE_BADDUALVIEW = C.DISP_CHANGE_BADDUALVIEW
	DISP_CHANGE_BADFLAGS    = C.DISP_CHANGE_BADFLAGS
	DISP_CHANGE_BADMODE     = C.DISP_CHANGE_BADMODE
	DISP_CHANGE_BADPARAM    = C.DISP_CHANGE_BADPARAM
	DISP_CHANGE_FAILED      = C.DISP_CHANGE_FAILED
	DISP_CHANGE_NOTUPDATED  = C.DISP_CHANGE_NOTUPDATED
	DISP_CHANGE_RESTART     = C.DISP_CHANGE_RESTART

	CDS_TEST           = C.CDS_TEST
	CDS_UPDATEREGISTRY = C.CDS_UPDATEREGISTRY
)

// http://msdn.microsoft.com/en-us/library/windows/desktop/dd183413(v=vs.85).aspx
func ChangeDisplaySettingsEx(lpszDeviceName string, lpDevMode *DEVMODE, dwFlags DWORD, lParam *VIDEOPARAMETERS) (ret LONG) {
	cDeviceName := StringToLPTSTR(lpszDeviceName)
	defer C.free(unsafe.Pointer(cDeviceName))

	ret = LONG(C.ChangeDisplaySettingsEx(cDeviceName, (C.LPDEVMODE)(unsafe.Pointer(lpDevMode)), nil, C.DWORD(dwFlags), (C.LPVOID)(unsafe.Pointer(lParam))))
	return
}

// http://msdn.microsoft.com/en-us/library/windows/desktop/dd372194(v=vs.85).aspx
func SetDeviceGammaRamp(device HDC, ramp [3][256]WORD) (ret bool) {
	ret = C.SetDeviceGammaRamp(C.HDC(device), (C.LPVOID)(&ramp[0])) != 0
	return
}

// http://msdn.microsoft.com/en-us/library/windows/desktop/dd316946(v=vs.85).aspx
func GetDeviceGammaRamp(dc HDC) (ret bool, ramp [3][256]WORD) {
	ret = C.GetDeviceGammaRamp(C.HDC(dc), (C.LPVOID)(&ramp[0])) != 0
	return
}

// nIndex possible values:
const (
	HORZSIZE      = C.HORZSIZE      // mm width
	VERTSIZE      = C.VERTSIZE      // mm height
	HORZRES       = C.HORZRES       // px width
	VERTRES       = C.VERTRES       // px height
	VREFRESH      = C.VREFRESH      // current refresh rate
	CM_GAMMA_RAMP = C.CM_GAMMA_RAMP // supports gamma ramps
)

//
func GetDeviceCaps(dc HDC, nIndex Int) (ret Int) {
	ret = Int(C.GetDeviceCaps(C.HDC(dc), C.int(nIndex)))
	return
}

func CreateDC(lpszDriver, lpszDevice string, lpInitData *DEVMODE) (dc HDC) {
	var cDriver, cDevice C.LPTSTR
	if len(lpszDriver) > 0 {
		cDriver = StringToLPTSTR(lpszDriver)
		defer C.free(unsafe.Pointer(cDriver))
	}

	if len(lpszDevice) > 0 {
		cDevice = StringToLPTSTR(lpszDevice)
		defer C.free(unsafe.Pointer(cDevice))
	}

	dc = HDC(C.CreateDC(cDriver, cDevice, nil, (*C.DEVMODEW)(lpInitData)))
	return
}

// http://msdn.microsoft.com/en-us/library/windows/desktop/dd183533(v=vs.85).aspx
func DeleteDC(dc HDC) (ret bool) {
	ret = C.DeleteDC(C.HDC(dc)) != 0
	return
}

type HWND C.HWND

// http://msdn.microsoft.com/en-us/library/windows/desktop/dd162920(v=vs.85).aspx
// just returns bool even though docs say int.. stupid!
func ReleaseDC(wnd HWND, dc HDC) (ret bool) {
	ret = C.ReleaseDC(C.HWND(wnd), C.HDC(dc)) != 0
	return
}

// http://msdn.microsoft.com/en-us/library/windows/desktop/ms679360(v=vs.85).aspx
func GetLastError() (ret DWORD) {
	ret = DWORD(C.GetLastError())
	return
}

// Not an actual win32 api function
func GetLastErrorString() (ret string) {
	err := DWORD(C.GetLastError())

	var lpMsgBuf C.LPVOID
	defer C.LocalFree((C.HLOCAL)(lpMsgBuf))

	C.FormatMessage(C.FORMAT_MESSAGE_ALLOCATE_BUFFER|C.FORMAT_MESSAGE_FROM_SYSTEM|C.FORMAT_MESSAGE_IGNORE_INSERTS, nil, C.DWORD(err), C.DWORD(C.win32_MAKELANGID(C.LANG_NEUTRAL, C.SUBLANG_DEFAULT)), (C.LPTSTR)(unsafe.Pointer(&lpMsgBuf)), 0, nil)

	ret = LPTSTRToString((C.LPTSTR)(lpMsgBuf))
	return
}

type CREATESTRUCT C.CREATESTRUCT

type HMENU C.HMENU

type HINSTANCE C.HINSTANCE

type WPARAM C.WPARAM

func (c WPARAM) HIWORD() uint16 {
	return uint16((uint32(c) >> 16) & 0xFFFF)
}
func (c WPARAM) LOWORD() uint16 {
	return uint16(c)
}

type POINT C.POINT

func (c *POINT) SetX(x LONG) {
	c.x = C.LONG(x)
}
func (c *POINT) SetY(y LONG) {
	c.y = C.LONG(y)
}
func (c *POINT) X() LONG {
	return LONG(c.x)
}
func (c *POINT) Y() LONG {
	return LONG(c.y)
}

func ScreenToClient(hwnd HWND, point *POINT) bool {
	return C.ScreenToClient(C.HWND(hwnd), (C.LPPOINT)(unsafe.Pointer(point))) != 0
}

func ClientToScreen(hwnd HWND, point *POINT) bool {
	return C.ClientToScreen(C.HWND(hwnd), (C.LPPOINT)(unsafe.Pointer(point))) != 0
}

func WindowFromDC(hdc HDC) HWND {
	return HWND(C.WindowFromDC(C.HDC(hdc)))
}

type MINMAXINFO C.MINMAXINFO

/*
typedef struct tagMINMAXINFO {
  POINT ptReserved;
  POINT ptMaxSize;
  POINT ptMaxPosition;
  POINT ptMinTrackSize;
  POINT ptMaxTrackSize;
} MINMAXINFO, *PMINMAXINFO, *LPMINMAXINFO;
*/
func (c *MINMAXINFO) PtMinTrackSize() *POINT {
	return (*POINT)(&c.ptMinTrackSize)
}
func (c *MINMAXINFO) PtMaxTrackSize() *POINT {
	return (*POINT)(&c.ptMaxTrackSize)
}

type LPARAM C.LPARAM

//#define LOBYTE(w) ((BYTE)(w))
//#define HIBYTE(w) ((BYTE)(((WORD)(w) >> 8) & 0xFF))

//#define LOWORD(l) ((WORD)(l))
//#define HIWORD(l) ((WORD)(((DWORD)(l) >> 16) & 0xFFFF))
func (c LPARAM) HIWORD() uint16 {
	return uint16((uint32(c) >> 16) & 0xFFFF)
}
func (c LPARAM) LOWORD() uint16 {
	return uint16(c)
}

func (c LPARAM) MINMAXINFO() *MINMAXINFO {
	return (*MINMAXINFO)(unsafe.Pointer(uintptr(c)))
}

func (c LPARAM) RECT() *RECT {
	return (*RECT)(unsafe.Pointer(uintptr(c)))
}

type HICON C.HICON

type HCURSOR C.HCURSOR

type HBRUSH C.HBRUSH

func IntToHBRUSH(v Int) HBRUSH {
	return (HBRUSH)(unsafe.Pointer(&v))
}

type COLORREF C.COLORREF

func CreateSolidBrush(crColor COLORREF) HBRUSH {
	return HBRUSH(C.CreateSolidBrush(C.COLORREF(crColor)))
}

type HRGN C.HRGN

func CreateRectRgn(nLeftRect, nTopRect, nRightRect, nBottomRect int) HRGN {
	return HRGN(C.CreateRectRgn(C.int(nLeftRect), C.int(nTopRect), C.int(nRightRect), C.int(nBottomRect)))
}

type HMONITOR C.HMONITOR

func MonitorFromWindow(hwnd HWND, dwFlags DWORD) HMONITOR {
	return HMONITOR(C.MonitorFromWindow(C.HWND(hwnd), C.DWORD(dwFlags)))
}

type MONITORINFOEX struct {
	CbSize    DWORD
	RcMonitor RECT
	RcWork    RECT
	DwFlags   DWORD
	szDevice  [C.CCHDEVICENAME]uint16
}

func (m *MONITORINFOEX) SetSize() {
	m.CbSize = DWORD(unsafe.Sizeof(MONITORINFOEX{}))
}

func (m *MONITORINFOEX) Device() string {
	ar := make([]uint16, len(m.szDevice)+1)
	for i, c := range m.szDevice {
		ar[i] = c
	}
	return LPTSTRToString((C.LPTSTR)(unsafe.Pointer(&ar[0])))
}

func GetMonitorInfo(hMonitor HMONITOR, lpmi *MONITORINFOEX) bool {
	return C.GetMonitorInfo(C.HMONITOR(hMonitor), (C.LPMONITORINFO)(unsafe.Pointer(lpmi))) != 0
}

type MonitorEnumProc func(hMonitor HMONITOR, hdcMonitor HDC, lprcMonitor *RECT, dwData LPARAM) bool

var monitorEnumProcCallback MonitorEnumProc

/*
BOOL CALLBACK MonitorEnumProc(
  _In_  HMONITOR hMonitor,
  _In_  HDC hdcMonitor,
  _In_  LPRECT lprcMonitor,
  _In_  LPARAM dwData
);
*/

//export MonitorEnumProcCallback
func MonitorEnumProcCallback(hMonitor HMONITOR, hdcMonitor HDC, lprcMonitor *RECT, dwData LPARAM) C.WINBOOL {
	if monitorEnumProcCallback(hMonitor, hdcMonitor, lprcMonitor, dwData) {
		return C.WINBOOL(1)
	}
	return C.WINBOOL(0)
}

func EnumDisplayMonitors(hdc HDC, lprcClip *RECT, fn MonitorEnumProc, dwData LPARAM) bool {
	monitorEnumProcCallback = fn
	ret := C.EnumDisplayMonitors(C.HDC(hdc), (C.LPCRECT)(unsafe.Pointer(lprcClip)), C.win32_MonitorEnumProcCallbackHandle, C.LPARAM(dwData)) != 0
	monitorEnumProcCallback = nil
	return ret
}

/*
BOOL EnumDisplayMonitors(
  _In_  HDC hdc,
  _In_  LPCRECT lprcClip,
  _In_  MONITORENUMPROC lpfnEnum,
  _In_  LPARAM dwData
);
*/

type LowLevelKeyboardHookProc func(nCode Int, wParam WPARAM, lParam LPARAM) LRESULT

/*
LRESULT CALLBACK LowLevelKeyboardProc(
  _In_  int nCode,
  _In_  WPARAM wParam,
  _In_  LPARAM lParam
);
*/

var lowLevelKeyboardHookCallback LowLevelKeyboardHookProc

//export LowLevelKeyboardHookCallback
func LowLevelKeyboardHookCallback(nCode C.int, wParam C.WPARAM, lParam C.LPARAM) C.LRESULT {
	return C.LRESULT(lowLevelKeyboardHookCallback(Int(nCode), WPARAM(wParam), LPARAM(lParam)))
}

type HHOOK C.HHOOK

func SetLowLevelKeyboardHook(fn LowLevelKeyboardHookProc, hMod HINSTANCE, dwThreadId DWORD) HHOOK {
	lowLevelKeyboardHookCallback = fn
	return HHOOK(C.SetWindowsHookEx(C.WH_KEYBOARD_LL, C.win32_LowLevelKeyboardHookCallbackHandle, C.HINSTANCE(hMod), C.DWORD(dwThreadId)))
}

func UnhookWindowsHookEx(hook HHOOK) bool {
	lowLevelKeyboardHookCallback = nil
	return C.UnhookWindowsHookEx(C.HHOOK(hook)) != 0
}

type KBDLLHOOKSTRUCT struct {
	VkCode      DWORD
	ScanCode    DWORD
	Flags       DWORD
	Time        DWORD
	DwExtraInfo ULONG_PTR
}

const (
	HC_ACTION = C.HC_ACTION
)

func CallNextHookEx(hhk HHOOK, nCode Int, wParam WPARAM, lParam LPARAM) LRESULT {
	return LRESULT(C.CallNextHookEx(C.HHOOK(hhk), C.int(nCode), C.WPARAM(wParam), C.LPARAM(lParam)))
}

type WNDCLASSEX C.WNDCLASSEX

func NewWNDCLASSEX() *WNDCLASSEX {
	w := WNDCLASSEX{}
	w.cbSize = C.UINT(unsafe.Sizeof(w))
	return &w
}

func (w *WNDCLASSEX) SetStyle(style UINT) {
	w.style = C.UINT(style)
}

var callbacks = make(map[HWND]func(HWND, UINT, WPARAM, LPARAM) LRESULT)
var callbacksAccess sync.RWMutex

//export WndProcCallback
func WndProcCallback(hwnd C.HWND, msg C.UINT, wParam C.WPARAM, lParam C.LPARAM) C.LRESULT {
	// This gets called from C
	//fmt.Println("HWND", hwnd)

	callbacksAccess.RLock()
	callback, ok := callbacks[HWND(hwnd)]
	callbacksAccess.RUnlock()

	if ok {
		return C.LRESULT(callback(HWND(hwnd), UINT(msg), WPARAM(wParam), LPARAM(lParam)))
	}
	return C.DefWindowProc(hwnd, msg, wParam, lParam)
}

func (w *WNDCLASSEX) SetLpfnWndProc() {

	w.lpfnWndProc = C.win32_WndProcWrapperHandle
}

func RegisterWndProc(hwnd HWND, fn func(HWND, UINT, WPARAM, LPARAM) LRESULT) {
	callbacksAccess.Lock()
	defer callbacksAccess.Unlock()
	callbacks[hwnd] = fn
}

func UnregisterWndProc(hwnd HWND) {
	callbacksAccess.Lock()
	defer callbacksAccess.Unlock()
	delete(callbacks, hwnd)

}

func (w *WNDCLASSEX) SetCbClsExtra(v Int) {
	w.cbClsExtra = C.int(v)
}

func (w *WNDCLASSEX) SetCbWndExtra(v Int) {
	w.cbWndExtra = C.int(v)
}

func (w *WNDCLASSEX) SetHInstance(instance HINSTANCE) {
	w.hInstance = C.HINSTANCE(instance)
}

func (w *WNDCLASSEX) SetHIcon(icon HICON) {
	w.hIcon = C.HICON(icon)
}

func (w *WNDCLASSEX) SetHIconSm(icon HICON) {
	w.hIconSm = C.HICON(icon)
}

func (w *WNDCLASSEX) SetHCursor(cursor HCURSOR) {
	w.hCursor = C.HCURSOR(cursor)
}

const COLOR_WINDOW = C.COLOR_WINDOW

func (w *WNDCLASSEX) SetHbrBackground(v HBRUSH) {
	w.hbrBackground = C.HBRUSH(v)
}

func (w *WNDCLASSEX) SetLpszMenuName(v string) {
	cstr := StringToLPTSTR(v)
	defer C.free(unsafe.Pointer(cstr))

	w.lpszMenuName = (C.LPCWSTR)(cstr)
}

func (w *WNDCLASSEX) SetLpszClassName(v string) {
	cClassName := (C.LPCWSTR)(unsafe.Pointer(StringToLPTSTR(v)))

	// Never free here, this is free'd with UnregisterClass() calls
	//defer C.free(unsafe.Pointer(cClassName))

	w.lpszClassName = cClassName
}

// http://msdn.microsoft.com/en-us/library/windows/desktop/ms632600(v=vs.85).aspx
const (
	// We don't actually, use, we just hog the DC always.
	CS_OWNDC = C.CS_OWNDC

	CS_NOCLOSE = C.CS_NOCLOSE

	CS_PARENTDC = C.CS_PARENTDC

	// The window has a thin-line border.
	WS_BORDER = C.WS_BORDER

	// The window has a title bar (includes the WS_BORDER style).
	WS_CAPTION = C.WS_CAPTION

	// The window is a child window. A window with this style cannot have a menu bar. This style
	// cannot be used with the WS_POPUP style.
	WS_CHILD = C.WS_CHILD

	// Same as the WS_CHILD style.
	WS_CHILDWINDOW = C.WS_CHILDWINDOW

	// Excludes the area occupied by child windows when drawing occurs within the parent window.
	// This style is used when creating the parent window.
	WS_CLIPCHILDREN = C.WS_CLIPCHILDREN

	// Clips child windows relative to each other; that is, when a particular child window receives
	// a WM_PAINT message, the WS_CLIPSIBLINGS style clips all other overlapping child windows out
	// of the region of the child window to be updated. If WS_CLIPSIBLINGS is not specified and
	// child windows overlap, it is possible, when drawing within the client area of a child window
	// , to draw within the client area of a neighboring child window.
	WS_CLIPSIBLINGS = C.WS_CLIPSIBLINGS

	// The window is initially disabled. A disabled window cannot receive input from the user. To
	// change this after a window has been created, use the EnableWindow function.
	WS_DISABLED = C.WS_DISABLED

	// The window has a border of a style typically used with dialog boxes. A window with this
	// style cannot have a title bar.
	WS_DLGFRAME = C.WS_DLGFRAME

	// The window is the first control of a group of controls. The group consists of this first
	// control and all controls defined after it, up to the next control with the WS_GROUP style.
	// The first control in each group usually has the WS_TABSTOP style so that the user can move
	// from group to group. The user can subsequently change the keyboard focus from one control in
	// the group to the next control in the group by using the direction keys.
	// You can turn this style on and off to change dialog box navigation. To change this style
	// after a window has been created, use the SetWindowLong function.
	WS_GROUP = C.WS_GROUP

	// The window has a horizontal scroll bar.
	WS_HSCROLL = C.WS_HSCROLL

	// The window is initially minimized. Same as the WS_MINIMIZE style.
	WS_ICONIC = C.WS_ICONIC

	// The window is initially maximized.
	WS_MAXIMIZE = C.WS_MAXIMIZE

	// The window has a maximize button. Cannot be combined with the WS_EX_CONTEXTHELP style. The
	// WS_SYSMENU style must also be specified.
	WS_MAXIMIZEBOX = C.WS_MAXIMIZEBOX

	// The window is initially minimized. Same as the WS_ICONIC style.
	WS_MINIMIZE = C.WS_MINIMIZE

	// The window has a minimize button. Cannot be combined with the WS_EX_CONTEXTHELP style. The
	// WS_SYSMENU style must also be specified.
	WS_MINIMIZEBOX = C.WS_MINIMIZEBOX

	// The window is an overlapped window. An overlapped window has a title bar and a border. Same
	// as the WS_TILED style.
	WS_OVERLAPPED = C.WS_OVERLAPPED

	// The window is an overlapped window. Same as the WS_TILEDWINDOW style.
	WS_OVERLAPPEDWINDOW = C.WS_OVERLAPPEDWINDOW

	// The windows is a pop-up window. This style cannot be used with the WS_CHILD style.
	WS_POPUP = C.WS_POPUP

	// The window is a pop-up window. The WS_CAPTION and WS_POPUPWINDOW styles must be combined to
	// make the window menu visible.
	WS_POPUPWINDOW = C.WS_POPUPWINDOW

	// The window has a sizing border. Same as the WS_THICKFRAME style.
	WS_SIZEBOX = C.WS_SIZEBOX

	// The window has a window menu on its title bar. The WS_CAPTION style must also be specified.
	WS_SYSMENU = C.WS_SYSMENU

	// The window is a control that can receive the keyboard focus when the user presses the TAB
	// key. Pressing the TAB key changes the keyboard focus to the next control with the WS_TABSTOP
	// style.
	// You can turn this style on and off to change dialog box navigation. To change this style
	// after a window has been created, use the SetWindowLong function. For user-created windows
	// and modeless dialogs to work with tab stops, alter the message loop to call the
	// IsDialogMessage function.
	WS_TABSTOP = C.WS_TABSTOP

	// The window has a sizing border. Same as the WS_SIZEBOX style.
	WS_THICKFRAME = C.WS_THICKFRAME

	// The window is an overlapped window. An overlapped window has a title bar and a border. Same
	// as the WS_OVERLAPPED style.
	WS_TILED = C.WS_TILED

	// The window is an overlapped window. Same as the WS_OVERLAPPEDWINDOW style.
	WS_TILEDWINDOW = C.WS_TILEDWINDOW

	// The window is initially visible.
	// This style can be turned on and off by using the ShowWindow or SetWindowPos function.
	WS_VISIBLE = C.WS_VISIBLE

	// The window has a vertical scroll bar.
	WS_VSCROLL = C.WS_VSCROLL

	WS_EX_OVERLAPPEDWINDOW = C.WS_EX_OVERLAPPEDWINDOW

	WS_EX_COMPOSITED = C.WS_EX_COMPOSITED
)

const (
	GWL_EXSTYLE = C.GWL_EXSTYLE
	GWL_ID      = C.GWL_ID
	GWL_STYLE   = C.GWL_STYLE
)

func SetWindowLongPtr(hwnd HWND, nIndex Int, dwNewLong LONG_PTR) (ret LONG_PTR) {
	ret = LONG_PTR(C.SetWindowLongPtr(C.HWND(hwnd), C.int(nIndex), C.LONG_PTR(dwNewLong)))
	return
}

func GetWindowLongPtr(hwnd HWND, nIndex Int) (ret LONG_PTR) {
	ret = LONG_PTR(C.LONG_PTR(C.GetWindowLongPtr(C.HWND(hwnd), C.int(nIndex))))
	return
}

func CreateWindowEx(dwExStyle uint32, lpClassName, lpWindowName string, dwStyle DWORD, x, y, nWidth, nHeight Int, hWndParent HWND, hMenu HMENU, hInstance HINSTANCE, createStruct *CREATESTRUCT) (ret HWND) {
	cClassName := (*C.WCHAR)(unsafe.Pointer(StringToLPTSTR(lpClassName)))
	defer C.free(unsafe.Pointer(cClassName))

	cWindowName := (*C.WCHAR)(unsafe.Pointer(StringToLPTSTR(lpWindowName)))
	defer C.free(unsafe.Pointer(cWindowName))

	ret = HWND(C.CreateWindowEx(
		C.DWORD(dwExStyle),
		cClassName,
		cWindowName,
		C.DWORD(dwStyle),
		C.int(x), C.int(y),
		C.int(nWidth), C.int(nHeight),
		C.HWND(hWndParent),
		C.HMENU(hMenu),
		C.HINSTANCE(hInstance),
		C.LPVOID(createStruct),
	))
	return
}

func DestroyWindow(hwnd HWND) (ret bool) {
	ret = C.DestroyWindow(C.HWND(hwnd)) != 0
	return
}

const (
	// Minimizes a window, even if the thread that owns the window is not responding. This flag
	// should only be used when minimizing windows from a different thread.
	SW_FORCEMINIMIZE = C.SW_FORCEMINIMIZE

	// Hides the window and activates another window.
	SW_HIDE = C.SW_HIDE

	// Maximizes the specified window.
	SW_MAXIMIZE = C.SW_MAXIMIZE

	// Minimizes the specified window and activates the next top-level window in the Z order.
	SW_MINIMIZE = C.SW_MINIMIZE

	// Activates and displays the window. If the window is minimized or maximized, the system
	// restores it to its original size and position. An application should specify this flag when
	// restoring a minimized window.
	SW_RESTORE = C.SW_RESTORE

	// Activates the window and displays it in its current size and position.
	SW_SHOW = C.SW_SHOW

	// Sets the show state based on the SW_ value specified in the STARTUPINFO structure passed to
	// the CreateProcess function by the program that started the application.
	SW_SHOWDEFAULT = C.SW_SHOWDEFAULT

	// Activates the window and displays it as a maximized window.
	SW_SHOWMAXIMIZED = C.SW_SHOWMAXIMIZED

	// Activates the window and displays it as a minimized window.
	SW_SHOWMINIMIZED = C.SW_SHOWMINIMIZED

	// Displays the window as a minimized window. This value is similar to SW_SHOWMINIMIZED, except
	// the window is not activated.
	SW_SHOWMINNOACTIVE = C.SW_SHOWMINNOACTIVE

	// Displays the window in its current size and position. This value is similar to SW_SHOW,
	// except that the window is not activated.
	SW_SHOWNA = C.SW_SHOWNA

	// Displays a window in its most recent size and position. This value is similar to
	// SW_SHOWNORMAL, except that the window is not activated.
	SW_SHOWNOACTIVATE = C.SW_SHOWNOACTIVATE

	// Activates and displays a window. If the window is minimized or maximized, the system
	// restores it to its original size and position. An application should specify this flag when
	// displaying the window for the first time.
	SW_SHOWNORMAL = C.SW_SHOWNORMAL
)

func ShowWindow(hwnd HWND, nCmdShow Int) (wasShownBefore bool) {
	return C.ShowWindow(C.HWND(hwnd), C.int(nCmdShow)) != 0
}

type HMODULE C.HMODULE

func GetModuleHandle(lpModuleName string) (ret HMODULE) {
	cModuleName := StringToLPTSTR(lpModuleName)
	defer C.free(unsafe.Pointer(cModuleName))

	ret = HMODULE(C.GetModuleHandle((C.LPTSTR)(cModuleName)))
	return
}

type ATOM C.ATOM

func RegisterClassEx(lpwcx *WNDCLASSEX) (class ATOM) {
	class = ATOM(C.RegisterClassEx((*C.WNDCLASSEXW)(lpwcx)))
	return
}

func UnregisterClass(class string, instance HINSTANCE) (ret bool) {
	cClass := StringToLPTSTR(class)
	defer C.free(unsafe.Pointer(cClass))
	ret = C.UnregisterClass(cClass, C.HINSTANCE(instance)) != 0
	return
}

func DefWindowProc(hwnd HWND, msg UINT, wParam WPARAM, lParam LPARAM) (ret LRESULT) {
	ret = LRESULT(C.DefWindowProc(C.HWND(hwnd), C.UINT(msg), C.WPARAM(wParam), C.LPARAM(lParam)))
	return
}

type MSG C.MSG

/*
  HWND   hwnd;
  UINT   message;
  WPARAM wParam;
  LPARAM lParam;
  DWORD  time;
  POINT  pt;
*/

func (c *MSG) Hwnd() HWND {
	return HWND(c.hwnd)
}

func (c *MSG) Message() UINT {
	return UINT(c.message)
}

func (c *MSG) WParam() WPARAM {
	return WPARAM(c.wParam)
}

func (c *MSG) LParam() LPARAM {
	return LPARAM(c.lParam)
}

func (c *MSG) Time() DWORD {
	return DWORD(c.time)
}

func DispatchMessage(msg *MSG) (ret LRESULT) {
	ret = LRESULT(C.DispatchMessage((*C.MSG)(msg)))
	return
}

func SendMessage(hwnd HWND, msg UINT, wParam WPARAM, lParam LPARAM) LRESULT {
	return LRESULT(C.SendMessage(C.HWND(hwnd), C.UINT(msg), C.WPARAM(wParam), C.LPARAM(lParam)))
}

func SetCursor(cursor HCURSOR) HCURSOR {
	return HCURSOR(C.SetCursor(C.HCURSOR(cursor)))
}

func SetCapture(hwnd HWND) HWND {
	return HWND(C.SetCapture(C.HWND(hwnd)))
}

func ReleaseCapture() bool {
	return C.ReleaseCapture() != 0
}

const (
	PM_NOREMOVE       = C.PM_NOREMOVE
	PM_REMOVE         = C.PM_REMOVE
	PM_NOYIELD        = C.PM_NOYIELD
	PM_QS_INPUT       = C.PM_QS_INPUT
	PM_QS_POSTMESSAGE = C.PM_QS_POSTMESSAGE
	PM_QS_PAINT       = C.PM_QS_PAINT
	PM_QS_SENDMESSAGE = C.PM_QS_SENDMESSAGE
)

func PeekMessage(hwnd HWND, wMsgFilterMin, wMsgFilterMax, wRemoveMsg UINT) (ret bool, msg *MSG) {
	var cMsg C.MSG
	ret = C.PeekMessage((C.LPMSG)(unsafe.Pointer(&cMsg)), C.HWND(hwnd), C.UINT(wMsgFilterMin), C.UINT(wMsgFilterMax), C.UINT(wRemoveMsg)) != 0
	msg = (*MSG)(&cMsg)
	return
}

func GetMessage(hwnd HWND, wMsgFilterMin, wMsgFilterMax UINT) (ret bool, msg *MSG) {
	var cMsg C.MSG
	ret = C.GetMessage((C.LPMSG)(unsafe.Pointer(&cMsg)), C.HWND(hwnd), C.UINT(wMsgFilterMin), C.UINT(wMsgFilterMax)) != 0
	msg = (*MSG)(&cMsg)
	return
}

func TranslateMessage(msg *MSG) (ret bool) {
	ret = C.TranslateMessage((*C.MSG)(msg)) != 0
	return
}

type TIMERPROC C.TIMERPROC

const (
	WM_TIMER         = C.WM_TIMER
	WM_GETMINMAXINFO = C.WM_GETMINMAXINFO

	WMSZ_BOTTOM      = C.WMSZ_BOTTOM
	WMSZ_BOTTOMLEFT  = C.WMSZ_BOTTOMLEFT
	WMSZ_BOTTOMRIGHT = C.WMSZ_BOTTOMRIGHT
	WMSZ_LEFT        = C.WMSZ_LEFT
	WMSZ_RIGHT       = C.WMSZ_RIGHT
	WMSZ_TOP         = C.WMSZ_TOP
	WMSZ_TOPLEFT     = C.WMSZ_TOPLEFT
	WMSZ_TOPRIGHT    = C.WMSZ_TOPRIGHT
	WM_SIZING        = C.WM_SIZING

	WM_PAINT      = C.WM_PAINT
	WM_ERASEBKGND = C.WM_ERASEBKGND

	WM_SETCURSOR = C.WM_SETCURSOR

	ICON_BIG    = C.ICON_BIG
	ICON_SMALL  = C.ICON_SMALL
	ICON_SMALL2 = C.ICON_SMALL2
	WM_GETICON  = C.WM_GETICON
	WM_SETICON  = C.WM_SETICON

	WM_SIZE        = C.WM_SIZE
	SIZE_MAXIMIZED = C.SIZE_MAXIMIZED
	SIZE_MINIMIZED = C.SIZE_MINIMIZED
	SIZE_RESTORED  = C.SIZE_RESTORED

	WM_CLOSE = C.WM_CLOSE

	WM_ACTIVATE    = C.WM_ACTIVATE
	WA_INACTIVE    = C.WA_INACTIVE
	WA_ACTIVE      = C.WA_ACTIVE
	WA_CLICKACTIVE = C.WA_CLICKACTIVE

	WM_MOUSEMOVE     = C.WM_MOUSEMOVE
	WM_LBUTTONDOWN   = C.WM_LBUTTONDOWN
	WM_LBUTTONUP     = C.WM_LBUTTONUP
	WM_LBUTTONDBLCLK = C.WM_LBUTTONDBLCLK
	WM_RBUTTONDOWN   = C.WM_RBUTTONDOWN
	WM_RBUTTONUP     = C.WM_RBUTTONUP
	WM_RBUTTONDBLCLK = C.WM_RBUTTONDBLCLK
	WM_MBUTTONDOWN   = C.WM_MBUTTONDOWN
	WM_MBUTTONUP     = C.WM_MBUTTONUP
	WM_MBUTTONDBLCLK = C.WM_MBUTTONDBLCLK
	WM_MOUSEWHEEL    = C.WM_MOUSEWHEEL

	MK_CONTROL       = 0x0008
	MK_LBUTTON       = 0x0001
	MK_MBUTTON       = 0x0010
	MK_RBUTTON       = 0x0002
	MK_SHIFT         = 0x0004
	MK_XBUTTON1      = 0x0020
	MK_XBUTTON2      = 0x0040
	WM_XBUTTONDOWN   = C.WM_XBUTTONDOWN
	WM_XBUTTONUP     = C.WM_XBUTTONUP
	WM_XBUTTONDBLCLK = C.WM_XBUTTONDBLCLK
	WM_MOUSELAST     = C.WM_MOUSELAST
	WM_MOUSEHWHEEL   = 0x020E

	// WM_MOUSEMOVE is WM_MOUSEENTER
	WM_MOUSELEAVE = C.WM_MOUSELEAVE

	WM_SYSKEYDOWN = C.WM_SYSKEYDOWN
	WM_SYSKEYUP   = C.WM_SYSKEYUP
	WM_KEYDOWN    = C.WM_KEYDOWN
	WM_KEYUP      = C.WM_KEYUP

	WM_CHAR = C.WM_CHAR

	WM_MOVE = C.WM_MOVE

	WM_INPUT = C.WM_INPUT

	GIDC_ARRIVAL           = 1
	GIDC_REMOVAL           = 2
	WM_INPUT_DEVICE_CHANGE = 0x00FE

	MONITOR_DEFAULTTONEAREST = C.MONITOR_DEFAULTTONEAREST
	WM_EXITSIZEMOVE          = C.WM_EXITSIZEMOVE

	HTNOWHERE     = C.HTNOWHERE
	HTTRANSPARENT = C.HTTRANSPARENT
	WM_NCHITTEST  = C.WM_NCHITTEST
)

func IsIconic(hwnd HWND) bool {
	return C.IsIconic(C.HWND(hwnd)) != 0
}

func SetCursorPos(x, y int32) bool {
	return C.SetCursorPos(C.int(x), C.int(y)) != 0
}

func SetTimer(hwnd HWND, nIDEvent UINT_PTR, uElapse UINT, lpTimerFunc TIMERPROC) (timer UINT_PTR) {
	timer = UINT_PTR(C.SetTimer(C.HWND(hwnd), C.UINT_PTR(nIDEvent), C.UINT(uElapse), C.TIMERPROC(lpTimerFunc)))
	return
}

func KillTimer(hwnd HWND, uIDEvent UINT_PTR) (ret bool) {
	ret = C.KillTimer(C.HWND(hwnd), C.UINT_PTR(uIDEvent)) != 0
	return
}

type RECT C.RECT

func (c *RECT) Left() LONG {
	return LONG(c.left)
}
func (c *RECT) Top() LONG {
	return LONG(c.top)
}
func (c *RECT) Right() LONG {
	return LONG(c.right)
}
func (c *RECT) Bottom() LONG {
	return LONG(c.bottom)
}

func (c *RECT) SetLeft(v LONG) {
	c.left = C.LONG(v)
}
func (c *RECT) SetTop(v LONG) {
	c.top = C.LONG(v)
}
func (c *RECT) SetRight(v LONG) {
	c.right = C.LONG(v)
}
func (c *RECT) SetBottom(v LONG) {
	c.bottom = C.LONG(v)
}

func GetUpdateRect(hwnd HWND, lpRect *RECT, bErase bool) bool {
	cbool := C.WINBOOL(0)
	if bErase {
		cbool = 1
	}
	if lpRect != nil {
		return C.GetUpdateRect(C.HWND(hwnd), (C.LPRECT)(unsafe.Pointer(lpRect)), cbool) != 0
	}
	return C.GetUpdateRect(C.HWND(hwnd), nil, cbool) != 0
}

func ValidateRect(hwnd HWND, rect *RECT) bool {
	if rect != nil {
		return C.ValidateRect(C.HWND(hwnd), (*C.RECT)(unsafe.Pointer(&rect))) != 0
	}
	return C.ValidateRect(C.HWND(hwnd), nil) != 0
}

func GetWindowRect(hwnd HWND) (status bool, r *RECT) {
	var cr C.RECT
	status = C.GetWindowRect(C.HWND(hwnd), (C.LPRECT)(unsafe.Pointer(&cr))) != 0
	r = (*RECT)(&cr)
	return
}

func GetClientRect(hwnd HWND) (status bool, r *RECT) {
	var cr C.RECT
	status = C.GetClientRect(C.HWND(hwnd), (C.LPRECT)(unsafe.Pointer(&cr))) != 0
	r = (*RECT)(&cr)
	return
}

func ClipCursor(rect *RECT) bool {
	return C.ClipCursor((*C.RECT)(unsafe.Pointer(rect))) != 0
}

func GetClipCursor() (clip *RECT, ok bool) {
	var c RECT
	ret := C.GetClipCursor((C.LPRECT)(unsafe.Pointer(&c)))
	if ret != 0 {
		return &c, true
	}
	return nil, false
}

const (
	SM_CYCAPTION   = C.SM_CYCAPTION // Title bar width
	SM_CXSIZEFRAME = C.SM_CXSIZEFRAME
	SM_CYSIZEFRAME = C.SM_CYSIZEFRAME
	SM_CXCURSOR    = C.SM_CXCURSOR
	SM_CYCURSOR    = C.SM_CYCURSOR
	SM_CXICON      = C.SM_CXICON
	SM_CYICON      = C.SM_CYICON
	SM_CXSMICON    = C.SM_CXSMICON
	SM_CYSMICON    = C.SM_CYSMICON
)

func GetSystemMetrics(nIndex Int) (ret Int) {
	ret = Int(C.GetSystemMetrics(C.int(nIndex)))
	return
}

const (
	SWP_ASYNCWINDOWPOS = C.SWP_ASYNCWINDOWPOS
	SWP_DEFERERASE     = C.SWP_DEFERERASE
	SWP_DRAWFRAME      = C.SWP_DRAWFRAME
	SWP_FRAMECHANGED   = C.SWP_FRAMECHANGED
	SWP_HIDEWINDOW     = C.SWP_HIDEWINDOW
	SWP_NOACTIVATE     = C.SWP_NOACTIVATE
	SWP_NOCOPYBITS     = C.SWP_NOCOPYBITS
	SWP_NOMOVE         = C.SWP_NOMOVE
	SWP_NOOWNERZORDER  = C.SWP_NOOWNERZORDER
	SWP_NOREDRAW       = C.SWP_NOREDRAW
	SWP_NOREPOSITION   = C.SWP_NOREPOSITION
	SWP_NOSENDCHANGING = C.SWP_NOSENDCHANGING
	SWP_NOSIZE         = C.SWP_NOSIZE
	SWP_NOZORDER       = C.SWP_NOZORDER
	SWP_SHOWWINDOW     = C.SWP_SHOWWINDOW
)

var (
	iHWND_TOP       = 0
	iHWND_BOTTOM    = 1
	iHWND_TOPMOST   = -1
	iHWND_NOTOPMOST = -2

	HWND_TOP       = *(*HWND)(unsafe.Pointer(&iHWND_TOP))
	HWND_BOTTOM    = *(*HWND)(unsafe.Pointer(&iHWND_BOTTOM))
	HWND_TOPMOST   = *(*HWND)(unsafe.Pointer(&iHWND_TOPMOST))
	HWND_NOTOPMOST = *(*HWND)(unsafe.Pointer(&iHWND_NOTOPMOST))
)

func SetWindowPos(hwnd, hwndInsertAfter HWND, X, Y, cx, cy Int, uFlags UINT) (ret bool) {
	ret = C.SetWindowPos(C.HWND(hwnd), C.HWND(hwndInsertAfter), C.int(X), C.int(Y), C.int(cx), C.int(cy), C.UINT(uFlags)) != 0
	return
}

func MoveWindow(hwnd HWND, x, y, width, height Int, repaint bool) bool {
	cbool := C.WINBOOL(0)
	if repaint {
		cbool = 1
	}
	return C.MoveWindow(C.HWND(hwnd), C.int(x), C.int(y), C.int(width), C.int(height), cbool) != 0
}

func EnableWindow(hwnd HWND, bEnable bool) (ret bool) {
	var enable C.WINBOOL
	if bEnable {
		enable = 1
	} else {
		enable = 0
	}
	ret = C.EnableWindow(C.HWND(hwnd), enable) != 0
	return
}

func FlashWindow(hwnd HWND, bInvert bool) (ret bool) {
	cInvert := 1
	if bInvert {
		cInvert = 0
	}
	return C.FlashWindow(C.HWND(hwnd), C.WINBOOL(cInvert)) != 0
}

func SetWindowText(hwnd HWND, lpString string) bool {
	cstr := StringToLPTSTR(lpString)
	defer C.free(unsafe.Pointer(cstr))
	return C.SetWindowText(C.HWND(hwnd), cstr) != 0
}

type HBITMAP unsafe.Pointer

type ICONINFO struct {
	FIcon    Int
	XHotspot DWORD
	YHotspot DWORD
	HbmMask  HBITMAP
	HbmColor HBITMAP
}

type BITMAPINFOHEADER struct {
	Size DWORD
	Width,
	Height LONG
	Planes,
	BitCount WORD
	Compression,
	SizeImage DWORD
	XPelsPerMeter,
	YPelsPerMeter LONG
	ClrUsed,
	ClrImportant DWORD
}

type RGBQUAD struct {
	RgbBlue,
	RgbGreen,
	RgbRed,
	RgbReserved uint8
}

type BITMAPINFO struct {
	BmiHeader BITMAPINFOHEADER
	BmiColors [1]RGBQUAD
}

const (
	DIB_RGB_COLORS = C.DIB_RGB_COLORS
	BI_RGB         = C.BI_RGB
)

type HGDIOBJ C.HGDIOBJ

func DeleteObject(object HGDIOBJ) bool {
	return C.DeleteObject(C.HGDIOBJ(object)) != 0
}

func SelectObject(hdc HDC, hgdiobj HGDIOBJ) HGDIOBJ {
	return HGDIOBJ(C.SelectObject(C.HDC(hdc), C.HGDIOBJ(hgdiobj)))
}

func CreateCompatibleDC(hdc HDC) HDC {
	return HDC(C.CreateCompatibleDC(C.HDC(hdc)))
}

const (
	NULL_BRUSH  = C.NULL_BRUSH
	BLACK_BRUSH = C.BLACK_BRUSH
)

func GetStockObject(fnObject Int) HGDIOBJ {
	return HGDIOBJ(C.GetStockObject(C.int(fnObject)))
}

func FillRect(hdc HDC, rect *RECT, hbr HBRUSH) bool {
	return C.FillRect(C.HDC(hdc), (*C.RECT)(unsafe.Pointer(rect)), C.HBRUSH(hbr)) != 0
}

type BLENDFUNCTION struct {
	BlendOp             BYTE
	BlendFlags          BYTE
	SourceConstantAlpha BYTE
	AlphaFormat         BYTE
}

const (
	AC_SRC_OVER  = C.AC_SRC_OVER
	AC_SRC_ALPHA = C.AC_SRC_ALPHA
)

func AlphaBlend(hdcDest HDC, xoriginDest, yoriginDest, wDest, hDest Int, hdcSrc HDC, xoriginSrc, yoriginSrc, wSrc, hSrc Int, ftn *BLENDFUNCTION) bool {
	return C.AlphaBlend(C.HDC(hdcDest), C.int(xoriginDest), C.int(yoriginDest), C.int(wDest), C.int(hDest), C.HDC(hdcSrc), C.int(xoriginSrc), C.int(yoriginSrc), C.int(wSrc), C.int(hSrc), *(*C.BLENDFUNCTION)(unsafe.Pointer(ftn))) != 0
}

func TransparentBlt(a HDC, b, c, d, e Int, f HDC, g, h, i, j Int, k uint32) bool {
	return C.TransparentBlt(C.HDC(a), C.int(b), C.int(c), C.int(d), C.int(e), C.HDC(f), C.int(g), C.int(h), C.int(i), C.int(j), C.UINT(k)) != 0
}

var (
	IDC_ARROW = LPTSTR(C.macro_MAKEINTRESOURCE(32512))
)

func LoadCursor(hinstance HINSTANCE, lpCursorName LPTSTR) HCURSOR {
	return HCURSOR(C.LoadCursor(C.HINSTANCE(hinstance), C.LPTSTR(lpCursorName)))
}

func DestroyCursor(cursor HCURSOR) bool {
	return C.DestroyCursor(C.HCURSOR(cursor)) != 0
}

func DestroyIcon(icon HICON) bool {
	return C.DestroyIcon(C.HICON(icon)) != 0
}

func CreateIconIndirect(piconinfo *ICONINFO) HICON {
	return HICON(C.CreateIconIndirect((C.PICONINFO)(unsafe.Pointer(piconinfo))))
}

func CreateCompatibleBitmap(hdc HDC, nWidth, nHeight Int) HBITMAP {
	return HBITMAP(C.CreateCompatibleBitmap(C.HDC(hdc), C.int(nWidth), C.int(nHeight)))
}

func SetDIBits(hdc HDC, hbmp HBITMAP, uStartScan, cScanLines UINT, lpvBits unsafe.Pointer, lpbmi *BITMAPINFO, fuColorUse UINT) Int {
	return Int(C.SetDIBits(C.HDC(hdc), C.HBITMAP(hbmp), C.UINT(uStartScan), C.UINT(cScanLines), lpvBits, (*C.BITMAPINFO)(unsafe.Pointer(lpbmi)), C.UINT(fuColorUse)))
}

// See: http://msdn.microsoft.com/en-us/library/ms724833(v=vs.85).aspx
type OSVERSIONINFOEX C.OSVERSIONINFOEX

/*
typedef struct _OSVERSIONINFOEX {
  DWORD dwOSVersionInfoSize;
  DWORD dwMajorVersion;
  DWORD dwMinorVersion;
  DWORD dwBuildNumber;
  DWORD dwPlatformId;
  TCHAR szCSDVersion[128];
  WORD  wServicePackMajor;
  WORD  wServicePackMinor;
  WORD  wSuiteMask;
  BYTE  wProductType;
  BYTE  wReserved;
} OSVERSIONINFOEX, *POSVERSIONINFOEX, *LPOSVERSIONINFOEX;
*/

func (c *OSVERSIONINFOEX) DwMajorVersion() DWORD {
	return DWORD(c.dwMajorVersion)
}

func (c *OSVERSIONINFOEX) DwMinorVersion() DWORD {
	return DWORD(c.dwMinorVersion)
}

func GetVersionEx() (ret bool, lpVersionInfo *OSVERSIONINFOEX) {
	var vi C.OSVERSIONINFOEX
	vi.dwOSVersionInfoSize = C.DWORD(unsafe.Sizeof(C.OSVERSIONINFOEX{}))

	return C.GetVersionEx((C.LPOSVERSIONINFO)(unsafe.Pointer(&vi))) != 0, (*OSVERSIONINFOEX)(&vi)
}

func GetDC(hwnd HWND) HDC {
	return HDC(C.GetDC(C.HWND(hwnd)))
}

type PIXELFORMATDESCRIPTOR C.PIXELFORMATDESCRIPTOR

/*
ct tagPIXELFORMATDESCRIPTOR {
  WORD  nSize;
  WORD  nVersion;
  DWORD dwFlags;
  BYTE  iPixelType;
  BYTE  cColorBits;
  BYTE  cRedBits;
  BYTE  cRedShift;
  BYTE  cGreenBits;
  BYTE  cGreenShift;
  BYTE  cBlueBits;
  BYTE  cBlueShift;
  BYTE  cAlphaBits;
  BYTE  cAlphaShift;
  BYTE  cAccumBits;
  BYTE  cAccumRedBits;
  BYTE  cAccumGreenBits;
  BYTE  cAccumBlueBits;
  BYTE  cAccumAlphaBits;
  BYTE  cDepthBits;
  BYTE  cStencilBits;
  BYTE  cAuxBuffers;
  BYTE  iLayerType;
  BYTE  bReserved;
  DWORD dwLayerMask;
  DWORD dwVisibleMask;
  DWORD dwDamageMask;
} PIXELFORMATDESCRIPTOR, *PPIXELFORMATDESCRIPTOR;
*/

const (
	PFD_DRAW_TO_WINDOW      = C.PFD_DRAW_TO_WINDOW
	PFD_DRAW_TO_BITMAP      = C.PFD_DRAW_TO_BITMAP
	PFD_SUPPORT_GDI         = C.PFD_SUPPORT_GDI
	PFD_SUPPORT_OPENGL      = C.PFD_SUPPORT_OPENGL
	PFD_GENERIC_ACCELERATED = C.PFD_GENERIC_ACCELERATED
	PFD_GENERIC_FORMAT      = C.PFD_GENERIC_FORMAT
	PFD_NEED_PALETTE        = C.PFD_NEED_PALETTE
	PFD_NEED_SYSTEM_PALETTE = C.PFD_NEED_SYSTEM_PALETTE
	PFD_DOUBLEBUFFER        = C.PFD_DOUBLEBUFFER
	PFD_STEREO              = C.PFD_STEREO
	PFD_SWAP_LAYER_BUFFERS  = C.PFD_SWAP_LAYER_BUFFERS

	PFD_SUPPORT_COMPOSITION = 0x00008000
)

func (c *PIXELFORMATDESCRIPTOR) DwFlags() DWORD {
	return DWORD(c.dwFlags)
}

const (
	PFD_TYPE_RGBA       = C.PFD_TYPE_RGBA
	PFD_TYPE_COLORINDEX = C.PFD_TYPE_COLORINDEX
)

func (c *PIXELFORMATDESCRIPTOR) IPixelType() BYTE {
	return BYTE(c.iPixelType)
}
func (c *PIXELFORMATDESCRIPTOR) CRedBits() BYTE {
	return BYTE(c.cRedBits)
}
func (c *PIXELFORMATDESCRIPTOR) CGreenBits() BYTE {
	return BYTE(c.cGreenBits)
}
func (c *PIXELFORMATDESCRIPTOR) CBlueBits() BYTE {
	return BYTE(c.cBlueBits)
}
func (c *PIXELFORMATDESCRIPTOR) CAlphaBits() BYTE {
	return BYTE(c.cAlphaBits)
}
func (c *PIXELFORMATDESCRIPTOR) CAccumRedBits() BYTE {
	return BYTE(c.cAccumRedBits)
}
func (c *PIXELFORMATDESCRIPTOR) CAccumGreenBits() BYTE {
	return BYTE(c.cAccumGreenBits)
}
func (c *PIXELFORMATDESCRIPTOR) CAccumBlueBits() BYTE {
	return BYTE(c.cAccumBlueBits)
}
func (c *PIXELFORMATDESCRIPTOR) CAccumAlphaBits() BYTE {
	return BYTE(c.cAccumAlphaBits)
}
func (c *PIXELFORMATDESCRIPTOR) CDepthBits() BYTE {
	return BYTE(c.cDepthBits)
}
func (c *PIXELFORMATDESCRIPTOR) CStencilBits() BYTE {
	return BYTE(c.cStencilBits)
}
func (c *PIXELFORMATDESCRIPTOR) CAuxBuffers() BYTE {
	return BYTE(c.cAuxBuffers)
}

func DescribePixelFormat(hdc HDC, iPixelFormat Int) (Int, *PIXELFORMATDESCRIPTOR) {
	nBytes := unsafe.Sizeof(C.PIXELFORMATDESCRIPTOR{})
	pfd := &PIXELFORMATDESCRIPTOR{}
	return Int(C.DescribePixelFormat(C.HDC(hdc), C.int(iPixelFormat), C.UINT(nBytes), (C.LPPIXELFORMATDESCRIPTOR)(unsafe.Pointer(pfd)))), pfd
}

func SetPixelFormat(hdc HDC, iPixelFormat Int, ppfd *PIXELFORMATDESCRIPTOR) bool {
	return C.SetPixelFormat(C.HDC(hdc), C.int(iPixelFormat), (*C.PIXELFORMATDESCRIPTOR)(ppfd)) != 0
}

func SwapBuffers(hdc HDC) bool {
	return C.SwapBuffers(C.HDC(hdc)) != 0
}

const (
	HID_USAGE_PAGE_GENERIC  USHORT = 0x01
	HID_USAGE_GENERIC_MOUSE USHORT = 0x02
	RIDEV_INPUTSINK                = C.RIDEV_INPUTSINK

	RID_INPUT     = C.RID_INPUT
	RIM_TYPEMOUSE = C.RIM_TYPEMOUSE
)

type RAWINPUTHEADER struct {
	Type   DWORD
	Size   DWORD
	Device unsafe.Pointer
	Param  WPARAM
}

type RAWMOUSE C.RAWMOUSE

/*
typedef struct tagRAWMOUSE {
  USHORT usFlags;
  union {
    ULONG  ulButtons;
    struct {
      USHORT usButtonFlags;
      USHORT usButtonData;
    };
  };
  ULONG  ulRawButtons;
  LONG   lLastX;
  LONG   lLastY;
  ULONG  ulExtraInformation;
} RAWMOUSE, *PRAWMOUSE, *LPRAWMOUSE;
*/
func (c *RAWMOUSE) LastX() LONG {
	return LONG(c.lLastX)
}
func (c *RAWMOUSE) LastY() LONG {
	return LONG(c.lLastY)
}

type RAWINPUT C.RAWINPUT

/*
typedef struct tagRAWINPUT {
  RAWINPUTHEADER header;
  union {
    RAWMOUSE    mouse;
    RAWKEYBOARD keyboard;
    RAWHID      hid;
  } data;
} RAWINPUT, *PRAWINPUT, *LPRAWINPUT;
*/

func (c *RAWINPUT) Header() *RAWINPUTHEADER {
	return (*RAWINPUTHEADER)(unsafe.Pointer(&c.header))
}

func (c *RAWINPUT) Mouse() *RAWMOUSE {
	x := (*RAWMOUSE)(unsafe.Pointer(&c.data[0]))
	return x
}

type RAWINPUTDEVICE struct {
	UsagePage USHORT
	Usage     USHORT
	Flags     DWORD
	Target    HWND
}

func RegisterRawInputDevices(pRawInputDevices *RAWINPUTDEVICE, uiNumDevices UINT, cbSize UINT) bool {
	return C.RegisterRawInputDevices((C.PCRAWINPUTDEVICE)(unsafe.Pointer(pRawInputDevices)), C.UINT(uiNumDevices), C.UINT(cbSize)) != 0
}

func GetRegisteredRawInputDevices(pRawInputDevices *RAWINPUTDEVICE, puiNumDevices *UINT, cbSize UINT) UINT {
	return UINT(C.GetRegisteredRawInputDevices((C.PRAWINPUTDEVICE)(unsafe.Pointer(pRawInputDevices)), (C.PUINT)(unsafe.Pointer(puiNumDevices)), C.UINT(cbSize)))
}

type HRAWINPUT C.HRAWINPUT

func GetRawInputData(hRawInput HRAWINPUT, uiCommand UINT, pData unsafe.Pointer, pcbSize *UINT, cbSizeHeader UINT) UINT {
	return UINT(C.GetRawInputData(C.HRAWINPUT(hRawInput), C.UINT(uiCommand), (C.LPVOID)(pData), (C.PUINT)(unsafe.Pointer(pcbSize)), C.UINT(cbSizeHeader)))
}

func GetAsyncKeyState(vKey Int) int16 {
	return int16(C.GetAsyncKeyState(C.int(vKey)))
}

func GetKeyState(vKey Int) int16 {
	return int16(C.GetKeyState(C.int(vKey)))
}

const (
	VK_LBUTTON  = 0x01
	VK_RBUTTON  = 0x02
	VK_CANCEL   = 0x03
	VK_MBUTTON  = 0x04
	VK_XBUTTON1 = 0x05
	VK_XBUTTON2 = 0x06
	VK_BACK     = 0x08
	VK_TAB      = 0x09
	VK_CLEAR    = 0x0C
	VK_RETURN   = 0x0D
	VK_SHIFT    = 0x10
	VK_CONTROL  = 0x11
	VK_MENU     = 0x12
	VK_PAUSE    = 0x13
	VK_CAPITAL  = 0x14

	VK_KANA    = 0x15
	VK_HANGUEL = 0x15
	VK_HANGUL  = 0x15

	VK_JUNJA = 0x17
	VK_FINAL = 0x18

	VK_HANJA = 0x19
	VK_KANJI = 0x19

	VK_ESCAPE     = 0x1B
	VK_CONVERT    = 0x1C
	VK_NONCONVERT = 0x1D
	VK_ACCEPT     = 0x1E

	VK_MODECHANGE = 0x1F
	VK_SPACE      = 0x20
	VK_PRIOR      = 0x21
	VK_NEXT       = 0x22
	VK_END        = 0x23
	VK_HOME       = 0x24
	VK_LEFT       = 0x25
	VK_UP         = 0x26
	VK_RIGHT      = 0x27
	VK_DOWN       = 0x28
	VK_SELECT     = 0x29
	VK_PRINT      = 0x2A
	VK_EXECUTE    = 0x2B
	VK_SNAPSHOT   = 0x2C
	VK_INSERT     = 0x2D
	VK_DELETE     = 0x2E
	VK_HELP       = 0x2F

	VK_UNDEF_0 = 0x30
	VK_UNDEF_1 = 0x31
	VK_UNDEF_2 = 0x32
	VK_UNDEF_3 = 0x33
	VK_UNDEF_4 = 0x34
	VK_UNDEF_5 = 0x35
	VK_UNDEF_6 = 0x36
	VK_UNDEF_7 = 0x37
	VK_UNDEF_8 = 0x38
	VK_UNDEF_9 = 0x39

	VK_UNDEF_A = 0x41
	VK_UNDEF_B = 0x42
	VK_UNDEF_C = 0x43
	VK_UNDEF_D = 0x44
	VK_UNDEF_E = 0x45
	VK_UNDEF_F = 0x46
	VK_UNDEF_G = 0x47
	VK_UNDEF_H = 0x48
	VK_UNDEF_I = 0x49
	VK_UNDEF_J = 0x4A
	VK_UNDEF_K = 0x4B
	VK_UNDEF_L = 0x4C
	VK_UNDEF_M = 0x4D
	VK_UNDEF_N = 0x4E
	VK_UNDEF_O = 0x4F
	VK_UNDEF_P = 0x50
	VK_UNDEF_Q = 0x51
	VK_UNDEF_R = 0x52
	VK_UNDEF_S = 0x53
	VK_UNDEF_T = 0x54
	VK_UNDEF_U = 0x55
	VK_UNDEF_V = 0x56
	VK_UNDEF_W = 0x57
	VK_UNDEF_X = 0x58
	VK_UNDEF_Y = 0x59
	VK_UNDEF_Z = 0x5A

	VK_LWIN      = 0x5B
	VK_RWIN      = 0x5C
	VK_APPS      = 0x5D
	VK_SLEEP     = 0x5F
	VK_NUMPAD0   = 0x60
	VK_NUMPAD1   = 0x61
	VK_NUMPAD2   = 0x62
	VK_NUMPAD3   = 0x63
	VK_NUMPAD4   = 0x64
	VK_NUMPAD5   = 0x65
	VK_NUMPAD6   = 0x66
	VK_NUMPAD7   = 0x67
	VK_NUMPAD8   = 0x68
	VK_NUMPAD9   = 0x69
	VK_MULTIPLY  = 0x6A
	VK_ADD       = 0x6B
	VK_SEPARATOR = 0x6C
	VK_SUBTRACT  = 0x6D
	VK_DECIMAL   = 0x6E
	VK_DIVIDE    = 0x6F
	VK_F1        = 0x70
	VK_F2        = 0x71
	VK_F3        = 0x72
	VK_F4        = 0x73
	VK_F5        = 0x74
	VK_F6        = 0x75
	VK_F7        = 0x76
	VK_F8        = 0x77
	VK_F9        = 0x78
	VK_F10       = 0x79
	VK_F11       = 0x7A
	VK_F12       = 0x7B
	VK_F13       = 0x7C
	VK_F14       = 0x7D
	VK_F15       = 0x7E
	VK_F16       = 0x7F
	VK_F17       = 0x80
	VK_F18       = 0x81
	VK_F19       = 0x82
	VK_F20       = 0x83
	VK_F21       = 0x84
	VK_F22       = 0x85
	VK_F23       = 0x86
	VK_F24       = 0x87

	VK_NUMLOCK  = 0x90
	VK_SCROLL   = 0x91
	VK_LSHIFT   = 0xA0
	VK_RSHIFT   = 0xA1
	VK_LCONTROL = 0xA2
	VK_RCONTROL = 0xA3
	VK_LMENU    = 0xA4

	VK_RMENU               = 0xA5
	VK_BROWSER_BACK        = 0xA6
	VK_BROWSER_FORWARD     = 0xA7
	VK_BROWSER_REFRESH     = 0xA8
	VK_BROWSER_STOP        = 0xA9
	VK_BROWSER_SEARCH      = 0xAA
	VK_BROWSER_FAVORITES   = 0xAB
	VK_BROWSER_HOME        = 0xAC
	VK_VOLUME_MUTE         = 0xAD
	VK_VOLUME_DOWN         = 0xAE
	VK_VOLUME_UP           = 0xAF
	VK_MEDIA_NEXT_TRACK    = 0xB0
	VK_MEDIA_PREV_TRACK    = 0xB1
	VK_MEDIA_STOP          = 0xB2
	VK_MEDIA_PLAY_PAUSE    = 0xB3
	VK_LAUNCH_MAIL         = 0xB4
	VK_LAUNCH_MEDIA_SELECT = 0xB5
	VK_LAUNCH_APP1         = 0xB6
	VK_LAUNCH_APP2         = 0xB7
	VK_OEM_1               = 0xBA
	VK_OEM_PLUS            = 0xBB
	VK_OEM_COMMA           = 0xBC
	VK_OEM_MINUS           = 0xBD
	VK_OEM_PERIOD          = 0xBE
	VK_OEM_2               = 0xBF
	VK_OEM_3               = 0xC0
	VK_OEM_4               = 0xDB
	VK_OEM_5               = 0xDC
	VK_OEM_6               = 0xDD
	VK_OEM_7               = 0xDE
	VK_OEM_8               = 0xDF
	VK_OEM_102             = 0xE2
	VK_PROCESSKEY          = 0xE5
	VK_PACKET              = 0xE7
	VK_ATTN                = 0xF6
	VK_CRSEL               = 0xF7
	VK_EXSEL               = 0xF8
	VK_EREOF               = 0xF9
	VK_PLAY                = 0xFA
	VK_ZOOM                = 0xFB
	VK_PA1                 = 0xFD
	VK_OEM_CLEAR           = 0xFE
)
