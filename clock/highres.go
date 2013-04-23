// +build !windows

package clock

// In here we simply fallback to the standard time package for systems that already support high
// resolution timers.

func Time() {
	return highResTimeFallback()
}
