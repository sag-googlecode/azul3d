// +build !no_xf86vm

package x11

/*
#include <stdlib.h>
#include <X11/Xlib.h>
#include <X11/extensions/xf86vmode.h>

#cgo LDFLAGS: -lXxf86vm

XF86VidModeModeInfo* XF86VidModeModeInfo_Index(XF86VidModeModeInfo** ptr, int index) {
    return ptr[index];
}

unsigned short* newGammaRampArray(int size) {
    int sSize = sizeof(unsigned short*);

    return malloc(sSize * size);
}

void setGammaRampArrayAtIndex(unsigned short* array, int index, unsigned short color) {
    array[index] = color;
}

unsigned short getGammaRampArrayAtIndex(unsigned short* array, int index) {
    return array[index];
}
*/
import "C"

import (
	"unsafe"
)

type (
	XF86VidModeMonitor  C.XF86VidModeMonitor
	XF86VidModeModeInfo C.XF86VidModeModeInfo
	XF86VidModeModeLine C.XF86VidModeModeLine
)

func (m *XF86VidModeMonitor) VsyncHi() float32 {
	return float32(m.vsync.hi)
}

func (m *XF86VidModeMonitor) Free() {
	C.free(unsafe.Pointer(m.vendor))
	C.free(unsafe.Pointer(m.model))
	C.free(unsafe.Pointer(m.hsync))
	C.free(unsafe.Pointer(m.vsync))
}

func (c *XF86VidModeModeInfo) Dotclock() uint {
	return uint(c.dotclock)
}

func (c *XF86VidModeModeInfo) Vtotal() uint16 {
	return uint16(c.vtotal)
}

func (c *XF86VidModeModeInfo) Htotal() uint16 {
	return uint16(c.htotal)
}

func (c *XF86VidModeModeInfo) Flags() uint {
	return uint(c.flags)
}

func (c *XF86VidModeModeInfo) Hdisplay() uint16 {
	return uint16(c.hdisplay)
}

func (c *XF86VidModeModeInfo) Vdisplay() uint16 {
	return uint16(c.vdisplay)
}

// returns status, major, minor
func XF86VidModeQueryVersion(display *Display) (int, int, int) {
	var major, minor C.int
	ret := C.XF86VidModeQueryVersion((*C.Display)(display), &major, &minor)
	return int(ret), int(major), int(minor)
}

// returns status, event_base, error_base
func XF86VidModeQueryExtension(display *Display) (int, int, int) {
	var event, error C.int
	ret := C.XF86VidModeQueryExtension((*C.Display)(display), &event, &error)
	return int(ret), int(event), int(error)
}

func XF86VidModeGetMonitor(display *Display, screen int) (int, *XF86VidModeMonitor) {
	var monitor C.XF86VidModeMonitor
	ret := C.XF86VidModeGetMonitor((*C.Display)(display), C.int(screen), &monitor)
	return int(ret), (*XF86VidModeMonitor)(&monitor)
}

// returns status, dotclock, modeline
func XF86VidModeGetModeLine(display *Display, screen int) (int, int, *XF86VidModeModeLine) {
	var modeline C.XF86VidModeModeLine
	var dotclock C.int
	ret := C.XF86VidModeGetModeLine((*C.Display)(display), C.int(screen), &dotclock, &modeline)
	gmodeline := XF86VidModeModeLine(modeline)
	return int(ret), int(dotclock), &gmodeline
}

func XF86VidModeValidateModeLine(display *Display, screen int, modeline *XF86VidModeModeInfo) int {
	return int(C.XF86VidModeValidateModeLine((*C.Display)(display), C.int(screen), (*C.XF86VidModeModeInfo)(modeline)))
}

func XF86VidModeGetAllModeLines(display *Display, screen int) (int, []*XF86VidModeModeInfo, func()) {
	var nModes C.int
	var first **C.XF86VidModeModeInfo

	ret := C.XF86VidModeGetAllModeLines((*C.Display)(display), C.int(screen), &nModes, &first)

	var slice []*XF86VidModeModeInfo

	if ret != 0 {
		slice = make([]*XF86VidModeModeInfo, nModes)
		for i := 0; i < int(nModes); i++ {
			slice[i] = (*XF86VidModeModeInfo)(C.XF86VidModeModeInfo_Index(first, C.int(i)))
		}
	}

	return int(ret), slice, func() {
		C.XFree(unsafe.Pointer(first))
	}
}

func XF86VidModeSwitchToMode(display *Display, screen int, modeline *XF86VidModeModeInfo) int {
	return int(C.XF86VidModeSwitchToMode((*C.Display)(display), C.int(screen), (*C.XF86VidModeModeInfo)(modeline)))
}

func XF86VidModeSetViewPort(display *Display, screen, x, y int) int {
	return int(C.XF86VidModeSetViewPort((*C.Display)(display), C.int(screen), C.int(x), C.int(y)))
}

func XF86VidModeSetGammaRamp(display *Display, screen int, red, green, blue []uint16) int {
	size := C.int(len(red))

	cRed := C.newGammaRampArray(size)
	cGreen := C.newGammaRampArray(size)
	cBlue := C.newGammaRampArray(size)
	defer C.free(unsafe.Pointer(cRed))
	defer C.free(unsafe.Pointer(cGreen))
	defer C.free(unsafe.Pointer(cBlue))

	for i, color := range red {
		C.setGammaRampArrayAtIndex(cRed, C.int(i), C.ushort(color))
	}

	for i, color := range green {
		C.setGammaRampArrayAtIndex(cGreen, C.int(i), C.ushort(color))
	}

	for i, color := range blue {
		C.setGammaRampArrayAtIndex(cBlue, C.int(i), C.ushort(color))
	}

	return int(C.XF86VidModeSetGammaRamp((*C.Display)(display), C.int(screen), size, cRed, cGreen, cBlue))
}

func XF86VidModeGetGammaRampSize(display *Display, screen int) (status int, size int) {
	var csize C.int
	status = int(C.XF86VidModeGetGammaRampSize((*C.Display)(display), C.int(screen), &csize))
	size = int(csize)
	return
}

func XF86VidModeGetGammaRamp(display *Display, screen int) (status int, red, green, blue []uint16) {
	// Umm, we expect whatever size XF86VidMode has for us? seems obvious.
	status, size := XF86VidModeGetGammaRampSize(display, screen)
	if status == 0 {
		return
	}

	cSize := C.int(size)
	cRed := C.newGammaRampArray(cSize)
	cGreen := C.newGammaRampArray(cSize)
	cBlue := C.newGammaRampArray(cSize)
	defer C.free(unsafe.Pointer(cRed))
	defer C.free(unsafe.Pointer(cGreen))
	defer C.free(unsafe.Pointer(cBlue))
	status = int(C.XF86VidModeGetGammaRamp((*C.Display)(display), C.int(screen), cSize, cRed, cGreen, cBlue))

	if status != 0 {
		red = make([]uint16, size)
		green = make([]uint16, size)
		blue = make([]uint16, size)

		for i := 0; i < size; i++ {
			red[i] = uint16(C.getGammaRampArrayAtIndex(cRed, C.int(i)))
			green[i] = uint16(C.getGammaRampArrayAtIndex(cGreen, C.int(i)))
			blue[i] = uint16(C.getGammaRampArrayAtIndex(cBlue, C.int(i)))
		}
	}
	return
}
