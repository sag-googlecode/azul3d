package aio

/*
#cgo LDFLAGS: -ldsound

#define UNICODE
#include <windows.h>

#define CINTERFACE
#include <dsound.h>

HRESULT mIDirectSound_Initialize(IDirectSound* i, LPCGUID a);
HRESULT mIDirectSound_SetCooperativeLevel(IDirectSound* i, HWND a, DWORD b);
void mIDirectSound_AddRef(IDirectSound* i);
void mIDirectSound_Release(IDirectSound* i);
HRESULT mIDirectSound_GetCaps(IDirectSound* i, LPDSCAPS a);
HRESULT mIDirectSound_CreateSoundBuffer(IDirectSound* i, LPCDSBUFFERDESC, LPDIRECTSOUNDBUFFER*, LPUNKNOWN);
HRESULT mIDirectSoundBuffer_Stop(IDirectSoundBuffer* i);


void mIDirectSoundBuffer_AddRef(IDirectSoundBuffer* i);
void mIDirectSoundBuffer_Release(IDirectSoundBuffer* i);
HRESULT mIDirectSoundBuffer_Lock(IDirectSoundBuffer* i, DWORD a, DWORD b, LPVOID* c, LPDWORD d, LPVOID* e, LPDWORD f, DWORD g);
HRESULT mIDirectSoundBuffer_Unlock(IDirectSoundBuffer* i, LPVOID a, DWORD b, LPVOID c, DWORD d);
HRESULT mIDirectSoundBuffer_Play(IDirectSoundBuffer* i, DWORD a, DWORD b, DWORD c);
HRESULT mIDirectSoundBuffer_GetStatus(IDirectSoundBuffer* i, LPDWORD a);
HRESULT mIDirectSoundBuffer_SetCurrentPosition(IDirectSoundBuffer* i, DWORD a);
HRESULT mIDirectSoundBuffer_GetCurrentPosition(IDirectSoundBuffer* i, LPDWORD a, LPDWORD b);

BOOL audioDSEnumCallback(LPGUID, LPTSTR, LPTSTR, LPVOID);
*/
import "C"

import (
	"azul3d.org/v0/audio"
	"errors"
	"fmt"
	"reflect"
	"runtime"
	"time"
	"unicode/utf16"
	"unsafe"
)

// Decodes a UTF-16 encoded C.LPTSTR/C.LPWSTR to a UTF-8 encoded Go string.
//
// if cstr == nil: "" is returned
//
// This function does not touch/free the memory held by the cstr parameter.
func lPTSTRToString(cstr C.LPTSTR) string {
	if cstr == nil {
		return ""
	}
	strlen := int(C.wcslen((*C.wchar_t)(unsafe.Pointer(cstr))))

	var wstr []uint16
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&wstr))
	sliceHeader.Cap = strlen
	sliceHeader.Len = strlen
	sliceHeader.Data = uintptr(unsafe.Pointer(cstr))

	return string(utf16.Decode(wstr))
}

// Encodes a UTF-8 encoded Go string to a UTF-16 encoded C.LPTSTR/C.LPWSTR.
//
// if len(g) == 0: nil is returned.
//
// Note: The returned C.LPTSTR should be free'd at some point; it is malloc'd
func stringToLPTSTR(g string) C.LPTSTR {
	if len(g) == 0 {
		return nil
	}

	u16 := utf16.Encode([]rune(g))

	// u16 is uint16 type
	nBytes := C.size_t(len(u16) * 2)

	// Allocate a buffer
	cstr := (C.LPTSTR)(C.calloc(1, nBytes+2)) // +2 for uint16 NULL terminator

	// Memcpy the UTF-16 encoded string into the buffer
	C.memcpy(unsafe.Pointer(cstr), unsafe.Pointer(&u16[0]), nBytes)

	return cstr
}

