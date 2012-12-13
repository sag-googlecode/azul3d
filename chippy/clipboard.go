package chippy

// SetClipboard sets the contents of the clipboard, in most modern applications,
// this occurs when the user right clicks something and hits copy, or presses the
// keyboard shortcut Control + C to copy text to the clipboard.
//
// This is any valid UTF-8 Go string, but is restricted to the font the user is using
// on their computer, which may be incapable of displaying said UTF-8 string.
func SetClipboard(contents string) {
    setClipboard(contents)
}

// Clipboard returns the contents of the clipboard, in most modern applications,
// this occurs when the user right clicks something and hits paste, or presses the
// keyboard shortcut Control + V to paste text into another application
//
// This will be a valid UTF-8 Go string
func Clipboard() (string, error) {
    return clipboard()
}


