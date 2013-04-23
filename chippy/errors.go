// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package chippy

import (
	"errors"
)

var (
	// Could happen due to lack of hardware support for the screen mode, or the driver may reject
	// the screen mode as well, really this could be generically anything but is typically an
	// hardware or driver issue using the screen mode.
	//
	// You should never take away the screen mode as an option from the user; as it is possible the
	// user may change an configuration setting of some sort with their operating system that will
	// allow the screen mode to be used properly.
	ErrBadScreenMode = errors.New("unable to switch screen mode; hardware or drivers do not support the mode.")

	// Only Microsoft Windows will ever return this error, systems that use DualView will sometimes
	// ignore screen mode change requests.
	ErrDualViewCapable = errors.New("unable to switch screen mode; the system is DualView capable.")

	// If there is no support for gamma ramps due to either lack of hardware or software support
	// for them.
	ErrGammaRampsNotSupported = errors.New("unable to use gamma ramps; the hardware has no support for them.")
)
