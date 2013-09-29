// Package aio provides access to audio input and output devices.
package aio

import (
	"io"
	"io/ioutil"
	"log"
	"sync"
)

var (
	theLoggerAccess sync.RWMutex
	theLogger       *log.Logger

	cacheAccess                             sync.RWMutex
	cachedInputs                            []*Input
	cachedOutputs                           []*Output
	cachedDefaultInput, cachedDefaultOutput int
)

func logger() *log.Logger {
	theLoggerAccess.RLock()
	defer theLoggerAccess.RUnlock()

	return theLogger
}

// SetDebugOutput specifies the io.Writer that debug output will be written to
// (ioutil.Discard by default).
func SetDebugOutput(w io.Writer) {
	theLoggerAccess.Lock()
	defer theLoggerAccess.Unlock()

	theLogger = log.New(w, "chippy: ", log.Ltime|log.Lshortfile)
}

// Refresh reloads the cached list of inputs and outputs returned by the
// Inputs() and Outputs() functions.
//
// After a call to this function; Inputs() and Outputs() may return newly
// connected or disconnected input/output devices.
//
// This function is automatically invoked (once) at this package's init.
func Refresh() {
	cacheAccess.Lock()
	defer cacheAccess.Unlock()

	cachedInputs, cachedOutputs, cachedDefaultInput, cachedDefaultOutput = backend_Devices()
}

// DefaultInput returns the default audio input device.
func DefaultInput() *Input {
	cacheAccess.RLock()
	defer cacheAccess.RUnlock()

	if len(cachedInputs) == 0 {
		return nil
	}
	return cachedInputs[cachedDefaultInput]
}

// Inputs returns a copy of the cached list of audio input devices.
func Inputs() []*Input {
	cacheAccess.RLock()
	defer cacheAccess.RUnlock()

	cpy := make([]*Input, len(cachedInputs))
	copy(cpy, cachedInputs)
	return cpy
}

// DefaultOutput returns the default audio output device.
func DefaultOutput() *Output {
	cacheAccess.RLock()
	defer cacheAccess.RUnlock()

	if len(cachedOutputs) == 0 {
		return nil
	}
	return cachedOutputs[cachedDefaultOutput]
}

// Outputs returns a copy of the cached list of audio output devices.
func Outputs() []*Output {
	cacheAccess.RLock()
	defer cacheAccess.RUnlock()

	cpy := make([]*Output, len(cachedOutputs))
	copy(cpy, cachedOutputs)
	return cpy
}

func init() {
	SetDebugOutput(ioutil.Discard)
	Refresh()
}
