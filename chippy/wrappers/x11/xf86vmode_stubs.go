// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

// +build linux,no_xf86vm

package x11

type (
	XF86VidModeMonitor  struct{}
	XF86VidModeModeInfo struct{}
	XF86VidModeModeLine struct{}
)

func (m *XF86VidModeMonitor) VsyncHi() float32 {
	return 0
}

func (m *XF86VidModeMonitor) Free() {
}

func (c *XF86VidModeModeInfo) Dotclock() uint {
	return 0
}

func (c *XF86VidModeModeInfo) Vtotal() uint16 {
	return 0
}

func (c *XF86VidModeModeInfo) Htotal() uint16 {
	return 0
}

func (c *XF86VidModeModeInfo) Flags() uint {
	return 0
}

func (c *XF86VidModeModeInfo) Hdisplay() uint16 {
	return 0
}

func (c *XF86VidModeModeInfo) Vdisplay() uint16 {
	return 0
}

func XF86VidModeQueryVersion(display *Display) (int, int, int) {
	return 0, 0, 0
}

func XF86VidModeQueryExtension(display *Display) (int, int, int) {
	return 0, 0, 0
}

func XF86VidModeGetMonitor(display *Display, screen int) (int, *XF86VidModeMonitor) {
	return 0, &XF86VidModeMonitor{}
}

func XF86VidModeGetModeLine(display *Display, screen int) (int, int, *XF86VidModeModeLine) {
	return 0, 0, nil
}

func XF86VidModeValidateModeLine(display *Display, screen int, modeline *XF86VidModeModeInfo) int {
	return 0
}

func XF86VidModeGetAllModeLines(display *Display, screen int) (int, []*XF86VidModeModeInfo, func()) {
	return 0, []*XF86VidModeModeInfo{}, func() {}
}

func XF86VidModeSwitchToMode(display *Display, screen int, modeline *XF86VidModeModeInfo) int {
	return 0
}

func XF86VidModeSetViewPort(display *Display, screen, x, y int) int {
	return 0
}

func XF86VidModeSetGammaRamp(display *Display, screen int, red, green, blue []uint16) int {
	return 0
}

func XF86VidModeGetGammaRampSize(display *Display, screen int) (status int, size int) {
	return 0, 0
}

func XF86VidModeGetGammaRamp(display *Display, screen int) (status int, red, green, blue []uint16) {
	return
}
