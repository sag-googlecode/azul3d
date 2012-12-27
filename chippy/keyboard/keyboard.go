package keyboard

import(
    "fmt"
)

type Code uint8



type State uint8

const(
    // The physical button is being held down by the user
    Pressed State = iota

    // The physical button was released and is no longer being held down by the user
    Released
)

func (s State) String() string {
    switch s{
        case Pressed:
            return "Pressed"
        case Released:
            return "Released"
    }
    return fmt.Sprintf("State(%d)", s)
}


type Button struct {
	// The unicode rune that this Button represents -- Only some Button's have an valid unicode representation,
	// for instance, Caps Lock has no valid unicode representation, so do take note of cases like these.
	//
	// If this Button is unable to be represented by an unicode rune, like in the case of Caps Lock, this rune
	// will be zero (0), and you may identify this Button using the Code member below.
	Rune rune

	// The Code that this Button represents, this Code will be zero (0) if this Button has an valid an Rune (non-zero), that is
	// if this Button has an valid unicode representation.
	//
	// There are some cases where you might want to make a distinction between two different keyboard buttons, that represent the
	// same valid unicode rune, for instance the + key, there is + on the main part of most keyboards, and also an + on keyboards
	// that have number pads.
	//
	// In the above case, Rune will be an valid unicode rune (+), and Code will be a valid Code (Plus or NumberPadPlus, respectively)
	Code Code

    // The state that the Button is in
    State State

    // Special OS-specific event identifier, this really shouldn't be used directly.
    // See constant "Unknown" for more information.
    //
    // Instead use the Code member, which is an OS-independant Chippy key code.
    Unique int
}

func (b *Button) String() string {
    return fmt.Sprintf("Button(Rune = (" + string(b.Rune) + ") %U, Code = " + b.Code.String() + ", State = " + b.State.String() + ", Unique = %d)", b.Rune, b.Unique)
}

// These are all layed out in a nice neat order, but that's just because I typed them this way
// When adding support for other language keyboards (this is currently just support for US keyboards)
// the order and pretty-ness is less important, godoc formats these anyway

const (
    // It should never be Unknown, but if it is it means the following:
    //
    // The user did /something/ with the mouse.
    //
    // We have no idea what that /something/ is.
    //
    // Support for that /something/ should be added to Chippy ASAP.
    //
    // An special OS-specific identifier exists as Button.Unique, but ensure that you never expect this to
    // be the same across multiple operating systems. This is a low level event id. This is helpful for
    // telling is which event we should impliment inside Chippy.
    Unknown Code = iota

	// Function keys
	F1
	F2
	F3
	F4
	F5
	F6
	F7
	F8
	F9
	F10
	F11
	F12
	F13
	F14
	F15
	F16
	F17
	F18
	F19
	F20
	F21
	F22
	F23
	F24
	F25
	F26
	F27
	F28
	F29
	F30
	F31
	F32
	F33
	F34
	F35
	// More could be added later, if need be, it still seems crazy. Also it should be noted that
	// basically nobody has a keyboard with 35 function keys... Normal keyboards have F1-F12 keys
	// Some laptops only have F1-F10 keys

	// All other non-unicode rune representable keyboard buttons
	Escape
	Tab
	CapsLock
    ScrollLock
	PrintScreen
    Backspace
	Menu
	Pause

	LeftShift
	RightShift
	LeftControl
	RightControl
	LeftAlt
	RightAlt
	LeftSuper
	RightSuper

	// Number Pad duplicates
	Number0
	Number1
	Number2
	Number3
	Number4
	Number5
	Number6
	Number7
	Number8
	Number9
	ForwardSlash
	Asterisk
	Plus
	HyphenMinus
	Delete
	Insert
	NumLock
	Home
	PageUp
	PageDown
	End
	Enter
	Up
	Down
	Left
	Right
    Period

	NumPadNumber0
	NumPadNumber1
	NumPadNumber2
	NumPadNumber3
	NumPadNumber4
	NumPadNumber5
	NumPadNumber6
	NumPadNumber7
	NumPadNumber8
	NumPadNumber9
	NumPadForwardSlash
	NumPadAsterisk
	NumPadPlus
	NumPadHyphenMinus
	NumPadDelete
	NumPadInsert
	NumPadHome
	NumPadPageUp
	NumPadPageDown
	NumPadEnd
	NumPadEnter
	NumPadUp
	NumPadDown
	NumPadLeft
	NumPadRight
	NumPadPeriod
)

// Aliases provided as convienance for developers who are used to other operating systems
const(
	// Aliases for LeftSuper
	LeftWindows = LeftSuper
	LeftCommand = LeftSuper

	// Aliases for RightSuper
	RightWindows = RightSuper
	RightCommand = RightSuper
)

