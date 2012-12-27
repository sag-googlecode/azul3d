package chippy

import(
    "fmt"
    "sync"
)

// We store all the global screen objects here
var globalDefaultScreen *Screen
var globalScreens []*Screen

// Screen represents a physical screen device, such as an physical monitor attatched to the computer
type Screen struct {
    backend_screen
	number                                uint16
	resolution                            *Resolution
	origResolution                        *Resolution
	origGammaRamp                         *Ramp
	autoRestoreOriginalGamma              bool
	autoRestoreOriginalGammaCallback      *callback
	autoRestoreOriginalResolution         bool
	autoRestoreOriginalResolutionCallback *callback
	access                                sync.RWMutex
}

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
// Chippy will automatically call this for you when you call Destroy()
//
// You can turn off this default behavior with screen.SetAutoRestoreOriginalResolution(bool)
func (s *Screen) RestoreOriginalResolution() error {
	s.access.RLock()
	oldResolution := s.origResolution
	s.access.RUnlock() // Important to release before call to SetResolution

	return s.SetResolution(oldResolution)
}

// SetAutoRestoreOriginalResolution if restore is true, chippy will automatically restore the resolution of this screen
// as it was when the application started. If restore is false, then chippy will instead leave the resolution as-is
// when the application finally calls Destroy()
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
// Chippy will automatically call this for you when you call Destroy()
//
// You can turn off this default behavior with screen.SetAutoRestoreOriginalGamma(bool)
func (s *Screen) RestoreOriginalGamma() error {
	s.access.RLock()
	oldRamp := s.origGammaRamp
	s.access.RUnlock()
	return s.SetGammaRamp(oldRamp)
}

// SetAutoRestoreOriginalGamma if restore is true, chippy will automatically restore the gamma of this screen
// as it was when Chippy was initialized. If restore is false, then chippy will instead leave the gamma as-is
// when the application calls Destroy().
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

// Screens returns all available Screens that where plugged into the computer and regonized by the
// window manager at the time of Init(), at the current point in time there is no way to re-scan
// for available screens. (For instance, ones that where plugged in after Init() was called)
func Screens() ([]*Screen, error) {
	// accessing a global
	chippyAccess.Lock()
	defer chippyAccess.Unlock()

	err := getInitError()
	if err != nil {
		return nil, err
	}

    screens, err := backend_screens()
    if err != nil {
        return nil, err
    }

    return screens, nil
}

// DefaultScreen returns the default Screen that was determined to be the appropriate screen to open a new
// Window by the window manager. Unless you explicitly wish to open a Window on a specific Screen, then the
// Screen returned by this is appropriate to use for all new Window's you wish to open.
func DefaultScreen() (*Screen, error) {
	// accessing a global
	chippyAccess.Lock()
	defer chippyAccess.Unlock()

	err := getInitError()
	if err != nil {
		return nil, err
	}

    screen, err := backend_defaultScreen()
    if err != nil {
        return nil, err
    }
    return screen, nil
}

