package chippy

// The following constants are event buffer sizes
const(
    eventBufferLarge = 32  // These events happen alot, higher buffer is lower latency at a small memory cost
    eventBufferSmall = 8   // These events happen rarely, so we can use smaller buffers
    eventBufferTiny  = 2   // These events are mostly toggle's, so we can use tiny buffers
)

// Platform specific configurations exist here. Even though these are platform specific,
// they're to be exposed to the end user, and because of this are available on all platforms
// e.g. SetLinuxDisplay exists even on Windows, and just does absolutely nil.
var linuxDisplayName string

var cleanupConfig = &callback{func() {
    chippyAccess.Lock()
    defer chippyAccess.Unlock()
	linuxDisplayName = ""
}}

// This will be called before the backend is initialized
func initConfig() {
	addDestroyCallback(cleanupConfig)
}

// SetLinuxDisplayName sets the display_name string that will be passed into
// XOpenDisplay (See http://tronche.com/gui/x/xlib/display/opening.html) on
// Linux operating systems.
// After a call to Destroy() this will be reset to an empty string
//
// The string is similar to the DISPLAY environment variable that X11 uses on POSIX compliant systems.
func SetLinuxDisplayName(displayName string) error {
	chippyAccess.Lock()
	defer chippyAccess.Unlock()
	err := getInitError()
	if err != nil {
		return err
	}

	linuxDisplayName = displayName
	return nil
}

// LinuxDisplayName returns the string previously set by SetLinuxDisplayName
func LinuxDisplayName() (string, error) {
	chippyAccess.Lock()
	defer chippyAccess.Unlock()
	err := getInitError()
	if err != nil {
		return "", err
	}

	return linuxDisplayName, nil
}
