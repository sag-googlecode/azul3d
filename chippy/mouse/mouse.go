package mouse

import (
	"fmt"
)

type State uint8

const (
	InvalidState State = iota
	Down
	Up
	ScrollDown
	ScrollUp
	ScrollLeft
	ScrollRight
)

func (s State) String() string {
	switch s {
	case InvalidState:
		return "InvalidState"
	case Down:
		return "Down"
	case Up:
		return "Up"
	case ScrollDown:
		return "ScrollDown"
	case ScrollUp:
		return "ScrollUp"
	case ScrollLeft:
		return "ScrollLeft"
	case ScrollRight:
		return "ScrollRight"
	}
	return ""
}

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

func (b Button) String() string {
	switch b {
	case Invalid:
		return "Invalid"
	case Left:
		return "Left"
	case Right:
		return "Right"
	case Wheel:
		return "Wheel"
	case Four:
		return "Four"
	case Five:
		return "Five"
	case Six:
		return "Six"
	case Seven:
		return "Seven"
	case Eight:
		return "Eight"
	}
	return ""
}

const (
	Left  = One
	Right = Two
	Wheel = Three
)

// Event represents an single mouse event, such as pushing an button, or using the scroller, etc.
type Event struct {
	Button Button
	State  State
}

func (e *Event) String() string {
	return fmt.Sprintf("Event(Button=%s, State=%s)", e.Button.String(), e.State.String())
}
