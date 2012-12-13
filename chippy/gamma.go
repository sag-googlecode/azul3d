package chippy

import "math"

// These two functions are Go re-writes of the SDL functions
// that apply a gamma correction curve, you can find lots of
// information regarding what gamma correction is online
func calculateGammaRamp(gamma float32) [256]uint16 {
    ramp := [256]uint16{}
    if gamma <= 0.0 {
        // 0.0 gamma is all black
        for i := 0; i < len(ramp); i++ {
            ramp[i] = 0
        }
        return ramp
    } else if gamma == 1.0 {
        // 1.0 is identity
        for i := 0; i < len(ramp); i++ {
            ramp[i] = uint16((i << 8) | i)
        }
        return ramp
    }

    // Calculate a real gamma ramp
    gamma = 1.0 / gamma
    for i := 0; i < len(ramp); i++ {
        value := int32(math.Pow(float64(i) / 256.0, float64(gamma)) * 65535.0 + 0.5)
        if value > 65535 {
            value = 65535
        }
        ramp[i] = uint16(value)
    }
    return ramp
}

// Read comments above
func calculateGammaFromRamp(ramp [256]uint16) float32 {
    gamma := 1.0
    sum := 0.0
    count := 0

    for i := 1; i < len(ramp); i++ {
        if (ramp[i] != 0) && (ramp[i] != 65535) {
            B := float64(i) / 256.0
            A := float64(ramp[i]) / 65535.0
            sum += math.Log(A) / math.Log(B)
            count += 1
        }
    }

    if count != 0 && sum > 0.0 {
        gamma = 1.0 / (sum / float64(count))
    }
    return float32(gamma)
}



// Ramp represents a gamma ramp
type Ramp struct {
    Red [256]uint16
    Green [256]uint16
    Blue [256]uint16
}

// SetGammaRamp sets the gamma to the ramp specified, note
// that this sets the gamma for all attatched monitors, since
// in most cases there is no way to access a per-monitor gamma
func (s *Screen) SetGammaRamp(ramp *Ramp) error {
    s.access.Lock()
    defer s.access.Unlock()

    // Calling C things, get the lock
    chippyAccess.Lock()
    defer chippyAccess.Unlock()

    err := getInitError()
    if err != nil {
        return err
    }

    return setGammaRamp(s, ramp)
}

// GammaRamp returns the currently in use gamma Ramp for all monitors
func (s *Screen) GammaRamp() (*Ramp, error) {
    s.access.Lock()
    defer s.access.Unlock()

    // Calling C things, get the lock
    chippyAccess.Lock()
    defer chippyAccess.Unlock()

    err := getInitError()
    if err != nil {
        return nil, err
    }

    return getGammaRamp(s)
}

// SetGammaRgb sets the gamma of all monitors as rgb float32
func (s *Screen) SetGammaRgb(r, g, b float32) error {
    ramp := Ramp{}

    ramp.Red = calculateGammaRamp(r)
    ramp.Green = calculateGammaRamp(g)
    ramp.Blue = calculateGammaRamp(b)

    return s.SetGammaRamp(&ramp)
}

// GammaRgb returns the gamma of all monitors as rgb float32
func (s *Screen) GammaRgb() (float32, float32, float32, error) {
    ramp, err := s.GammaRamp()
    if err != nil {
        return 0, 0, 0, err
    }

    r := calculateGammaFromRamp(ramp.Red)
    g := calculateGammaFromRamp(ramp.Green)
    b := calculateGammaFromRamp(ramp.Blue)
    return r, g, b, nil
}

// SetGamma sets the rgb gamma of all monitors as an float32
func (s *Screen) SetGamma(gamma float32) error {
    return s.SetGammaRgb(gamma, gamma, gamma)
}

// Gamma returns the gamma of all monitors as an float32
func (s *Screen) Gamma() (float32, error) {
    r, g, b, err := s.GammaRgb()
    if err != nil {
        return 0.0, err
    }
    return (r + g + b) / 3.0, nil
}

