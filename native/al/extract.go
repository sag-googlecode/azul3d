// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package al

/*
#include <inttypes.h>

extern uint8_t openal_data[] asm("_binary_openal_start");
extern uint8_t openal_data_size[] asm("_binary_openal_size");

long get_openal_data(void) {
	return (long)openal_data;
}

long get_openal_data_size(void) {
	return (long)openal_data_size;
}
*/
import "C"

import (
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path"
	"path/filepath"
	"reflect"
	"unsafe"
)

var (
	blob        []byte
	libraryPath string
)

func init() {
	// Initialize blob slice using blob data
	sz := int(C.get_openal_data_size())
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&blob))
	sliceHeader.Data = uintptr(C.get_openal_data())
	sliceHeader.Len = sz
	sliceHeader.Cap = sz

	// Determine where the library should actually be placed on the system
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	libraryPath = filepath.Join(usr.HomeDir, ".openal")
	libraryPath = filepath.Join(libraryPath, blobFileName)

	// Check if the library already exists at that location -- if it does it
	// means we've already extracted it there or the user has placed their own
	// implementation of the library there (per the LGPL restrictions).
	_, err = os.Stat(libraryPath)
	if err != nil {
		err := os.MkdirAll(path.Dir(libraryPath), 0777)
		if err != nil {
			log.Fatal(err)
		}

		// There is no dynamic library at that location, we can extract our
		// copy of it then.
		err = ioutil.WriteFile(libraryPath, blob, 0777)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Load the library
	err = loadLibrary(libraryPath)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize pointers
	alInit()
	alcInit()
}
