// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "_cgo_export.h"

WORD win32_MAKELANGID(USHORT usPrimaryLanguage, USHORT usSubLanguage) {
    return MAKELANGID(usPrimaryLanguage, usSubLanguage);
}

LRESULT CALLBACK win32_WndProcWrapper(HWND hwnd, UINT msg, WPARAM wParam, LPARAM lParam) {
    return WndProcCallback(hwnd, msg, wParam, lParam);
}

WNDPROC win32_WndProcWrapperHandle = win32_WndProcWrapper;

DWORD win32_DEVMODE_dmDisplayFixedOutput(DEVMODE* dm) {
	return dm->dmDisplayFixedOutput;
}

POINTL win32_DEVMODE_dmPosition(DEVMODE* dm) {
	return dm->dmPosition;
}

LPTSTR macro_MAKEINTRESOURCE(WORD wInteger) {
	return MAKEINTRESOURCE(wInteger);
}

MONITORENUMPROC win32_MonitorEnumProcCallbackHandle = MonitorEnumProcCallback;

HOOKPROC win32_LowLevelKeyboardHookCallbackHandle = LowLevelKeyboardHookCallback;

