// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

// Package mouse implements various mouse related data types.
package mouse

import (
	"fmt"
)

// State represents an single mouse state, such as Up, Down, or an scroll
// direction.
type State uint8

const (
	InvalidState State = iota
	Down
	Up
	ScrollForward
	ScrollBack
	ScrollLeft
	ScrollRight
)

// String returns an string representation of the mouse state.
func (s State) String() string {
	switch s {
	case InvalidState:
		return "mouse.InvalidState"
	case Down:
		return "mouse.Down"
	case Up:
		return "mouse.Up"
	case ScrollForward:
		return "mouse.ScrollForward"
	case ScrollBack:
		return "mouse.ScrollBack"
	case ScrollLeft:
		return "mouse.ScrollLeft"
	case ScrollRight:
		return "mouse.ScrollRight"
	}
	return fmt.Sprintf("mouse.State(%d)")
}

// Button represents an single mouse button.
type Button uint8

const (
	Invalid Button = iota
	One
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
)

// String returns an string representation of the mouse button.
func (b Button) String() string {
	switch b {
	case Invalid:
		return "mouse.Invalid"
	case Left:
		return "mouse.Left"
	case Right:
		return "mouse.Right"
	case Wheel:
		return "mouse.Wheel"
	case Four:
		return "mouse.Four"
	case Five:
		return "mouse.Five"
	case Six:
		return "mouse.Six"
	case Seven:
		return "mouse.Seven"
	case Eight:
		return "mouse.Eight"
	}
	return fmt.Sprintf("mouse.Button(%d)", b)
}

const (
	Left  = One
	Right = Two
	Wheel = Three
)
