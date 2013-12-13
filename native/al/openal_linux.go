// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package al

/*
#include <stdlib.h>
#include <dlfcn.h>
#cgo LDFLAGS: -ldl

*/
import "C"

import (
	"errors"
	"unsafe"
)

var (
	lib unsafe.Pointer
)

func loadLibrary(name string) error {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	lib = C.dlopen(cname, C.RTLD_NOW|C.RTLD_GLOBAL)
	cerr := C.dlerror()
	if lib == nil || cerr != nil {
		err := C.GoString(cerr)
		return errors.New(err)
	}
	return nil
}

func symbol(name string) unsafe.Pointer {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	return C.dlsym(lib, cname)
}
