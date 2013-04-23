// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package chippy

import (
	"sync"
)

// Screen represents an single physical screen device. It is only possible to get an screen from
// either the Screens or DefaultScreen functions, creating an Screen struct yourself is strictly
// forbidden due to the fact that it represents physically attatched and active screen devices.
type Screen interface {
	// String returns an nice string representation of this Screen
	String() string

	// Name returns an formatted string of the screens name, this is something that the user
	// should be able to relate on their own to the actual physical screen device, this typically
	// includes device brand name or model etc..
	Name() string

	// PhysicalSize returns the physical width and height of this screen, in millimeters, or 0 as
	// both width and height in the event there is no way to determine the physical size of this
	// screen.
	PhysicalSize() (width float32, height float32)

	// OriginalScreenMode returns the original screen mode of this screen, as it was when this
	// screen was created.
	OriginalScreenMode() ScreenMode

	// ScreenModes returns all available screen modes on this screen, sorted by highest resolution,
	// then highest bytes per pixel, then highest refresh rate.
	ScreenModes() []ScreenMode

	// SetScreenMode switches this screen to the specified screen mode, or returns an error in the
	// event that we where unable to switch the screen to the specified screen mode.
	//
	// The newMode parameter must be an screen mode that originally came from one of the functions
	// ScreenModes, ScreenMode, or OriginalScreenModeone or else this function will panic.
	//
	// If an error is returned, it will be either ErrBadScreenMode, or ErrDualViewCapable.
	SetScreenMode(newMode ScreenMode) error

	// ScreenMode returns the current screen mode in use by this screen, this will be either the
	// last screen mode set via SetScreenMode, or the original screen mode from OriginalScreenMode
	// in the event that no screen mode was previously set on this screen.
	ScreenMode() ScreenMode

	// OriginalGammaRamp returns the original gamma ramp of this screen, as it was when this screen
	// was created, or ErrNoGammaRampSupport in the event that there is no support (due to hardware
	// or software) for gamma ramps on this screen.
	OriginalGammaRamp() (*GammaRamp, error)

	// SetGammaRamp sets the screen's gamma ramp (color correction lookup table / LUT) to the
	// specified gamma ramp, or returns ErrNoGammaRampSupport in the event that there is no support
	// (due to hardware or software) for gamma ramps on this screen.
	//
	// The ramp argument must be an gamma ramp where each slice inside the GammaRamp struct is of
	// the exact length of whatever is returned by GammaRampSize() or else an panic will occur.
	SetGammaRamp(ramp *GammaRamp) error

	// GammaRamp returns the screen's current gamma ramp (color correction lookup table / LUT) or
	// an ErrNoGammaRampSupport in the event that there is no support (due to hardware or software)
	// for gamma ramps on this screen.
	GammaRamp() (*GammaRamp, error)

	// GammaRampSize returns the size of each per-color array that an GammaRamp must have in order
	// for it to be allowed for use on this screen.
	//
	// This function may return 0, which can be considered eqivilent of ErrNoGammaRampSupport,
	// meaning that there is no support (due to hardware or software) for gamma ramps on this
	// screen.
	//
	// Even though GammaRamp has Red, Green, and Blue defined as slices, they actually must contain
	// an exact number of elements (that is, the number returned by this function).
	GammaRampSize() uint
}

var (
	screenCacheLock       sync.RWMutex
	cachedScreens         []Screen
	queriedScreensAlready bool
)

// Screens returns all available, attatched, and activated screens possible. Once this function is
// called, the result is cached such that future calls to this function are faster and return the
// cached result.
//
// To update the internal screen cache, see the RefreshScreens function.
func Screens() []Screen {
	screenCacheLock.RLock()
	screenCacheLock.RUnlock()

	if !queriedScreensAlready {
		RefreshScreens()
	}
	return cachedScreens
}

// RefreshScreens queries for all available screens, and updates the internal cache returned by the
// Screens() function, such that the Screens() function returns newly attatched or detatched Screen
// devices.
func RefreshScreens() {
	screenCacheLock.Lock()
	defer screenCacheLock.Unlock()

	queriedScreensAlready = true
	cachedScreens = backend_Screens()
}

// DefaultScreen returns the 'default' screen, this is determined by either the window manager
// itself (as per an user's person computer setup and configuration) or will be guessed by Chippy.
//
// It is possible for this function to return nil, in the unlikely event that Screens() returns no
// screens at all, due to an user having none plugged in or activated.
func DefaultScreen() Screen {
	return backend_DefaultScreen()
}
