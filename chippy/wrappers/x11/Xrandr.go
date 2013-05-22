// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

// +build linux,!no_xrandr

package x11

/*
#include <stdlib.h>
#include <X11/Xlib.h>
#include <X11/extensions/Xrandr.h>

#cgo LDFLAGS: -lXrandr

short shortAtIndex(short* array, int index) {
    return array[index];
}

XRRScreenSize XRRScreenSizeAtIndex(XRRScreenSize* array, int index) {
    return array[index];
}

unsigned short ushortAtIndex(unsigned short* array, int index) {
    return array[index];
}

void ushortSetAtIndex(unsigned short* array, int index, unsigned short value) {
    array[index] = value;
}

RRCrtc RRCrtcAtIndex(RRCrtc* array, int index) {
    return array[index];
}

RROutput* newRROutputArray(int size) {
    return malloc(sizeof(RROutput*) * size);
}

void RROutputSetAtIndex(RROutput* array, int index, RROutput value) {
    array[index] = value;
}

RROutput RROutputAtIndex(RROutput* array, int index) {
    return array[index];
}

RRMode RRModeAtIndex(RRMode* array, int index) {
    return array[index];
}

XRRModeInfo XRRModeInfoAtIndex(XRRModeInfo* array, int index) {
    return array[index];
}
*/
import "C"

// See the oh-so-important document describing this protocol at:
// http://cgit.freedesktop.org/xorg/proto/randrproto/tree/randrproto.txt
//
// "What is an CRTC?":
// http://en.wikipedia.org/wiki/Video_Display_Controller

import (
	"unsafe"
)

type (
	RROutput     C.RROutput
	RRCrtc       C.RRCrtc
	RRMode       C.RRMode
	Rotation     C.Rotation
	SizeID       C.SizeID
	XRRModeFlags C.XRRModeFlags
)

const (
	RR_HSyncPositive  XRRModeFlags = C.RR_HSyncPositive
	RR_HSyncNegative  XRRModeFlags = C.RR_HSyncNegative
	RR_VSyncPositive  XRRModeFlags = C.RR_VSyncPositive
	RR_VSyncNegative  XRRModeFlags = C.RR_VSyncNegative
	RR_Interlace      XRRModeFlags = C.RR_Interlace
	RR_DoubleScan     XRRModeFlags = C.RR_DoubleScan
	RR_CSync          XRRModeFlags = C.RR_CSync
	RR_CSyncPositive  XRRModeFlags = C.RR_CSyncPositive
	RR_CSyncNegative  XRRModeFlags = C.RR_CSyncNegative
	RR_HSkewPresent   XRRModeFlags = C.RR_HSkewPresent
	RR_BCast          XRRModeFlags = C.RR_BCast
	RR_PixelMultiplex XRRModeFlags = C.RR_PixelMultiplex
	RR_DoubleClock    XRRModeFlags = C.RR_DoubleClock
	RR_ClockDivideBy2 XRRModeFlags = C.RR_ClockDivideBy2
)

type XRRScreenSize C.XRRScreenSize

func (c *XRRScreenSize) Width() int {
	return int(c.width)
}

func (c *XRRScreenSize) Height() int {
	return int(c.height)
}

func (c *XRRScreenSize) Mwidth() int {
	return int(c.mwidth)
}

func (c *XRRScreenSize) Mheight() int {
	return int(c.mheight)
}

