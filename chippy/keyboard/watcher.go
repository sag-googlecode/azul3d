// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package keyboard

import (
	"bytes"
	"sync"
	"fmt"
)

type watcherKey struct {
	key Key
	os  OS
}

// Watcher watches the state of various keyboard keys.
type Watcher struct {
	access   sync.RWMutex
	states   map[Key]State
	osStates map[OS]State
}

// String returns a multi-line string representation of this keyboard watcher
// and it's associated states (but not OS ones).
func (w *Watcher) String() string {
	bb := new(bytes.Buffer)
	fmt.Fprintf(bb, "keyboard.Watcher(\n")
	for k, s := range w.States() {
		fmt.Fprintf(bb, "    %v: %v\n", k, s)
	}
	fmt.Fprintf(bb, ")")
	return bb.String()
}

// SetState specifies the current state of the specified key.
func (w *Watcher) SetState(key Key, state State) {
	w.access.Lock()
	defer w.access.Unlock()

	w.states[key] = state
}

// States returns an copy of the internal key state map used by this watcher.
func (w *Watcher) States() map[Key]State {
	w.access.RLock()
	defer w.access.RUnlock()

	copy := make(map[Key]State)
	for key, state := range w.states {
		copy[key] = state
	}
	return copy
}

// State returns the current state of the specified key.
func (w *Watcher) State(key Key) State {
	w.access.Lock()
	defer w.access.Unlock()

	state, ok := w.states[key]
	if !ok {
		w.states[key] = Up
	}
	return state
}

// Down tells weather the specified key is currently in the down state.
func (w *Watcher) Down(key Key) bool {
	return w.State(key) == Down
}

// Up tells weather the specified key is currently in the up state.
func (w *Watcher) Up(key Key) bool {
	return w.State(key) == Up
}

// SetOSState specifies the current state of the specified OS key value.
func (w *Watcher) SetOSState(os OS, state State) {
	w.access.Lock()
	defer w.access.Unlock()

	w.osStates[os] = state
}

// OSStates returns an copy of the internal OS key state map used by this watcher.
func (w *Watcher) OSStates() map[OS]State {
	w.access.RLock()
	defer w.access.RUnlock()

	copy := make(map[OS]State)
	for os, state := range w.osStates {
		copy[os] = state
	}
	return copy
}

// OSState returns the current state of the specified OS key value.
func (w *Watcher) OSState(os OS) State {
	w.access.Lock()
	defer w.access.Unlock()

	state, ok := w.osStates[os]
	if !ok {
		w.osStates[os] = Up
	}
	return state
}

// OSDown tells weather the specified OS key value is currently in the down state.
func (w *Watcher) OSDown(os OS) bool {
	return w.OSState(os) == Down
}

// OSUp tells weather the specified OS key value is currently in the up state.
func (w *Watcher) OSUp(os OS) bool {
	return w.OSState(os) == Up
}

// NewWatcher returns a new, initialized, watcher.
func NewWatcher() *Watcher {
	w := new(Watcher)
	w.states = make(map[Key]State)
	w.osStates = make(map[OS]State)
	return w
}
