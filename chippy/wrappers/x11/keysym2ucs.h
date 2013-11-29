/*
 * This module converts keysym values into the corresponding ISO 10646-1
 * (UCS, Unicode) values.
 */

#include <xcb/xcb.h>

long keysym2ucs(xcb_keysym_t keysym);
