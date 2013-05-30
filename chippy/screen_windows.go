// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package chippy

import (
	"code.google.com/p/azul3d/chippy/wrappers/win32"
	"runtime"
	"fmt"
	"math"
	"sort"
	"sync"
)

type w32ScreenMode struct {
	valid         bool
	isCurrentMode bool
	width, height uint
	refreshRate   float32

	w32Bpp  win32.DWORD
	w32Mode *win32.DEVMODE
}

func newScreenMode(screen Screen) *w32ScreenMode {
	m := &w32ScreenMode{}
	//m.screen = screen
	m.valid = true
	return m
}

func (m *w32ScreenMode) panicUnlessValid() {
	if !m.valid {
		panic("ScreenMode came from an incorrect source; you cannot create it yourself!")
	}
}

func (m *w32ScreenMode) String() string {
	m.panicUnlessValid()
	w, h := m.Resolution()
	return fmt.Sprintf("ScreenMode(%d by %dpx, %.1fhz, %dbpp)", w, h, m.RefreshRate(), m.BytesPerPixel())
}

func (m *w32ScreenMode) Equals(other ScreenMode) bool {
	m.panicUnlessValid()
	width, height := m.Resolution()
	otherWidth, otherHeight := other.Resolution()

	return (width == otherWidth) && (height == otherHeight) && (m.RefreshRate() == other.RefreshRate()) && (m.BytesPerPixel() == other.BytesPerPixel())
}

func (m *w32ScreenMode) Resolution() (width, height uint) {
	m.panicUnlessValid()
	return m.width, m.height
}

func (m *w32ScreenMode) RefreshRate() float32 {
	m.panicUnlessValid()
	return m.refreshRate
}

func (m *w32ScreenMode) BytesPerPixel() uint {
	m.panicUnlessValid()
	return uint(m.w32Bpp)
}

type w32Screen struct {
	access sync.RWMutex
	valid  bool

	name                           string
	physicalWidth, physicalHeight  float32
	originalGammaRamp, gammaRamp   *GammaRamp
	originalScreenMode, screenMode ScreenMode
	screenModes                    sortedScreenModes

	gammaRampSize  uint
	gammaRampError error

	screenModeChanged, gammaRampChanged          bool

	isDefaultScreen                             bool
	w32MonitorDeviceName, w32GraphicsDeviceName string
	dc                                          win32.HDC
	w32Position                                 win32.RECT
}

func newScreen() *w32Screen {
	s := &w32Screen{}
	s.valid = true

	runtime.SetFinalizer(s, func(s *w32Screen) {
		dispatchNoWait(func() {
			// Do screen related cleanup here..
			if !win32.DeleteDC(s.dc) {
				logger.Println("Cannot delete DC; DeleteDC():", win32.GetLastErrorString())
			}
		})
	})

	return s
}

func (s *w32Screen) String() string {
	w, h := s.PhysicalSize()
	return fmt.Sprintf("Screen(\"%s\", %.0f by %.0fmm)", s.Name(), w, h)
}

func (w *w32Screen) Equals(s Screen) bool {
	if w.String() == s.String() {
		otherModes := s.ScreenModes()
		for i, mode := range w.ScreenModes() {
			if !mode.Equals(otherModes[i]) {
				return false
			}
		}
		return true
	}

	return false
}

func (s *w32Screen) Name() string {
	s.access.RLock()
	defer s.access.RUnlock()
	return s.name
}

func (s *w32Screen) PhysicalSize() (width float32, height float32) {
	s.access.RLock()
	defer s.access.RUnlock()
	return s.physicalWidth, s.physicalHeight
}

func (s *w32Screen) OriginalScreenMode() ScreenMode {
	s.access.RLock()
	defer s.access.RUnlock()
	return s.originalScreenMode
}

func (s *w32Screen) ScreenModes() []ScreenMode {
	s.access.RLock()
	defer s.access.RUnlock()
	return s.screenModes
}

