// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

#include "_cgo_export.h"

WORD win32_MAKELANGID(USHORT usPrimaryLanguage, USHORT usSubLanguage) {
    return MAKELANGID(usPrimaryLanguage, usSubLanguage);
}

void win32_SetLPTSTRAtIndex(LPTSTR array, int index, WORD v) {
    array[index] = v;
}

WORD win32_LPTSTRAtIndex(LPTSTR array, int index) {
    return array[index];
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


MONITORENUMPROC win32_MonitorEnumProcCallbackHandle = MonitorEnumProcCallback;

