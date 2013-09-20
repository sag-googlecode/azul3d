// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

// Package event implements event-based communication management.
//
// It is important to understand the event dispatch model used. There are both
// event handlers (callback functions) and event notifiers (channels).
//
// When an event is sent using the Send() function, the event handlers are each
// invoked in a random order (you cannot depend on the order in which they are
// executed). In addition, the event is sent to each event notifier.
//
// The Send() function does not return until all event handlers have been
// executed and the event sent to all event notifiers. This defines the
// importance of both event handlers and notifiers.
//
// An event handler should be used when you need to complete something before
// the event has finished sending. For example, in the case of Azul3D's
// "pre-frame" event, an event handler can be used to ensure code executes
// strictly *before* a frame is rendered.
//
// An event notifier should be used when you do not need to complete something
// strictly before the event has finished sending. For example, in the case of
// Azul3D's "pre-frame" event, an event notifier may receive the event before
// or *after* the frame has actually been rendered.
package event

import (
	"math/rand"
	"sync"
	"time"
)

var (
	access    sync.RWMutex
	handlers  = make(map[interface{}][]*func(*Event), 32)
	notifiers = make(map[interface{}][]chan *Event, 32)
)

// Send sends the specified named event and it's associated data to all
// registered event handler functions and event notifier channels.
//
// The eventName parameter is of type interface{}, not string, so that event
// names can be anything (event a unique pointer, for instance).
//
// Data is the data that will be stored in the created Event struct. It is also
// of type interface{}, and may be any associated data with the named event.
func Send(eventName, data interface{}) {
	access.RLock()
	defer access.RUnlock()

	ev := &Event{
		Name: eventName,
		Data: data,
		Time: time.Now(),
	}

	h, ok := handlers[eventName]
	if ok {
		cpy := make([]*func(*Event), len(h))
		copy(cpy, h)

		access.RUnlock()

		p := rand.Perm(len(cpy))
		for _, index := range p {
			fn := *cpy[index]
			fn(ev)
		}

		access.RLock()
	}

	n, ok := notifiers[eventName]
	if ok {
		cpy := make([]chan *Event, len(n))
		copy(cpy, n)

		access.RUnlock()

		p := rand.Perm(len(cpy))
		for _, index := range p {
			ch := cpy[index]
			ch <- ev
		}

		access.RLock()
	}
}

// Handle registers and event handler function to be called when an event with
// the specific event name is sent via the Send() function.
//
// The returned stop() function will unregister the event handler, causing it
// to stop being called.
//
// It is important to unregister event handlers by calling stop() when you are
// done using them, or else you are leaking memory (an internal reference to
// the function is held).
func Handle(eventName interface{}, handler func(*Event)) (stop func()) {
	access.Lock()
	defer access.Unlock()

	handlerPtr := &handler

	h := handlers[eventName]
	for _, fn := range handlers[eventName] {
		if fn == handlerPtr {
			return
		}
	}
	h = append(h, handlerPtr)
	handlers[eventName] = h

	stopped := false
	return func() {
		access.Lock()
		defer access.Unlock()

		if !stopped {
			stopped = true
			var (
				index int
				fn    *func(*Event)
			)
			h := handlers[eventName]
			for index, fn = range h {
				if fn == handlerPtr {
					break
				}
			}
			h = append(h[:index], h[index+1:]...)
			handlers[eventName] = h
		}
	}
}

// Define defines several event handler functions and returns a single stop()
// function which will unregister all of the handlers.
//
// It is important to unregister event handlers by calling stop() when you are
// done using them, or else you are leaking memory (an internal reference to
// the function is held).
//
// This is a simple wrapper around the Handle() function for convenience.
func Define(h Handlers) (stop func()) {
	if len(h) == 0 {
		return
	}
	stoppers := make([]func(), len(h))
	i := 0
	for eventName, handler := range h {
		stoppers[i] = Handle(eventName, handler)
		i++
	}

	return func() {
		for _, stop := range stoppers {
			stop()
		}
	}
}

// NotifyBuffer returns an event notification channel for the specified event
// name using the specified buffer size for the channel.
//
// The returned function unregisters the returned event notification channel so
// that it will stop receiving event notifications.
//
// It is important to unregister event notifiers by calling stop() when you are
// done using them, or else you are leaking memory (an internal reference to
// the channel is held).
//
// If the buffer size is less than zero, a panic will occur.
func NotifyBuffer(eventName interface{}, buffer int) (ch chan *Event, stop func()) {
	if buffer < 0 {
		panic("NotifyBuffer(): Cannot create event channel with buffer size < 0 !")
	}

	ch = make(chan *Event, buffer)

	access.Lock()
	defer access.Unlock()

	n := notifiers[eventName]
	n = append(n, ch)
	notifiers[eventName] = n

	stopped := false
	return ch, func() {
		access.Lock()
		defer access.Unlock()

		if !stopped {
			stopped = true
			var (
				index int
				noti  chan *Event
			)
			n := notifiers[eventName]
			for index, noti = range n {
				if noti == ch {
					break
				}
			}
			n = append(n[:index], n[index+1:]...)
			notifiers[eventName] = n
		}
	}
}

// Notify is an convenience function for:
//
//  return NotifyBuffer(eventName, 32)
//
func Notify(eventName interface{}) (ch chan *Event, stop func()) {
	return NotifyBuffer(eventName, 32)
}
