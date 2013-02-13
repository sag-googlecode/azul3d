// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package x11

/*
#include <stdlib.h>
#include <X11/Xlib.h>
#include "keysym2ucs.h"

#cgo LDFLAGS: -lX11
*/
import "C"

func Keysym2ucs(keysym KeySym) int {
	return int(C.keysym2ucs(C.KeySym(keysym)))
}
