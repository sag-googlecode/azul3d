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

#cgo LDFLAGS: -luser32 -lgdi32 -lkernel32

WORD macro_MAKELANGID(USHORT usPrimaryLanguage, USHORT usSubLanguage);
void SetLPTSTRAtIndex(LPTSTR array, int index, WORD v);
WORD LPTSTRAtIndex(LPTSTR array, int index);

LRESULT CALLBACK WndProcWrapper(HWND hwnd, UINT msg, WPARAM wParam, LPARAM lParam);

WNDPROC WndProcWrapperHandle;

DWORD DEVMODE_dmDisplayFixedOutput(DEVMODE* dm);
*/
import "C"

import(
    "unicode/utf16"
	"unsafe"
)

// UTF16ToString returns the UTF-8 encoding of the UTF-16 sequence s,
// with a terminating NUL removed.
func UTF16ToString(s []uint16) string {
    for i, v := range s {
        if v == 0 {
            s = s[0:i]
            break
        }
    }
    return string(utf16.Decode(s))
}

func LPTSTRToString(cstr C.LPTSTR) string {
    if cstr == nil {
        return ""
    }
    strlen := int(C.wcslen((*C.wchar_t)(unsafe.Pointer(cstr))))

    slice := make([]uint16, strlen)
    for i := 0; i < strlen; i++ {
        slice[i] = uint16(C.LPTSTRAtIndex(cstr, C.int(i)))
    }
    return UTF16ToString(slice)
}

