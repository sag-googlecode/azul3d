package chippy

import (
	"code.google.com/p/azul3d/chippy/wrappers/win32"
	"errors"
	"fmt"
	"math"
)

type backend_ScreenMode struct {
	isCurrentMode bool
	w32Bpp        uint32
	w32Mode       *win32.DEVMODE
}

type backend_Screen struct {
	isDefaultScreen                             bool
	w32MonitorDeviceName, w32GraphicsDeviceName string
}

func (s *Screen) setGammaRamp(gammaRamp *GammaRamp) error {
	_, err := s.gammaRampSize()
	if err != nil {
		return err
	}

	// Yay WinDOS! So there are huge gamma ramp restrictions, which you can only get rid of if you
	// are willing to add an registry key and run your process under an specific account, so I
	// guess the lesson to be learned here is that your gamma ramps will look much worse on windows
	// (tiny tiny range, basically only enough for controlling the screen brightness) than they
	// will on any other modern operating system.. sorry folks!
	//
	// Also, to top it off, if you do the above hacky fix, windows will always report that setting
	// the gamma ramp failed (see: https://bugs.launchpad.net/redshift/+bug/850865)
	//
	// Google: "GdiIcmGammaRange" for more information
	// See this blog: http://jonls.dk/2010/09/windows-gamma-adjustments/
	//

	dc := win32.CreateDC(s.w32GraphicsDeviceName, "", nil)
	if dc != nil {
		defer win32.DeleteDC(dc)

		var ramp [3][256]uint16
		for i := 0; i < 256; i++ {
			ramp[0][i] = uint16(float32(i) * (128.0 + (gammaRamp.Red[i] * 128.0)))
			ramp[1][i] = uint16(float32(i) * (128.0 + (gammaRamp.Green[i] * 128.0)))
			ramp[2][i] = uint16(float32(i) * (128.0 + (gammaRamp.Blue[i] * 128.0)))
		}

		worked := win32.SetDeviceGammaRamp(dc, ramp)
		if !worked {
			// Despite what it looks like, this error isin't helpful at all
			// It's always just incorrect argument blah blah.
			//logger.Println("error:", win32.GetLastErrorString())

			return errors.New("SetDeviceGammaRamp() call failed! Unable to set gamma ramp!")
		}
		return nil
	}
	return errors.New("CreateDC() failed! Unable to set gamma ramp!")
}

func (s *Screen) gammaRamp() (*GammaRamp, error) {
	_, err := s.gammaRampSize()
	if err != nil {
		return nil, err
	}

	dc := win32.CreateDC(s.w32GraphicsDeviceName, "", nil)
	if dc != nil {
		defer win32.DeleteDC(dc)

		ret, deviceRamp := win32.GetDeviceGammaRamp(dc)
		if ret == false {
			return nil, errors.New("GetDeviceGammaRamp() call failed! Unable to get gamma ramp!")
		}

		ramp := &GammaRamp{}
		ramp.Red = make([]float32, 256)
		ramp.Green = make([]float32, 256)
		ramp.Blue = make([]float32, 256)
		for i := 0; i < 256; i++ {
			ramp.Red[i] = float32(deviceRamp[0][i]) / float32(math.MaxUint16)
			ramp.Green[i] = float32(deviceRamp[1][i]) / float32(math.MaxUint16)
			ramp.Blue[i] = float32(deviceRamp[1][i]) / float32(math.MaxUint16)
		}

		return ramp, nil

	}
	return nil, errors.New("CreateDC() failed! Unable to get gamma ramp!")
}

func (s *Screen) gammaRampSize() (int, error) {
	dc := win32.CreateDC(s.w32GraphicsDeviceName, "", nil)
	if dc != nil {
		defer win32.DeleteDC(dc)

		if win32.GetDeviceCaps(dc, win32.CM_GAMMA_RAMP) != 0 {
			return 256, nil
		}
		return 0, errors.New("GetDeviceCaps(CM_GAMMA_RAMP) reports no support for gamma ramps on this device.")
	}
	return 0, errors.New("CreateDC() failed! Unable to get gamma ramp size!")
}

func (s *Screen) setScreenMode() {
	mode := s.screenMode.w32Mode

	//mode.SetDmFields(win32.DM_PELSWIDTH & win32.DM_PELSHEIGHT & win32.DM_DISPLAYFREQUENCY)

	ret := win32.ChangeDisplaySettingsEx(s.w32GraphicsDeviceName, mode, win32.CDS_TEST, nil)
	if ret != 0 {
		logger.Println("Unable to set screen mode; CDS_TEST reports bad mode.")
		return
	}

	ret = win32.ChangeDisplaySettingsEx(s.w32GraphicsDeviceName, mode, 0, nil)

	if ret == win32.DISP_CHANGE_BADDUALVIEW {
		logger.Println("Unable to set screen mode; because the system is DualView capable.")
	}
	if ret == win32.DISP_CHANGE_BADFLAGS {
		logger.Println("Unable to set screen mode; An invalid set of flags was passed in.")
	}
	if ret == win32.DISP_CHANGE_BADMODE {
		logger.Println("Unable to set screen mode; The graphics mode is not supported.")
	}
	if ret == win32.DISP_CHANGE_BADPARAM {
		logger.Println("Unable to set screen mode; Invalid parameter or invalid flag (or combination of)")
	}
	if ret == win32.DISP_CHANGE_FAILED {
		logger.Println("Unable to set screen mode; Display driver failed the specified graphics mode.")
	}
	if ret == win32.DISP_CHANGE_NOTUPDATED {
		logger.Println("Unable to set screen mode; Unable to write settings to the registry.")
	}
	if ret == win32.DISP_CHANGE_RESTART {
		logger.Println("Unable to set screen mode; Windows requires restart to achieve specific mode.")
	}
}

