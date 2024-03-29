// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package wav

import (
	"azul3d.org/v1/audio"
	"encoding/binary"
	"errors"
	"io"
	"sync"
	"unsafe"
)

const (
	// Data format codes

	// PCM
	wave_FORMAT_PCM = 0x0001

	// IEEE float
	wave_FORMAT_IEEE_FLOAT = 0x0003

	// 8-bit ITU-T G.711 A-law
	wave_FORMAT_ALAW = 0x0006

	// 8-bit ITU-T G.711 µ-law
	wave_FORMAT_MULAW = 0x0007

	// Determined by SubFormat
	wave_FORMAT_EXTENSIBLE = 0xFFFE
)

type decoder struct {
	access sync.RWMutex

	format, bitsPerSample   uint16
	chunkSize, currentCount uint32
	dataChunkBegin          int32

	r      interface{}
	rd     io.Reader
	config *audio.Config
}

func (d *decoder) bRead(data interface{}, sz uintptr) error {
	d.currentCount += uint32(sz)
	if d.chunkSize > 0 {
		if d.currentCount > d.chunkSize {
			return audio.EOS
		}
	} else {
		d.dataChunkBegin += int32(sz)
	}
	return binary.Read(d.rd, binary.LittleEndian, data)
}

// Reads and returns the next RIFF chunk, note that always len(ident) == 4
// E.g.
//
//  "fmt " (notice space).
//
// Length is length of chunk data.
//
// Returns any read errors.
func (d *decoder) nextChunk() (ident string, length uint32, err error) {
	// Read chunk identity, like "RIFF" or "fmt "
	var chunkIdent [4]byte
	err = d.bRead(&chunkIdent, unsafe.Sizeof(chunkIdent))
	if err != nil {
		return "", 0, err
	}
	ident = string(chunkIdent[:])

	// Read chunk length
	err = d.bRead(&length, unsafe.Sizeof(length))
	if err != nil {
		return "", 0, err
	}
	return
}

