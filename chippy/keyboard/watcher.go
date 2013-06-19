package keyboard

import (
	"sync"
)

type StateWatcherInterface interface {
	SetState(key Key, state State)
	State(key Key) State
	Down(key Key) bool
	Up(key Key) bool
}

type StateWatcher struct {
	access sync.RWMutex
	states map[Key]State
}

// SetState specifies the current state of the specified key.
func (s *StateWatcher) SetState(key Key, state State) {
	s.access.Lock()
	defer s.access.Unlock()

	s.states[key] = state
}

// State returns the current state of the specified key.
func (s *StateWatcher) State(key Key) State {
	s.access.RLock()
	defer s.access.RUnlock()

	state, ok := s.states[key]
	if !ok {
		s.states[key] = Down
	}
	return state
}

// Down tells weather the specified key is currently in the down state.
func (s *StateWatcher) Down(key Key) bool {
	return s.State(key) == Down
}

// Up tells weather the specified key is currently in the up state.
func (s *StateWatcher) Up(key Key) bool {
	return s.State(key) == Up
}

func NewStateWatcher() *StateWatcher {
	s := new(StateWatcher)
	s.states = make(map[Key]State)
	return s
}
