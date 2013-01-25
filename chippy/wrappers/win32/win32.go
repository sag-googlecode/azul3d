package win32

// Windows Type Information:
//  http://en.wikibooks.org/wiki/Windows_Programming/Handles_and_Data_Types
//
// Most windows types are in:
//  C:\mingw\x86_64-w64-mingw32\include\windef.h
//
// Also look at windows.h in the same folder, and wingdi.h for the gdi32 headers.

/*
#include <windows.h>
#cgo LDFLAGS: -luser32 -lgdi32 -lkernel32


WORD macro_MAKELANGID(USHORT usPrimaryLanguage, USHORT usSubLanguage) {
    return MAKELANGID(usPrimaryLanguage, usSubLanguage);
}
*/
import "C"

import (
	"unsafe"
)

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

// Flags:
const (
	DISPLAY_DEVICE_ACTIVE           = C.DISPLAY_DEVICE_ACTIVE
	DISPLAY_DEVICE_MIRRORING_DRIVER = C.DISPLAY_DEVICE_MIRRORING_DRIVER
	DISPLAY_DEVICE_MODESPRUNED      = C.DISPLAY_DEVICE_MODESPRUNED
	DISPLAY_DEVICE_PRIMARY_DEVICE   = C.DISPLAY_DEVICE_PRIMARY_DEVICE
	DISPLAY_DEVICE_REMOVABLE        = C.DISPLAY_DEVICE_REMOVABLE
	DISPLAY_DEVICE_VGA_COMPATIBLE   = C.DISPLAY_DEVICE_VGA_COMPATIBLE
)

func (c *DISPLAY_DEVICE) GetStateFlags() uint32 {
	return uint32(c.StateFlags)
}

// http://msdn.microsoft.com/en-us/library/windows/desktop/dd162609(v=vs.85).aspx
func EnumDisplayDevices(lpDevice string, iDevNum, dwFlags uint32) (bool, *DISPLAY_DEVICE) {
	var dd C.DISPLAY_DEVICE
	dd.cb = C.DWORD(unsafe.Sizeof(dd))

	var cstr *C.CHAR
	if len(lpDevice) > 0 {
		cstr = (*C.CHAR)(C.CString(lpDevice))
		defer C.free(unsafe.Pointer(cstr))
	} else {
		cstr = nil
	}
	ret := C.EnumDisplayDevices(cstr, C.DWORD(iDevNum), (C.PDISPLAY_DEVICE)(unsafe.Pointer(&dd)), C.DWORD(dwFlags))
	return ret != 0, (*DISPLAY_DEVICE)(&dd)
}

type DEVMODE C.DEVMODE

func NewDEVMODE() *DEVMODE {
	m := DEVMODE{}
	m.dmSize = C.WORD(unsafe.Sizeof(m))
	return &m
}

func (m *DEVMODE) DmDeviceName() string {
	slice := []rune{}
	for _, c := range m.dmDeviceName {
		if c != 0 {
			slice = append(slice, rune(c))
		}
	}
	return string(slice)
}

func (m *DEVMODE) DmBitsPerPel() uint32 {
	return uint32(m.dmBitsPerPel)
}

func (m *DEVMODE) DmPelsWidth() uint32 {
	return uint32(m.dmPelsWidth)
}

func (m *DEVMODE) DmPelsHeight() uint32 {
	return uint32(m.dmPelsHeight)
}

func (m *DEVMODE) DmDisplayFrequency() uint32 {
	return uint32(m.dmDisplayFrequency)
}

const (
	DM_BITSPERPEL       = C.DM_BITSPERPEL
	DM_PELSWIDTH        = C.DM_PELSWIDTH
	DM_PELSHEIGHT       = C.DM_PELSHEIGHT
	DM_DISPLAYFLAGS     = C.DM_DISPLAYFLAGS
	DM_DISPLAYFREQUENCY = C.DM_DISPLAYFREQUENCY
	DM_POSITION         = C.DM_POSITION
)

func (m *DEVMODE) SetDmFields(fields uint32) {
	m.dmFields = C.DWORD(fields)
}

func (m *DEVMODE) SetDmPelsWidth(value uint32) {
	m.dmPelsWidth = C.DWORD(value)
}

func (m *DEVMODE) SetDmPelsHeight(value uint32) {
	m.dmPelsHeight = C.DWORD(value)
}

func (m *DEVMODE) SetDmBitsPerPel(value uint32) {
	m.dmBitsPerPel = C.DWORD(value)
}

