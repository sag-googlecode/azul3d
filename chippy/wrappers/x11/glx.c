#include <X11/Xlib-xcb.h>
#include <GL/glx.h>

#include "_cgo_export.h"

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