/*
//  Events.

typedef struct {
    int type;			/* event base
    unsigned long serial;	/* # of last request processed by server
    Bool send_event;		/* true if this came from a SendEvent request
    Display *display;		/* Display the event was read from
    Window window;		/* window which selected for this event
    Window root;		/* Root window for changed screen
    Time timestamp;		/* when the screen change occurred
    Time config_timestamp;	/* when the last configuration change
    SizeID size_index;
    SubpixelOrder subpixel_order;
    Rotation rotation;
    int width;
    int height;
    int mwidth;
    int mheight;
} XRRScreenChangeNotifyEvent;

typedef struct {
    int type;			/* event base
    unsigned long serial;	/* # of last request processed by server
    Bool send_event;		/* true if this came from a SendEvent request
    Display *display;		/* Display the event was read from
    Window window;		/* window which selected for this event
    int subtype;		/* RRNotify_ subtype
} XRRNotifyEvent;

typedef struct {
    int type;			/* event base
    unsigned long serial;	/* # of last request processed by server
    Bool send_event;		/* true if this came from a SendEvent request
    Display *display;		/* Display the event was read from
    Window window;		/* window which selected for this event
    int subtype;		/* RRNotify_OutputChange
    RROutput output;		/* affected output
    RRCrtc crtc;	    	/* current crtc (or None)
    RRMode mode;	    	/* current mode (or None)
    Rotation rotation;		/* current rotation of associated crtc
    Connection connection;	/* current connection status
    SubpixelOrder subpixel_order;
} XRROutputChangeNotifyEvent;

typedef struct {
    int type;			/* event base
    unsigned long serial;	/* # of last request processed by server
    Bool send_event;		/* true if this came from a SendEvent request
    Display *display;		/* Display the event was read from
    Window window;		/* window which selected for this event
    int subtype;		/* RRNotify_CrtcChange
    RRCrtc crtc;    		/* current crtc (or None)
    RRMode mode;	    	/* current mode (or None)
    Rotation rotation;		/* current rotation of associated crtc
    int x, y;			/* position
    unsigned int width, height;	/* size
} XRRCrtcChangeNotifyEvent;

typedef struct {
    int type;			/* event base
    unsigned long serial;	/* # of last request processed by server
    Bool send_event;		/* true if this came from a SendEvent request
    Display *display;		/* Display the event was read from
    Window window;		/* window which selected for this event
    int subtype;		/* RRNotify_OutputProperty
    RROutput output;		/* related output
    Atom property;		/* changed property
    Time timestamp;		/* time of change
    int state;			/* NewValue, Deleted
} XRROutputPropertyNotifyEvent;

*/
type XRRScreenConfiguration C.XRRScreenConfiguration

// returns status, event_base, error_base
func XRRQueryExtension(display *Display) (int, int, int) {
	var event, error C.int
	ret := C.XRRQueryExtension((*C.Display)(display), &event, &error)
	return int(ret), int(event), int(error)
}

// returns status, major, minor
func XRRQueryVersion(display *Display) (int, int, int) {
	var major, minor C.int
	ret := C.XRRQueryVersion((*C.Display)(display), &major, &minor)
	return int(ret), int(major), int(minor)
}

// Safe to call always, even without RandR
// Rotation XRRRotations(Display *dpy, int screen, Rotation *current_rotation);

// Safe to call always, even without RandR
// XRRScreenSize *XRRSizes(Display *dpy, int screen, int *nsizes);

// Safe to call always, even without RandR
func XRRSizes(display *Display, screen int) ([]*XRRScreenSize, func()) {
	var nSizes C.int
	sizes := C.XRRSizes((*C.Display)(display), C.int(screen), &nSizes)

	slice := make([]*XRRScreenSize, nSizes)
	for i, _ := range slice {
		size := (XRRScreenSize)(C.XRRScreenSizeAtIndex(sizes, C.int(i)))
		slice[i] = &size
	}
	return slice, func() {
		C.XFree(unsafe.Pointer(sizes))
	}
}

// Safe to call always, even without RandR
func XRRRates(display *Display, screen int, sizeID int) []int16 {
	var nRates C.int

	rates := C.XRRRates((*C.Display)(display), C.int(screen), C.int(sizeID), &nRates)
	slice := make([]int16, int(nRates))
	for i, _ := range slice {
		slice[i] = int16(C.shortAtIndex(rates, C.int(i)))
	}
	return slice
}

// Safe to call always, even without RandR
// Time XRRTimes (Display *dpy, int screen, Time *config_timestamp);

// 1.0
// void XRRFreeScreenConfigInfo (XRRScreenConfiguration *config);

// 1.0
func XRRGetScreenInfo(display *Display, window Window) (*XRRScreenConfiguration, func()) {
	info := C.XRRGetScreenInfo((*C.Display)(display), C.Window(window))
	return (*XRRScreenConfiguration)(info), func() {
		C.XRRFreeScreenConfigInfo(info)
	}
}

// 1.0
func XRRSetScreenConfig(display *Display, config *XRRScreenConfiguration, draw Drawable, size_index int, rotation Rotation, timestamp Time) int {
	return int(C.XRRSetScreenConfig((*C.Display)(display), (*C.XRRScreenConfiguration)(config), C.Drawable(draw), C.int(size_index), C.Rotation(rotation), C.Time(timestamp)))
}

