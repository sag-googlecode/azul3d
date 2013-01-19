// Copyright 2012 Lightpoke. All rights reserved.
// Use of this source code is governed by an BSD
// license found in the License.txt file

package chippy

import(
    "sync"
    "fmt"
    "sort"
)

type sortedScreenModes []*ScreenMode
func (s sortedScreenModes) Len() int { return len(s) }
func (s sortedScreenModes) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s sortedScreenModes) Less(i, j int) bool {
    iWeight := float32(s[i].height) + float32(s[i].width) + s[i].refreshRate
    jWeight := float32(s[j].height) + float32(s[j].width) + s[j].refreshRate
    return iWeight > jWeight
}


type Screen struct {
    backend_Screen

    valid bool
    access sync.RWMutex

    autoRestoreOriginalGamma bool
    restoreGammaRampCallback *callback
    originalGammaRamp *GammaRamp
    brightness, contrast, gamma float32

    screenMode *ScreenMode
    physicalWidth, physicalHeight float32
}

func newScreen() *Screen {
    s := &Screen{}
    s.valid = true
    s.physicalWidth = -1
    s.physicalHeight = -1
    s.restoreGammaRampCallback = &callback{func() {
        original := s.OriginalGammaRamp()
        if original != nil {
            s.SetGammaRamp(original)
        }
    }}
    addDestroyCallback(s.restoreGammaRampCallback)
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
    return fmt.Sprintf("Screen(ScreenMode=%s, physicalWidth=%.1fmm, physicalHeight=%.1fmm)", s.screenMode.String(), s.physicalWidth, s.physicalHeight)
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

// SetScreenMode changes this Screen's ScreenMode to the one specified, this operation
// may take a few seconds on most window managers, and it also *could* fail, so you should
// confirm with the user somehow that the new ScreenMode is appropriate and looks good.
func (s *Screen) SetScreenMode(m *ScreenMode) {
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
func (s *Screen) ScreenModes() []*ScreenMode {
    var sorted sortedScreenModes = s.screenModes()
    sort.Sort(sorted)
    return sorted
}

// SetGammaRamp sets the specified gamma ramp to be active on this Screen, note that not all video
// cards support gamma ramps.
//
// An error is returned if we where unable to set the GammaRamp due to some reason (such as lack of
// support by the hardware, or current extensions on Linux/X11, etc)
func (s *Screen) SetGammaRamp(gammaRamp *GammaRamp) error {
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
// SetGammaRamp on this Screen, or 0 and an error in the event that there is no support for setting
// GammaRamp's on this Screen for some reason (such as the user missing X extensions required on
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

// SetAutoRestoreOriginalGamma changes weather when chippy.Destroy() is called, this Screen will
// auto restore the original gamma (as returned by OriginalGammaRamp()), or will leave it (causing
// the gamma to stay active after the program exits).
//
// By default this is already on (and I reccomend against turning it off, and potentially angering
// users, unless your application is an monitor settings application)
func (s *Screen) SetAutoRestoreOriginalGamma(restore bool) {
    s.access.Lock()
    defer s.access.Unlock()
    s.autoRestoreOriginalGamma = restore
    if restore {
        addDestroyCallback(s.restoreGammaRampCallback)
    } else {
        removeDestroyCallback(s.restoreGammaRampCallback)
    }
}

// AutoRestoreOriginalGamma tells weather this Screen is currently set to auto restore the original
// gamma ramp once chippy.Destroy() is called.
//
// See SetAutoRestoreOriginalGamma() for more information.
func (s *Screen) AutoRestoreOriginalGamma() bool {
    s.access.RLock()
    defer s.access.RUnlock()
    return s.autoRestoreOriginalGamma
}

// SetBrightnessContrastGamma creates an new GammaRamp, using the specified parameters, brightness,
// contrast, and gamma, respectively, and then applies it using SetGammaRamp().
//
// An error is returned if there was an problem setting the gamma ramp (and thus being unable to
// set the brightness, contrast, or gamma)
//
// An error is returned if brightness is outside the range of -1.0 to 1.0.
//
// An error is returned if contrast is an negative number.
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

// Screens returns all possible Screen's
func Screens() []*Screen {
    return backend_Screens()
}

// DefaultScreen returns the default Screen, that is, the default Screen to create new Window's on,
// according to what the window manager says.
func DefaultScreen() *Screen {
    return backend_DefaultScreen()
}

