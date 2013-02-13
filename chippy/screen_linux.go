// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package chippy

import(
    "code.google.com/p/azul3d/chippy/wrappers/x11"
    "errors"
    "math"
)

type backend_ScreenMode struct {
    isCurrentMode bool
    xf86ModeLine int
    xrandrModeInfoId x11.RRMode
}

type backend_Screen struct {
    xScreenNumber int
    xrandrCrtc x11.RRCrtc
}


func xrandrRefreshRateFromModeInfo(modeInfo *x11.XRRModeInfo) float32 {
    // Calculate refresh rate (confusing stuff incoming!)
    dotclock := float32(modeInfo.DotClock())
    vtotal := float32(modeInfo.VTotal())
    flags := modeInfo.ModeFlags()
    if (flags & x11.RR_Interlace) > 0 {
        dotclock *= 2
    }
    if (flags & x11.RR_DoubleScan) > 0 {
        vtotal *= 2
    }
    if (flags & x11.RR_ClockDivideBy2) > 0 {
        dotclock /= 2
    }

    return dotclock / (float32(modeInfo.HTotal()) * vtotal)
}

func float32Fromuint16Slice(slice []uint16) []float32 {
    new := make([]float32, len(slice))
    for i, v := range slice {
        new[i] = float32(v) / float32(math.MaxUint16)
    }
    return new
}

func uint16Fromfloat32Slice(slice []float32) []uint16 {
    new := make([]uint16, len(slice))
    for i, v := range slice {
        new[i] = uint16(v * math.MaxUint16)
    }
    return new
}

// Since so much xf86vm and xrandr code is in here, it makes since for this to be as well
func (s *Screen) setGammaRamp(gammaRamp *GammaRamp) error {
    if xrandr >= xrandrMinimum {
        gamma := &x11.XRRCrtcGamma{}
        gamma.Size = len(gammaRamp.Red)
        gamma.Red = uint16Fromfloat32Slice(gammaRamp.Red)
        gamma.Green = uint16Fromfloat32Slice(gammaRamp.Green)
        gamma.Blue = uint16Fromfloat32Slice(gammaRamp.Blue)

        x11.XRRSetCrtcGamma(xDisplay, s.xrandrCrtc, gamma)
        x11.XFlush(xDisplay)

    } else if xf86vm >= xf86vmMinimum {
        red := uint16Fromfloat32Slice(gammaRamp.Red)
        green := uint16Fromfloat32Slice(gammaRamp.Green)
        blue := uint16Fromfloat32Slice(gammaRamp.Blue)
        ret := x11.XF86VidModeSetGammaRamp(xDisplay, s.xScreenNumber, red, green, blue)
        if ret == 0 {
            return errors.New("XF86VidModeSetGammaRamp() call failed! Unable to set gamma ramp!")
        }
        x11.XFlush(xDisplay)
    }
    return nil
}

func (s *Screen) gammaRamp() (*GammaRamp, error) {
    if xrandr >= xrandrMinimum {
        gamma := x11.XRRGetCrtcGamma(xDisplay, s.xrandrCrtc)
        r := &GammaRamp{}
        r.Red = float32Fromuint16Slice(gamma.Red)
        r.Green = float32Fromuint16Slice(gamma.Green)
        r.Blue = float32Fromuint16Slice(gamma.Blue)
        return r, nil

    } else if xf86vm >= xf86vmMinimum {
        ret, red, green, blue := x11.XF86VidModeGetGammaRamp(xDisplay, s.xScreenNumber)
        if ret == 0 {
            return nil, errors.New("XF86VidModeGetGammaRamp() call failed! Unable to get gamma ramp!")
        }
        r := &GammaRamp{}
        r.Red = float32Fromuint16Slice(red)
        r.Green = float32Fromuint16Slice(green)
        r.Blue = float32Fromuint16Slice(blue)
        return r, nil
    }

    return nil, errors.New("Missing both xrandr and xf86vm extensions! Unable to get gamma ramp!")
}

func (s *Screen) gammaRampSize() (int, error) {
    if xrandr >= xrandrMinimum {
        return x11.XRRGetCrtcGammaSize(xDisplay, s.xrandrCrtc), nil

    } else if xf86vm >= xf86vmMinimum {
        ret, size := x11.XF86VidModeGetGammaRampSize(xDisplay, s.xScreenNumber)
        if ret == 0 {
            return 0, errors.New("XF86VidModeGetGammaRampSize() call failed! Unable to get gamma ramp size!")
        }
        return size, nil
    }

    return 0, errors.New("Missing both xrandr and xf86vm extensions! Unable to get gamma ramp size!")
}

