// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

#include "_cgo_export.h"

void init() {
    XSetErrorHandler((XErrorHandler)(fakeXErrorHandler));
}

