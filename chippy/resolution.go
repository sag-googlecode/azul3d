package chippy

import(
    "fmt"
    "sync"
)

// Resolution represents a Screen's Resolution
type Resolution struct {
    backend_resolution
	width, height uint16
	screen        *Screen
	access        sync.RWMutex
}

// String returns a string representation of this Resolution
func (r *Resolution) String() string {
	r.access.RLock()
	defer r.access.RUnlock()
	return fmt.Sprintf("Resolution(width=%d, height=%d)", r.width, r.height)
}

// Use simply calls r.Screen().SetResolution(r) for you, where r is this Resolution, thus setting this resolution
// to be the active resolution on the Screen.
//
// (See SetResolution() for more information about changing Screen resolutions)
func (r *Resolution) Use() error {
	return r.Screen().SetResolution(r)
}

// Screen returns the Screen that this resolution came from and is attatched to.
func (r *Resolution) Screen() *Screen {
	r.access.RLock()
	defer r.access.RUnlock()
	return r.screen
}

// Width returns the width in pixels of this Resolution
func (r *Resolution) Width() uint16 {
	r.access.RLock()
	defer r.access.RUnlock()
	return r.width
}

// Height returns the height in pixels of this Resolution
func (r *Resolution) Height() uint16 {
	r.access.RLock()
	defer r.access.RUnlock()
	return r.height
}

