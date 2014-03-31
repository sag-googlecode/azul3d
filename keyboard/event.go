// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package keyboard

import (
	"fmt"
	"time"
)

// StateEvent represents an event when an keyboard key changes state, that is,
// being pushed down when it was previously up, or being toggled on when it was
// previously off, etc.
//
// A StateEvent may have an Invalid key; in which case the OS unique identifier
// may be used to identify the key uniquely. (E.g. for special keys and non-US
// keyboard keys).
//
// The StateEvent must always have a unique OS value which uniquely represents
// the key, even if the Key value is set to Invalid.
type StateEvent struct {
	T     time.Time
	Key   Key
	OS    OS
	State State
}

// Implements the chippy.Event interface.
func (e StateEvent) Time() time.Time {
	return e.T
}

// String returns an string representation of this event.
func (e StateEvent) String() string {
	return fmt.Sprintf("keyboard.StateEvent(Key=%v, OS=%v, State=%v, Time=%v)", e.Key.String(), e.OS, e.State.String(), e.T)
}

// TypedEvent represents an event where some sort of user input has generated
// an input character which should be considered input.
type TypedEvent struct {
	T    time.Time
	Rune rune
}

// Implements the chippy.Event interface.
func (e TypedEvent) Time() time.Time {
	return e.T
}

// String returns an string representation of this event.
func (e TypedEvent) String() string {
	return fmt.Sprintf("keyboard.TypedEvent(Rune=%U %q, Time=%v)", e.Rune, string(e.Rune), e.T)
}
