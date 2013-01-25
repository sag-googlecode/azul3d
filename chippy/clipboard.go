package chippy

import (
	"fmt"
)

type ClipboardFormat uint8

const (
	// Clipboard format identifier for text content types
	CLIPBOARD_TEXT ClipboardFormat = iota

	// TODO: Add support for more clipboard formats.
)

func panicUnlessValidFormat(format ClipboardFormat) {
	switch format {
	case CLIPBOARD_TEXT:
		return
	}
	panic(fmt.Sprintf("Invalid clipboard format constant used as argument: %d", format))
}

// ClipboardHasData tells weather the clipboard currently has some sort of data inside of it.
func ClipboardHasData() bool {
	return backend_ClipboardHasData()
}

// IsClipboardFormatSupported tells weather the specified clipboard format is supported, meaning it
// is an valid format to be used in an call to SetClipboardData().
//
// The format parameter must be one of the ClipboardFormat constants defined by this package, or
// calling this function will panic.
func IsClipboardFormatSupported(format ClipboardFormat) bool {
	panicUnlessValidFormat(format)
	return backend_IsClipboardFormatSupported(format)
}

// ClipboardData returns the format of the data, as well as an slice of bytes (that is, the actual
// data itself).
//
// An error is returned in the event that we are unable to get the current clipboard contents for
// some reason.
//
// The format parameter must be one of the ClipboardFormat constants defined by this package, or
// calling this function will panic.
func ClipboardData() ([]byte, ClipboardFormat, error) {
	return backend_ClipboardData()
}

// SetClipboardData sets the data in the clipboard to the specified slice of bytes, and clipboard format.
//
// An error is returned in the event that we are unable to set the clipboard data for some reason,
// or if the clipboard format is unavailable for any reason.
//
// The format parameter must be one of the ClipboardFormat constants defined by this package, or
// calling this function will panic.
func SetClipboardData(data []byte, format ClipboardFormat) error {
	panicUnlessValidFormat(format)
	return backend_SetClipboardData(data, format)
}

// ClearClipboard clears the contents of the clipboard, so that the data in the clipboard no longer
// exists.
//
// An error is returned in the event that we are unable to clear the contents of the clipboard for
// some reason.
func ClearClipboard() error {
	return backend_ClearClipboard()
}

// ClipboardString returns the data from ClipboardData() as an string, so long as the format of the
// data returned by ClipboardData() is CLIPBOARD_TEXT.
//
// If the format returned by ClipboardData() is something else, an empty string, and no error, is returned.
//
// If ClipboardData() returns an error due to retrieving the clipboard contents, then an empty
// string, and that error is returned.
func ClipboardString() (string, error) {
	data, format, err := ClipboardData()
	if err != nil {
		return "", err
	}

	if format == CLIPBOARD_TEXT {
		return string(data), nil
	}
	return "", nil
}

// SetClipboardString sets the data in the clipboard to the specified string, this simply calls the
// SetClipboardData() function for you with an clipboard format of CLIPBOARD_TEXT.
func SetClipboardString(data string) error {
	return SetClipboardData([]byte(data), CLIPBOARD_TEXT)
}
