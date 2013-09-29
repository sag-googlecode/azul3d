// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

// Package audio implements various audio types and interfaces.
//
// This package aims to be like the 'image' package, except for audio.
package audio

// Reader is a generic interface which describes any type who can have audio
// samples read from it into an audio buffer.
type Reader interface {
	// Read tries to read into the audio buffer, b, filling it with at max
	// b.Len() audio samples.
	//
	// Returned is the number of samples that where read into the buffer, and
	// an error if any occured.
	//
	// It is possible for the number of samples read to be non-zero; and for an
	// error to be returned at the same time (E.g. read 300 audio samples, but
	// also encountered io.EOF).
	Read(b Buffer) (read int, e error)
}

// ReadSeeker is the generic seekable audio reader interface.
type ReadSeeker interface {
	Reader

	// Seek seeks to the specified sample number, relative to the start of the
	// stream. As such, subsequent Read() calls on the Reader, begin reading at
	// the specified sample.
	//
	// If any error is returned, it means it was impossible to seek to the
	// specified audio sample for some reason, and that the current playhead is
	// unchanged.
	Seek(sample uint64) error
}

// Writer is a generic interface which describes any type who can have audio
// samples written from an audio buffer into it.
type Writer interface {
	// Write attempts to write all, b.Len(), samples in the buffer to the
	// writer.
	//
	// Returned is the number of samples from the buffer that where wrote to
	// the writer, and an error if any occured.
	//
	// The number of samples wrote may be less than buf.Len(), in which case
	// you should subsequently write b.Slice(wrote, b.Len()) until you have
	// finished sending all data or an error occurs.
	//
	// If any error is returned, it should be considered as fatal to the
	// writer, no more data can subsequently be wrote to the writer.
	Write(b Buffer) (wrote int, err error)
}