// Returns nil for no error; returns a LONG error string for an error
func errString(c C.HRESULT) (err error) {
	switch c {
	case C.DS_OK:
		return nil
	case C.DS_NO_VIRTUALIZATION:
		return errors.New("The buffer was created, but another 3D algorithm was substituted.")
	case C.DS_INCOMPLETE:
		return errors.New("The method succeeded, but not all the optional effects were obtained.")
	case C.DSERR_ACCESSDENIED:
		return errors.New("The request failed because access was denied.")
	case C.DSERR_ALLOCATED:
		return errors.New("The request failed because resources, such as a priority level, were already in use by another caller.")
	case C.DSERR_ALREADYINITIALIZED:
		return errors.New("The object is already initialized.")
	case C.DSERR_BADFORMAT:
		return errors.New("The specified wave format is not supported.")
	case C.DSERR_BADSENDBUFFERGUID:
		return errors.New("The GUID specified in an audiopath file does not match a valid mix-in buffer.")
	case C.DSERR_BUFFERLOST:
		return errors.New("The buffer memory has been lost and must be restored.")
	case C.DSERR_BUFFERTOOSMALL:
		return errors.New("The buffer size is not great enough to enable effects processing.")
	case C.DSERR_CONTROLUNAVAIL:
		return errors.New("The buffer control (volume, pan, and so on) requested by the caller is not available. Controls must be specified when the buffer is created, using the dwFlags member of case C.DSBUFFERDESC.")
	case C.DSERR_DS8_REQUIRED:
		return errors.New("A DirectSound object of class CLSID_DirectSound or later is required for the requested functionality. For more information, see IDirectSound Interface.")
	case C.DSERR_FXUNAVAILABLE:
		return errors.New("The effects requested could not be found on the system, or they are in the wrong order or in the wrong location; for example, an effect expected in hardware was found in software.")
	case C.DSERR_GENERIC:
		return errors.New("An undetermined error occurred inside the DirectSound subsystem.")
	case C.DSERR_INVALIDCALL:
		return errors.New("This function is not valid for the current state of this object.")
	case C.DSERR_INVALIDPARAM:
		return errors.New("An invalid parameter was passed to the returning function.")
	case C.DSERR_NOAGGREGATION:
		return errors.New("The object does not support aggregation.")
	case C.DSERR_NODRIVER:
		return errors.New("No sound driver is available for use, or the given GUID is not a valid DirectSound output ID.")
	case C.DSERR_NOINTERFACE:
		return errors.New("The requested COM interface is not available.")
	case C.DSERR_OBJECTNOTFOUND:
		return errors.New("The requested object was not found.")
	case C.DSERR_OTHERAPPHASPRIO:
		return errors.New("Another application has a higher priority level, preventing this call from succeeding.")
	case C.DSERR_OUTOFMEMORY:
		return errors.New("The DirectSound subsystem could not allocate sufficient memory to complete the caller's request.")
	case C.DSERR_PRIOLEVELNEEDED:
		return errors.New("A cooperative level of DSSCL_PRIORITY or higher is required.")
	case C.DSERR_SENDLOOP:
		return errors.New("A circular loop of send effects was detected.")
	case C.DSERR_UNINITIALIZED:
		return errors.New("The IDirectSound::Initialize method has not been called or has not been called successfully before other methods were called.")
	case C.DSERR_UNSUPPORTED:
		return errors.New("The function called is not supported at this time.")
	}
	return fmt.Errorf("Unknown error 0x%X", c)
}

type enumContext struct {
	outputs       []*Output
	defaultOutput int
}

