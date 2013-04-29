package keyboard

// Note: constants defined here should be spelled out in plain english.
//
// If the key would generate specific symbols, then include those symbols in an comment after the line
//
// Including links to where you found such keys, is helpful.

const (
	Invalid Key = iota

	// http://en.wikipedia.org/wiki/File:KB_United_States-NoAltGr.svg
	Tilde        // "~"
	Dash         // "-"
	Equals       // "="
	Semicolon    // ";"
	Apostrophe   // "'"
	Comma        // ","
	Period       // "."
	ForwardSlash // "/"
	Backspace
	Tab // "\t"
	CapsLock
	Space // " "
	Menu
	Enter // "\r", "\n", "\r\n"
	Escape
	Insert
	PrintScreen
	Delete
	PageUp
	PageDown
	Home
	End

	// Arrow keys
	ArrowLeft
	ArrowRight
	ArrowDown
	ArrowUp

	// Lefties
	LeftBracket // [
	LeftShift
	LeftCtrl
	LeftSuper
	LeftAlt

	// Righties
	RightBracket // ]
	RightShift
	RightCtrl
	RightSuper
	RightAlt

	// Numbers
	Zero  // "0"
	One   // "1"
	Two   // "2"
	Three // "3"
	Four  // "4"
	Five  // "5"
	Six   // "6"
	Seven // "7"
	Eight // "8"
	Nine  // "9"

	// Functions
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

	// English characters
	A // "a"
	B // "b"
	C // "c"
	D // "d"
	E // "e"
	F // "f"
	G // "g"
	H // "h"
	I // "i"
	J // "j"
	K // "k"
	L // "l"
	M // "m"
	N // "n"
	O // "o"
	P // "p"
	Q // "q"
	R // "r"
	S // "s"
	T // "t"
	U // "u"
	V // "v"
	W // "w"
	X // "x"
	Y // "y"
	Z // "z"

	// Number pads
	NumLock
	NumMultiply // "*"
	NumDivide   // "/"
	NumAdd      // "+"
	NumSubtract // "-"
	NumZero     // "0"
	NumOne      // "1"
	NumTwo      // "2"
	NumThree    // "3"
	NumFour     // "4"
	NumFive     // "5"
	NumSix      // "6"
	NumSeven    // "7"
	NumEight    // "8"
	NumNine     // "9"
	NumDecimal  // "."
	NumEnter
)
