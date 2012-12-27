package chippy

import(
    "fmt"
)

// FrameBufferConfig represents options to configure the frame buffer at render target creation time
type FrameBufferConfig struct {
    backend_frameBufferConfig

    // The number of bits for each color per pixel, for instance, consider the following:
    // RedBits=8, GreenBits=8, BlueBits=8, AlphaBits=8, that's a 32 bits per pixel frame
    // buffer, meaning this is True Color (24bpp RGB, with 8 bits of alpha), according to
    // Wikipedia: http://en.wikipedia.org/wiki/Color_depth
	RedBits, GreenBits, BlueBits, AlphaBits uint8

    // An number representing the accumulation buffer bits for each color, note that most of
    // modern GPU's (Nvidia/ATI/Intel) have no support for hardware accelerated accumulation
    // buffers meaning it will be all CPU, and very slow to use.
	AccumRedBits, AccumGreenBits, AccumBlueBits, AccumAlphaBits uint8

    // An number representing how many bits make up each pixel inside the Depth Buffer    
	DepthBits uint8

    // An number representing how many bits make up each pixel inside the Stencil Buffer
    StencilBits uint8

    // An number representing how many samples this configuration has (See: http://www.opengl.org/wiki/Multisampling)
    // like 2, 4, 6, 8, 16, for 2x, 4x, 6x, 8x, or 16x multisample, respectively
    Samples uint8

    // An number representing how many auxillary buffers this OpenGL context will have
    AuxBuffers uint8

    // true/false representing weather this configuration is double buffered (avoids tearing)
	DoubleBuffered bool

    // true/false representing weather this configuration is capable of steroscopic 3d
    StereoScopic bool
}

// String returns a string representation of this FrameBufferConfig
func (f *FrameBufferConfig) String() string {
	return fmt.Sprintf("FrameBufferConfig(RedBits=%d, greenBits=%d, BlueBits=%d, AlphaBits=%d, AccumRedBits=%d, AccumGreenBits=%d, AccumBlueBits=%d, AccumAlphaBits=%d, DepthBits=%d, StencilBits=%d, Samples=%d, AuxBuffers=%d, DoubleBuffered=%t, StereoScopic=%t)", f.RedBits, f.GreenBits, f.BlueBits, f.AlphaBits, f.AccumRedBits, f.AccumGreenBits, f.AccumBlueBits, f.AccumAlphaBits, f.DepthBits, f.StencilBits, f.Samples, f.AuxBuffers, f.DoubleBuffered, f.StereoScopic)
}



// FrameBufferConfigs returns all available frame buffer configurations, you may either choose a specific one and use that
// or you may use the ChooseConfig() function which can be very helpful in choosing appropriate frame buffer configurations.
//
// It's possible for frame buffer configurations to be different depending on the Screen, so you must pass in which Screen
// you intend to use this configuration on. This will only return a list of configurations for that Screen.
func FrameBufferConfigs(screen *Screen) ([]*FrameBufferConfig, error) {
    return backend_frameBufferConfigs(screen)
}