func (m *DEVMODE) SetDmDisplayFrequency(value uint32) {
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

func EnumDisplaySettings(lpszDeviceName string, iModeNum uint32) (bool, *DEVMODE) {
	var mode C.DEVMODE

	var cDeviceName C.LPCTSTR
	if len(lpszDeviceName) > 0 {
		cstr := C.CString(lpszDeviceName)
		defer C.free(unsafe.Pointer(cstr))
		cDeviceName = (C.LPCTSTR)(unsafe.Pointer(cstr))
	}
	ret := C.EnumDisplaySettings(cDeviceName, C.DWORD(iModeNum), (C.LPDEVMODE)(unsafe.Pointer(&mode)))
	return ret != 0, (*DEVMODE)(&mode)
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
func ChangeDisplaySettingsEx(lpszDeviceName string, lpDevMode *DEVMODE, dwFlags uint32, lParam *VIDEOPARAMETERS) int32 {
	var cDeviceName C.LPCTSTR
	if len(lpszDeviceName) > 0 {
		cstr := C.CString(lpszDeviceName)
		defer C.free(unsafe.Pointer(cstr))
		cDeviceName = (C.LPCTSTR)(unsafe.Pointer(cstr))
	}
	r := C.ChangeDisplaySettingsEx(cDeviceName, (C.LPDEVMODE)(unsafe.Pointer(lpDevMode)), nil, C.DWORD(dwFlags), (C.LPVOID)(&lParam))
	return int32(r)
}

// http://msdn.microsoft.com/en-us/library/windows/desktop/dd372194(v=vs.85).aspx
func SetDeviceGammaRamp(device HDC, ramp [3][256]uint16) bool {
	return C.SetDeviceGammaRamp(C.HDC(device), (C.LPVOID)(&ramp[0])) != 0
}

// http://msdn.microsoft.com/en-us/library/windows/desktop/dd316946(v=vs.85).aspx
func GetDeviceGammaRamp(dc HDC) (bool, [3][256]uint16) {
	var ramp [3][256]uint16
	ret := C.GetDeviceGammaRamp(C.HDC(dc), (C.LPVOID)(&ramp[0]))
	return ret != 0, ramp
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
func GetDeviceCaps(dc HDC, nIndex int) int32 {
	return int32(C.GetDeviceCaps(C.HDC(dc), C.int(nIndex)))
}

func CreateDC(lpszDriver, lpszDevice string, lpInitData *DEVMODE) HDC {
	var cDriver *C.CHAR
	if len(lpszDriver) > 0 {
		cDriver = (*C.CHAR)(C.CString(lpszDriver))
		defer C.free(unsafe.Pointer(cDriver))
	} else {
		cDriver = nil
	}

	var cDevice *C.CHAR
	if len(lpszDevice) > 0 {
		cDevice = (*C.CHAR)(C.CString(lpszDevice))
		defer C.free(unsafe.Pointer(cDevice))
	} else {
		cDevice = nil
	}

	return HDC(C.CreateDC(cDriver, cDevice, nil, (*C.DEVMODEA)(lpInitData)))
}

// http://msdn.microsoft.com/en-us/library/windows/desktop/dd183533(v=vs.85).aspx
func DeleteDC(dc HDC) bool {
	return C.DeleteDC(C.HDC(dc)) != 0
}

type HWND C.HWND

// http://msdn.microsoft.com/en-us/library/windows/desktop/dd162920(v=vs.85).aspx
// just returns bool even though docs say int.. stupid!
func ReleaseDC(wnd HWND, dc HDC) bool {
	return C.ReleaseDC(C.HWND(wnd), C.HDC(dc)) != 0
}

// http://msdn.microsoft.com/en-us/library/windows/desktop/ms679360(v=vs.85).aspx
func GetLastError() uint32 {
	return uint32(C.GetLastError())
}

// Not an actual win32 api function
func GetLastErrorString() string {
	err := GetLastError()

	var lpMsgBuf C.LPVOID
	C.FormatMessage(C.FORMAT_MESSAGE_ALLOCATE_BUFFER|C.FORMAT_MESSAGE_FROM_SYSTEM|C.FORMAT_MESSAGE_IGNORE_INSERTS, nil, C.DWORD(err), C.DWORD(C.macro_MAKELANGID(C.LANG_NEUTRAL, C.SUBLANG_DEFAULT)), (C.LPTSTR)(unsafe.Pointer(&lpMsgBuf)), 0, nil)

	defer C.LocalFree((C.HLOCAL)(lpMsgBuf))
	return C.GoString((*C.char)(lpMsgBuf))
}