// 1.0
// Rotation XRRConfigRotations(XRRScreenConfiguration *config, Rotation *current_rotation);

// 1.0
// Time XRRConfigTimes (XRRScreenConfiguration *config, Time *config_timestamp);

// 1.0
func XRRConfigSizes(config *XRRScreenConfiguration) []*XRRScreenSize {
	var nSizes C.int

	sizes := C.XRRConfigSizes((*C.XRRScreenConfiguration)(config), &nSizes)
	slice := make([]*XRRScreenSize, nSizes)
	for i, _ := range slice {
		size := (XRRScreenSize)(C.XRRScreenSizeAtIndex(sizes, C.int(i)))
		slice[i] = &size
	}
	return slice
}

// 1.0
func XRRConfigCurrentConfiguration(config *XRRScreenConfiguration) (SizeID, Rotation) {
	var rotation C.Rotation
	ret := C.XRRConfigCurrentConfiguration((*C.XRRScreenConfiguration)(config), &rotation)
	return SizeID(ret), Rotation(rotation)
}

// 1.0
func XRRRootToScreen(display *Display, root Window) int {
	return int(C.XRRRootToScreen((*C.Display)(display), C.Window(root)))
}

// 1.0
func XRRSelectInput(display *Display, window Window, mask int) {
	C.XRRSelectInput((*C.Display)(display), C.Window(window), C.int(mask))
}

// 1.1
func XRRSetScreenConfigAndRate(display *Display, config *XRRScreenConfiguration, draw Drawable, size_index int, rotation Rotation, rate int16, timestamp Time) int {
	return int(C.XRRSetScreenConfigAndRate((*C.Display)(display), (*C.XRRScreenConfiguration)(config), C.Drawable(draw), C.int(size_index), C.Rotation(rotation), C.short(rate), C.Time(timestamp)))
}

// 1.1
func XRRConfigRates(config *XRRScreenConfiguration, sizeID int) []int16 {
	var nRates C.int

	rates := C.XRRConfigRates((*C.XRRScreenConfiguration)(config), C.int(sizeID), &nRates)

	slice := make([]int16, int(nRates))
	for i, _ := range slice {
		slice[i] = int16(C.shortAtIndex(rates, C.int(i)))
	}
	C.XFree(unsafe.Pointer(rates))
	return slice
}

// 1.1
func XRRConfigCurrentRate(config *XRRScreenConfiguration) int16 {
	return int16(C.XRRConfigCurrentRate((*C.XRRScreenConfiguration)(config)))
}

// 1.2
type XRRModeInfo C.XRRModeInfo

/*
typedef struct _XRRModeInfo {
    RRMode		id;
    unsigned int	width;
    unsigned int	height;
    unsigned long	dotClock;
    unsigned int	hSyncStart;
    unsigned int	hSyncEnd;
    unsigned int	hTotal;
    unsigned int	hSkew;
    unsigned int	vSyncStart;
    unsigned int	vSyncEnd;
    unsigned int	vTotal;
    char		*name;
    unsigned int	nameLength;
    XRRModeFlags	modeFlags;
} XRRModeInfo;
*/
func (c *XRRModeInfo) Id() RRMode {
	return RRMode(c.id)
}

func (c *XRRModeInfo) Width() uint {
	return uint(c.width)
}

func (c *XRRModeInfo) Height() uint {
	return uint(c.height)
}

func (c *XRRModeInfo) DotClock() uint {
	return uint(c.dotClock)
}

func (c *XRRModeInfo) HTotal() uint {
	return uint(c.hTotal)
}

func (c *XRRModeInfo) VTotal() uint {
	return uint(c.vTotal)
}

func (c *XRRModeInfo) ModeFlags() XRRModeFlags {
	return XRRModeFlags(c.modeFlags)
}

// 1.2
type XRRScreenResources C.XRRScreenResources

/*
typedef struct _XRRScreenResources {
    Time	timestamp;
    Time	configTimestamp;
    int		ncrtc;
    RRCrtc	*crtcs;
    int		noutput;
    RROutput	*outputs;
    int		nmode;
    XRRModeInfo	*modes;
} XRRScreenResources;
*/

func (c *XRRScreenResources) Crtcs() []RRCrtc {
	slice := make([]RRCrtc, int(c.ncrtc))
	for i, _ := range slice {
		slice[i] = RRCrtc(C.RRCrtcAtIndex(c.crtcs, C.int(i)))
	}
	return slice
}

