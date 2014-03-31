// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

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
		return "keyboard.InvalidState"
	case Down:
		return "keyboard.Down"
	case Up:
		return "keyboard.Up"
	}
	return fmt.Sprintf("keyboard.State(%d)", s)
}
