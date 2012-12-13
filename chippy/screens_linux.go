package chippy

import "sync"
import "errors"

// Resolution represents a Screen's Resolution
type Resolution struct {
    width, height uint16
    screen *Screen
    access sync.RWMutex

    // Platform
    videoMode *c_XF86VidModeModeInfo
}

// Screen represents a physical screen device
type Screen struct {
    number uint16
    resolution *Resolution
    origResolution *Resolution
    origGammaRamp *Ramp
    autoRestoreOriginalGamma bool
    autoRestoreOriginalGammaCallback *callback
    autoRestoreOriginalResolution bool
    autoRestoreOriginalResolutionCallback *callback
    access sync.RWMutex

    // Platform
    xScreen *c_Screen
    xScreenNumber int32
}

// Helper to create a new screen struct
func newScreen(xScreen *c_Screen) (*Screen, error) {
    s := Screen{}
    s.xScreen = xScreen
    s.xScreenNumber = c_XScreenNumberOfScreen(s.xScreen)
    s.number = uint16(s.xScreenNumber + 1)

    // The largest video mode available is the screen
    modes, err := s.resolutions()
    if err != nil {
        return nil, err
    }
    var working *Resolution
    for i := 0; i < len(modes); i++ {
        current := modes[i]
        if int32(current.Width()) == c_XWidthOfScreen(s.xScreen) && int32(current.Height()) == c_XHeightOfScreen(s.xScreen) {
            working = current
            break
        }
    }
    if working == nil {
        return nil, errors.New("Unable to locate current screen resolution; does glxinfo report a valid fbconfig?")
    }

    s.resolution = working
    s.origResolution = working
    s.autoRestoreOriginalResolutionCallback = &callback{func(){
        s.RestoreOriginalResolution()
    }}
    s.autoRestoreOriginalResolution = true
    addDestroyCallback(s.autoRestoreOriginalResolutionCallback)

    s.origGammaRamp, _ = getGammaRamp(&s)
    s.autoRestoreOriginalGammaCallback = &callback{func(){
        s.RestoreOriginalGamma()
    }}
    s.autoRestoreOriginalGamma = true
    addDestroyCallback(s.autoRestoreOriginalGammaCallback)
    return &s, nil
}

func (s *Screen) resolutions() ([]*Resolution, error) {
    modes := c_XF86VidModeGetAllModeLines(xDisplay, s.xScreenNumber)
    resolutions := []*Resolution{}
    for i := 0; i < len(modes); i++ {
        mode := modes[i]

        resolution := Resolution{}
        resolution.screen = s
        resolution.width = uint16(mode.hdisplay)
        resolution.height = uint16(mode.vdisplay)
        resolution.videoMode = mode
        resolutions = append(resolutions, &resolution)
    }
    return resolutions, nil
}

func (s *Screen) setResolution() error {
    // Warp the pointer to the upper left corner, this is
    // necessary as the: XF86VidModeSetViewPort() call
    // below does not seem to do anything on newer Xorg
    // servers, instead the viewport appears to be always
    // centered at the last mouse position
    c_XWarpPointer(xDisplay, c_Window(c_None), c_XDefaultRootWindow(xDisplay), 0, 0, 0, 0, 0, 0)

    err := c_XF86VidModeSwitchToMode(xDisplay, s.xScreenNumber, s.resolution.videoMode)
    if err != nil {
        return err
    }
    c_XF86VidModeSetViewPort(xDisplay, s.xScreenNumber, 0, 0)
    c_XSync(xDisplay, false)
    return nil
}

func screens() ([]*Screen, error) {
    xDisplayAccess.RLock()
    defer xDisplayAccess.RUnlock()

    screens := []*Screen{}
    screenCount := c_XScreenCount(xDisplay)
    for i := int32(0); i < screenCount; i++ {
        screen := c_XScreenOfDisplay(xDisplay, i)
        s, err := newScreen(screen)
        if err != nil {
            return nil, err
        }
        screens = append(screens, s)
    }
    return screens, nil
}

func defaultScreen() (*Screen, error) {
    xDisplayAccess.RLock()
    defer xDisplayAccess.RUnlock()

    return newScreen(c_XDefaultScreenOfDisplay(xDisplay))
}

func initScreens() error {
    var err error
    platformDefaultScreen, err = defaultScreen()
    if err != nil {
        return err
    }
    platformScreens, err = screens()
    if err != nil {
        return err
    }
    return nil
}

func destroyScreens() {
    platformDefaultScreen = nil
    platformScreens = nil
}

