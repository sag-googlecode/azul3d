// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

// Package audio implements various audio types and interfaces.
//
// This package aims to be like the 'image' package, except for audio.
package audio

import (
	"errors"
	"fmt"
	"time"
)

// ErrInvalidData represents an error for decoding input data that is invalidd
// or corrupted for some reason.
var ErrInvalidData = errors.New("audio: input data is invalid or corrupt")

// Config represents an audio stream's configuration, like it's sample rate
// and number of interleaved channels.
type Config struct {
	// SampleRate is the number of samples per second at which the audio stream
	// is played or recorded at (I.e. 44100 would be compact disc quality).
	SampleRate int

	// Channels is the number of interleaved channels for the audio stream's
	// data.
	Channels int
}

// String returns an string representation of this audio config.
func (c *Config) String() string {
	return fmt.Sprintf("Config(SampleRate=%v, Channels=%v)", c.SampleRate, c.Channels)
}

// Reader is the generic audio reader interface.
type Reader interface {
	// Read reads at max n audio samples into the internal buffer and returns a
	// slice representing the audio samples.
	//
	// The slice returned is backed by an array that is re-used upon multiple
	// calls to this function (for efficieny). As such, the data in the
	// returned slice is only valid until the next call to Read().
	//
	// The length of the returned slice should be considered before the error,
	// (I.e. the slice may contain audio samples AND io.EOF may be returned.)
	//
	// The behavior of Read(0) is invalid and the result undefined.
	Read(n int) (buf Samples, e error)
}

// ReadSeeker is the generic seekable audio reader interface.
type ReadSeeker interface {
	Reader

	// Seek should seek to the specified point in time, relative to the start.
	Seek(d time.Duration)
}

// Decoder is the generic audio decoder interface, for use with the 
// RegisterFormat() function.
type Decoder interface {
	ReadSeeker

	// Config returns the configuration of this decoder.
	//
	// This function must never return nil (instead, ErrInvalidData should be
	// returned at creation of the new decoder).
	Config() *Config
}