func (c *XRRScreenResources) Modes() []*XRRModeInfo {
	slice := make([]*XRRModeInfo, int(c.nmode))
	for i, _ := range slice {
		info := XRRModeInfo(C.XRRModeInfoAtIndex(c.modes, C.int(i)))
		slice[i] = &info
	}
	return slice
}

// 1.2
func XRRGetScreenResources(display *Display, window Window) (*XRRScreenResources, func()) {
	resources := C.XRRGetScreenResources((*C.Display)(display), C.Window(window))
	return (*XRRScreenResources)(resources), func() {
		C.XRRFreeScreenResources(resources)
	}
}

// 1.2
// void XRRFreeScreenResources (XRRScreenResources *resources);

/*
typedef struct _XRROutputInfo {
    Time	    timestamp;
    RRCrtc	    crtc;
    char	    *name;
    int		    nameLen;
    unsigned long   mm_width;
    unsigned long   mm_height;
    Connection	    connection;
    SubpixelOrder   subpixel_order;
    int		    ncrtc;
    RRCrtc	    *crtcs;
    int		    nclone;
    RROutput	    *clones;
    int		    nmode;
    int		    npreferred;
    RRMode	    *modes;
} XRROutputInfo;
*/

// 1.2
type XRROutputInfo C.XRROutputInfo

func (c *XRROutputInfo) Mm_width() uint {
	return uint(c.mm_width)
}

func (c *XRROutputInfo) Mm_height() uint {
	return uint(c.mm_height)
}

func (c *XRROutputInfo) Modes() []RRMode {
	slice := make([]RRMode, c.nmode)
	for i, _ := range slice {
		slice[i] = RRMode(C.RRModeAtIndex(c.modes, C.int(i)))
	}
	return slice
}

// 1.2
func XRRGetOutputInfo(display *Display, resources *XRRScreenResources, output RROutput) (*XRROutputInfo, func()) {
	info := C.XRRGetOutputInfo((*C.Display)(display), (*C.XRRScreenResources)(resources), C.RROutput(output))
	return (*XRROutputInfo)(info), func() {
		C.XRRFreeOutputInfo(info)
	}
}

/*
Atom *
XRRListOutputProperties (Display *dpy, RROutput output, int *nprop);

typedef struct {
    Bool    pending;
    Bool    range;
    Bool    immutable;
    int	    num_values;
    long    *values;
} XRRPropertyInfo;

XRRPropertyInfo *
XRRQueryOutputProperty (Display *dpy, RROutput output, Atom property);

XRRModeInfo *
XRRAllocModeInfo (char *name, int nameLength);

RRMode
XRRCreateMode (Display *dpy, Window window, XRRModeInfo *modeInfo);

void
XRRDestroyMode (Display *dpy, RRMode mode);

void
XRRAddOutputMode (Display *dpy, RROutput output, RRMode mode);

void
XRRDeleteOutputMode (Display *dpy, RROutput output, RRMode mode);

void
XRRFreeModeInfo (XRRModeInfo *modeInfo);
*/

type XRRCrtcInfo C.XRRCrtcInfo

/*
&{377577560 0 0 0 0 0 1 [8 0] 0 0x11d9eda0 63 [223 3] 4 0x11d9eda0}

typedef struct _XRRCrtcInfo {
    Time	    timestamp;
    int		    x, y;
    unsigned int    width, height;
    RRMode	    mode;
    Rotation	    rotation;
    int		    noutput;
    RROutput	    *outputs;
    Rotation	    rotations;
    int		    npossible;
    RROutput	    *possible;
} XRRCrtcInfo;
*/

func (c *XRRCrtcInfo) Free() {
	C.XRRFreeCrtcInfo((*C.XRRCrtcInfo)(c))
}

func (c *XRRCrtcInfo) X() int {
	return int(c.x)
}

func (c *XRRCrtcInfo) Y() int {
	return int(c.y)
}

func (c *XRRCrtcInfo) Width() uint {
	return uint(c.width)
}

func (c *XRRCrtcInfo) Height() uint {
	return uint(c.height)
}

func (c *XRRCrtcInfo) Mode() RRMode {
	return RRMode(c.mode)
}

func (c *XRRCrtcInfo) Rotation() Rotation {
	return Rotation(c.rotation)
}

