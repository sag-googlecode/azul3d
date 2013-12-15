// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package al

import (
	"syscall"
	"unsafe"
)

var (
	lib *syscall.DLL
)

func loadLibrary(name string) error {
	var err error
	lib, err = syscall.LoadDLL(name)
	if err != nil {
		return err
	}
	return nil
}

func symbol(name string) unsafe.Pointer {
	proc := lib.MustFindProc(name)
	return unsafe.Pointer(proc.Addr())
}
