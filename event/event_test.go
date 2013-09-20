package event

import (
	"testing"
)

var (
	ok, secondFail bool
)

func someEventHandler(ev *Event) {
	if ev.Name.(string) == "someEvent" || ev.Data.(int) == 5 {
		if ok {
			secondFail = true
		}
		ok = true
	}
}

func TestEvent(t *testing.T) {
	stop := Handle("someEvent", someEventHandler)
	Send("someEvent", 5)
	stop()
	Send("someEvent", 5)

	if !ok {
		t.Fail()
	}
	if secondFail {
		t.Log("Got second event despite previous Stop().")
		t.Fail()
	}
}