func (c *XRRCrtcInfo) Outputs() []RROutput {
	slice := make([]RROutput, int(c.noutput))
	for i, _ := range slice {
		slice[i] = RROutput(C.RROutputAtIndex(c.outputs, C.int(i)))
	}
	return slice
}

// 1.2
func XRRGetCrtcInfo(display *Display, resources *XRRScreenResources, crtc RRCrtc) *XRRCrtcInfo {
	return (*XRRCrtcInfo)(C.XRRGetCrtcInfo((*C.Display)(display), (*C.XRRScreenResources)(resources), C.RRCrtc(crtc)))
}

// 1.2
// void XRRFreeCrtcInfo (XRRCrtcInfo *crtcInfo);

// 1.2
func XRRSetCrtcConfig(display *Display, resources *XRRScreenResources, crtc RRCrtc, timestamp Time, x, y int, mode RRMode, rotation Rotation, outputs []RROutput) int {
	cOutputs := C.newRROutputArray(C.int(len(outputs)))

	for i, v := range outputs {
		C.RROutputSetAtIndex(cOutputs, C.int(i), C.RROutput(v))
	}

	return int(C.XRRSetCrtcConfig((*C.Display)(display), (*C.XRRScreenResources)(resources), C.RRCrtc(crtc), C.Time(timestamp), C.int(x), C.int(y), C.RRMode(mode), C.Rotation(rotation), cOutputs, C.int(len(outputs))))
}

// 1.2
func XRRGetCrtcGammaSize(display *Display, crtc RRCrtc) int {
	return int(C.XRRGetCrtcGammaSize((*C.Display)(display), C.RRCrtc(crtc)))
}

// 1.2
type XRRCrtcGamma struct {
	Size             int
	Red, Green, Blue []uint16
}

func (gamma *XRRCrtcGamma) C() (*C.XRRCrtcGamma, func()) {
	c := C.XRRAllocGamma(C.int(gamma.Size))

	// Copy
	for i := 0; i < gamma.Size; i++ {
		C.ushortSetAtIndex(c.red, C.int(i), C.ushort(gamma.Red[i]))
		C.ushortSetAtIndex(c.green, C.int(i), C.ushort(gamma.Green[i]))
		C.ushortSetAtIndex(c.blue, C.int(i), C.ushort(gamma.Blue[i]))
	}

	return c, func() {
		C.XRRFreeGamma((*C.XRRCrtcGamma)(c))
	}
}

func CToGoXRRCrtcGamma(c *C.XRRCrtcGamma) *XRRCrtcGamma {
	gamma := &XRRCrtcGamma{}
	size := int(c.size)
	gamma.Red = make([]uint16, size)
	gamma.Green = make([]uint16, size)
	gamma.Blue = make([]uint16, size)

	for i := 0; i < size; i++ {
		gamma.Red[i] = uint16(C.ushortAtIndex(c.red, C.int(i)))
		gamma.Green[i] = uint16(C.ushortAtIndex(c.green, C.int(i)))
		gamma.Blue[i] = uint16(C.ushortAtIndex(c.blue, C.int(i)))
	}
	return gamma
}

// 1.2
// XRRCrtcGamma* XRRGetCrtcGamma (Display *dpy, RRCrtc crtc);
func XRRGetCrtcGamma(display *Display, crtc RRCrtc) *XRRCrtcGamma {
	cgamma := C.XRRGetCrtcGamma((*C.Display)(display), C.RRCrtc(crtc))
	defer C.XRRFreeGamma(cgamma)
	return CToGoXRRCrtcGamma(cgamma)
}

// 1.2
// XRRCrtcGamma* XRRAllocGamma (int size);

// 1.2
func XRRSetCrtcGamma(display *Display, crtc RRCrtc, gamma *XRRCrtcGamma) {
	cgamma, free := gamma.C()
	defer free()
	C.XRRSetCrtcGamma((*C.Display)(display), C.RRCrtc(crtc), (*C.XRRCrtcGamma)(cgamma))
}

// 1.2
// void XRRFreeGamma (XRRCrtcGamma *gamma);

// 1.3
func XRRGetScreenResourcesCurrent(display *Display, window Window) (*XRRScreenResources, func()) {
	resources := C.XRRGetScreenResourcesCurrent((*C.Display)(display), C.Window(window))
	return (*XRRScreenResources)(resources), func() {
		C.XRRFreeScreenResources(resources)
	}
}