func (s *Screen) setScreenMode() {
    if xrandr >= xrandrMinimum {
        root := x11.XRootWindow(xDisplay, s.xScreenNumber)

        var resources *x11.XRRScreenResources
        var free func()
        if xrandr >= 1.3 {
            resources, free = x11.XRRGetScreenResourcesCurrent(xDisplay, root) // 1.3
        } else {
            resources, free = x11.XRRGetScreenResources(xDisplay, root) // 1.2
        }
        defer free()

        resourceModes := resources.Modes()

        info := x11.XRRGetCrtcInfo(xDisplay, resources, s.xrandrCrtc) // 1.2
        //crtcModeId := info.Mode()

        for _, mode := range resourceModes {
            if mode.Id() == s.screenMode.xrandrModeInfoId {
                // It's this one!
                x11.XRRSetCrtcConfig(xDisplay, resources, s.xrandrCrtc, x11.CurrentTime, info.X(), info.Y(), mode.Id(), info.Rotation(), info.Outputs())
            }
        }

    } else if xf86vm >= xf86vmMinimum {
        // Locate the proper modeLine in array
        var modeLine *x11.XF86VidModeModeInfo

        ret, modeLines, free := x11.XF86VidModeGetAllModeLines(xDisplay, s.xScreenNumber)
        if ret == 0 {
            logger.Println("Warning: XF86VidModeGetAllModeLines() failed! Unable to set screen mode!")
            return
        }
        defer free()
        modeLine = modeLines[s.screenMode.xf86ModeLine]


        // Now switch to the mode, and use XWarpPointer below to fix an bug with some Xorg servers

        // Warp the pointer to the upper left corner, this is
        // necessary as the: XF86VidModeSetViewPort() call
        // below does not seem to do anything on newer Xorg
        // servers, instead the viewport appears to be always
        // centered at the last mouse position
        root := x11.XRootWindow(xDisplay, s.xScreenNumber)
        x11.XWarpPointer(xDisplay, x11.Window(x11.None), root, 0, 0, 0, 0, 0, 0)

        ret = x11.XF86VidModeSwitchToMode(xDisplay, s.xScreenNumber, modeLine)
        if ret == 0 {
            logger.Println("Warning: XF86VidModeSwitchToMode() failed! Unable to set screen mode!")
        }

        x11.XF86VidModeSetViewPort(xDisplay, s.xScreenNumber, 0, 0)
        x11.XFlush(xDisplay)
    } else {
        // No Xrandr or xf86vm extensions means no way to change screen mode.
        logger.Println("Missing both xrandr and xf86vm extensions! Unable to set screen mode!")
    }
}

func (s *Screen) screenModes() []*ScreenMode {
    modes := []*ScreenMode{}

    if xrandr >= xrandrMinimum {
        root := x11.XRootWindow(xDisplay, s.xScreenNumber)

        var resources *x11.XRRScreenResources
        var free func()
        if xrandr >= 1.3 {
            resources, free = x11.XRRGetScreenResourcesCurrent(xDisplay, root) // 1.3
        } else {
            resources, free = x11.XRRGetScreenResources(xDisplay, root) // 1.2
        }
        defer free()

        resourceModes := resources.Modes()

        info := x11.XRRGetCrtcInfo(xDisplay, resources, s.xrandrCrtc) // 1.2
        crtcModeId := info.Mode()

        // The output's are the actual physical monitor devices
        for _/*outputIndex*/, output := range info.Outputs() {
            outputInfo, free := x11.XRRGetOutputInfo(xDisplay, resources, output) // 1.2
            defer free()

            for _, actualMode := range outputInfo.Modes() {
                for _, modeInfo := range resourceModes {
                    if actualMode == modeInfo.Id() {
                        // This mode is related to output

                        screenMode := newScreenMode(s)
                        screenMode.width = modeInfo.Width()
                        screenMode.height = modeInfo.Height()
                        screenMode.refreshRate = xrandrRefreshRateFromModeInfo(modeInfo)
                        screenMode.xrandrModeInfoId = modeInfo.Id()

                        if actualMode == crtcModeId {
                            screenMode.isCurrentMode = true
                        }
                        modes = append(modes, screenMode)
                    }
                }
            }
        }


    } else if xf86vm >= xf86vmMinimum {
        // We are missing Xrandr support, so use xf86vm if we have it
        ret, modeLines, free := x11.XF86VidModeGetAllModeLines(xDisplay, s.xScreenNumber)
        defer free()

        if ret != 0 { // No failure accepted!
            for i, modeLine := range modeLines {
                mode := newScreenMode(s)
                mode.xf86ModeLine = i
                mode.width = uint(modeLine.Hdisplay())
                mode.height = uint(modeLine.Vdisplay())

                // Calculate refresh rate (confusing stuff incoming!)
                dotclock := float32(modeLine.Dotclock())
                vtotal := float32(modeLine.Vtotal())
                if (modeLine.Flags() & 0x0010) > 0 { // interlaced
                    dotclock *= 2
                }
                if (modeLine.Flags() & 0x0020) > 0 { // Double Scan
                    vtotal *= 2
                }
                mode.refreshRate = 1000.0 * dotclock / (float32(modeLine.Htotal()) * vtotal)

                if i == 0 {
                    // First element is always current one
                    mode.isCurrentMode = true
                }

                modes = append(modes, mode)
            }
        }
    }

    if len(modes) == 0 {
        // We have no way to determine the screen mode through just Xlib alone.. we need
        // the extensions, the only thing we can do is say there is only an single screen mode
        // available.
        mode := newScreenMode(s)
        xScreen := x11.XScreenOfDisplay(xDisplay, s.xScreenNumber)
        mode.width = uint(x11.XWidthOfScreen(xScreen))
        mode.height = uint(x11.XHeightOfScreen(xScreen))
        mode.isCurrentMode = true
        modes = append(modes, mode)
    }

    return modes
}




