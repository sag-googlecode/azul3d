package chippy

import (
	"code.google.com/p/azul3d/chippy/keyboard"
	"code.google.com/p/azul3d/chippy/mouse"
	"reflect"
	"sync"
)

// One-way channel with unlimited buffering
func elasticBuffer(write, read interface{}) (lengthQuery chan int) {
	buffer := []reflect.Value{}

	lengthQuery = make(chan int)
	valueLengthQuery := reflect.ValueOf(lengthQuery)
	valueIn := reflect.ValueOf(write)
	valueOut := reflect.ValueOf(read)
	go func() {
		for {
			if len(buffer) > 0 {
				// select{
				//     case v := <-valueIn:
				//         buffer = append(buffer, v)
				//
				//     case valueOut <- buffer[0]:
				//         buffer = buffer[1:len(buffer)]
				//
				//     case valueLengthQuery <- len(buffer):
				//         continue
				// }

				cases := []reflect.SelectCase{
					reflect.SelectCase{
						Dir:  reflect.SelectRecv,
						Chan: valueIn,
					},

					reflect.SelectCase{
						Dir:  reflect.SelectSend,
						Chan: valueOut,
						Send: buffer[0],
					},

					reflect.SelectCase{
						Dir:  reflect.SelectSend,
						Chan: valueLengthQuery,
						Send: reflect.ValueOf(len(buffer)),
					},
				}

				chosen, recv, recvOk := reflect.Select(cases)
				if chosen == 0 {
					if !recvOk {
						return
					}
					buffer = append(buffer, recv)
				} else if chosen == 1 {
					buffer = buffer[1:len(buffer)]
				}

			} else {
				// select{
				//     case v := <- valueIn:
				//         buffer = append(buffer, v)
				//     case valueLengthQuery <- len(buffer):
				//         continue
				// }

				cases := []reflect.SelectCase{
					reflect.SelectCase{
						Dir:  reflect.SelectRecv,
						Chan: valueIn,
					},

					reflect.SelectCase{
						Dir:  reflect.SelectSend,
						Chan: valueLengthQuery,
						Send: reflect.ValueOf(len(buffer)),
					},
				}

				chosen, recv, recvOk := reflect.Select(cases)
				if chosen == 0 {
					if !recvOk {
						return
					}
					buffer = append(buffer, recv)
				} else if chosen == 1 {
					continue
				}
			}
		}
	}()
	return
}

type PaintEventBuffer struct {
	Read, write chan bool
	lengthQuery chan int
}

func (b *PaintEventBuffer) Length() int {
	return <-b.lengthQuery
}


type CloseEventBuffer struct {
	Read, write chan bool
	lengthQuery chan int
}

func (b *CloseEventBuffer) Length() int {
	return <-b.lengthQuery
}


type CursorPositionEventBuffer struct {
	Read, write chan []float64
	lengthQuery chan int
}

func (b *CursorPositionEventBuffer) Length() int {
	return <-b.lengthQuery
}


type CursorWithinEventBuffer struct {
	Read, write chan bool
	lengthQuery chan int
}

func (b *CursorWithinEventBuffer) Length() int {
	return <-b.lengthQuery
}


type KeyboardStateEventBuffer struct {
	Read, write chan *keyboard.StateEvent
	lengthQuery chan int
}

func (b *KeyboardStateEventBuffer) Length() int {
	return <-b.lengthQuery
}


type KeyboardTypedEventBuffer struct {
	Read, write chan *keyboard.TypedEvent
	lengthQuery chan int
}

func (b *KeyboardTypedEventBuffer) Length() int {
	return <-b.lengthQuery
}


type ResizeEventBuffer struct {
	Read, write chan []uint
	lengthQuery chan int
}

func (b *ResizeEventBuffer) Length() int {
	return <-b.lengthQuery
}


type MaximizedEventBuffer struct {
	Read, write chan bool
	lengthQuery chan int
}

func (b *MaximizedEventBuffer) Length() int {
	return <-b.lengthQuery
}


type MinimizedEventBuffer struct {
	Read, write chan bool
	lengthQuery chan int
}

func (b *MinimizedEventBuffer) Length() int {
	return <-b.lengthQuery
}


type MouseEventBuffer struct {
	Read, write chan *mouse.Event
	lengthQuery chan int
}

func (b *MouseEventBuffer) Length() int {
	return <-b.lengthQuery
}


type FocusedEventBuffer struct {
	Read, write chan bool
	lengthQuery chan int
}

func (b *FocusedEventBuffer) Length() int {
	return <-b.lengthQuery
}


type PositionEventBuffer struct {
	Read, write chan []int
	lengthQuery chan int
}

func (b *PositionEventBuffer) Length() int {
	return <-b.lengthQuery
}


type SizeEventBuffer struct {
	Read, write chan []uint
	lengthQuery chan int
}

func (b *SizeEventBuffer) Length() int {
	return <-b.lengthQuery
}


type ScreenChangedEventBuffer struct {
	Read, write chan Screen
	lengthQuery chan int
}

func (b *ScreenChangedEventBuffer) Length() int {
	return <-b.lengthQuery
}