func (d *decoder) Seek(sample uint64) error {
	rs, ok := d.r.(io.ReadSeeker)
	if ok {
		offset := int64(sample * (uint64(d.bitsPerSample) / 8))
		_, err := rs.Seek(int64(d.dataChunkBegin)+offset, 0)
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *decoder) readPCM8(b audio.Slice) (read int, err error) {
	bb, bbOk := b.(audio.PCM8Samples)

	for read = 0; read < b.Len(); read++ {
		// Pull one sample from the data stream
		var sample audio.PCM8

		err = d.bRead(&sample, unsafe.Sizeof(sample))
		if err != nil {
			return
		}

		if bbOk {
			bb[read] = sample
		} else {
			f64 := audio.PCM8ToF64(sample)
			b.Set(read, f64)
		}
	}

	return
}

func (d *decoder) readPCM16(b audio.Slice) (read int, err error) {
	bb, bbOk := b.(audio.PCM16Samples)

	for read = 0; read < b.Len(); read++ {
		// Pull one sample from the data stream
		var sample audio.PCM16

		err = d.bRead(&sample, unsafe.Sizeof(sample))
		if err != nil {
			return
		}

		if bbOk {
			bb[read] = sample
		} else {
			f64 := audio.PCM16ToF64(sample)
			b.Set(read, f64)
		}
	}

	return
}

func (d *decoder) readPCM24(b audio.Slice) (read int, err error) {
	bb, bbOk := b.(audio.PCM32Samples)

	for read = 0; read < b.Len(); read++ {
		// Pull one sample from the data stream
		var sample [3]uint8

		err = d.bRead(&sample, unsafe.Sizeof(sample))
		if err != nil {
			return
		}

		var ss audio.PCM32
		if (ss & 0x800000) > 0 {
			ss |= ^0xffffff
		}

		if bbOk {
			bb[read] = ss
		} else {
			f64 := audio.PCM32ToF64(ss)
			b.Set(read, f64)
		}
	}

	return
}

func (d *decoder) readPCM32(b audio.Slice) (read int, err error) {
	bb, bbOk := b.(audio.PCM32Samples)

	for read = 0; read < b.Len(); read++ {
		// Pull one sample from the data stream
		var sample audio.PCM32

		err = d.bRead(&sample, unsafe.Sizeof(sample))
		if err != nil {
			return
		}

		if bbOk {
			bb[read] = sample
		} else {
			f64 := audio.PCM32ToF64(sample)
			b.Set(read, f64)
		}
	}

	return
}

func (d *decoder) readF32(b audio.Slice) (read int, err error) {
	bb, bbOk := b.(audio.F32Samples)

	for read = 0; read < b.Len(); read++ {
		// Pull one sample from the data stream
		var sample audio.F32

		err = d.bRead(&sample, unsafe.Sizeof(sample))
		if err != nil {
			return
		}

		if bbOk {
			bb[read] = sample
		} else {
			b.Set(read, audio.F64(sample))
		}
	}

	return
}

func (d *decoder) readF64(b audio.Slice) (read int, err error) {
	for read = 0; read < b.Len(); read++ {
		// Pull one sample from the data stream
		var sample audio.F64

		err = d.bRead(&sample, unsafe.Sizeof(sample))
		if err != nil {
			return
		}

		b.Set(read, sample)
	}

	return
}

func (d *decoder) readMuLaw(b audio.Slice) (read int, err error) {
	bb, bbOk := b.(audio.MuLawSamples)

	for read = 0; read < b.Len(); read++ {
		// Pull one sample from the data stream
		var sample audio.MuLaw

		err = d.bRead(&sample, unsafe.Sizeof(sample))
		if err != nil {
			return
		}

		if bbOk {
			bb[read] = sample
		} else {
			p16 := audio.MuLawToPCM16(sample)
			b.Set(read, audio.PCM16ToF64(p16))
		}
	}

	return
}

func (d *decoder) readALaw(b audio.Slice) (read int, err error) {
	bb, bbOk := b.(audio.ALawSamples)

	for read = 0; read < b.Len(); read++ {
		// Pull one sample from the data stream
		var sample audio.ALaw

		err = d.bRead(&sample, unsafe.Sizeof(sample))
		if err != nil {
			return
		}

		if bbOk {
			bb[read] = sample
		} else {
			p16 := audio.ALawToPCM16(sample)
			b.Set(read, audio.PCM16ToF64(p16))
		}
	}

	return
}

func (d *decoder) Read(b audio.Slice) (read int, err error) {
	if b.Len() == 0 {
		return
	}

	d.access.Lock()
	defer d.access.Unlock()

	switch d.format {
	case wave_FORMAT_PCM:
		switch d.bitsPerSample {
		case 8:
			return d.readPCM8(b)
		case 16:
			return d.readPCM16(b)
		case 24:
			return d.readPCM24(b)
		case 32:
			return d.readPCM32(b)
		}

	case wave_FORMAT_IEEE_FLOAT:
		switch d.bitsPerSample {
		case 32:
			return d.readF32(b)
		case 64:
			return d.readF64(b)
		}

	case wave_FORMAT_MULAW:
		return d.readMuLaw(b)

	case wave_FORMAT_ALAW:
		return d.readALaw(b)
	}
	return
}

func (d *decoder) Config() audio.Config {
	d.access.RLock()
	defer d.access.RUnlock()

	return *d.config
}

// ErrUnsupported defines an error for decoding wav data that is valid (by the
// wave specification) but not supported by the decoder in this package.
//
// This error only happens for audio files containing extensible wav data.
var ErrUnsupported = errors.New("wav: data format is valid but not supported by decoder")

func to32(b []byte) uint32 {
	return binary.LittleEndian.Uint32(b)
}

func to16(b []byte) uint16 {
	return binary.LittleEndian.Uint16(b)
}

// NewDecoder returns a new initialized audio decoder for the io.Reader or
// io.ReadSeeker, r.
func newDecoder(r interface{}) (audio.Decoder, error) {
	d := new(decoder)
	d.r = r

	switch t := r.(type) {
	case io.Reader:
		d.rd = t
	case io.ReadSeeker:
		d.rd = io.Reader(t)
	default:
		panic("NewDecoder(): Invalid reader type; must be io.Reader or io.ReadSeeker!")
	}

	var (
		complete bool

		c16 fmtChunk16
		c18 fmtChunk18
		c40 fmtChunk40
	)
	for !complete {
		ident, length, err := d.nextChunk()
		if err != nil {
			return nil, err
		}

		switch ident {
		case "RIFF":
			var format [4]byte
			err = d.bRead(&format, unsafe.Sizeof(format))
			if string(format[:]) != "WAVE" {
				return nil, audio.ErrInvalidData
			}

		case "fmt ":
			// Always contains the 16-byte chunk
			err = d.bRead(&c16, unsafe.Sizeof(c16))
			if err != nil {
				return nil, err
			}
			d.bitsPerSample = to16(c16.BitsPerSample[:])

			// Sometimes contains extensive 18/40 total byte chunks
			switch length {
			case 18:
				err = d.bRead(&c18, unsafe.Sizeof(c18))
				if err != nil {
					return nil, err
				}
			case 40:
				err = d.bRead(&c40, unsafe.Sizeof(c40))
				if err != nil {
					return nil, err
				}
			}

			// Verify format tag
			ft := to16(c16.FormatTag[:])
			switch {
			case ft == wave_FORMAT_PCM && (d.bitsPerSample == 8 || d.bitsPerSample == 16 || d.bitsPerSample == 24 || d.bitsPerSample == 32):
				break
			case ft == wave_FORMAT_IEEE_FLOAT && (d.bitsPerSample == 32 || d.bitsPerSample == 64):
				break
			case ft == wave_FORMAT_ALAW && d.bitsPerSample == 8:
				break
			case ft == wave_FORMAT_MULAW && d.bitsPerSample == 8:
				break
			// We don't support extensible wav files
			//case wave_FORMAT_EXTENSIBLE:
			//	break
			default:
				return nil, ErrUnsupported
			}

			// Assign format tag for later (See Read() method)
			d.format = to16(c16.FormatTag[:])

			// We now have enough information to build the audio configuration
			d.config = &audio.Config{
				Channels:   int(to16(c16.Channels[:])),
				SampleRate: int(to32(c16.SamplesPerSec[:])),
			}

		case "fact":
			// We need to scan fact chunk first.
			var fact factChunk
			err = d.bRead(&fact, unsafe.Sizeof(fact))
			if err != nil {
				return nil, err
			}

		case "data":
			// Read the data chunk header now
			d.chunkSize = length
			complete = true
		}
	}

	return d, nil
}

func init() {
	audio.RegisterFormat("wav", "RIFF", newDecoder)
}