func (s *w32Screen) SetScreenMode(newMode ScreenMode) (err error) {
	s.access.Lock()
	defer s.access.Unlock()

	if s.screenMode.Equals(newMode) {
		// We're already using this mode -- avoid flicker.
		return nil
	}

	s.screenMode = newMode

	dispatch(func() {
		mode := s.screenMode.(*w32ScreenMode).w32Mode

		//mode.SetDmFields(win32.DM_PELSWIDTH & win32.DM_PELSHEIGHT & win32.DM_DISPLAYFREQUENCY)

		ret := win32.ChangeDisplaySettingsEx(s.w32GraphicsDeviceName, mode, win32.CDS_TEST, nil)
		if ret != 0 {
			logger.Println("Unable to set screen mode; ChangeDisplaySettingsEx(,,CDS_TEST,) reports bad mode.")
			err = ErrBadScreenMode
			return
		}

		ret = win32.ChangeDisplaySettingsEx(s.w32GraphicsDeviceName, mode, 0, nil)

		if ret == win32.DISP_CHANGE_BADDUALVIEW {
			logger.Println("Unable to set screen mode; Because the system is DualView capable.")
			err = ErrDualViewCapable
			return
		}
		if ret == win32.DISP_CHANGE_BADMODE {
			logger.Println("Unable to set screen mode; The graphics mode is not supported.")
			err = ErrBadScreenMode
			return
		}

		// highly unlikely errors:

		if ret == win32.DISP_CHANGE_BADFLAGS {
			logger.Println("Unable to set screen mode; An invalid set of flags was passed in.")
			err = ErrBadScreenMode
			return
		}
		if ret == win32.DISP_CHANGE_BADPARAM {
			logger.Println("Unable to set screen mode; Invalid parameter or invalid flag (or combination of)")
			err = ErrBadScreenMode
			return
		}
		if ret == win32.DISP_CHANGE_NOTUPDATED {
			logger.Println("Unable to set screen mode; Unable to write settings to the registry.")
			err = ErrBadScreenMode
			return
		}
		if ret == win32.DISP_CHANGE_RESTART {
			logger.Println("Unable to set screen mode; Windows requires restart to achieve specific mode.")
			err = ErrBadScreenMode
			return
		}
	})
	return
}

func (s *w32Screen) ScreenMode() ScreenMode {
	s.access.RLock()
	defer s.access.RUnlock()
	return s.screenMode
}

func (s *w32Screen) OriginalGammaRamp() (*GammaRamp, error) {
	s.access.RLock()
	defer s.access.RUnlock()
	return s.originalGammaRamp, s.gammaRampError
}

func (s *w32Screen) SetGammaRamp(gammaRamp *GammaRamp) (err error) {
	gammaRampSize := int(s.GammaRampSize())
	if gammaRampSize == 0 {
		return ErrGammaRampsNotSupported
	}

	if len(gammaRamp.Red) != gammaRampSize || len(gammaRamp.Green) != gammaRampSize || len(gammaRamp.Blue) != gammaRampSize {
		panic("Incorrect gamma ramp size; gamma ramp size must be of the size returned by GammaRampSize()")
	}

	s.access.Lock()
	defer s.access.Unlock()

	s.gammaRamp = gammaRamp

	// Yay WinDOS! So there are huge gamma ramp restrictions, which you can only get rid of if you
	// are willing to add an registry key and run your process under an specific account, so I
	// guess the lesson to be learned here is that your gamma ramps will look much worse on windows
	// (tiny tiny range, basically only enough for controlling the screen brightness..) than they
	// will on any other operating system.. sorry folks!
	//
	// Also, to top it off, if you do the above hacky fix, windows will always report that setting
	// the gamma ramp failed (see: https://bugs.launchpad.net/redshift/+bug/850865)
	//
	// Google: "GdiIcmGammaRange" for more information
	// See this blog: http://jonls.dk/2010/09/windows-gamma-adjustments/
	//

	dispatch(func() {
		// Note: WORD == uint16
		var ramp [3][256]win32.WORD
		for i := 0; i < 256; i++ {

			// Convert from float32 to 16 bit unsigned int (WORD)
			fromFloat := func(v float32) win32.WORD {
				v = float32(math.Min(math.Max(float64(v), 0.0), 1.0))

				maxValue := float32(i+1) * 256.0
				minValue := float32(i+1) * 128.0
				rangeValue := maxValue - minValue
				v = minValue + (rangeValue * v)

				v = float32(math.Min(math.Max(float64(v), float64(minValue)), float64(maxValue-1)))
				return win32.WORD(v)
			}

			ramp[0][i] = fromFloat(gammaRamp.Red[i])
			ramp[1][i] = fromFloat(gammaRamp.Green[i])
			ramp[2][i] = fromFloat(gammaRamp.Blue[i])
		}

		win2kOrBelow := (w32VersionMajor <= 5) || (w32VersionMinor <= 0)

		worked := win32.SetDeviceGammaRamp(s.dc, ramp)

		// On windows 2000, sometimes SetDeviceGammaRamp will return false, even though it worked.
		if (win2kOrBelow && win32.GetLastError() != 0) || (!win2kOrBelow && !worked) {
			logger.Println(fmt.Sprintf("Unable to set gamma ramp on %s; SetDeviceGammaRamp(): %s", s.name, win32.GetLastErrorString()))
			err = ErrGammaRampsNotSupported
		}
		return
	})
	return

}

