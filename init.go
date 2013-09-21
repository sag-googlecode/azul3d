// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

// Package Azul3D is a 3D game engine.
//
// This package is a very thin wrapper around the various other sub-packages
// found inside this directory.
//
// Most full applications will probably not use this package, but use it as a
// reference instead.
package azul3d

import (
	"code.google.com/p/azul3d/chippy"
	"code.google.com/p/azul3d/clock"
	"io"
	"io/ioutil"
	"log"
	"sync"
)

var (
	globalLock  sync.RWMutex
	debugOutput io.Writer

	// The default clock stats.
	Stats = clock.NewStats()

	// The Logger used for debug output.
	Logger *log.Logger
)

// SetDebugOutput specifies where debug output of Azul3D should be written to.
func SetDebugOutput(output io.Writer) {
	chippy.SetDebugOutput(output)

	globalLock.Lock()
	defer globalLock.Unlock()

	Logger = log.New(output, "Azul3D: ", 0)
	debugOutput = output
}

// DebugOutput returns where debug output of Azul3D will be written to.
func DebugOutput() io.Writer {
	globalLock.RLock()
	defer globalLock.RUnlock()

	return debugOutput
}

func init() {
	Stats.SetEnabled(false)
	SetDebugOutput(ioutil.Discard)
}

// Init initializes Azul3D, or returns an error if we cannot initialize Azul3D.
func Init() error {
	err := chippy.Init()
	if err != nil {
		return err
	}
	return nil
}

// Exit exits the main loop.
func Exit() {
	chippy.Exit()
}

// MainLoop enters the main loop.
//
// This function *must* be called on the main thread (due to the restrictions
// that some platforms place on us).
//
// It's best to place this function inside either your init or main function.
//
// This function will not return until azul3d.Exit() is called.
//
// If azul3d is not initialized (via an previous call to the Init() function)
// then an panic will occur.)
func MainLoop() {
	chippy.MainLoop()
}

// MainLoopFrames returns an channel of functions which return an boolean
// status as to weather you should continue running the 'main loop'.
//
// Typically you would not use this function and would instead use the
// MainLoop() function.
//
// This is for advanced users where the main loop is required to be shared with
// some other external library. I.e. this allows for communicative main loop
// handling.
//
// See the MainLoop() function source code in chippy/chippy.go for an example of using
// this function properly.
//
// This function should only really be called once (the same channel is always
// returned).
func MainLoopFrames() chan func() bool {
	return chippy.MainLoopFrames()
}
