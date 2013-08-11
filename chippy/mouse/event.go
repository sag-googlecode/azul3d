// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package mouse

import (
	"fmt"
	"time"
)

// Event represents an single mouse event, such as pushing an button, or using the scroller, etc.
type Event struct {
	T      time.Time
	Button Button
	State  State
}

// Implements the chippy.Event interface.
func (e *Event) Time() time.Time {
	return e.T
}

// String returns an string representation of this event.
func (e *Event) String() string {
	return fmt.Sprintf("mouse.Event(Button=%s, State=%s, Time=%v)", e.Button.String(), e.State.String(), e.T)
}