//export audioDSEnumCallback
func audioDSEnumCallback(guid C.LPGUID, cdescription, cmodule C.LPTSTR, contextPtr unsafe.Pointer) C.BOOL {
	ctx := (*enumContext)(contextPtr)

	description := lPTSTRToString(cdescription)
	//module := lPTSTRToString(cmodule)

	var ds *C.IDirectSound

	err := errString(C.DirectSoundCreate(guid, (*C.LPDIRECTSOUND)(unsafe.Pointer(&ds)), nil))
	if err != nil {
		logger().Printf("DirectSoundCreate() failed (for %q)!\n", description)
		logger().Println(err)
		return C.TRUE
	}

	// See the following:
	//
	// http://stackoverflow.com/questions/6714111/how-to-provide-a-hwnd-to-directsound-setcooperativelevel-in-a-console-program
	//
	hwnd := C.GetDesktopWindow()
	err = errString(C.mIDirectSound_SetCooperativeLevel(ds, hwnd, C.DSSCL_PRIORITY))
	if err != nil {
		logger().Printf("IDirectSound_SetCooperativeLevel() failed (for %q)!\n", description)
		logger().Println(err)
		return C.TRUE
	}

	// We need to specify the size of the struct in dwSize (used to handle
	// different version's of DirectSound on Windows).
	var dsCaps C.DSCAPS
	dsCaps.dwSize = C.DWORD(unsafe.Sizeof(dsCaps))
	caps := (C.LPDSCAPS)(unsafe.Pointer(&dsCaps))
	err = errString(C.mIDirectSound_GetCaps(ds, caps))
	if err != nil {
		logger().Printf("IDirectSound_GetCaps() failed (for %q)!\n", description)
		logger().Println(err)
		return C.TRUE
	}

	if (dsCaps.dwFlags & C.DSCAPS_EMULDRIVER) > 0 {
		logger().Println("DirectSound reports that driver is emulated. OK.")
	}

	// Determine audio type
	var (
		audioType     audio.Type
		bitsPerSample int
	)
	if (dsCaps.dwFlags & C.DSCAPS_PRIMARY16BIT) > 0 {
		// Looks like we have 16-bit support.
		audioType = audio.TYPE_PCM16
		bitsPerSample = 16

	} else if (dsCaps.dwFlags & C.DSCAPS_PRIMARY8BIT) > 0 {
		// We have 8-bit support then.
		audioType = audio.TYPE_PCM8
		bitsPerSample = 8
	} else {
		// We don't have support for any audio formats? That's odd. Maybe it's
		// just a buggy driver. We assume 16-bit support.
		logger().Println("DirectSound reports no support 8/16 bit data. Using 16-bit regardless.")
		audioType = audio.TYPE_PCM16
		bitsPerSample = 16
	}

	output := newOutput(description, audioType)
	output.nativeOutput.audioType = audioType
	output.nativeOutput.bitsPerSample = bitsPerSample
	output.nativeOutput.ds = ds
	output.nativeOutput.wBuffer = make(chan audio.Buffer)
	output.nativeOutput.wSamples = make(chan int)
	output.nativeOutput.wError = make(chan error)

	// Add finalizer now.
	runtime.SetFinalizer(output, func(d *Output) {
		// Decrease reference count
		C.mIDirectSound_Release(ds)
	})

	// For later, we need the GUID, but make sure that we make a copy of it (as
	// it's memory is local in context to this callback).
	if guid != nil {
		// Note: first 'default' guid is nil (which works on DirectSound
		// functions as the 'default' device).
		cpy := *guid
		output.nativeOutput.guid = &cpy
	}

	if guid == nil {
		// This is the default output device. We store the index of it here.
		ctx.defaultOutput = len(ctx.outputs)
	}

	// Append to output (only if creating output above worked).
	ctx.outputs = append(ctx.outputs, output)

	return C.TRUE
}

func backend_Devices() (inputs []*Input, outputs []*Output, defaultInput int, defaultOutput int) {
	// Find output devices
	callback := (*[0]byte)(unsafe.Pointer(C.audioDSEnumCallback))

	context := &enumContext{
		outputs: outputs,
	}
	if C.DirectSoundEnumerate(callback, C.LPVOID(unsafe.Pointer(context))) != 0 {
		logger().Println("Failed to query outputs; DirectSoundEnumerate() failed!")
		return
	}
	outputs = context.outputs
	defaultOutput = context.defaultOutput
	return
}

type nativeInput struct {
}

func (d *nativeInput) equals(other *nativeInput) bool {
	return false
}

type nativeOutput struct {
	ds            *C.IDirectSound
	guid          *C.GUID
	bitsPerSample int
	audioType     audio.Type

	configError error
	config      *audio.Config
	dsBuffer    *C.IDirectSoundBuffer

	// Used for write() calls.
	bufferSize, bufferSizeBytes int
	wBuffer                     chan audio.Buffer
	wSamples                    chan int
	wError                      chan error
}

func (d *nativeOutput) equals(other *nativeOutput) bool {
	if (d.guid == nil && other.guid != nil) || (d.guid != nil && other.guid == nil) {
		return false
	}
	if d.guid == nil && other.guid == nil {
		return true
	}
	if *d.guid == *other.guid {
		return true
	}
	return false
}

