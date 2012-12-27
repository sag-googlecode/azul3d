package chippy

// SetClipboard sets the contents of the clipboard to the specified UTF-8 Go string
func SetClipboard(contents string) error {
	chippyAccess.Lock()
	defer chippyAccess.Unlock()
	err := getInitError()
	if err != nil {
		return err
	}

	return backend_setClipboard(contents)
}

// Clipboard returns the contents of the clipboard as an valid UTF-8 Go string
func Clipboard() (string, error) {
	chippyAccess.Lock()
	defer chippyAccess.Unlock()
	err := getInitError()
	if err != nil {
		return "", err
	}

	return backend_clipboard()
}
