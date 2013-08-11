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
type StateEvent struct {
	T     time.Time
	Key   Key
	State State
}

// Implements the chippy.Event interface.
func (e *StateEvent) Time() time.Time {
	return e.T
}

// String returns an string representation of this event.
func (e *StateEvent) String() string {
	return fmt.Sprintf("keyboard.StateEvent(Key=%v, State=%v, Time=%v)", e.Key.String(), e.State.String(), e.T)
}

// TypedEvent represents an event where some sort of user input has generated
// an input character which should be considered input.
type TypedEvent struct {
	T    time.Time
	Rune rune
}

// Implements the chippy.Event interface.
func (e *TypedEvent) Time() time.Time {
	return e.T
}

// String returns an string representation of this event.
func (e *TypedEvent) String() string {
	return fmt.Sprintf("keyboard.TypedEvent(Rune=%U %q, Time=%v)", e.Rune, string(e.Rune), e.T)
}
