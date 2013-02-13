// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

#include "_cgo_export.h"

WORD macro_MAKELANGID(USHORT usPrimaryLanguage, USHORT usSubLanguage) {
    return MAKELANGID(usPrimaryLanguage, usSubLanguage);
}

void SetLPTSTRAtIndex(LPTSTR array, int index, WORD v) {
    array[index] = v;
}

WORD LPTSTRAtIndex(LPTSTR array, int index) {
    return array[index];
}

LRESULT CALLBACK WndProcWrapper(HWND hwnd, UINT msg, WPARAM wParam, LPARAM lParam) {
    return WndProcCallback(hwnd, msg, wParam, lParam);
}

WNDPROC WndProcWrapperHandle = WndProcWrapper;

DWORD DEVMODE_dmDisplayFixedOutput(DEVMODE* dm) {
	return dm->dmDisplayFixedOutput;
}

