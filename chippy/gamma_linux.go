package chippy

import "errors"

func backend_setGammaRamp(screen *Screen, ramp *Ramp) error {
	// Could have been non-256 before, so no setting gama
	if ramp == nil {
		return errors.New("Unable to set gamma ramp; Gamma ramp is nil")
	}

	err := c_XF86VidModeSetGammaRamp(xDisplay, screen.xScreenNumber, ramp.Red, ramp.Green, ramp.Blue)
	if err != nil {
		return errors.New("Unable to set gamma ramp; XF86VidModeSetGammaRamp() failed")
	}
	return nil
}

func backend_gammaRamp(screen *Screen) (*Ramp, error) {
	var err error
	ramp := Ramp{}
	ramp.Red, ramp.Green, ramp.Blue, err = c_XF86VidModeGetGammaRamp(xDisplay, screen.xScreenNumber)
	if err != nil {
		return nil, err
	}
	return &ramp, nil
}
