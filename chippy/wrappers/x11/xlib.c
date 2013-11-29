#include <X11/Xlib-xcb.h>

#include "_cgo_export.h"

XIC chippy_CreateIC(XIM xim, Display* d, Window w) {
    XIMStyles *xim_styles;
    XIMStyle xim_style = 0;
    char *imvalret;
	int i;

	imvalret = XGetIMValues(xim, XNQueryInputStyle, &xim_styles, NULL);
	if (imvalret != NULL || xim_styles == NULL) {
		return NULL;
	}

	if (xim_styles) {
		xim_style = 0;
		for (i = 0; i < xim_styles->count_styles; i++) {
			if (xim_styles->supported_styles[i] == (XIMPreeditNothing | XIMStatusNothing)) {
				xim_style = xim_styles->supported_styles[i];
				break;
			}
		}

		if (xim_style == 0) {
			return NULL;
		}
		XFree(xim_styles);
	}

	return XCreateIC(
		xim,
		XNInputStyle, xim_style,
		XNClientWindow, w,
		XNFocusWindow, w,
		NULL
	);
}


int chippy_xlib_error(Display* d, XErrorEvent* e) {
	chippy_xlib_error_callback(d, e);
	return 0;
}

