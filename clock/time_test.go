package clock

import (
	"testing"
	"time"
)

func TestHighResolutionTime(t *testing.T) {
	lrStart := time.Now()
	hrStart := Time()

	var diffTotal time.Duration
	for i := 0; i < 10; i++ {
		lrDiff := time.Since(lrStart)
		hrDiff := Time() - hrStart

		diffTotal += hrDiff
		t.Logf("%d.\ttime.Since()=%d\tclock.Time()=%d", i, lrDiff, hrDiff)

		lrStart = time.Now()
		hrStart = Time()
	}

	if diffTotal <= 0 {
		t.Fail()
	}
}
