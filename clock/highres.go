// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

// +build !windows

package clock

import (
	"time"
)

// In here we simply fallback to the standard time package for systems that already support high
// resolution timers.

func Time() time.Duration {
	return highResTimeFallback()
}