func (s *w32Screen) GammaRamp() (*GammaRamp, error) {
	s.access.RLock()
	defer s.access.RUnlock()
	return s.gammaRamp, s.gammaRampError
}

func (s *w32Screen) GammaRampSize() uint {
	s.access.RLock()
	defer s.access.RUnlock()
	return s.gammaRampSize
}

func (s *w32Screen) Restore() {
	s.SetScreenMode(s.OriginalScreenMode())

	ramp, err := s.OriginalGammaRamp()
	if err != nil {
		logger.Println(err.Error())
		return
	}
	err = s.SetGammaRamp(ramp)
	if err != nil {
		logger.Println(err.Error())
		return
	}
}

func (s *w32Screen) setGammaRampSize() {
	if win32.GetDeviceCaps(s.dc, win32.CM_GAMMA_RAMP) != 0 {
		s.gammaRampSize = 256
		return
	}
	logger.Println("GetDeviceCaps(CM_GAMMA_RAMP) reports no support for gamma ramps on device:", s.name)
}

func (s *w32Screen) setCurrentGammaRamp() {
	if s.GammaRampSize() == 0 {
		logger.Println("GetDeviceCaps(CM_GAMMA_RAMP) reports no support for gamma ramps on device:", s.name)
		s.gammaRampError = ErrGammaRampsNotSupported
		return
	}

	ret, deviceRamp := win32.GetDeviceGammaRamp(s.dc)
	if ret == false {
		s.gammaRampError = ErrGammaRampsNotSupported
		logger.Println(fmt.Sprintf("Unable to get current gamma ramp on %s; GetDeviceGammaRamp(): %s", s.name, win32.GetLastErrorString()))
		return
	}

	s.gammaRamp = new(GammaRamp)

	ramp := s.gammaRamp
	ramp.Red = make([]float32, 256)
	ramp.Green = make([]float32, 256)
	ramp.Blue = make([]float32, 256)
	for i := 0; i < 256; i++ {

		// Convert from 16 bit unsigned int (WORD) to float32
		fromWORD := func(v win32.WORD) float32 {
			maxValue := float32(i+1) * 256.0
			minValue := float32(i+1) * 128.0
			rangeValue := maxValue - minValue

			// Get our float value back
			x := (float32(v) - minValue) / rangeValue
			//logger.Printf("%.3f, ", x)
			return x
		}

		ramp.Red[i] = fromWORD(deviceRamp[0][i])
		ramp.Green[i] = fromWORD(deviceRamp[1][i])
		ramp.Blue[i] = fromWORD(deviceRamp[2][i])
	}

	// Make sure this is an copy otherwise they might change it by accident
	s.originalGammaRamp = s.gammaRamp.Copy()
}

func (s *w32Screen) setScreenModes() {
	hasCurrentMode, mode := win32.EnumDisplaySettings(s.w32GraphicsDeviceName, win32.ENUM_CURRENT_SETTINGS)

	var currentScreenMode *w32ScreenMode
	if hasCurrentMode {
		currentScreenMode = newScreenMode(s)
		currentScreenMode.width = uint(mode.DmPelsWidth())
		currentScreenMode.height = uint(mode.DmPelsHeight())
		currentScreenMode.refreshRate = float32(mode.DmDisplayFrequency())
		currentScreenMode.w32Bpp = mode.DmBitsPerPel()
		//currentScreenMode.w32Mode = mode // See: 'Assign it here' below
		currentScreenMode.isCurrentMode = true
		s.screenModes = append(s.screenModes, currentScreenMode)
	}

	hasNext := true
	i := 0
	for hasNext {
		var mode *win32.DEVMODE
		hasNext, mode = win32.EnumDisplaySettings(s.w32GraphicsDeviceName, win32.DWORD(i))
		i++
		if hasNext {
			// This one is an good one

			if mode.DmDisplayFixedOutput() != win32.DMDFO_STRETCH {
				// Skip all modes that are not specified to stretch across the screen (there is
				// always at least one for each resolution specified as stretching, so no worries)
				continue
			}

			screenMode := newScreenMode(s)
			screenMode.width = uint(mode.DmPelsWidth())
			screenMode.height = uint(mode.DmPelsHeight())
			screenMode.refreshRate = float32(mode.DmDisplayFrequency())
			screenMode.w32Bpp = mode.DmBitsPerPel()
			screenMode.w32Mode = mode

			if hasCurrentMode {
				if screenMode.width == currentScreenMode.width && screenMode.height == currentScreenMode.height && screenMode.refreshRate == currentScreenMode.refreshRate && screenMode.w32Bpp == currentScreenMode.w32Bpp {

					// Assign it here, avoid issues later on with comparison just in case we ever use it.
					currentScreenMode.w32Mode = mode
					continue // We already appended this before
				}
			}

			s.screenModes = append(s.screenModes, screenMode)
		}
	}

	sort.Sort(s.screenModes)

	if !hasCurrentMode {
		currentScreenMode = s.screenModes[0].(*w32ScreenMode)
		currentScreenMode.isCurrentMode = true
	}

	s.screenMode = currentScreenMode
	s.originalScreenMode = currentScreenMode
}

