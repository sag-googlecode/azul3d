// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package chippy

func backend_ClipboardHasData() bool {
	panic("ClipboardHasData")
}

func backend_IsClipboardFormatSupported(format ClipboardFormat) bool {
	panic("IsClipboardFormatSupported")
}

func backend_ClipboardData() ([]byte, ClipboardFormat, error) {
	panic("ClipboardData")
}

func backend_SetClipboardData(data []byte, format ClipboardFormat) error {
	panic("SetClipboardData")
}

func backend_ClearClipboard() error {
	panic("ClearClipboard")
}
