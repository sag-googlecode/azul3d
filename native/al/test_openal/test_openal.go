// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.
// +build tests

package main

import (
	"code.google.com/p/azul3d/native/al"
	"log"
)

func main() {
	defaultDevice := al.AlcGetString(nil, al.ALC_DEFAULT_DEVICE_SPECIFIER)
	log.Println("Default Device:", defaultDevice)

	defaultCapture := al.AlcGetString(nil, al.ALC_CAPTURE_DEFAULT_DEVICE_SPECIFIER)
	log.Println("Default Capture Device:", defaultCapture)

	allDevices := al.AlcGetString(nil, al.ALC_DEVICE_SPECIFIER)
	log.Println("All Devices:", allDevices)

	captureDevices := al.AlcGetString(nil, al.ALC_CAPTURE_DEVICE_SPECIFIER)
	log.Println("All Capture Devices:", captureDevices)

	extensions := al.AlcGetString(nil, al.ALC_EXTENSIONS)
	log.Println("Extensions:", extensions)

	device, err := al.OpenDevice("", nil)
	if err != nil {
		log.Fatal(err)
	}
	defer device.Close()

	haveCapture := device.AlcIsExtensionPresent("ALC_EXT_CAPTURE")

	var maxSources int32
	device.AlcGetIntegerv(al.ALC_MONO_SOURCES, 1, &maxSources)
	log.Println("Maximum sources:", maxSources)

	if haveCapture {
		log.Println("Have the ALC_EXT_CAPTURE extension.")
		err = device.InitCapture(44100, al.FORMAT_MONO16, 44100/2)
		if err != nil {
			log.Fatal(err)
		}
	}
}