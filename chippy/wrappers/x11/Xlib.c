#include "_cgo_export.h"

void init() {
    XSetErrorHandler((XErrorHandler)(fakeXErrorHandler));
}

