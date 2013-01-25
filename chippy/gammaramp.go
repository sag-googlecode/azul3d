package chippy

import (
	"errors"
	"math"
)

// GammaRamp represents an slice of pixel color values, for each color: Red, Green, and Blue
//
// Each element in one of Red, Green, or Blue arrays, represent the intensity of the color once it
// is displayed on the screen, that is, if:
//  Red[255] = 0
// then any pixel whose Red value was 255, will now be 0 instead.
//
// Likewise, in the following case:
//  Red[255] = 0.5
// then any pixel whose Red value was 0, will now be 128.
//
// It should be noted, that different hardware supports different array length of gamma ramps, so
// each color (Red, Green, Blue) will have to be of the specific length the hardware supports, take
// an look at GammaRampSize() to obtain the specific size that the hardware supports.
//
// This is particularly useful for gamma correction (See: http://en.wikipedia.org/wiki/Gamma_correction)
type GammaRamp struct {
	Red, Green, Blue []float32
}

// SetAsGammaBrightnessContrast sets each respective color in the GammaRamp to one calculated from
// respective gamma, brightness, and contrast values.
//
// If contrast is an negative number, an error is returned.
//
// If brightness is outside the inclusive range of -1.0 to 1.0, an error is returned.
func (r *GammaRamp) SetAsBrightnessContrastGamma(rampSize int, brightness, contrast, gamma float32) error {
	if contrast < 0.0 {
		return errors.New("contrast must be an positive number!")
	}

	if brightness < -1.0 || brightness > 1.0 {
		return errors.New("brightness must be inside range -1.0 to 1.0")
	}

	r.Red = make([]float32, rampSize)
	r.Green = make([]float32, rampSize)
	r.Blue = make([]float32, rampSize)

	for i := 0; i < rampSize; i++ {
		intensity := float32(i) / (float32(rampSize) - 1.0)

		value := float32(math.Pow(float64(intensity), float64(gamma))) // gamma
		value += brightness                                            // brightness
		value = (value-0.5)*contrast + 0.5                             // contrast

		if value > 1.0 {
			value = 1.0
		} else if value < 0.0 {
			value = 0.0
		}

		r.Red[i] = value
		r.Green[i] = value
		r.Blue[i] = value
	}
	return nil
}
