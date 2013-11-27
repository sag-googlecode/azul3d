// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package mouse

import (
	"bytes"
	"fmt"
	"sync"
)

// Watcher watches the state of various mouse buttons.
type Watcher struct {
	access sync.RWMutex
	states map[Button]State
}

// String returns a multi-line string representation of this mouse watcher and
// it's associated states.
func (w *Watcher) String() string {
	bb := new(bytes.Buffer)
	fmt.Fprintf(bb, "mouse.Watcher(\n")
	for b, s := range w.States() {
		fmt.Fprintf(bb, "    %v: %v\n", b, s)
	}
	fmt.Fprintf(bb, ")")
	return bb.String()
}

// SetState specifies the current state of the specified mouse button.
func (s *Watcher) SetState(button Button, state State) {
	s.access.Lock()
	defer s.access.Unlock()

	s.states[button] = state
}

// States returns an copy of the internal mouse button state map used by this watcher.
func (s *Watcher) States() map[Button]State {
	s.access.RLock()
	defer s.access.RUnlock()

	copy := make(map[Button]State)
	for button, state := range s.states {
		copy[button] = state
	}
	return copy
}

// State returns the current state of the specified mouse button.
func (s *Watcher) State(button Button) State {
	s.access.Lock()
	defer s.access.Unlock()

	state, ok := s.states[button]
	if !ok {
		s.states[button] = Up
	}
	return state
}

// Down tells weather the specified mouse button is currently in the down state.
func (s *Watcher) Down(button Button) bool {
	return s.State(button) == Down
}

// Up tells weather the specified mouse button is currently in the up state.
func (s *Watcher) KeyUp(button Button) bool {
	return s.State(button) == Up
}

// NewWatcher returns a new, initialized, watcher.
func NewWatcher() *Watcher {
	s := new(Watcher)
	s.states = make(map[Button]State)
	return s
}
