package mouse

type Code uint8
type State uint8

type Button struct {
    // The actual Code that this Button represents, see constants for possibilities
    Code Code

    // The state that the Button is in
    State State

    // Special OS-specific event identifier, this really shouldn't be used directly.
    // See constant "Unknown" for more information.
    //
    // Instead use the Code member, which is an OS-independant Chippy key code.
    Unique int
}

const (
    // The physical button is being held down by the user
    Pressed = State(1)

    // The physical button was released and is no longer being held down by the user
    Released = State(2)

    // The physical mouse wheel is rolling.
    Scrolling = State(3)


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
    Unknown = Code(0)

	// Mouse specific codes
	MouseWheelUp    = Code(1)
	MouseWheelDown  = Code(2)
	MouseWheelLeft  = Code(3)
	MouseWheelRight = Code(4)

	// Gaming mice have lots of buttons, it's doubtful that some platforms support this many, but you never know
	MouseButton1  = Code(5)
	MouseButton2  = Code(6)
	MouseButton3  = Code(7)
	MouseButton4  = Code(8)
	MouseButton5  = Code(9)
	MouseButton6  = Code(10)
	MouseButton7  = Code(11)
	MouseButton8  = Code(12)
	MouseButton9  = Code(13)
	MouseButton10 = Code(14)
	MouseButton11 = Code(15)
	MouseButton12 = Code(16)
	MouseButton13 = Code(17)
	MouseButton14 = Code(18)
	MouseButton15 = Code(19)
	MouseButton16 = Code(20)
	MouseButton17 = Code(21)
	MouseButton18 = Code(22)
	MouseButton19 = Code(23)
	MouseButton20 = Code(24)
	MouseButton21 = Code(25)
	MouseButton22 = Code(26)
	MouseButton23 = Code(27)
	MouseButton24 = Code(28)
	MouseButton25 = Code(29)
	MouseButton26 = Code(30)
	MouseButton27 = Code(31)
	MouseButton28 = Code(32)
	MouseButton29 = Code(33)
	MouseButton30 = Code(34)
	// I've yet to find a mouse with more than 20 buttons, I've listed 30 possibilities here just in case and for
	// future expansion, but I doubt more than 20 button mice will exist, it seems crazy (just a tad).

	// Aliases for MouseButton1, MouseButton2, and MouseButton3
	MouseButtonLeft   = MouseButton1
	MouseButtonMiddle = MouseButton2
	MouseButtonRight  = MouseButton3
)
