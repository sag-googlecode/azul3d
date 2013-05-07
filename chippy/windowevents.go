package chippy

import (
	"code.google.com/p/azul3d/chippy/keyboard"
	"code.google.com/p/azul3d/chippy/mouse"
	"reflect"
)

// One-way channel with unlimited bufferwriteg
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

type CloseEventBuffer struct {
	Read, write chan bool
	lengthQuery chan int
}

func (b *CloseEventBuffer) Length() int {
	return <-b.lengthQuery
}

type CursorPositionEventBuffer struct {
	Read, write chan []int
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

type KeyboardEventBuffer struct {
	Read, write chan *keyboard.Event
	lengthQuery chan int
}

func (b *KeyboardEventBuffer) Length() int {
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

type eventDispatcher struct {
	closeEvents          []*CloseEventBuffer
	cursorPositionEvents []*CursorPositionEventBuffer
	cursorWithinEvents   []*CursorWithinEventBuffer
	keyboardEvents       []*KeyboardEventBuffer
	maximizedEvents      []*MaximizedEventBuffer
	minimizedEvents      []*MinimizedEventBuffer
	mouseEvents          []*MouseEventBuffer
	focusedEvents        []*FocusedEventBuffer
	positionEvents       []*PositionEventBuffer
	sizeEvents           []*SizeEventBuffer
}

func (e *eventDispatcher) CloseEvents() *CloseEventBuffer {
	buf := new(CloseEventBuffer)
	buf.Read = make(chan bool)
	buf.write = make(chan bool)
	buf.lengthQuery = elasticBuffer(buf.write, buf.Read)
	e.closeEvents = append(e.closeEvents, buf)
	return buf
}
func (e *eventDispatcher) addCloseEvent() bool {
	if len(e.closeEvents) == 0 {
		return false
	}

	for _, buf := range e.closeEvents {
		buf.write <- true
	}
	return true
}

func (e *eventDispatcher) CursorPositionEvents() *CursorPositionEventBuffer {
	buf := new(CursorPositionEventBuffer)
	buf.Read = make(chan []int)
	buf.write = make(chan []int)
	buf.lengthQuery = elasticBuffer(buf.write, buf.Read)
	e.cursorPositionEvents = append(e.cursorPositionEvents, buf)
	return buf
}
func (e *eventDispatcher) addCursorPositionEvent(v []int) {
	for _, buf := range e.cursorPositionEvents {
		buf.write <- v
	}
}

func (e *eventDispatcher) CursorWithinEvents() *CursorWithinEventBuffer {
	buf := new(CursorWithinEventBuffer)
	buf.Read = make(chan bool)
	buf.write = make(chan bool)
	buf.lengthQuery = elasticBuffer(buf.write, buf.Read)
	e.cursorWithinEvents = append(e.cursorWithinEvents, buf)
	return buf
}
func (e *eventDispatcher) addCursorWithinEvent(v bool) {
	for _, buf := range e.cursorWithinEvents {
		buf.write <- v
	}
}

func (e *eventDispatcher) KeyboardEvents() *KeyboardEventBuffer {
	buf := new(KeyboardEventBuffer)
	buf.Read = make(chan *keyboard.Event)
	buf.write = make(chan *keyboard.Event)
	buf.lengthQuery = elasticBuffer(buf.write, buf.Read)
	e.keyboardEvents = append(e.keyboardEvents, buf)
	return buf
}
func (e *eventDispatcher) addKeyboardEvent(v *keyboard.Event) {
	for _, buf := range e.keyboardEvents {
		buf.write <- v
	}
}

func (e *eventDispatcher) MaximizedEvents() *MaximizedEventBuffer {
	buf := new(MaximizedEventBuffer)
	buf.Read = make(chan bool)
	buf.write = make(chan bool)
	buf.lengthQuery = elasticBuffer(buf.write, buf.Read)
	e.maximizedEvents = append(e.maximizedEvents, buf)
	return buf
}
func (e *eventDispatcher) addMaximizedEvent(v bool) {
	for _, buf := range e.maximizedEvents {
		buf.write <- v
	}
}

func (e *eventDispatcher) MinimizedEvents() *MinimizedEventBuffer {
	buf := new(MinimizedEventBuffer)
	buf.Read = make(chan bool)
	buf.write = make(chan bool)
	buf.lengthQuery = elasticBuffer(buf.write, buf.Read)
	e.minimizedEvents = append(e.minimizedEvents, buf)
	return buf
}
func (e *eventDispatcher) addMinimizedEvent(v bool) {
	for _, buf := range e.minimizedEvents {
		buf.write <- v
	}
}

func (e *eventDispatcher) MouseEvents() *MouseEventBuffer {
	buf := new(MouseEventBuffer)
	buf.Read = make(chan *mouse.Event)
	buf.write = make(chan *mouse.Event)
	buf.lengthQuery = elasticBuffer(buf.write, buf.Read)
	e.mouseEvents = append(e.mouseEvents, buf)
	return buf
}
func (e *eventDispatcher) addMouseEvent(v *mouse.Event) {
	for _, buf := range e.mouseEvents {
		buf.write <- v
	}
}

func (e *eventDispatcher) FocusedEvents() *FocusedEventBuffer {
	buf := new(FocusedEventBuffer)
	buf.Read = make(chan bool)
	buf.write = make(chan bool)
	buf.lengthQuery = elasticBuffer(buf.write, buf.Read)
	e.focusedEvents = append(e.focusedEvents, buf)
	return buf
}
func (e *eventDispatcher) addFocusedEvent(v bool) {
	for _, buf := range e.focusedEvents {
		buf.write <- v
	}
}

func (e *eventDispatcher) PositionEvents() *PositionEventBuffer {
	buf := new(PositionEventBuffer)
	buf.Read = make(chan []int)
	buf.write = make(chan []int)
	buf.lengthQuery = elasticBuffer(buf.write, buf.Read)
	e.positionEvents = append(e.positionEvents, buf)
	return buf
}
func (e *eventDispatcher) addPositionEvent(v []int) {
	for _, buf := range e.positionEvents {
		buf.write <- v
	}
}

func (e *eventDispatcher) SizeEvents() *SizeEventBuffer {
	buf := new(SizeEventBuffer)
	buf.Read = make(chan []uint)
	buf.write = make(chan []uint)
	buf.lengthQuery = elasticBuffer(buf.write, buf.Read)
	e.sizeEvents = append(e.sizeEvents, buf)
	return buf
}
func (e *eventDispatcher) addSizeEvent(v []uint) {
	for _, buf := range e.sizeEvents {
		buf.write <- v
	}
}
