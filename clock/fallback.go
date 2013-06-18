// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package clock

import (
	"time"
)

var (
	fallbackStart time.Time
)

func init() {
	fallbackStart = time.Now()
}

func highResTimeFallback() time.Duration {
	return time.Since(fallbackStart)
}