type eventDispatcher struct {
	eventsAccess sync.RWMutex
	paintEvents          []*PaintEventBuffer
	closeEvents          []*CloseEventBuffer
	cursorPositionEvents []*CursorPositionEventBuffer
	cursorWithinEvents   []*CursorWithinEventBuffer
	keyboardStateEvents  []*KeyboardStateEventBuffer
	keyboardTypedEvents  []*KeyboardTypedEventBuffer
	maximizedEvents      []*MaximizedEventBuffer
	minimizedEvents      []*MinimizedEventBuffer
	mouseEvents          []*MouseEventBuffer
	focusedEvents        []*FocusedEventBuffer
	positionEvents       []*PositionEventBuffer
	sizeEvents           []*SizeEventBuffer
	screenChangedEvents  []*ScreenChangedEventBuffer
}

func (e *eventDispatcher) PaintEvents() *PaintEventBuffer {
	e.eventsAccess.Lock()
	defer e.eventsAccess.Unlock()

	buf := new(PaintEventBuffer)
	buf.Read = make(chan bool)
	buf.write = make(chan bool)
	buf.lengthQuery = elasticBuffer(buf.write, buf.Read)
	e.paintEvents = append(e.paintEvents, buf)
	return buf
}
func (e *eventDispatcher) addPaintEvent() bool {
	e.eventsAccess.RLock()
	defer e.eventsAccess.RUnlock()

	if len(e.paintEvents) == 0 {
		return false
	}

	for _, buf := range e.paintEvents {
		buf.write <- true
	}
	return true
}

func (e *eventDispatcher) CloseEvents() *CloseEventBuffer {
	e.eventsAccess.Lock()
	defer e.eventsAccess.Unlock()

	buf := new(CloseEventBuffer)
	buf.Read = make(chan bool)
	buf.write = make(chan bool)
	buf.lengthQuery = elasticBuffer(buf.write, buf.Read)
	e.closeEvents = append(e.closeEvents, buf)
	return buf
}
func (e *eventDispatcher) addCloseEvent() bool {
	e.eventsAccess.RLock()
	defer e.eventsAccess.RUnlock()

	if len(e.closeEvents) == 0 {
		return false
	}

	for _, buf := range e.closeEvents {
		buf.write <- true
	}
	return true
}

func (e *eventDispatcher) CursorPositionEvents() *CursorPositionEventBuffer {
	e.eventsAccess.Lock()
	defer e.eventsAccess.Unlock()

	buf := new(CursorPositionEventBuffer)
	buf.Read = make(chan []float64)
	buf.write = make(chan []float64)
	buf.lengthQuery = elasticBuffer(buf.write, buf.Read)
	e.cursorPositionEvents = append(e.cursorPositionEvents, buf)
	return buf
}
func (e *eventDispatcher) addCursorPositionEvent(v []float64) {
	e.eventsAccess.RLock()
	defer e.eventsAccess.RUnlock()

	for _, buf := range e.cursorPositionEvents {
		buf.write <- v
	}
}

func (e *eventDispatcher) CursorWithinEvents() *CursorWithinEventBuffer {
	e.eventsAccess.Lock()
	defer e.eventsAccess.Unlock()

	buf := new(CursorWithinEventBuffer)
	buf.Read = make(chan bool)
	buf.write = make(chan bool)
	buf.lengthQuery = elasticBuffer(buf.write, buf.Read)
	e.cursorWithinEvents = append(e.cursorWithinEvents, buf)
	return buf
}
func (e *eventDispatcher) addCursorWithinEvent(v bool) {
	e.eventsAccess.RLock()
	defer e.eventsAccess.RUnlock()

	for _, buf := range e.cursorWithinEvents {
		buf.write <- v
	}
}

func (e *eventDispatcher) KeyboardStateEvents() *KeyboardStateEventBuffer {
	e.eventsAccess.Lock()
	defer e.eventsAccess.Unlock()

	buf := new(KeyboardStateEventBuffer)
	buf.Read = make(chan *keyboard.StateEvent)
	buf.write = make(chan *keyboard.StateEvent)
	buf.lengthQuery = elasticBuffer(buf.write, buf.Read)
	e.keyboardStateEvents = append(e.keyboardStateEvents, buf)
	return buf
}
func (e *eventDispatcher) addKeyboardStateEvent(v *keyboard.StateEvent) {
	e.eventsAccess.RLock()
	defer e.eventsAccess.RUnlock()

	for _, buf := range e.keyboardStateEvents {
		buf.write <- v
	}
}

func (e *eventDispatcher) KeyboardTypedEvents() *KeyboardTypedEventBuffer {
	e.eventsAccess.Lock()
	defer e.eventsAccess.Unlock()

	buf := new(KeyboardTypedEventBuffer)
	buf.Read = make(chan *keyboard.TypedEvent)
	buf.write = make(chan *keyboard.TypedEvent)
	buf.lengthQuery = elasticBuffer(buf.write, buf.Read)
	e.keyboardTypedEvents = append(e.keyboardTypedEvents, buf)
	return buf
}
func (e *eventDispatcher) addKeyboardTypedEvent(v *keyboard.TypedEvent) {
	e.eventsAccess.RLock()
	defer e.eventsAccess.RUnlock()

	for _, buf := range e.keyboardTypedEvents {
		buf.write <- v
	}
}

