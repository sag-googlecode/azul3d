// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package keyboard

import (
	"sync"
)

type StateWatcherInterface interface {
	SetKeyState(key Key, state State)
	KeyStates() map[Key]State
	KeyState(key Key) State
	KeyDown(key Key) bool
	KeyUp(key Key) bool
}

type StateWatcher struct {
	access sync.RWMutex
	states map[Key]State
}

// SetKeyState specifies the current state of the specified key.
func (s *StateWatcher) SetKeyState(key Key, state State) {
	s.access.Lock()
	defer s.access.Unlock()

	s.states[key] = state
}

// KeyStates returns an copy of the internal key state map used by this StateWatcher.
func (s *StateWatcher) KeyStates() map[Key]State {
	s.access.RLock()
	defer s.access.RUnlock()

	copy := make(map[Key]State)
	for key, state := range s.states {
		copy[key] = state
	}
	return copy
}

// KeyState returns the current state of the specified key.
func (s *StateWatcher) KeyState(key Key) State {
	s.access.Lock()
	defer s.access.Unlock()

	state, ok := s.states[key]
	if !ok {
		s.states[key] = Down
	}
	return state
}

// KeyDown tells weather the specified key is currently in the down state.
func (s *StateWatcher) KeyDown(key Key) bool {
	return s.KeyState(key) == Down
}

// KeyUp tells weather the specified key is currently in the up state.
func (s *StateWatcher) KeyUp(key Key) bool {
	return s.KeyState(key) == Up
}

func NewStateWatcher() *StateWatcher {
	s := new(StateWatcher)
	s.states = make(map[Key]State)
	return s
}
