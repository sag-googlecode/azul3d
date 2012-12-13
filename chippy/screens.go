package chippy

import "fmt"

// String returns a string representation of this Resolution
func (r *Resolution) String() string {
    r.access.RLock()
    defer r.access.RUnlock()
    return fmt.Sprintf("Resolution(width=%d, height=%d)", r.width, r.height)
}

// Use This simply calls Resolution.Screen().SetResolution() for you, setting this resolution
// to be active on the Screen.
func (r *Resolution) Use() error {
    return r.Screen().SetResolution(r)
}

// Screen returns the Screen that this resolution came from
func (r *Resolution) Screen() *Screen {
    r.access.RLock()
    defer r.access.RUnlock()
    return r.screen
}

// Width returns the width in pixels of this Resolution
func (r *Resolution) Width() uint16 {
    r.access.RLock()
    defer r.access.RUnlock()
    return r.width
}

// Height returns the height in pixels of this Resolution
func (r *Resolution) Height() uint16 {
    r.access.RLock()
    defer r.access.RUnlock()
    return r.height
}



// We store all the global screen objects here
var platformDefaultScreen *Screen
var platformScreens []*Screen

// String returns a formatted string representation of this Screen
func (s *Screen) String() string {
    s.access.RLock()
    defer s.access.RUnlock()
    return fmt.Sprintf("Screen(%d, %dx%d)", s.number, s.resolution.Width(), s.resolution.Height())
}

// Number returns the number (id) of this screen, like 1, 2, 3 (starting at 1)
func (s *Screen) Number() uint16 {
    s.access.RLock()
    defer s.access.RUnlock()
    return s.number
}

// Width returns the width in pixels of this Screen's Resolution
func (s *Screen) Width() uint16 {
    s.access.RLock()
    defer s.access.RUnlock()
    return s.resolution.Width()
}

// Height returns the height in pixels of this Screen's Resolution
func (s *Screen) Height() uint16 {
    s.access.RLock()
    defer s.access.RUnlock()
    return s.resolution.Height()
}

// RestoreOriginalResolution restores the original resolution of this screen
// just like it was when the application started
//
// Chippy will automatically call this for you when you call chippy.Destroy()
//
// You can turn off this default behavior with screen.SetAutoRestoreOriginalResolution(bool)
func (s *Screen) RestoreOriginalResolution() error {
    s.access.RLock()
    oldResolution := s.origResolution
    s.access.RUnlock()// Important to release before call to SetResolution

    return s.SetResolution(oldResolution)
}


// SetAutoRestoreOriginalResolution if restore is true, chippy will automatically restore the resolution of this screen
// as it was when the application started. If restore is false, then chippy will instead leave the resolution as-is
// when the application finally calls chippy.Destroy()
func (s *Screen) SetAutoRestoreOriginalResolution(restore bool) error {
    s.access.Lock()
    defer s.access.Unlock()

    if restore != s.autoRestoreOriginalResolution {
        s.autoRestoreOriginalResolution = restore

        // Calling add/remove DestroyCallback
        chippyAccess.Lock()
        defer chippyAccess.Unlock()

        err := getInitError()
        if err != nil {
            return err
        }

        if restore == true {
            addDestroyCallback(s.autoRestoreOriginalResolutionCallback)
        } else {
            removeDestroyCallback(s.autoRestoreOriginalResolutionCallback)
        }
    }
    return nil
}

// SetResolution sets the current Resolution of this Screen
func (s *Screen) SetResolution(resolution *Resolution) error {
    s.access.Lock()
    defer s.access.Unlock()

    if resolution != s.resolution {
        s.resolution = resolution

        // Calling C things, so get lock
        chippyAccess.Lock()
        defer chippyAccess.Unlock()

        err := getInitError()
        if err != nil {
            return err
        }

        return s.setResolution()
    }
    return nil
}

// Resolution returns the current, in use, Resolution of this Screen
func (s *Screen) Resolution() *Resolution {
    s.access.RLock()
    defer s.access.RUnlock()
    return s.resolution
}

// OriginalResolution returns the original Resolution of this Screen, as it was when Chippy
// was initialized
func (s *Screen) OriginalResolution() *Resolution {
    s.access.RLock()
    defer s.access.RUnlock()
    return s.origResolution
}

// Resolutions returns all available Resolution's for this Screen
func (s *Screen) Resolutions() ([]*Resolution, error) {
    s.access.Lock()
    defer s.access.Unlock()

    // Calling C things, so get lock
    chippyAccess.Lock()
    defer chippyAccess.Unlock()

    err := getInitError()
    if err != nil {
        return nil, err
    }

    return s.resolutions()
}



// RestoreOriginalGamma restores the original gamma ramp of this screen
// just like it was when Chippy was intialized.
//
// Chippy will automatically call this for you when you call chippy.Destroy()
//
// You can turn off this default behavior with screen.SetAutoRestoreOriginalGamma(bool)
func (s *Screen) RestoreOriginalGamma() error {
    s.access.RLock()
    oldRamp := s.origGammaRamp
    s.access.RUnlock()
    return s.SetGammaRamp(oldRamp)
}

// SetAutoRestoreOriginalGamma if restore is true, chippy will automatically restore the gamma of this screen
// as it was when the application started. If restore is false, then chippy will instead leave the gamma as-is
// when the application finally calls chippy.Destroy()
func (s *Screen) SetAutoRestoreOriginalGamma(restore bool) error {
    s.access.Lock()
    defer s.access.Unlock()

    // Calling add/remove DestroyCallback
    chippyAccess.Lock()
    defer chippyAccess.Unlock()

    err := getInitError()
    if err != nil {
        return err
    }

    if restore != s.autoRestoreOriginalGamma {
        s.autoRestoreOriginalGamma = restore
        if restore == true {
            addDestroyCallback(s.autoRestoreOriginalGammaCallback)
        } else {
            removeDestroyCallback(s.autoRestoreOriginalGammaCallback)
        }
    }
    return nil
}

// OriginalGamma returns the gamma Ramp that was in use when Chippy was intialized
func (s *Screen) OriginalGamma() *Ramp {
    s.access.RLock()
    defer s.access.RUnlock()
    return s.origGammaRamp
}

// Screens returns all available screens
func Screens() ([]*Screen, error) {
    // accessing a global
    chippyAccess.Lock()
    defer chippyAccess.Unlock()

    err := getInitError()
    if err != nil {
        return nil, err
    }

    return platformScreens, nil
}

// DefaultScreen returns the default screen
func DefaultScreen() (*Screen, error) {
    // accessing a global
    chippyAccess.Lock()
    defer chippyAccess.Unlock()

    err := getInitError()
    if err != nil {
        return nil, err
    }

    return platformDefaultScreen, nil
}

