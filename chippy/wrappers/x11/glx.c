#include <X11/Xlib-xcb.h>
#include <GL/glx.h>

#include "_cgo_export.h"

// See the following:
//
// http://wrl.illest.net/post/45342765813/code-tip-glx-and-xcbownseventqueue
int chippy_handle_glx_secret_event(Display* d, xcb_generic_event_t* ev) {
	unsigned int rt = ev->response_type & ~0x80;
	int (*setWireToEventProc)(Display*, XEvent*, xEvent*);

	XLockDisplay(d);

	setWireToEventProc = XESetWireToEvent(d, rt, 0);
	if(setWireToEventProc) {
		XESetWireToEvent(d, rt, setWireToEventProc);
		XEvent xev;
		ev->sequence = LastKnownRequestProcessed(d);
		setWireToEventProc(d, &xev, (xEvent*)ev);
	}

	XUnlockDisplay(d);

	if(setWireToEventProc) {
		return 1;
	}
	return 0;
}

typedef GLXContext (*chippy_p_glXCreateContextAttribsARB) (Display* dpy, GLXFBConfig config, GLXContext share, Bool direct, const int* attribs);
GLXContext chippy_glXCreateContextAttribsARB(void* p, Display* dpy, GLXFBConfig config, GLXContext share, Bool direct, const int* attribs) {
	chippy_p_glXCreateContextAttribsARB fn = (chippy_p_glXCreateContextAttribsARB)p;
	return fn(dpy, config, share, direct, attribs);
}

typedef void (*chippy_p_glXSwapIntervalEXT) (Display* dpy, GLXDrawable d, int interval);
void chippy_glXSwapIntervalEXT(void* p, Display* dpy, GLXDrawable d, int interval) {
	chippy_p_glXSwapIntervalEXT fn = (chippy_p_glXSwapIntervalEXT)p;
	fn(dpy, d, interval);
}

typedef int (*chippy_p_glXSwapIntervalMESA) (int interval);
int chippy_glXSwapIntervalMESA(void* p, int interval) {
	chippy_p_glXSwapIntervalMESA fn = (chippy_p_glXSwapIntervalMESA)p;
	return fn(interval);
}

typedef int (*chippy_p_glXSwapIntervalSGI) (int interval);
int chippy_glXSwapIntervalSGI(void* p, int interval) {
	chippy_p_glXSwapIntervalSGI fn = (chippy_p_glXSwapIntervalSGI)p;
	return fn(interval);
}