func StringToLPTSTR(g string) C.LPTSTR {
    if len(g) == 0 {
        return nil
    }

    encoded := utf16.Encode([]rune(g))
    cstr := (C.LPTSTR)(C.malloc(C.size_t(len(encoded) * 16)))

    for i := 0; i < len(encoded)+1; i++ {
        if i == len(encoded) {
            C.SetLPTSTRAtIndex(cstr, C.int(i), C.WORD(0))
            break
        }

        C.SetLPTSTRAtIndex(cstr, C.int(i), C.WORD(encoded[i]))
    }

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

// Wonder if there is an better string conversion for these?
// FIXME
func (c *DISPLAY_DEVICE) GetDeviceName() string {
	slice := []rune{}
	for i := 0; i < len(c.DeviceName); i++ {
		c := c.DeviceName[i]
		if c != 0 {
			slice = append(slice, rune(c))
		}
	}
	return string(slice)
}

// FIXME
func (c *DISPLAY_DEVICE) GetDeviceString() string {
	slice := []rune{}
	for i := 0; i < len(c.DeviceString); i++ {
		c := c.DeviceString[i]
		if c != 0 {
			slice = append(slice, rune(c))
		}
	}
	return string(slice)
}

const (
	DISPLAY_DEVICE_ACTIVE           = C.DISPLAY_DEVICE_ACTIVE
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
    var dd C.DISPLAY_DEVICE
    dd.cb = C.DWORD(unsafe.Sizeof(dd))

    cDevice := StringToLPTSTR(lpDevice)
    defer C.free(unsafe.Pointer(cDevice))

    ret = C.EnumDisplayDevices(cDevice, C.DWORD(iDevNum), (C.PDISPLAY_DEVICE)(unsafe.Pointer(&dd)), C.DWORD(dwFlags)) != 0
    displayDevice = (*DISPLAY_DEVICE)(&dd)
    return
}

type DEVMODE C.DEVMODE

func NewDEVMODE() *DEVMODE {
	m := DEVMODE{}
	m.dmSize = C.WORD(unsafe.Sizeof(m))
	return &m
}

func (m *DEVMODE) DmDeviceName() string {
    // FIXME
	slice := []rune{}
	for _, c := range m.dmDeviceName {
		if c != 0 {
			slice = append(slice, rune(c))
		}
	}
	return string(slice)
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

const(
	DMDFO_DEFAULT = C.DMDFO_DEFAULT
	DMDFO_CENTER = C.DMDFO_CENTER
	DMDFO_STRETCH = C.DMDFO_STRETCH
)


func (m *DEVMODE) DmDisplayFixedOutput() DWORD {
	return DWORD(C.DEVMODE_dmDisplayFixedOutput((*C.DEVMODE)(m)))
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
    cDriver := StringToLPTSTR(lpszDriver)
    defer C.free(unsafe.Pointer(cDriver))

    cDevice := StringToLPTSTR(lpszDevice)
    defer C.free(unsafe.Pointer(cDevice))

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

    C.FormatMessage(C.FORMAT_MESSAGE_ALLOCATE_BUFFER|C.FORMAT_MESSAGE_FROM_SYSTEM|C.FORMAT_MESSAGE_IGNORE_INSERTS, nil, C.DWORD(err), C.DWORD(C.macro_MAKELANGID(C.LANG_NEUTRAL, C.SUBLANG_DEFAULT)), (C.LPTSTR)(unsafe.Pointer(&lpMsgBuf)), 0, nil)

    ret = LPTSTRToString((C.LPTSTR)(lpMsgBuf))
    return
}

type CREATESTRUCT C.CREATESTRUCT

type HMENU C.HMENU

type HINSTANCE C.HINSTANCE

type WPARAM C.WPARAM

type POINT C.POINT
func (c *POINT) SetX(x int32) {
    c.x = C.LONG(x)
}
func (c *POINT) SetY(y int32) {
    c.y = C.LONG(y)
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
func (c LPARAM) MINMAXINFO() *MINMAXINFO {
    return (*MINMAXINFO)(unsafe.Pointer(uintptr(c)))
}

type HICON C.HICON

type HCURSOR C.HCURSOR

type HBRUSH C.HBRUSH

func IntToHBRUSH(v Int) HBRUSH {
    return (HBRUSH)(unsafe.Pointer(&v))
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


var callbacks = make(map[HWND]func(HWND, UINT, WPARAM, LPARAM)(LRESULT))

//export WndProcCallback
func WndProcCallback(hwnd C.HWND, msg C.UINT, wParam C.WPARAM, lParam C.LPARAM) C.LRESULT {
    // This gets called from C
    //fmt.Println("HWND", hwnd)

    callback, ok := callbacks[HWND(hwnd)]
    if ok {
        return C.LRESULT(callback(HWND(hwnd), UINT(msg), WPARAM(wParam), LPARAM(lParam)))
    }
    return C.DefWindowProc(hwnd, msg, wParam, lParam)
}

func (w *WNDCLASSEX) SetLpfnWndProc() {
    w.lpfnWndProc = C.WndProcWrapperHandle
}

func RegisterWndProc(hwnd HWND, fn func(HWND, UINT, WPARAM, LPARAM)(LRESULT)) {
    callbacks[hwnd] = fn
}

func UnregisterWndProc(hwnd HWND) {
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
const(
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
)


const(
    GWL_EXSTYLE = C.GWL_EXSTYLE
    GWL_ID = C.GWL_ID
    GWL_STYLE = C.GWL_STYLE
)

func SetWindowLongPtr(hwnd HWND, nIndex Int, dwNewLong LONG_PTR) (ret LONG_PTR) {
    ret = LONG_PTR(C.SetWindowLongPtr(C.HWND(hwnd), C.int(nIndex), C.LONG(dwNewLong)))
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

    type arguments struct{
        dwExStyle C.DWORD
        lpClassName, lpWindowName *C.WCHAR
        dwStyle C.DWORD
        x, y, nWidth, nHeight C.int
        hWndParent C.HWND
        hMenu C.HMENU
        hInstance C.HINSTANCE
        createStruct C.LPVOID
    }

    args := arguments{
        C.DWORD(dwExStyle),
        cClassName,
        cWindowName,
        C.DWORD(dwStyle),
        C.int(x), C.int(y), C.int(nWidth), C.int(nHeight),
        C.HWND(hWndParent),
        C.HMENU(hMenu),
        C.HINSTANCE(hInstance),
        C.LPVOID(&createStruct.lpCreateParams),
    }

    ret = HWND(C.CreateWindowEx(args.dwExStyle, args.lpClassName, args.lpWindowName, args.dwStyle, args.x, args.y, args.nWidth, args.nHeight, args.hWndParent, args.hMenu, args.hInstance, args.createStruct))
    return
}

func DestroyWindow(hwnd HWND) (ret bool) {
    ret = C.DestroyWindow(C.HWND(hwnd)) != 0
    return
}


const(
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

func ShowWindowAsync(hwnd HWND, nCmdShow Int) (wasShownBefore bool) {
    return C.ShowWindowAsync(C.HWND(hwnd), C.int(nCmdShow)) != 0
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

const(
    PM_NOREMOVE = C.PM_NOREMOVE
    PM_REMOVE = C.PM_REMOVE
    PM_NOYIELD = C.PM_NOYIELD
    PM_QS_INPUT = C.PM_QS_INPUT
    PM_QS_POSTMESSAGE = C.PM_QS_POSTMESSAGE
    PM_QS_PAINT = C.PM_QS_PAINT
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

const(
    WM_TIMER = C.WM_TIMER
    WM_GETMINMAXINFO = C.WM_GETMINMAXINFO

    WM_SIZE = C.WM_SIZE
    WM_SIZING = C.WM_SIZING
)

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


const(
    SM_CYCAPTION = C.SM_CYCAPTION // Title bar width
    SM_CXSIZEFRAME = C.SM_CXSIZEFRAME
    SM_CYSIZEFRAME = C.SM_CYSIZEFRAME
)
func GetSystemMetrics(nIndex Int) (ret Int) {
    ret = Int(C.GetSystemMetrics(C.int(nIndex)))
    return
}

const(
    SWP_ASYNCWINDOWPOS = C.SWP_ASYNCWINDOWPOS
    SWP_DEFERERASE = C.SWP_DEFERERASE
    SWP_DRAWFRAME = C.SWP_DRAWFRAME
    SWP_FRAMECHANGED = C.SWP_FRAMECHANGED
    SWP_HIDEWINDOW = C.SWP_HIDEWINDOW
    SWP_NOACTIVATE = C.SWP_NOACTIVATE
    SWP_NOCOPYBITS = C.SWP_NOCOPYBITS
    SWP_NOMOVE = C.SWP_NOMOVE
    SWP_NOOWNERZORDER = C.SWP_NOOWNERZORDER
    SWP_NOREDRAW = C.SWP_NOREDRAW
    SWP_NOREPOSITION = C.SWP_NOREPOSITION
    SWP_NOSENDCHANGING = C.SWP_NOSENDCHANGING
    SWP_NOSIZE = C.SWP_NOSIZE
    SWP_NOZORDER = C.SWP_NOZORDER
    SWP_SHOWWINDOW = C.SWP_SHOWWINDOW
)

func SetWindowPos(hwnd, hwndInsertAfter HWND, X, Y, cx, cy Int, uFlags UINT) (ret bool) {
    ret = C.SetWindowPos(C.HWND(hwnd), C.HWND(hwndInsertAfter), C.int(X), C.int(Y), C.int(cx), C.int(cy), C.UINT(uFlags)) != 0
    return
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

