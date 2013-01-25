package chippy

import (
	"fmt"
	"sync"
)

// ScreenMode represents an single mode that an Screen is capable of being in, it specifies an
// resolution, and refresh rate.
//
// On platforms that support multiple different bpp (bytes per pixel) choices for the same
// resolution, you will only ever see the largest bitrate available for that resolution.
type ScreenMode struct {
	backend_ScreenMode

	valid  bool
	access sync.RWMutex

	screen        *Screen
	width, height uint
	refreshRate   float32
}

func newScreenMode(screen *Screen) *ScreenMode {
	m := &ScreenMode{}
	m.valid = true
	m.screen = screen
	m.refreshRate = -1
	return m
}

func (m *ScreenMode) panicUnlessValid() {
	if !m.valid {
		panic("ScreenMode is invalid! It must have came from Chippy, not made by you!")
	}
}

// String returns an string representation of this ScreenMode
func (m *ScreenMode) String() string {
	return fmt.Sprintf("ScreenMode(width=%dpx, height=%dpx, refreshRate=%.2fhz)", m.width, m.height, m.refreshRate)
}

// Equals tells weather this ScreenMode is equal to other ScreenMode, this just does comparison
// between refresh rate, width, and height
func (m *ScreenMode) Equals(other *ScreenMode) bool {
	if m.RefreshRate() != other.RefreshRate() {
		return false
	}
	if m.Width() != other.Width() {
		return false
	}
	if m.Height() != other.Height() {
		return false
	}
	return true
}

// Screen returns the screen that this ScreenMode came from, that is to say that this ScreenMode is
// an screen mode from the returned Screen.
func (m *ScreenMode) Screen() *Screen {
	m.access.RLock()
	defer m.access.RUnlock()
	return m.screen
}

// Width returns the width (in pixels) of this ScreenMode
func (m *ScreenMode) Width() uint {
	m.access.RLock()
	defer m.access.RUnlock()
	return m.width
}

// Height returns the height (in pixels) of this ScreenMode
func (m *ScreenMode) Height() uint {
	m.access.RLock()
	defer m.access.RUnlock()
	return m.height
}

// RefreshRate returns the refresh rate (in hertz) of this ScreenMode
func (m *ScreenMode) RefreshRate() float32 {
	m.access.RLock()
	defer m.access.RUnlock()
	return m.refreshRate
}

// Use is an helper and just calls m.Screen().SetScreenMode(m), take a look
// at Screen.SetScreenMode for more information.
func (m *ScreenMode) Use() {
	m.Screen().SetScreenMode(m)
}
