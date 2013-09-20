// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package wav

import (
	"bytes"
	"code.google.com/p/azul3d/audio"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"sync"
	"time"
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

	format, bitsPerSample, chunkSize uint16
	buffer                           audio.Samples

	r      interface{}
	rd     io.Reader
	config *audio.Config
}

func (d *decoder) Seek(t time.Duration) {
}

func (d *decoder) readPCM8(n int) (buf audio.Samples, err error) {
	if d.buffer == nil || d.buffer.Len() < n {
		d.buffer = make(audio.PCM8Samples, n)
	}
	bb := d.buffer.(audio.PCM8Samples)

	for i := 0; i < n; i++ {
		// Pull one sample from the PCM data stream
		var sample uint8

		err = binary.Read(d.rd, binary.LittleEndian, &sample)
		if err != nil {
			return bb[:i], err
		}

		bb[i] = audio.PCM8(sample)
	}

	return bb[:n], nil
}

func (d *decoder) readPCM16(n int) (buf audio.Samples, err error) {
	if d.buffer == nil || d.buffer.Len() < n {
		d.buffer = make(audio.PCM16Samples, n)
	}
	bb := d.buffer.(audio.PCM16Samples)

	for i := 0; i < n; i++ {
		// Pull one sample from the PCM data stream
		var sample audio.PCM16

		err = binary.Read(d.rd, binary.LittleEndian, &sample)
		if err != nil {
			return bb[:i], err
		}

		bb[i] = sample
	}

	return bb[:n], nil
}

func (d *decoder) readPCM32(n int) (buf audio.Samples, err error) {
	if d.buffer == nil || d.buffer.Len() < n {
		d.buffer = make(audio.PCM32Samples, n)
	}
	bb := d.buffer.(audio.PCM32Samples)

	for i := 0; i < n; i++ {
		// Pull one sample from the PCM data stream
		var sample audio.PCM32

		err = binary.Read(d.rd, binary.LittleEndian, &sample)
		if err != nil {
			return bb[:i], err
		}

		bb[i] = sample
	}

	return bb[:n], nil
}

func (d *decoder) readFloat32(n int) (buf audio.Samples, err error) {
	if d.buffer == nil || d.buffer.Len() < n {
		d.buffer = make(audio.F32Samples, n)
	}
	bb := d.buffer.(audio.F32Samples)

	for i := 0; i < n; i++ {
		// Pull one sample from the PCM data stream
		var sample audio.F32

		err = binary.Read(d.rd, binary.LittleEndian, &sample)
		if err != nil {
			return bb[:i], err
		}

		bb[i] = sample
	}

	return bb[:n], nil
}

func (d *decoder) readFloat64(n int) (buf audio.Samples, err error) {
	if d.buffer == nil || d.buffer.Len() < n {
		d.buffer = make(audio.F64Samples, n)
	}
	bb := d.buffer.(audio.F64Samples)

	for i := 0; i < n; i++ {
		// Pull one sample from the PCM data stream
		var sample audio.F64

		err = binary.Read(d.rd, binary.LittleEndian, &sample)
		if err != nil {
			return bb[:i], err
		}

		bb[i] = sample
	}

	return bb[:n], nil
}

func (d *decoder) readMuLaw(n int) (buf audio.Samples, err error) {
	if d.buffer == nil || d.buffer.Len() < n {
		d.buffer = make(audio.MuLawSamples, n)
	}
	bb := d.buffer.(audio.MuLawSamples)

	for i := 0; i < n; i++ {
		// Pull one sample from the MuLaw data stream
		var sample audio.MuLaw

		err = binary.Read(d.rd, binary.LittleEndian, &sample)
		if err != nil {
			return bb[:i], err
		}

		bb[i] = sample
	}

	return bb[:n], nil
}

func (d *decoder) readALaw(n int) (buf audio.Samples, err error) {
	if d.buffer == nil || d.buffer.Len() < n {
		d.buffer = make(audio.ALawSamples, n)
	}
	bb := d.buffer.(audio.ALawSamples)

	for i := 0; i < n; i++ {
		// Pull one sample from the ALaw data stream
		var sample audio.ALaw

		err = binary.Read(d.rd, binary.LittleEndian, &sample)
		if err != nil {
			return bb[:i], err
		}

		bb[i] = sample
	}

	return bb[:n], nil
}

