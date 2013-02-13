package chippy

// We use this type for sorting the ScreenModes in backends
type sortedScreenModes []ScreenMode

func (s sortedScreenModes) Len() int {
    return len(s)
}

func (s sortedScreenModes) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

func (s sortedScreenModes) Less(i, j int) bool {
    iWidth, iHeight := s[i].Resolution()
	iResolution := iWidth + iHeight
    iRefreshRate := s[i].RefreshRate()
    iBytesPerPixel := s[i].BytesPerPixel()

    jWidth, jHeight := s[j].Resolution()
	jResolution := jWidth + jHeight
    jRefreshRate := s[j].RefreshRate()
    jBytesPerPixel := s[j].BytesPerPixel()

	// if resolution and bpp are the same, sort by refresh rate
	if iResolution == jResolution && iBytesPerPixel == jBytesPerPixel {
		return iRefreshRate > jRefreshRate

	// Second case, if resolution is the same, sort by bpp
	} else if iResolution == jResolution {
		return iBytesPerPixel > jBytesPerPixel
	}

	// First case, sort by resolution
	return iResolution > jResolution

	/*
	// Multiply by 100.0 to get rid of smaller inconsitencies (where 1366+768+60+8 less than 1360+768+60+32)
	e := 100.0
    iWeight := (float32(iWidth) * e) + (float32(iHeight) * e) + (float32(iBytesPerPixel) * e) + (iRefreshRate * e)
    jWeight := (float32(jWidth) * e) + (float32(jHeight) * e) + (float32(jBytesPerPixel) * e) + (jRefreshRate * e)
	return iWeight > jWeight
	*/
}

// ScreenMode represents an single, unique, screen mode, with an resolution, refresh rate, and bpp.
//
// It is possible for multiple different ScreenMode's to exist with the same resolution, each with
// different refresh rates or bytes per pixel, respectively.
type ScreenMode interface {
    // String returns an nice string representing this ScreenMode
    String() string

    // Equals compares two ScreenMode(s) for equality. It does this by comparing resolutions,
    // refresh rates, and bytes per pixels.
    Equals(other ScreenMode) bool

    // Resolution returns the width and height of this ScreenMode, in pixels.
    Resolution() (width, height uint)

    // RefreshRate returns the refresh rate of this ScreenMode, in hertz, or 0 if the refresh rate
    // is unable to be determined.
    RefreshRate() float32

    // BytesPerPixel returns the number of bytes that represent an single pixel of this ScreenMode,
    // or 0 if the bytes per pixel is unable to be determined.
    BytesPerPixel() uint
}


/*
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

*/
