package event

import (
	"fmt"
	"time"
)

// Handlers defines a map of event names and associated event handlers, used
// for convenience functions.
//
// For example:
//
//  event.Handlers{
//      "eventName": func(ev *event.Event) {
//          ...closure...
//      },
//      "eventName2": func(ev *event.Event) {
//          ...closure...
//      },
//      "eventName3": someHandlerFunction,
//  }
//
type Handlers map[interface{}]func(ev *Event)

// Event is the genericly used event structure. When an event is sent using
// Send(eventName, data), event handlers and notifiers will recieve one of
// these with the specified event name, and data.
//
// The time is automatically filled out at the time Send() is called.
type Event struct {
	Name, Data interface{}
	Time       time.Time
}

// String returns a string representation of this event.
func (e *Event) String() string {
	return fmt.Sprintf("Event(Name=%v, Data=%v, Time=%v)", e.Name, e.Data, e.Time)
}

// UniqueName can represent a unique named event (as it is a pointer type).
type UniqueName struct {
	name string
}

// String returns the event name as passed in at the call to Unique()
func (u *UniqueName) String() string {
	return u.name
}

// Unique returns an intialized *UniqueName struct given the specified event
// name.
//
// You could for instance define an unique event as a public variable of a
// package, and send the event like so:
//
//  var MyEvent = event.Unique("someNameForPrinting")
//  ...
//  Send(mypackage.MyEvent, data)
//
// Since it's type is a pointer, it cannot collide with any other events.
func Unique(name string) *UniqueName {
	return &UniqueName{
		name: name,
	}
}