func (e *eventDispatcher) MaximizedEvents() *MaximizedEventBuffer {
	e.eventsAccess.Lock()
	defer e.eventsAccess.Unlock()

	buf := new(MaximizedEventBuffer)
	buf.Read = make(chan bool)
	buf.write = make(chan bool)
	buf.lengthQuery = elasticBuffer(buf.write, buf.Read)
	e.maximizedEvents = append(e.maximizedEvents, buf)
	return buf
}
func (e *eventDispatcher) addMaximizedEvent(v bool) {
	e.eventsAccess.RLock()
	defer e.eventsAccess.RUnlock()

	for _, buf := range e.maximizedEvents {
		buf.write <- v
	}
}

func (e *eventDispatcher) MinimizedEvents() *MinimizedEventBuffer {
	e.eventsAccess.Lock()
	defer e.eventsAccess.Unlock()

	buf := new(MinimizedEventBuffer)
	buf.Read = make(chan bool)
	buf.write = make(chan bool)
	buf.lengthQuery = elasticBuffer(buf.write, buf.Read)
	e.minimizedEvents = append(e.minimizedEvents, buf)
	return buf
}
func (e *eventDispatcher) addMinimizedEvent(v bool) {
	e.eventsAccess.RLock()
	defer e.eventsAccess.RUnlock()

	for _, buf := range e.minimizedEvents {
		buf.write <- v
	}
}

func (e *eventDispatcher) MouseEvents() *MouseEventBuffer {
	e.eventsAccess.Lock()
	defer e.eventsAccess.Unlock()

	buf := new(MouseEventBuffer)
	buf.Read = make(chan *mouse.Event)
	buf.write = make(chan *mouse.Event)
	buf.lengthQuery = elasticBuffer(buf.write, buf.Read)
	e.mouseEvents = append(e.mouseEvents, buf)
	return buf
}
func (e *eventDispatcher) addMouseEvent(v *mouse.Event) {
	e.eventsAccess.RLock()
	defer e.eventsAccess.RUnlock()

	for _, buf := range e.mouseEvents {
		buf.write <- v
	}
}

func (e *eventDispatcher) FocusedEvents() *FocusedEventBuffer {
	e.eventsAccess.Lock()
	defer e.eventsAccess.Unlock()

	buf := new(FocusedEventBuffer)
	buf.Read = make(chan bool)
	buf.write = make(chan bool)
	buf.lengthQuery = elasticBuffer(buf.write, buf.Read)
	e.focusedEvents = append(e.focusedEvents, buf)
	return buf
}
func (e *eventDispatcher) addFocusedEvent(v bool) {
	e.eventsAccess.RLock()
	defer e.eventsAccess.RUnlock()

	for _, buf := range e.focusedEvents {
		buf.write <- v
	}
}

func (e *eventDispatcher) PositionEvents() *PositionEventBuffer {
	e.eventsAccess.Lock()
	defer e.eventsAccess.Unlock()

	buf := new(PositionEventBuffer)
	buf.Read = make(chan []int)
	buf.write = make(chan []int)
	buf.lengthQuery = elasticBuffer(buf.write, buf.Read)
	e.positionEvents = append(e.positionEvents, buf)
	return buf
}
func (e *eventDispatcher) addPositionEvent(v []int) {
	e.eventsAccess.RLock()
	defer e.eventsAccess.RUnlock()

	for _, buf := range e.positionEvents {
		buf.write <- v
	}
}

func (e *eventDispatcher) SizeEvents() *SizeEventBuffer {
	e.eventsAccess.Lock()
	defer e.eventsAccess.Unlock()

	buf := new(SizeEventBuffer)
	buf.Read = make(chan []uint)
	buf.write = make(chan []uint)
	buf.lengthQuery = elasticBuffer(buf.write, buf.Read)
	e.sizeEvents = append(e.sizeEvents, buf)
	return buf
}
func (e *eventDispatcher) addSizeEvent(v []uint) {
	e.eventsAccess.RLock()
	defer e.eventsAccess.RUnlock()

	for _, buf := range e.sizeEvents {
		buf.write <- v
	}
}

func (e *eventDispatcher) ScreenChangedEvents() *ScreenChangedEventBuffer {
	e.eventsAccess.Lock()
	defer e.eventsAccess.Unlock()

	buf := new(ScreenChangedEventBuffer)
	buf.Read = make(chan Screen)
	buf.write = make(chan Screen)
	buf.lengthQuery = elasticBuffer(buf.write, buf.Read)
	e.screenChangedEvents = append(e.screenChangedEvents, buf)
	return buf
}
func (e *eventDispatcher) addScreenChangedEvent(v Screen) {
	e.eventsAccess.RLock()
	defer e.eventsAccess.RUnlock()

	for _, buf := range e.screenChangedEvents {
		buf.write <- v
	}
}

