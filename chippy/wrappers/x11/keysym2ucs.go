// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.
// +build linux

// Do not use - Wrappers around very specific x11 APIs.
package x11

/*
#include <stdlib.h>
#include <xcb/xcb.h>
#include "keysym2ucs.h"

#cgo LDFLAGS: -lxcb
*/
import "C"

func Keysym2ucs(keysym Keysym) rune {
	return rune(C.keysym2ucs(C.xcb_keysym_t(keysym)))
}
