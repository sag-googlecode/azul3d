// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package keyboard

import (
	"sync"
)

// Watcher watches the state of various mouse buttons.
type Watcher struct {
	access sync.RWMutex
	states map[Key]State
}

// SetState specifies the current state of the specified key.
func (s *Watcher) SetState(key Key, state State) {
	s.access.Lock()
	defer s.access.Unlock()

	s.states[key] = state
}

// States returns an copy of the internal key state map used by this watcher.
func (s *Watcher) States() map[Key]State {
	s.access.RLock()
	defer s.access.RUnlock()

	copy := make(map[Key]State)
	for key, state := range s.states {
		copy[key] = state
	}
	return copy
}

// State returns the current state of the specified key.
func (s *Watcher) State(key Key) State {
	s.access.Lock()
	defer s.access.Unlock()

	state, ok := s.states[key]
	if !ok {
		s.states[key] = Down
	}
	return state
}

// Down tells weather the specified key is currently in the down state.
func (s *Watcher) Down(key Key) bool {
	return s.State(key) == Down
}

// Up tells weather the specified key is currently in the up state.
func (s *Watcher) Up(key Key) bool {
	return s.State(key) == Up
}

// NewWatcher returns a new, initialized, watcher.
func NewWatcher() *Watcher {
	s := new(Watcher)
	s.states = make(map[Key]State)
	return s
}