// ChooseConfig returns an appropriate frame buffer configuration for you to use, based off a lowest and highest
// target. You must pass in the configs for ChooseConfig to choose from, this can be gotten from FrameBufferConfigs()
//
// minConfig is an FrameBufferConfig representing the lowest configuration you're willing to accept.
//
// maxConfig is an FrameBufferConfig representing the highest configuration you're willing to accept.
//
// If there is no FrameBufferConfig that is both above minConfig and below maxConfig, then nil will be returned.
// This means that there is no appropriate FrameBufferConfig for what you requested.
//
// If there are multiple FrameBufferConfig's that are both above minConfig and below maxConfig, then the best one
// within that criteria will be chosen and returned.
func ChooseConfig(configs []*FrameBufferConfig, minConfig, maxConfig *FrameBufferConfig) *FrameBufferConfig {
    if len(configs) == 0 {
        return nil
    }

    // Firstly, we remove any configurations that are over maxConfig, and we remove any configurations that are under minConfig
    newConfigs := []*FrameBufferConfig{}

	for i := 0; i < len(configs); i++ {
		config := configs[i]
		if config.RedBits > maxConfig.RedBits || config.RedBits < minConfig.RedBits {
			continue
		}
		if config.GreenBits > maxConfig.GreenBits || config.GreenBits < minConfig.GreenBits {
			continue
		}
		if config.BlueBits > maxConfig.BlueBits || config.BlueBits < minConfig.BlueBits {
			continue
		}
		if config.AlphaBits > maxConfig.AlphaBits || config.AlphaBits < minConfig.AlphaBits {
			continue
		}

		if config.AccumRedBits > maxConfig.AccumRedBits || config.AccumRedBits < minConfig.AccumRedBits {
			continue
		}
		if config.AccumGreenBits > maxConfig.AccumGreenBits || config.AccumGreenBits < minConfig.AccumGreenBits {
			continue
		}
		if config.AccumBlueBits > maxConfig.AccumBlueBits || config.AccumBlueBits < minConfig.AccumBlueBits {
			continue
		}
		if config.AccumAlphaBits > maxConfig.AccumAlphaBits || config.AccumAlphaBits < minConfig.AccumAlphaBits {
			continue
		}

		if config.DepthBits > maxConfig.DepthBits || config.DepthBits < minConfig.DepthBits {
			continue
		}
		if config.StencilBits > maxConfig.StencilBits || config.StencilBits < minConfig.StencilBits {
			continue
		}
		if config.Samples > maxConfig.Samples || config.Samples < minConfig.Samples {
			continue
		}
		if config.AuxBuffers > maxConfig.AuxBuffers || config.AuxBuffers < minConfig.AuxBuffers {
			continue
		}

		if config.DoubleBuffered && !maxConfig.DoubleBuffered || !config.DoubleBuffered && minConfig.DoubleBuffered {
			continue
		}
		if config.StereoScopic && !maxConfig.StereoScopic || !config.StereoScopic && minConfig.StereoScopic {
			continue
		}

		newConfigs = append(newConfigs, config)
	}
    configs = newConfigs

	// Now accumulate the frame buffers
	accumulated := make(map[*FrameBufferConfig]int32)
	for i := 0; i < len(configs); i++ {
		config := configs[i]
		a := int32(0)
		a += int32(config.RedBits + config.GreenBits + config.BlueBits + config.AlphaBits)
		a += int32(config.AccumRedBits + config.AccumGreenBits + config.AccumBlueBits + config.AccumAlphaBits)
		a += int32(config.DepthBits + config.StencilBits + config.Samples + config.AuxBuffers)
		if config.DoubleBuffered {
			a += 1
		}
		if config.StereoScopic {
			a += 1
		}
		accumulated[config] = a
	}

	// Now grab the largest (best) one
	var best *FrameBufferConfig
	var bestValue int32
	for k, v := range accumulated {
		if v > bestValue {
			bestValue = v
			best = k
		}
	}

	return best
}


// BestConfig represents the best (and impossible) FrameBufferConfig that you could possibly have,
// pass this into ChooseConfig() as the maxConfig parameter and you will get the *best* possible frame
// buffer configuration available
//
// Note: This is specifically an FrameBufferConfig with *no* accumulation buffer, due to the lack of
// hardware accelerated accumulation buffers on modern hardware, and the fact that they're really never
// used in modern programming as well.
//
// If you want an accumulation buffer, you'll need to use your own version of BestConfig in place of this.
var BestConfig = &FrameBufferConfig{
	RedBits: 255, GreenBits: 255, BlueBits: 255, AlphaBits: 255,
	AccumRedBits: 0, AccumGreenBits: 0, AccumBlueBits: 0, AccumAlphaBits: 0,
	DepthBits: 255, StencilBits: 255, Samples: 255, AuxBuffers: 255,
	DoubleBuffered: true, StereoScopic: true,
}

// WorstConfig represents the worst (and always possible) FrameBufferConfiguration that you could possibly have,
// pass this into ChooseConfig() as the minConfig parameter and you will get at least *some* frame buffer
// configuration capable of rendering.
var WorstConfig = &FrameBufferConfig{
	RedBits: 0, GreenBits: 0, BlueBits: 0, AlphaBits: 0,
	AccumRedBits: 0, AccumGreenBits: 0, AccumBlueBits: 0, AccumAlphaBits: 0,
	DepthBits: 0, StencilBits: 0, Samples: 0, AuxBuffers: 0,
	DoubleBuffered: false, StereoScopic: false,
}


