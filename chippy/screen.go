// Copyright 2012 Lightpoke. All rights reserved.
// Use of this source code is governed by an BSD
// license found in the License.txt file

package chippy

import(
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

var(
	screenCacheLock sync.RWMutex
	cachedScreens []Screen
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







/*
import (
	"fmt"
	"sort"
	"sync"
)

// Screen represents an physical screen device, such as an physical computer Monitor, or possibly
// an screen on an mobile phone etc.
type Screen struct {
	backend_Screen

	valid  bool
	access sync.RWMutex

	autoRestoreOriginalGammaRamp bool
	restoreGammaRampCallback     *callback
	originalGammaRamp            *GammaRamp
	brightness, contrast, gamma  float32

	autoRestoreOriginalScreenMode  bool
	restoreScreenModeCallback      *callback
	screenMode, originalScreenMode *ScreenMode

	name                          string
	physicalWidth, physicalHeight float32
}

func newScreen() *Screen {
	s := &Screen{}
	s.valid = true
	s.physicalWidth = -1
	s.physicalHeight = -1

	s.autoRestoreOriginalGammaRamp = true
	s.restoreGammaRampCallback = &callback{func() {
		original := s.OriginalGammaRamp()
		if original != nil {
			s.SetGammaRamp(original)
		}
	}}

	s.autoRestoreOriginalScreenMode = true
	s.restoreScreenModeCallback = &callback{func() {
		original := s.OriginalScreenMode()
		if original != nil {
			s.SetScreenMode(original)
		}
	}}
	return s
}

func (s *Screen) panicUnlessValid() {
	if !s.valid {
		panic("Screen is invalid! It must have came from Chippy, not made by you!")
	}
}

// String returns an string representation of this Screen
func (s *Screen) String() string {
	s.access.RLock()
	defer s.access.RUnlock()
	if s.screenMode == nil {
		panic("screen mode is nil, this should never ever happen!")
	}
	return fmt.Sprintf("Screen(name=%q, ScreenMode=%s, physicalWidth=%.1fmm, physicalHeight=%.1fmm)", s.name, s.screenMode.String(), s.physicalWidth, s.physicalHeight)
}

// Name returns an string name of this Screen, this is an name that the user will hopefully be able
// to recognize on their own. Typically this will be some combination of graphics card device name
// mixed with actual monitor name given by manufacturer, or monitor type, etc.
//
// If there is no known name, this function simply returns an empty string.
func (s *Screen) Name() string {
	s.access.RLock()
	defer s.access.RUnlock()
	return s.name
}

// PhysicalWidth returns the Screen's physical width (in millimeters), or -1 if there is no known
// physical width.
//
// Note: Some drivers, graphics cards, and monitors report incorrect information, or incorrect EDID
// information, making this incredibly inaccurate / sometimes scarilly wrong. Be careful using this
// without some sanity checks or any override.
func (s *Screen) PhysicalWidth() float32 {
	s.access.RLock()
	defer s.access.RUnlock()
	return s.physicalWidth
}

// PhysicalHeight returns the Screen's physical height (in millimeters), or -1 if there is no known
// physical width.
//
// Note: Some drivers, graphics cards, and monitors report incorrect information, or incorrect EDID
// information, making this incredibly inaccurate / sometimes scarilly wrong. Be careful using this
// without some sanity checks or any override.
func (s *Screen) PhysicalHeight() float32 {
	s.access.RLock()
	defer s.access.RUnlock()
	return s.physicalHeight
}

// OriginalScreenMode returns the original screen mode of this Screen, as it was when this Screen
// was created, or nil if if there is no original screen mode for some reason
func (s *Screen) OriginalScreenMode() *ScreenMode {
	s.access.RLock()
	defer s.access.RUnlock()
	return s.originalScreenMode
}

// SetAutoRestoreOriginalScreenMode changes weather when chippy.Destroy() is called, this Screen
// will auto restore the original screen mode (as returned by OriginalScreenMode()), or will leave
// it (causing the screen mode to stay active after the program exits, for instance).
//
// By default this is already on (and I recommend against turning it off, and potentially angering
// users).
func (s *Screen) SetAutoRestoreOriginalScreenMode(restore bool) {
	s.access.Lock()
	defer s.access.Unlock()
	s.autoRestoreOriginalScreenMode = restore
	if restore {
		addDestroyCallback(s.restoreScreenModeCallback)
	} else {
		removeDestroyCallback(s.restoreScreenModeCallback)
	}
}

// AutoRestoreOriginalScreenMode tells weather this Screen is currently set to auto restore the
// original screen mode once chippy.Destroy() is called.
//
// See SetAutoRestoreOriginalScreenMode() for more information.
func (s *Screen) AutoRestoreOriginalScreenMode() bool {
	s.access.RLock()
	defer s.access.RUnlock()
	return s.autoRestoreOriginalScreenMode
}

// SetScreenMode changes this Screen's ScreenMode to the one specified, this operation
// may take a few seconds on most window managers, and it also *could* fail, so you should
// confirm with the user somehow that the new ScreenMode is appropriate and looks good to them.
//
// Note: The original screen mode of this screen will be automatically restored when chippy.Destroy
// is finally called. If you would like to turn off this default behavior, then call SetAutoRestoreOriginalScreenMode(false).
func (s *Screen) SetScreenMode(m *ScreenMode) {
	if s.AutoRestoreOriginalScreenMode() {
		s.SetAutoRestoreOriginalScreenMode(true)
	}

	m.panicUnlessValid()
	s.access.Lock()
	defer s.access.Unlock()
	s.screenMode = m

	s.setScreenMode()
}

// ScreenMode returns this Screen's currently in use ScreenMode
func (s *Screen) ScreenMode() *ScreenMode {
	s.access.RLock()
	defer s.access.RUnlock()
	return s.screenMode
}

// ScreenModes returns all possible ScreenMode's for this Screen, sorted by highest screen
// resolution and refresh rate.
//
// This can return an empty slice in the event that there are no screens available at all.
func (s *Screen) ScreenModes() []*ScreenMode {
	var sorted sortedScreenModes = s.screenModes()
	sort.Sort(sorted)
	return sorted
}

// SetGammaRamp sets the specified GammaRamp to be active on this Screen. Each slice in the
// GammaRamp must be of the size returned by GammaRampSize() or an error is returned.
//
// Note: Only some Monitors, graphics cards, and video drivers support gamma ramps. An descriptive error is returned in this case, or if some other error occurs (missing extensions on Linux/X11 for instance).
//
// Note: Windows places very large restrictions on how much you may actually change the colors in
// an devices gamma ramp. To be precise, you are only allowed to change the last 128 values for
// each color intensity (For example, you have an max range of 128-256 possible values) there is no
// great workaround for this, other than modifying the registry and launching your application
// under an specific user account.
//
// See this blog for more information about this presumed Windows 'feature':
//
// http://jonls.dk/2010/09/windows-gamma-adjustments/
func (s *Screen) SetGammaRamp(gammaRamp *GammaRamp) error {
	if s.AutoRestoreOriginalGammaRamp() {
		s.SetAutoRestoreOriginalGammaRamp(true)
	}

	s.access.Lock()
	defer s.access.Unlock()
	return s.setGammaRamp(gammaRamp)
}

// GammaRamp returns the GammaRamp currently in use, provided by the operating system, or nil and
// an error in the event we are unable to obtain the GammaRamp in use due to some reason. (Such as
// the user missing X extensions required on Linux/X11 for example.)
func (s *Screen) GammaRamp() (*GammaRamp, error) {
	s.access.RLock()
	defer s.access.RUnlock()
	return s.gammaRamp()
}

// GammaRampSize returns the size that an GammaRamp must be for it to be accepted in an call to
// SetGammaRamp on this Screen, or 0 and an descriptive error as to why there is no support for
// settings the GammaRamp on this Screen (such as the user missing X extensions required on
// Linux/X11 for example).
func (s *Screen) GammaRampSize() (int, error) {
	s.access.RLock()
	defer s.access.RUnlock()
	return s.gammaRampSize()
}

// OriginalGammaRamp returns the gamma ramp of this Screen, as it was when this Screen was created,
// or nil if there is no gamma ramp support on this Screen for some reason.
func (s *Screen) OriginalGammaRamp() *GammaRamp {
	s.access.RLock()
	defer s.access.RUnlock()
	return s.originalGammaRamp
}

// SetAutoRestoreOriginalGammaRamp changes weather when chippy.Destroy() is called, this Screen will
// auto restore the original gamma (as returned by OriginalGammaRamp()), or will leave it (causing
// the gamma to stay active after the program exits).
//
// By default this is already on (and I recommend against turning it off, and potentially angering
// users).
func (s *Screen) SetAutoRestoreOriginalGammaRamp(restore bool) {
	s.access.Lock()
	defer s.access.Unlock()
	s.autoRestoreOriginalGammaRamp = restore
	if restore {
		addDestroyCallback(s.restoreGammaRampCallback)
	} else {
		removeDestroyCallback(s.restoreGammaRampCallback)
	}
}

// AutoRestoreOriginalGammaRamp tells weather this Screen is currently set to auto restore the original
// gamma ramp once chippy.Destroy() is called.
//
// See SetAutoRestoreOriginalGammaRamp() for more information.
func (s *Screen) AutoRestoreOriginalGammaRamp() bool {
	s.access.RLock()
	defer s.access.RUnlock()
	return s.autoRestoreOriginalGammaRamp
}

// SetBrightnessContrastGamma creates an new GammaRamp, using the specified parameters, brightness,
// contrast, and gamma, respectively, and then applies it using SetGammaRamp().
//
// An error is returned if brightness is outside the inclusive range of -1.0 to 1.0.
//
// An error is returned if contrast is an negative number.
//
// An error is returned if there was an problem setting the gamma ramp (and thus being unable to
// set the brightness, contrast, or gamma all together)
func (s *Screen) SetBrightnessContrastGamma(brightness, contrast, gamma float32) error {
	size, err := s.GammaRampSize()
	if err != nil {
		return err
	}

	r := &GammaRamp{}
	err = r.SetAsBrightnessContrastGamma(size, brightness, contrast, gamma)
	if err != nil {
		return err
	}

	err = s.SetGammaRamp(r)
	if err != nil {
		return err
	}

	s.access.Lock()
	defer s.access.Unlock()
	s.brightness = brightness
	s.contrast = contrast
	s.gamma = gamma
	return nil
}

// BrightnessContrastGamma returns the brightness, contrast, and gamma, as they where previously
// passed into an call to SetBrightnessContrastGamma (or zero if there was no previous call)
func (s *Screen) BrightnessContrastGamma() (brightness, contrast, gamma float32) {
	s.access.RLock()
	defer s.access.RUnlock()
	return s.brightness, s.contrast, s.gamma
}
*/
