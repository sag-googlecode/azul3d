package keyboard

import(
	"fmt"
)

type State uint8
type Key uint32
type OSKey uint64

const (
	InvalidState State = iota
	Down               // Being held down currently
	Up                 // No longer being held down
	On                 // Toggled to the on/active state, Caps Lock, Num Lock, etc
	Off                // Toggled to the off/inactive state, Caps Lock, Num Lock, etc
)

func (s State) String() string {
	switch(s) {
		case InvalidState:
			return "InvalidState"
		case Down:
			return "Down"
		case Up:
			return "Up"
		case On:
			return "On"
		case Off:
			return "Off"
	}
	return fmt.Sprintf("Unknown(%d)", s)
}

// Event represents an single keyboard event, such as pressing an key, releasing an key, etc.
type Event struct {
	// Key code, garanteed to be the same across different operating systems
	Key Key

	// OS key code, garanteed to be unique (and may identify keys not known by Chippy), but not the
	// same across multiple operating systems.
	OSKey OSKey

	// Unicode string that key would generate, useful for creating input boxes for instance.
	Unicode string

	// State the key is in, must be one of the predefined constants, Up, Down, On, Off, etc.
	// or Up (no longer being held down).
	State State
}

func (e *Event) String() string {
	return fmt.Sprintf("Event(Key=%v, OSKey=%v, Unicode=%v, State=%v)", e.Key.String(), e.OSKey, e.Unicode, e.State.String())
}