func (s *Screen) screenModes() []*ScreenMode {
	screenModes := []*ScreenMode{}

	hasCurrentMode, mode := win32.EnumDisplaySettings(s.w32GraphicsDeviceName, win32.ENUM_CURRENT_SETTINGS)

	screenMode := newScreenMode(s)
	screenMode.width = uint(mode.DmPelsWidth())
	screenMode.height = uint(mode.DmPelsHeight())
	screenMode.refreshRate = float32(mode.DmDisplayFrequency())
	screenMode.w32Bpp = mode.DmBitsPerPel()
	screenMode.w32Mode = mode
	screenMode.isCurrentMode = true
	screenModes = append(screenModes, screenMode)

	hasNext := true
	i := 0
	for hasNext {
		var mode *win32.DEVMODE
		hasNext, mode = win32.EnumDisplaySettings(s.w32GraphicsDeviceName, uint32(i))
		i++
		if hasNext {
			// This one is an good one

			// So, windows gives us multiple choices for monitor bpp, we do some magic to choose
			// only the highest bpp resolution, in the case of multiple duplicate resolutions.
			screenMode := newScreenMode(s)
			screenMode.width = uint(mode.DmPelsWidth())
			screenMode.height = uint(mode.DmPelsHeight())
			screenMode.refreshRate = float32(mode.DmDisplayFrequency())
			screenMode.w32Bpp = mode.DmBitsPerPel()
			screenMode.w32Mode = mode

			doAppend := true

			for i, other := range screenModes {
				if other.width == screenMode.width && other.height == screenMode.height && other.refreshRate == screenMode.refreshRate {

					// if other is current, we will have two as long as bpp is higher than current
					// If we're running 1680x1050/16bpp then we should give both 16bpp and 32bpp
					// options, because we strictly want to avoid forcing them to choose higher bpp
					if other.isCurrentMode && screenMode.w32Bpp > other.w32Bpp {
						break
					}

					doAppend = false

					// Determine which one has an higher bpp and decide to use that one only
					if screenMode.w32Bpp > other.w32Bpp {
						screenModes[i] = screenMode
						break
					}
				}
			}

			if doAppend {
				// It's an new one!
				screenModes = append(screenModes, screenMode)
			}

			/*
			   logger.Println(mode.DmDeviceName())
			   logger.Println(mode.DmBitsPerPel())
			   logger.Println(mode.DmPelsWidth())
			   logger.Println(mode.DmPelsHeight())
			   logger.Println(mode.DmDisplayFrequency())
			*/
		}
	}

	if !hasCurrentMode {
		// Hopefully this never happens, it really should never though, seriously
		//
		// Go out with an bang !
		screenModes[0].isCurrentMode = true
	}

	return screenModes
}

func backend_Screens() []*Screen {
	screens := []*Screen{}

	monitorNum := 0

	hasNext := true
	i := 0
	for hasNext {
		var dd *win32.DISPLAY_DEVICE
		hasNext, dd = win32.EnumDisplayDevices("", uint32(i), 0)
		i++
		if hasNext {
			// We're only interested in active devices (graphics cards)
			graphicsCardName := dd.GetDeviceName()
			graphicsCardString := dd.GetDeviceString()

			flags := dd.GetStateFlags()
			if (flags & win32.DISPLAY_DEVICE_ACTIVE) > 0 {
				hasMoreMonitors := true
				j := 0
				for hasMoreMonitors {
					hasMoreMonitors, dd = win32.EnumDisplayDevices(dd.GetDeviceName(), 0, 0)
					j++

					// We're only interested in active monitors
					flags := dd.GetStateFlags()
					if (flags & win32.DISPLAY_DEVICE_ACTIVE) > 0 {
						screen := newScreen()

						if (flags & win32.DISPLAY_DEVICE_PRIMARY_DEVICE) > 0 {
							screen.isDefaultScreen = true
						}

						screen.w32MonitorDeviceName = dd.GetDeviceName()
						screen.w32GraphicsDeviceName = graphicsCardName

						// Merge the two together, so we get something like the following:
						// "1. Generic PnP Monitor - Intel(R) HD Graphics Family"
						// "2. Generic PnP Monitor - Intel(R) HD Graphics Family"
						monitorNum++
						screen.name = fmt.Sprintf("%d. %s - %s", monitorNum, dd.GetDeviceString(), graphicsCardString)

						dc := win32.CreateDC(screen.w32GraphicsDeviceName, "", nil)
						if dc != nil {
							defer win32.DeleteDC(dc)

							screen.physicalWidth = float32(win32.GetDeviceCaps(dc, win32.HORZSIZE))
							screen.physicalHeight = float32(win32.GetDeviceCaps(dc, win32.VERTSIZE))

							var err error
							screen.originalGammaRamp, err = screen.gammaRamp()
							if err != nil {
								logger.Println(err)
							}

							for _, mode := range screen.screenModes() {
								if mode.isCurrentMode {
									screen.screenMode = mode
									screen.originalScreenMode = mode
									break
								}
							}

							screens = append(screens, screen)
						} else {
							// This hopefully never happens, but if it does that means
							// there is something wrong with this screen most likely or
							// an graphics driver bug or something else, who knows?
							logger.Println("CreateDC() failed! Unable to create device context!")
							logger.Println("Unable to determine screen physical size, gamma ramp!")
						}
					}
				}
			}
		}
	}

	return screens
}

func backend_DefaultScreen() *Screen {
	screens := backend_Screens()
	for _, screen := range screens {
		if screen.isDefaultScreen {
			return screen
		}
	}

	// Should never happen
	if len(screens) > 0 {
		return screens[0]
	}
	return nil
}