func (d *decoder) Read(n int) (buf audio.Samples, err error) {
	if n <= 0 {
		panic("Read(): n <= 0")
	}

	d.access.Lock()
	defer d.access.Unlock()

	fmt.Println("Read()", n, "samples")

	switch d.format {
	case wave_FORMAT_PCM:
		switch d.bitsPerSample {
		case 8:
			return d.readPCM8(n)
		case 16:
			return d.readPCM16(n)
		case 32:
			return d.readPCM32(n)
		}

	case wave_FORMAT_IEEE_FLOAT:
		switch d.bitsPerSample {
		case 32:
			return d.readFloat32(n)
		case 64:
			return d.readFloat64(n)
		}

	case wave_FORMAT_MULAW:
		return d.readMuLaw(n)

	case wave_FORMAT_ALAW:
		return d.readALaw(n)
	}
	return
}

func (d *decoder) Config() *audio.Config {
	d.access.RLock()
	defer d.access.RUnlock()

	return d.config
}

// ErrUnsupported defines an error for decoding wav data that is valid (by the
// wave specification) but not supported by the decoder in this package.
//
// This error only happens for audio files containing extensible wav data.
var ErrUnsupported = errors.New("wav: data format is valid but not supported by decoder")

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

	// Firstly, read RIFF chunk
	var riffChunkHeader chunkHeader
	err := binary.Read(d.rd, binary.LittleEndian, &riffChunkHeader)
	if err != nil {
		return nil, err
	}

	if !bytes.Equal(riffChunkHeader.ChunkID[:], []byte("RIFF")) {
		return nil, audio.ErrInvalidData
	}

	var format [4]byte
	err = binary.Read(d.rd, binary.LittleEndian, &format)
	if !bytes.Equal(format[:], []byte("WAVE")) {
		return nil, audio.ErrInvalidData
	}

	// Secondly, read the format chunk
	var fmtChunkHeader chunkHeader
	err = binary.Read(d.rd, binary.LittleEndian, &fmtChunkHeader)
	if err != nil {
		return nil, err
	}

	if !bytes.Equal(fmtChunkHeader.ChunkID[:], []byte("fmt ")) {
		return nil, audio.ErrInvalidData
	}

	var (
		c16 fmtChunk16
		c18 fmtChunk18
		c40 fmtChunk40
	)

	// Always contains the 16-byte chunk
	err = binary.Read(d.rd, binary.LittleEndian, &c16)
	if err != nil {
		return nil, err
	}
	d.bitsPerSample = to16(c16.BitsPerSample[:])

	// Sometimes contains extensive 18/40 byte chunks
	switch to16(fmtChunkHeader.ChunkSize[:]) {
	case 18:
		err = binary.Read(d.rd, binary.LittleEndian, &c18)
		if err != nil {
			return nil, err
		}
	case 40:
		err = binary.Read(d.rd, binary.LittleEndian, &c40)
		if err != nil {
			return nil, err
		}
	}

	// Verify format tag
	ft := to16(c16.FormatTag[:])
	switch {
	case ft == wave_FORMAT_PCM && (d.bitsPerSample == 8 || d.bitsPerSample == 16 || d.bitsPerSample == 32):
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
		SampleRate: int(to16(c16.SamplesPerSec[:])),
	}

	// We don't care about format, we just want to know what is next.
	var factOrData chunkHeader
	err = binary.Read(d.rd, binary.LittleEndian, &factOrData)
	if err != nil {
		return nil, err
	}

	var dataChunkHeader chunkHeader
	if bytes.Equal(factOrData.ChunkID[:], []byte("fact")) {
		// We need to scan fact chunk first.
		var fact factChunk
		err = binary.Read(d.rd, binary.LittleEndian, &fact)
		if err != nil {
			return nil, err
		}

		// Read the data chunk header now
		err = binary.Read(d.rd, binary.LittleEndian, &dataChunkHeader)
		if err != nil {
			return nil, err
		}

	} else if bytes.Equal(dataChunkHeader.ChunkID[:], []byte("data")) {
		// Read the data chunk header now
		err = binary.Read(d.rd, binary.LittleEndian, &dataChunkHeader)
		if err != nil {
			return nil, err
		}

	} else {
		return nil, audio.ErrInvalidData
	}

	d.chunkSize = to16(dataChunkHeader.ChunkSize[:])

	return d, nil
}

func init() {
	audio.RegisterFormat("wav", "RIFF", newDecoder)
}