func (o *nativeOutput) setConfig(bufferSize int, config *audio.Config) {
	o.bufferSize = bufferSize
	o.bufferSizeBytes = bufferSize * 2 * (o.bitsPerSample / 8)

	o.config = config
	var (
		description C.DSBUFFERDESC
		waveFormat  C.WAVEFORMATEX
	)
	waveFormat.wFormatTag = C.WAVE_FORMAT_PCM
	waveFormat.nSamplesPerSec = C.DWORD(config.SampleRate)
	waveFormat.wBitsPerSample = C.WORD(o.bitsPerSample)
	waveFormat.nChannels = C.WORD(config.Channels)
	waveFormat.nBlockAlign = C.WORD((o.bitsPerSample / 8) * config.Channels)
	waveFormat.nAvgBytesPerSec = waveFormat.nSamplesPerSec * C.DWORD(waveFormat.nBlockAlign)
	waveFormat.cbSize = 0

	description.dwSize = C.DWORD(unsafe.Sizeof(description))
	description.dwFlags |= C.DSBCAPS_GLOBALFOCUS
	description.dwBufferBytes = C.DWORD(o.bufferSizeBytes)
	description.dwReserved = 0
	description.lpwfxFormat = (C.LPWAVEFORMATEX)(unsafe.Pointer(&waveFormat))

	e := errString(C.mIDirectSound_CreateSoundBuffer(
		o.ds,
		(C.LPCDSBUFFERDESC)(unsafe.Pointer(&description)),
		(*C.LPDIRECTSOUNDBUFFER)(unsafe.Pointer(&o.dsBuffer)),
		nil,
	))
	if e != nil {
		o.configError = e
		return
	}

	go o.feedToDirectSound()

	e = errString(C.mIDirectSoundBuffer_Play(o.dsBuffer, 0, 0, C.DSBPLAY_LOOPING))
	if e != nil {
		logger().Println(e)
		//o.wSamples <- 0
		//o.wError <- e
		return
	}

	//runtime.SetFinalizer(b, func(b *HWBuffer) {
	//	C.mIDirectSoundBuffer_Release(b.nativeBuffer.ds)
	//})
}

func (o *nativeOutput) feedToDirectSound() {
	var (
		e        error
		typeSize = o.bitsPerSample / 8

		bPtr, bPtr2 C.LPVOID
		bSz, bSz2   C.DWORD

		samples unsafe.Pointer

		firstIteration                                         = true
		playCursor, safeWriteCursor, writeCursor, lastWriteEnd C.DWORD

		stopTime = 200 * time.Millisecond
	)

	for {
		// Wait for someone to write to this output.
		var buf audio.Buffer
		select {
		case buf = <-o.wBuffer:
			stopTime = 200 * time.Millisecond

			e = errString(C.mIDirectSoundBuffer_Play(o.dsBuffer, 0, 0, C.DSBPLAY_LOOPING))
			if e != nil {
				o.wSamples <- 0
				o.wError <- e
				return
			}

		case <-time.After(stopTime):
			// Wait a *really* long time untill the next samples come in over
			// the o.wBuffer channel.
			stopTime = 1 * time.Hour

			var tmpPlay, tmpWrite C.DWORD
			e = errString(C.mIDirectSoundBuffer_GetCurrentPosition(o.dsBuffer, &tmpPlay, &tmpWrite))
			if e != nil {
				logger().Println(e)
				return
			}
			if tmpPlay > writeCursor {
				C.mIDirectSoundBuffer_Stop(o.dsBuffer)
			}
			continue
		}

		bytesToCopy := C.DWORD(buf.Len() * typeSize)
		samplesWrote := buf.Len()

		// Before locking buffer; convert audio to correct format if we need
		// to.
		v := audio.Convert(buf, o.audioType)
		switch o.audioType {
		case audio.TYPE_PCM8:
			// Our buffer is PCM8 audio.
			samples = unsafe.Pointer(&v.(audio.PCM8Samples)[0])
		case audio.TYPE_PCM16:
			// Our buffer is PCM16 audio
			samples = unsafe.Pointer(&v.(audio.PCM16Samples)[0])
		default:
			panic("Buffer audio type unknown? This shouldn't happen.")
		}

		var writeEnd C.DWORD
		for {
			e = errString(C.mIDirectSoundBuffer_GetCurrentPosition(o.dsBuffer, &playCursor, &safeWriteCursor))
			if e != nil {
				o.wSamples <- 0
				o.wError <- e
				return
			}

			writeEnd = writeCursor + bytesToCopy

			if firstIteration {
				break
			}

			if lastWriteEnd < playCursor && writeEnd > playCursor {
				continue
			} else {
				break
			}
		}

		// Lock the buffer.
		e = errString(C.mIDirectSoundBuffer_Lock(o.dsBuffer, writeCursor, bytesToCopy, &bPtr, &bSz, &bPtr2, &bSz2, 0))
		if e != nil {
			o.wSamples <- 0
			o.wError <- e
			return
		}

		// Write to buffer
		C.memcpy(unsafe.Pointer(bPtr), samples, C.size_t(bSz))

		//samples = unsafe.Pointer(uintptr(samples) + uintptr(bSz))
		//C.memcpy(unsafe.Pointer(bPtr2), samples, C.size_t(bSz2))
		writeCursor += bytesToCopy
		writeCursor = C.DWORD(int(writeCursor) % o.bufferSizeBytes)

		// Unlock the buffer.
		e = errString(C.mIDirectSoundBuffer_Unlock(o.dsBuffer, bPtr, bSz, bPtr2, bSz2))
		if e != nil {
			o.wSamples <- 0
			o.wError <- e
			return
		}

		firstIteration = false
		lastWriteEnd = writeEnd % C.DWORD(o.bufferSizeBytes)
		o.wSamples <- samplesWrote
		o.wError <- e
	}
}

