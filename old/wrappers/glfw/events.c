#include "_cgo_export.h"

// ------------------------------------------------------------------------- //
// Error callback
// ------------------------------------------------------------------------- //
void error_cdoCallback(int error, char* description) {
    error_doCallback(error, description);
}

void error_initCallback(void) {
    glfwSetErrorCallback((GLFWerrorfun)error_cdoCallback);
}

// ------------------------------------------------------------------------- //
// Size callback
// ------------------------------------------------------------------------- //
void size_cdoCallback(GLFWwindow w, int width, int height) {
    size_doCallback(w, width, height);
}

void size_initCallback(GLFWwindow w) {
    glfwSetWindowSizeCallback(w, (GLFWwindowsizefun)size_cdoCallback);
}

// ------------------------------------------------------------------------- //
// Close callback
// ------------------------------------------------------------------------- //
int close_cdoCallback(GLFWwindow w) {
    return close_doCallback(w);
}

void close_initCallback(GLFWwindow w) {
    glfwSetWindowCloseCallback(w, (GLFWwindowclosefun)close_cdoCallback);
}

// ------------------------------------------------------------------------- //
// Refresh callback
// ------------------------------------------------------------------------- //
void refresh_cdoCallback(GLFWwindow w) {
    refresh_doCallback(w);
}

void refresh_initCallback(GLFWwindow w) {
    glfwSetWindowRefreshCallback(w, (GLFWwindowrefreshfun)refresh_cdoCallback);
}


// ------------------------------------------------------------------------- //
// Focus callback
// ------------------------------------------------------------------------- //
void focus_cdoCallback(GLFWwindow w, int activated) {
    focus_doCallback(w, activated);
}

void focus_initCallback(GLFWwindow w) {
    glfwSetWindowFocusCallback(w, (GLFWwindowfocusfun)focus_cdoCallback);
}


// ------------------------------------------------------------------------- //
// Iconify callback
// ------------------------------------------------------------------------- //
void iconify_cdoCallback(GLFWwindow w, int iconified) {
    iconify_doCallback(w, iconified);
}

void iconify_initCallback(GLFWwindow w) {
    glfwSetWindowIconifyCallback(w, (GLFWwindowiconifyfun)iconify_cdoCallback);
}


// ------------------------------------------------------------------------- //
// key callback
// ------------------------------------------------------------------------- //
void key_cdoCallback(GLFWwindow w, int key, int action) {
    key_doCallback(w, key, action);
}

void key_initCallback(GLFWwindow w) {
    glfwSetKeyCallback(w, (GLFWkeyfun)key_cdoCallback);
}


// ------------------------------------------------------------------------- //
// char callback
// ------------------------------------------------------------------------- //
void char_cdoCallback(GLFWwindow w, int character) {
    char_doCallback(w, character);
}

void char_initCallback(GLFWwindow w) {
    glfwSetCharCallback(w, (GLFWcharfun)char_cdoCallback);
}


// ------------------------------------------------------------------------- //
// mouseButton callback
// ------------------------------------------------------------------------- //
void mouseButton_cdoCallback(GLFWwindow w, int button, int action) {
    mouseButton_doCallback(w, button, action);
}

void mouseButton_initCallback(GLFWwindow w) {
    glfwSetMouseButtonCallback(w, (GLFWmousebuttonfun)mouseButton_cdoCallback);
}


// ------------------------------------------------------------------------- //
// cursorPos callback
// ------------------------------------------------------------------------- //
void cursorPos_cdoCallback(GLFWwindow w, int x, int y) {
    cursorPos_doCallback(w, x, y);
}

void cursorPos_initCallback(GLFWwindow w) {
    glfwSetCursorPosCallback(w, (GLFWcursorposfun)cursorPos_cdoCallback);
}


// ------------------------------------------------------------------------- //
// cursorEnter callback
// ------------------------------------------------------------------------- //
void cursorEnter_cdoCallback(GLFWwindow w, int entered) {
    cursorEnter_doCallback(w, entered);
}

void cursorEnter_initCallback(GLFWwindow w) {
    glfwSetCursorEnterCallback(w, (GLFWcursorenterfun)cursorEnter_cdoCallback);
}


// ------------------------------------------------------------------------- //
// scroll callback
// ------------------------------------------------------------------------- //
void scroll_cdoCallback(GLFWwindow w, double x, double y) {
    scroll_doCallback(w, x, y);
}

void scroll_initCallback(GLFWwindow w) {
    glfwSetScrollCallback(w, (GLFWscrollfun)scroll_cdoCallback);
}

