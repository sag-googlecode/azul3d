package chippy

import (
	"code.google.com/p/azul3d/chippy/keyboard"
	"code.google.com/p/azul3d/chippy/mouse"
	"reflect"
	"sync"
)

// One-way channel with unlimited buffering
func elasticBuffer(write, read interface{}, onClose func()) (lengthQuery chan int) {
	buffer := []reflect.Value{}

	lengthQuery = make(chan int)
	valueLengthQuery := reflect.ValueOf(lengthQuery)
	valueIn := reflect.ValueOf(write)
	valueOut := reflect.ValueOf(read)
	go func() {
		defer func() {
			onClose()
		}()

		for {
			if len(buffer) > 0 {
				// select{
				//     case <-valueCloseBuffer:
				//         return
				//
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
				//     case <-valueCloseBuffer:
				//         return
				//
				//     case v := <- valueIn:
				//         buffer = append(buffer, v)
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

type DestroyedEventBuffer struct {
	Read, write chan bool
	closeBuffer chan bool
	lengthQuery chan int
}

func (b *DestroyedEventBuffer) Length() int {
	return <-b.lengthQuery
}
func (b *DestroyedEventBuffer) Close() {
	close(b.write)
}

type PaintEventBuffer struct {
	Read, write chan bool
	closeBuffer chan bool
	lengthQuery chan int
}

func (b *PaintEventBuffer) Length() int {
	return <-b.lengthQuery
}
func (b *PaintEventBuffer) Close() {
	close(b.write)
}

type CloseEventBuffer struct {
	Read, write chan bool
	closeBuffer chan bool
	lengthQuery chan int
}

func (b *CloseEventBuffer) Length() int {
	return <-b.lengthQuery
}
func (b *CloseEventBuffer) Close() {
	close(b.write)
}

type CursorPositionEventBuffer struct {
	Read, write chan []float64
	closeBuffer chan bool
	lengthQuery chan int
}

func (b *CursorPositionEventBuffer) Length() int {
	return <-b.lengthQuery
}
func (b *CursorPositionEventBuffer) Close() {
	close(b.write)
}

type CursorWithinEventBuffer struct {
	Read, write chan bool
	closeBuffer chan bool
	lengthQuery chan int
}

func (b *CursorWithinEventBuffer) Length() int {
	return <-b.lengthQuery
}
func (b *CursorWithinEventBuffer) Close() {
	close(b.write)
}

type KeyboardStateEventBuffer struct {
	Read, write chan *keyboard.StateEvent
	closeBuffer chan bool
	lengthQuery chan int
}

func (b *KeyboardStateEventBuffer) Length() int {
	return <-b.lengthQuery
}
func (b *KeyboardStateEventBuffer) Close() {
	close(b.write)
}

type KeyboardTypedEventBuffer struct {
	Read, write chan *keyboard.TypedEvent
	closeBuffer chan bool
	lengthQuery chan int
}

func (b *KeyboardTypedEventBuffer) Length() int {
	return <-b.lengthQuery
}
func (b *KeyboardTypedEventBuffer) Close() {
	close(b.write)
}

type ResizeEventBuffer struct {
	Read, write chan []uint
	closeBuffer chan bool
	lengthQuery chan int
}

func (b *ResizeEventBuffer) Length() int {
	return <-b.lengthQuery
}
func (b *ResizeEventBuffer) Close() {
	close(b.write)
}

type MaximizedEventBuffer struct {
	Read, write chan bool
	closeBuffer chan bool
	lengthQuery chan int
}

func (b *MaximizedEventBuffer) Length() int {
	return <-b.lengthQuery
}
func (b *MaximizedEventBuffer) Close() {
	close(b.write)
}

type MinimizedEventBuffer struct {
	Read, write chan bool
	closeBuffer chan bool
	lengthQuery chan int
}

func (b *MinimizedEventBuffer) Length() int {
	return <-b.lengthQuery
}
func (b *MinimizedEventBuffer) Close() {
	close(b.write)
}

type MouseEventBuffer struct {
	Read, write chan *mouse.Event
	closeBuffer chan bool
	lengthQuery chan int
}

func (b *MouseEventBuffer) Length() int {
	return <-b.lengthQuery
}
func (b *MouseEventBuffer) Close() {
	close(b.write)
}

type FocusedEventBuffer struct {
	Read, write chan bool
	closeBuffer chan bool
	lengthQuery chan int
}

func (b *FocusedEventBuffer) Length() int {
	return <-b.lengthQuery
}
func (b *FocusedEventBuffer) Close() {
	close(b.write)
}

type PositionEventBuffer struct {
	Read, write chan []int
	closeBuffer chan bool
	lengthQuery chan int
}

func (b *PositionEventBuffer) Length() int {
	return <-b.lengthQuery
}
func (b *PositionEventBuffer) Close() {
	close(b.write)
}

type SizeEventBuffer struct {
	Read, write chan []uint
	closeBuffer chan bool
	lengthQuery chan int
}

func (b *SizeEventBuffer) Length() int {
	return <-b.lengthQuery
}
func (b *SizeEventBuffer) Close() {
	close(b.write)
}

type ScreenChangedEventBuffer struct {
	Read, write chan Screen
	closeBuffer chan bool
	lengthQuery chan int
}

func (b *ScreenChangedEventBuffer) Length() int {
	return <-b.lengthQuery
}
func (b *ScreenChangedEventBuffer) Close() {
	close(b.write)
}

type eventDispatcher struct {
	eventsAccess         sync.RWMutex
	destroyedEvents      map[*DestroyedEventBuffer]bool
	paintEvents          map[*PaintEventBuffer]bool
	closeEvents          map[*CloseEventBuffer]bool
	cursorPositionEvents map[*CursorPositionEventBuffer]bool
	cursorWithinEvents   map[*CursorWithinEventBuffer]bool
	keyboardStateEvents  map[*KeyboardStateEventBuffer]bool
	keyboardTypedEvents  map[*KeyboardTypedEventBuffer]bool
	maximizedEvents      map[*MaximizedEventBuffer]bool
	minimizedEvents      map[*MinimizedEventBuffer]bool
	mouseEvents          map[*MouseEventBuffer]bool
	focusedEvents        map[*FocusedEventBuffer]bool
	positionEvents       map[*PositionEventBuffer]bool
	sizeEvents           map[*SizeEventBuffer]bool
	screenChangedEvents  map[*ScreenChangedEventBuffer]bool
}

func (e *eventDispatcher) initEventDispatcher() {
	e.destroyedEvents = make(map[*DestroyedEventBuffer]bool)
	e.paintEvents = make(map[*PaintEventBuffer]bool)
	e.closeEvents = make(map[*CloseEventBuffer]bool)
	e.cursorPositionEvents = make(map[*CursorPositionEventBuffer]bool)
	e.cursorWithinEvents = make(map[*CursorWithinEventBuffer]bool)
	e.keyboardStateEvents = make(map[*KeyboardStateEventBuffer]bool)
	e.keyboardTypedEvents = make(map[*KeyboardTypedEventBuffer]bool)
	e.maximizedEvents = make(map[*MaximizedEventBuffer]bool)
	e.minimizedEvents = make(map[*MinimizedEventBuffer]bool)
	e.mouseEvents = make(map[*MouseEventBuffer]bool)
	e.focusedEvents = make(map[*FocusedEventBuffer]bool)
	e.positionEvents = make(map[*PositionEventBuffer]bool)
	e.sizeEvents = make(map[*SizeEventBuffer]bool)
	e.screenChangedEvents = make(map[*ScreenChangedEventBuffer]bool)
}

func (e *eventDispatcher) DestroyedEvents() *DestroyedEventBuffer {
	e.eventsAccess.Lock()
	defer e.eventsAccess.Unlock()

	buf := new(DestroyedEventBuffer)
	buf.Read = make(chan bool)
	buf.write = make(chan bool)

	buf.lengthQuery = elasticBuffer(buf.write, buf.Read, func() {
		e.eventsAccess.Lock()
		defer e.eventsAccess.Unlock()

		delete(e.destroyedEvents, buf)
	})

	e.destroyedEvents[buf] = true
	return buf
}
func (e *eventDispatcher) sendDestroyedEvent() {
	e.eventsAccess.RLock()
	defer e.eventsAccess.RUnlock()

	for buf, _ := range e.destroyedEvents {
		buf.write <- true
	}
}

func (e *eventDispatcher) PaintEvents() *PaintEventBuffer {
	e.eventsAccess.Lock()
	defer e.eventsAccess.Unlock()

	buf := new(PaintEventBuffer)
	buf.Read = make(chan bool)
	buf.write = make(chan bool)
	buf.lengthQuery = elasticBuffer(buf.write, buf.Read, func() {
		e.eventsAccess.Lock()
		defer e.eventsAccess.Unlock()

		delete(e.paintEvents, buf)
	})
	e.paintEvents[buf] = true
	return buf
}
func (e *eventDispatcher) addPaintEvent() bool {
	e.eventsAccess.RLock()
	defer e.eventsAccess.RUnlock()

	if len(e.paintEvents) == 0 {
		return false
	}

	for buf, _ := range e.paintEvents {
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
	buf.lengthQuery = elasticBuffer(buf.write, buf.Read, func() {
		e.eventsAccess.Lock()
		defer e.eventsAccess.Unlock()

		delete(e.closeEvents, buf)
	})
	e.closeEvents[buf] = true
	return buf
}
func (e *eventDispatcher) addCloseEvent() bool {
	e.eventsAccess.RLock()
	defer e.eventsAccess.RUnlock()

	if len(e.closeEvents) == 0 {
		return false
	}

	for buf, _ := range e.closeEvents {
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
	buf.lengthQuery = elasticBuffer(buf.write, buf.Read, func() {
		e.eventsAccess.Lock()
		defer e.eventsAccess.Unlock()

		delete(e.cursorPositionEvents, buf)
	})
	e.cursorPositionEvents[buf] = true
	return buf
}
func (e *eventDispatcher) addCursorPositionEvent(v []float64) {
	e.eventsAccess.RLock()
	defer e.eventsAccess.RUnlock()

	for buf, _ := range e.cursorPositionEvents {
		buf.write <- v
	}
}

func (e *eventDispatcher) CursorWithinEvents() *CursorWithinEventBuffer {
	e.eventsAccess.Lock()
	defer e.eventsAccess.Unlock()

	buf := new(CursorWithinEventBuffer)
	buf.Read = make(chan bool)
	buf.write = make(chan bool)
	buf.lengthQuery = elasticBuffer(buf.write, buf.Read, func() {
		e.eventsAccess.Lock()
		defer e.eventsAccess.Unlock()

		delete(e.cursorWithinEvents, buf)
	})
	e.cursorWithinEvents[buf] = true
	return buf
}
func (e *eventDispatcher) addCursorWithinEvent(v bool) {
	e.eventsAccess.RLock()
	defer e.eventsAccess.RUnlock()

	for buf, _ := range e.cursorWithinEvents {
		buf.write <- v
	}
}

func (e *eventDispatcher) KeyboardStateEvents() *KeyboardStateEventBuffer {
	e.eventsAccess.Lock()
	defer e.eventsAccess.Unlock()

	buf := new(KeyboardStateEventBuffer)
	buf.Read = make(chan *keyboard.StateEvent)
	buf.write = make(chan *keyboard.StateEvent)
	buf.lengthQuery = elasticBuffer(buf.write, buf.Read, func() {
		e.eventsAccess.Lock()
		defer e.eventsAccess.Unlock()

		delete(e.keyboardStateEvents, buf)
	})
	e.keyboardStateEvents[buf] = true
	return buf
}
func (e *eventDispatcher) addKeyboardStateEvent(v *keyboard.StateEvent) {
	e.eventsAccess.RLock()
	defer e.eventsAccess.RUnlock()

	for buf, _ := range e.keyboardStateEvents {
		buf.write <- v
	}
}

func (e *eventDispatcher) KeyboardTypedEvents() *KeyboardTypedEventBuffer {
	e.eventsAccess.Lock()
	defer e.eventsAccess.Unlock()

	buf := new(KeyboardTypedEventBuffer)
	buf.Read = make(chan *keyboard.TypedEvent)
	buf.write = make(chan *keyboard.TypedEvent)
	buf.lengthQuery = elasticBuffer(buf.write, buf.Read, func() {
		e.eventsAccess.Lock()
		defer e.eventsAccess.Unlock()

		delete(e.keyboardTypedEvents, buf)
	})
	e.keyboardTypedEvents[buf] = true
	return buf
}
func (e *eventDispatcher) addKeyboardTypedEvent(v *keyboard.TypedEvent) {
	e.eventsAccess.RLock()
	defer e.eventsAccess.RUnlock()

	for buf, _ := range e.keyboardTypedEvents {
		buf.write <- v
	}
}

func (e *eventDispatcher) MaximizedEvents() *MaximizedEventBuffer {
	e.eventsAccess.Lock()
	defer e.eventsAccess.Unlock()

	buf := new(MaximizedEventBuffer)
	buf.Read = make(chan bool)
	buf.write = make(chan bool)
	buf.lengthQuery = elasticBuffer(buf.write, buf.Read, func() {
		e.eventsAccess.Lock()
		defer e.eventsAccess.Unlock()

		delete(e.maximizedEvents, buf)
	})
	e.maximizedEvents[buf] = true
	return buf
}
func (e *eventDispatcher) addMaximizedEvent(v bool) {
	e.eventsAccess.RLock()
	defer e.eventsAccess.RUnlock()

	for buf, _ := range e.maximizedEvents {
		buf.write <- v
	}
}

func (e *eventDispatcher) MinimizedEvents() *MinimizedEventBuffer {
	e.eventsAccess.Lock()
	defer e.eventsAccess.Unlock()

	buf := new(MinimizedEventBuffer)
	buf.Read = make(chan bool)
	buf.write = make(chan bool)
	buf.lengthQuery = elasticBuffer(buf.write, buf.Read, func() {
		e.eventsAccess.Lock()
		defer e.eventsAccess.Unlock()

		delete(e.minimizedEvents, buf)
	})
	e.minimizedEvents[buf] = true
	return buf
}
func (e *eventDispatcher) addMinimizedEvent(v bool) {
	e.eventsAccess.RLock()
	defer e.eventsAccess.RUnlock()

	for buf, _ := range e.minimizedEvents {
		buf.write <- v
	}
}

func (e *eventDispatcher) MouseEvents() *MouseEventBuffer {
	e.eventsAccess.Lock()
	defer e.eventsAccess.Unlock()

	buf := new(MouseEventBuffer)
	buf.Read = make(chan *mouse.Event)
	buf.write = make(chan *mouse.Event)
	buf.lengthQuery = elasticBuffer(buf.write, buf.Read, func() {
		e.eventsAccess.Lock()
		defer e.eventsAccess.Unlock()

		delete(e.mouseEvents, buf)
	})
	e.mouseEvents[buf] = true
	return buf
}
func (e *eventDispatcher) addMouseEvent(v *mouse.Event) {
	e.eventsAccess.RLock()
	defer e.eventsAccess.RUnlock()

	for buf, _ := range e.mouseEvents {
		buf.write <- v
	}
}

func (e *eventDispatcher) FocusedEvents() *FocusedEventBuffer {
	e.eventsAccess.Lock()
	defer e.eventsAccess.Unlock()

	buf := new(FocusedEventBuffer)
	buf.Read = make(chan bool)
	buf.write = make(chan bool)
	buf.lengthQuery = elasticBuffer(buf.write, buf.Read, func() {
		e.eventsAccess.Lock()
		defer e.eventsAccess.Unlock()

		delete(e.focusedEvents, buf)
	})
	e.focusedEvents[buf] = true
	return buf
}
func (e *eventDispatcher) addFocusedEvent(v bool) {
	e.eventsAccess.RLock()
	defer e.eventsAccess.RUnlock()

	for buf, _ := range e.focusedEvents {
		buf.write <- v
	}
}

func (e *eventDispatcher) PositionEvents() *PositionEventBuffer {
	e.eventsAccess.Lock()
	defer e.eventsAccess.Unlock()

	buf := new(PositionEventBuffer)
	buf.Read = make(chan []int)
	buf.write = make(chan []int)
	buf.lengthQuery = elasticBuffer(buf.write, buf.Read, func() {
		e.eventsAccess.Lock()
		defer e.eventsAccess.Unlock()

		delete(e.positionEvents, buf)
	})
	e.positionEvents[buf] = true
	return buf
}
func (e *eventDispatcher) addPositionEvent(v []int) {
	e.eventsAccess.RLock()
	defer e.eventsAccess.RUnlock()

	for buf, _ := range e.positionEvents {
		buf.write <- v
	}
}

func (e *eventDispatcher) SizeEvents() *SizeEventBuffer {
	e.eventsAccess.Lock()
	defer e.eventsAccess.Unlock()

	buf := new(SizeEventBuffer)
	buf.Read = make(chan []uint)
	buf.write = make(chan []uint)
	buf.lengthQuery = elasticBuffer(buf.write, buf.Read, func() {
		e.eventsAccess.Lock()
		defer e.eventsAccess.Unlock()

		delete(e.sizeEvents, buf)
	})
	e.sizeEvents[buf] = true
	return buf
}
func (e *eventDispatcher) addSizeEvent(v []uint) {
	e.eventsAccess.RLock()
	defer e.eventsAccess.RUnlock()

	for buf, _ := range e.sizeEvents {
		buf.write <- v
	}
}

func (e *eventDispatcher) ScreenChangedEvents() *ScreenChangedEventBuffer {
	e.eventsAccess.Lock()
	defer e.eventsAccess.Unlock()

	buf := new(ScreenChangedEventBuffer)
	buf.Read = make(chan Screen)
	buf.write = make(chan Screen)
	buf.lengthQuery = elasticBuffer(buf.write, buf.Read, func() {
		e.eventsAccess.Lock()
		defer e.eventsAccess.Unlock()

		delete(e.screenChangedEvents, buf)
	})
	e.screenChangedEvents[buf] = true
	return buf
}
func (e *eventDispatcher) addScreenChangedEvent(v Screen) {
	e.eventsAccess.RLock()
	defer e.eventsAccess.RUnlock()

	for buf, _ := range e.screenChangedEvents {
		buf.write <- v
	}
}