func backend_doScreens() (screens []Screen) {
	win2kOrBelow := (w32VersionMajor <= 5) || (w32VersionMinor <= 0)

	monitorNum := 0
	hasNext := true
	i := 0
	for hasNext {
		var dd *win32.DISPLAY_DEVICE
		hasNext, dd = win32.EnumDisplayDevices("", win32.DWORD(i), 0)
		i++
		if hasNext {
			// We're only interested in active devices (graphics cards)
			graphicsCardName := dd.GetDeviceName()
			graphicsCardString := dd.GetDeviceString()

			gflags := dd.GetStateFlags()
			if (gflags & win32.DISPLAY_DEVICE_ACTIVE) > 0 {
				hasMoreMonitors := true
				j := 0
				for hasMoreMonitors {
					hasMoreMonitors, dd = win32.EnumDisplayDevices(dd.GetDeviceName(), 0, 0)
					j++

					// We're only interested in active monitors, but windows 2000 and below
					// never sets the DISPLAY_DEVICE_ACTIVE flag.
					//
					flags := dd.GetStateFlags()

					active := (flags & win32.DISPLAY_DEVICE_ACTIVE) > 0
					attached := (flags & win32.DISPLAY_DEVICE_ATTACHED) > 0
					if active || attached || win2kOrBelow {
						screen := newScreen()

						if (gflags & win32.DISPLAY_DEVICE_PRIMARY_DEVICE) > 0 && j == 1 {
							screen.isDefaultScreen = true
						}

						screen.w32MonitorDeviceName = dd.GetDeviceName()
						screen.w32GraphicsDeviceName = graphicsCardName

						// It's difficult to get monitor name or model, and it's only available on Windows 7+
						// eventually, we should try to fix this, but it seems mingw is missing the proper headers
						// with these definitions.
						//
						// See: http://msdn.microsoft.com/en-us/library/windows/hardware/ff553903(v=vs.85).aspx
						//
						monitorNum++
						screen.name = fmt.Sprintf("Monitor %d - %s", monitorNum, graphicsCardString)

						screen.dc = win32.CreateDC(screen.w32GraphicsDeviceName, "", nil)
						if screen.dc != nil {
							screen.physicalWidth = float32(win32.GetDeviceCaps(screen.dc, win32.HORZSIZE))
							screen.physicalHeight = float32(win32.GetDeviceCaps(screen.dc, win32.VERTSIZE))

							screen.setGammaRampSize()
							screen.setCurrentGammaRamp()
							screen.setScreenModes()

							screens = append(screens, screen)
						} else {
							// This hopefully never happens, but if it does that means
							// there is something wrong with this screen most likely or
							// an graphics driver bug or something else, who knows?
							logger.Println("CreateDC() on screen failed! Unable to create device context!")
							logger.Println("^ Screen will be ignored!")
						}
					}
				}
			}
		}
	}

	// Find the correct MONITORINFO struct for each screen and assign their w32Position properties
	proc := func(hMonitor win32.HMONITOR, hdcMonitor win32.HDC, lprcMonitor *win32.RECT, dwData win32.LPARAM) bool {
		mi := new(win32.MONITORINFOEX)
		mi.SetSize()
		if !win32.GetMonitorInfo(hMonitor, mi) {
			logger.Println("Unable to detect monitor position; GetMonitorInfo():", win32.GetLastErrorString())
		} else {
			for _, screen := range screens {
				if screen.(*w32Screen).w32GraphicsDeviceName == mi.Device() {
					screen.(*w32Screen).w32Position = mi.RcMonitor
				}
			}
		}

		return true
	}

	if !win32.EnumDisplayMonitors(nil, nil, proc, 0) {
		logger.Println("Unable to detect monitor positions; EnumDisplayMonitors():", win32.GetLastErrorString())
	}
	return
}

func backend_Screens() (screens []Screen) {
	dispatch(func() {
		screens = backend_doScreens()
	})
	return
}

func backend_DefaultScreen() Screen {
	screens := backend_Screens()
	for _, iScreen := range screens {
		screen, ok := iScreen.(*w32Screen)
		if ok {
			if screen.isDefaultScreen {
				return screen
			}
		}
	}

	// Should never happen
	if len(screens) > 0 {
		logger.Println("Unable to find default screen; falling back to first screen as default.")
		return screens[0]
	}
	logger.Println("No screens available!")
	return nil


}
