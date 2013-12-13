// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package al

import (
	"runtime"
)

// OpenAL is actually thread-safe by specification. But just to take extra
// caution to avoid any buggy OpenAL implentations, we only use a single OS
// thread.
//
// This doesn't really harm performance, because in practice OpenAL provides no
// performance benifit for multi-threaded users (e.g. OpenAL soft can only have
// one context per device, etc).

func init() {
	go dispatcher()
}

var dispatchChan = make(chan func())

func dispatcher() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	for {
		f := <-dispatchChan
		f()
		dispatchChan <- nil
	}
}

func dispatch(f func()) {
	dispatchChan <- f
	<-dispatchChan
}
