package keyboard

import (
	"fmt"
)

type State uint8

const (
	InvalidState State  = iota
	Down                // Being held down currently
	Up                  // No longer being held down (released)
	On           = Down // the on/active state (for lock keys; Caps Lock; Num Lock; etc..)
	Off          = Up   // the off/inactive state (for lock keys; Caps Lock; Num Lock; etc..)
)

func (s State) String() string {
	switch s {
	case InvalidState:
		return "InvalidState"
	case Down:
		return "Down"
	case Up:
		return "Up"
	}
	return fmt.Sprintf("Unknown(%d)", s)
}

// StateEvent represents an event when an keyboard key changes state, that is, being pushed down
// when it was previously up, or being toggled on when it was previously off, etc.
type StateEvent struct {
	Key   Key
	State State
}

func (e *StateEvent) String() string {
	return fmt.Sprintf("StateEvent(Key=%v, State=%v)", e.Key.String(), e.State.String())
}