func (o *nativeOutput) write(b audio.Buffer) (wrote int, err error) {
	if o.configError != nil {
		return 0, o.configError
	}

	o.wBuffer <- b
	wrote = <-o.wSamples
	err = <-o.wError
	return
}

/*
type nativeBuffer struct {
	ds *C.IDirectSoundBuffer
	cBufPtr C.LPVOID
	cBufPtrSize C.DWORD

	locked bool
}

func (b *nativeBuffer) setLocked() {
	b.locked = false
	if !b.locked {
		b.locked = true

	}
}

func (b *nativeBuffer) setUnlocked() {
	b.locked = true
	if b.locked {
		b.locked = false

		// Unlock buffer (also copies data to actual device memory).
		e := errString(C.mIDirectSoundBuffer_Unlock(b.ds, b.cBufPtr, b.cBufPtrSize, nil, 0))
		if e != nil {
			logger().Println(e)
		}
	}
}

func (o *nativeOutput) newBuffer(b *HWBuffer) (set func(index int, s audio.F64), err error) {
	return
}

func (o *nativeOutput) play(b *HWBuffer) {
	//b.nativeBuffer.setUnlocked()

	e := errString(C.mIDirectSoundBuffer_Play(b.nativeBuffer.ds, 0, 0, 0))
	if e != nil {
		logger().Println(e)
		return
	}

	s := float64(b.Len()) / float64(b.Config.SampleRate * b.Config.Channels)
	time.Sleep((time.Duration(s * 1000) * time.Millisecond))






	for {
		var status C.DWORD
		e := errString(C.mIDirectSoundBuffer_GetStatus(b.nativeBuffer.ds, &status))
		if e != nil {
			logger().Println(e)
			break
		}

		s := float64(b.Len()) / float64(b.Config.SampleRate * b.Config.Channels)
		time.Sleep(time.Duration(s * 1000) * time.Millisecond)

		break
		//playing := (status & C.DSBSTATUS_PLAYING) > 0
		//if !playing {
		//	break
		//}
	}








	e = errString(C.mIDirectSoundBuffer_SetCurrentPosition(b.nativeBuffer.ds, 0))
	if e != nil {
		logger().Println(e)
	}

	e = errString(C.mIDirectSoundBuffer_Stop(b.nativeBuffer.ds))
	if e != nil {
		logger().Println(e)
	}
}
*/