func backend_Screens() []*Screen {
    if xrandr >= xrandrMinimum {
        // Xrandr is an big virtual screen made up of physical monitors arranged on the virtual one
        screenCount := x11.XScreenCount(xDisplay)
        screens := []*Screen{}

        for screenNumber := 0; screenNumber < screenCount; screenNumber++ {
            root := x11.XRootWindow(xDisplay, screenNumber)

            var resources *x11.XRRScreenResources
            var free func()
            if xrandr >= 1.3 {
                resources, free = x11.XRRGetScreenResourcesCurrent(xDisplay, root) // 1.3
            } else {
                resources, free = x11.XRRGetScreenResources(xDisplay, root) // 1.2
            }
            defer free()

            resourceModes := resources.Modes()

            for _/*crtcIndex*/, crtc := range resources.Crtcs() {
                info := x11.XRRGetCrtcInfo(xDisplay, resources, crtc) // 1.2
                infoModeId := info.Mode()

                // Check for Disabled/Inactive
                if infoModeId != 0 {

                    // The output's are the actual physical monitor devices
                    for _/*outputIndex*/, output := range info.Outputs() {
                        outputInfo, free := x11.XRRGetOutputInfo(xDisplay, resources, output) // 1.2
                        defer free()

                        screen := newScreen()
                        screen.xScreenNumber = screenNumber
                        screen.physicalWidth = float32(outputInfo.Mm_width())
                        screen.physicalHeight = float32(outputInfo.Mm_height())
                        screen.xrandrCrtc = crtc

                        // Get screen mode
                        screenMode := newScreenMode(screen)
                        for _, modeInfo := range resourceModes {
                            if modeInfo.Id() == infoModeId {
                                screenMode.width = modeInfo.Width()
                                screenMode.height = modeInfo.Height()
                                screenMode.refreshRate = xrandrRefreshRateFromModeInfo(modeInfo)
                                screenMode.xrandrModeInfoId = modeInfo.Id()
                                screen.screenMode = screenMode

                                break
                            }
                        }

                        var err error
                        screen.originalGammaRamp, err = screen.gammaRamp()
                        if err != nil {
                            logger.Println(err)
                        }

                        screens = append(screens, screen)
                    }
                }
            }
        }

        return screens
    }



    // Otherwise we have no Xrandr, so we can only use Xlib
    screenCount := x11.XScreenCount(xDisplay)
    screens := make([]*Screen, screenCount)

    for i := 0; i < screenCount; i++ {
        screen := newScreen()
        screen.xScreenNumber = i

        modes := screen.screenModes()
        for _, mode := range modes {
            if mode.isCurrentMode {
                screen.screenMode = mode
                break
            }
        }

        xScreen := x11.XScreenOfDisplay(xDisplay, screen.xScreenNumber)
        screen.physicalWidth = float32(x11.XWidthMMOfScreen(xScreen))
        screen.physicalHeight = float32(x11.XHeightMMOfScreen(xScreen))

        var err error
        screen.originalGammaRamp, err = screen.gammaRamp()
        if err != nil {
            logger.Println(err)
        }

        screens[i] = screen
    }
    return screens
}

func backend_DefaultScreen() *Screen {
    screens := backend_Screens()
    defaultScreen := x11.XDefaultScreenOfDisplay(xDisplay)
    defaultScreenNumber := x11.XScreenNumberOfScreen(defaultScreen)

    for _, screen := range screens {
        if screen.xScreenNumber == defaultScreenNumber {
            return screen
        }
    }

    // For some reason we couldn't find the default one, this shouldn't happen
    // but in this case we should return something at least.
    logger.Println("Unable to detect default screen, this should never have happened!")
    if len(screens) >= 1 {
        return screens[0]
    }
    return nil
}

