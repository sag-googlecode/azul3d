#include <X11/Xlib-xcb.h>

#include "_cgo_export.h"

int chippy_xlib_error(Display* d, XErrorEvent* e) {
	chippy_xlib_error_callback(d, e);
	return 0;
}

