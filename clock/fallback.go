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