func (c Code) String() string {
    switch c{
        case F1:
            return "F1"
        case F2:
            return "F2"
        case F3:
            return "F3"
        case F4:
            return "F4"
        case F5:
            return "F5"
        case F6:
            return "F6"
        case F7:
            return "F7"
        case F8:
            return "F8"
        case F9:
            return "F9"
        case F10:
            return "F10"
        case F11:
            return "F11"
        case F12:
            return "F12"
        case F13:
            return "F13"
        case F14:
            return "F14"
        case F15:
            return "F15"
        case F16:
            return "F16"
        case F17:
            return "F17"
        case F18:
            return "F18"
        case F19:
            return "F19"
        case F20:
            return "F20"
        case F21:
            return "F21"
        case F22:
            return "F22"
        case F23:
            return "F23"
        case F24:
            return "F24"
        case F25:
            return "F25"
        case F26:
            return "F26"
        case F27:
            return "F27"
        case F28:
            return "F28"
        case F29:
            return "F29"
        case F30:
            return "F30"
        case F31:
            return "F31"
        case F32:
            return "F32"
        case F33:
            return "F33"
        case F34:
            return "F34"
        case F35:
            return "F35"

	    case Escape:
            return "Escape"
	    case Tab:
            return "Tab"
	    case CapsLock:
            return "CapsLock"
        case ScrollLock:
            return "ScrollLock"
	    case PrintScreen:
            return "PrintScreen"
        case Backspace:
            return "Backspace"
	    case Menu:
            return "Menu"
	    case Pause:
            return "Pause"

	    case LeftShift:
            return "LeftShift"
	    case RightShift:
            return "RightShift"
	    case LeftControl:
            return "LeftControl"
	    case RightControl:
            return "RightControl"
	    case LeftAlt:
            return "LeftAlt"
	    case RightAlt:
            return "RightAlt"
	    case LeftSuper:
            return "LeftSuper"
	    case RightSuper:
            return "RightSuper"

	    case Number0:
            return "Number0"
	    case Number1:
            return "Number1"
	    case Number2:
            return "Number2"
	    case Number3:
            return "Number3"
	    case Number4:
            return "Number4"
	    case Number5:
            return "Number5"
	    case Number6:
            return "Number6"
	    case Number7:
            return "Number7"
	    case Number8:
            return "Number8"
	    case Number9:
            return "Number9"
	    case ForwardSlash:
            return "ForwardSlash"
	    case Asterisk:
            return "Asterisk"
	    case Plus:
            return "Plus"
	    case HyphenMinus:
            return "HyphenMinus"
	    case Delete:
            return "Delete"
	    case Insert:
            return "Insert"
	    case NumLock:
            return "NumLock"
	    case Home:
            return "Home"
	    case PageUp:
            return "PageUp"
	    case PageDown:
            return "PageDown"
	    case End:
            return "End"
	    case Enter:
            return "Enter"
	    case Up:
            return "Up"
	    case Down:
            return "Down"
	    case Left:
            return "Left"
	    case Right:
            return "Right"
        case Period:
            return "Period"

	    case NumPadNumber0:
            return "NumPadNumber0"
	    case NumPadNumber1:
            return "NumPadNumber1"
	    case NumPadNumber2:
            return "NumPadNumber2"
	    case NumPadNumber3:
            return "NumPadNumber3"
	    case NumPadNumber4:
            return "NumPadNumber4"
	    case NumPadNumber5:
            return "NumPadNumber5"
	    case NumPadNumber6:
            return "NumPadNumber6"
	    case NumPadNumber7:
            return "NumPadNumber7"
	    case NumPadNumber8:
            return "NumPadNumber8"
	    case NumPadNumber9:
            return "NumPadNumber9"
	    case NumPadForwardSlash:
            return "NumPadForwardSlash"
	    case NumPadAsterisk:
            return "NumPadAsterisk"
	    case NumPadPlus:
            return "NumPadPlus"
	    case NumPadHyphenMinus:
            return "NumPadHyphenMinus"
	    case NumPadDelete:
            return "NumPadDelete"
	    case NumPadInsert:
            return "NumPadInsert"
	    case NumPadHome:
            return "NumPadHome"
	    case NumPadPageUp:
            return "NumPadPageUp"
	    case NumPadPageDown:
            return "NumPadPageDown"
	    case NumPadEnd:
            return "NumPadEnd"
	    case NumPadEnter:
            return "NumPadEnter"
	    case NumPadUp:
            return "NumPadUp"
	    case NumPadDown:
            return "NumPadDown"
	    case NumPadLeft:
            return "NumPadLeft"
	    case NumPadRight:
            return "NumPadRight"
	    case NumPadPeriod:
            return "NumPadPeriod"
    }
    return fmt.Sprintf("Code(%d)", c)
}

