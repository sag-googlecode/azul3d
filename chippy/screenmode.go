package chippy

// We use this type for sorting the ScreenModes in backends
type sortedScreenModes []ScreenMode

func (s sortedScreenModes) Len() int {
    return len(s)
}

func (s sortedScreenModes) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

func (s sortedScreenModes) Less(i, j int) bool {
    iWidth, iHeight := s[i].Resolution()
	iResolution := iWidth + iHeight
    iRefreshRate := s[i].RefreshRate()
    iBytesPerPixel := s[i].BytesPerPixel()

    jWidth, jHeight := s[j].Resolution()
	jResolution := jWidth + jHeight
    jRefreshRate := s[j].RefreshRate()
    jBytesPerPixel := s[j].BytesPerPixel()

	// if resolution and bpp are the same, sort by refresh rate
	if iResolution == jResolution && iBytesPerPixel == jBytesPerPixel {
		return iRefreshRate > jRefreshRate

	// Second case, if resolution is the same, sort by bpp
	} else if iResolution == jResolution {
		return iBytesPerPixel > jBytesPerPixel
	}

	// First case, sort by resolution
	return iResolution > jResolution

	/*
	// Multiply by 100.0 to get rid of smaller inconsitencies (where 1366+768+60+8 less than 1360+768+60+32)
	e := 100.0
    iWeight := (float32(iWidth) * e) + (float32(iHeight) * e) + (float32(iBytesPerPixel) * e) + (iRefreshRate * e)
    jWeight := (float32(jWidth) * e) + (float32(jHeight) * e) + (float32(jBytesPerPixel) * e) + (jRefreshRate * e)
	return iWeight > jWeight
	*/
}

// ScreenMode represents an single, unique, screen mode, with an resolution, refresh rate, and bpp.
//
// It is possible for multiple different ScreenMode's to exist with the same resolution, each with
// different refresh rates or bytes per pixel, respectively.
type ScreenMode interface {
    // String returns an nice string representing this ScreenMode
    String() string

    // Equals compares two ScreenMode(s) for equality. It does this by comparing resolutions,
    // refresh rates, and bytes per pixels.
    Equals(other ScreenMode) bool

    // Resolution returns the width and height of this ScreenMode, in pixels.
    Resolution() (width, height uint)

    // RefreshRate returns the refresh rate of this ScreenMode, in hertz, or 0 if the refresh rate
    // is unable to be determined.
    RefreshRate() float32

    // BytesPerPixel returns the number of bytes that represent an single pixel of this ScreenMode,
    // or 0 if the bytes per pixel is unable to be determined.
    BytesPerPixel() uint
}


