// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package audio

import (
	"errors"
	"fmt"
)

// ErrInvalidData represents an error for decoding input data that is invalidd
// or corrupted for some reason.
var ErrInvalidData = errors.New("audio: input data is invalid or corrupt")

// Config represents an audio stream's configuration, like it's sample rate
// and number of interleaved channels.
type Config struct {
	// SampleRate is the number of audio samples that the stream is played or
	// recorded at.
	//
	// E.g. 44100 would be compact disc quality.
	SampleRate int

	// Channels is the number of channels the stream contains.
	Channels int
}

// String returns an string representation of this audio config.
func (c Config) String() string {
	return fmt.Sprintf("Config(SampleRate=%v, Channels=%v)", c.SampleRate, c.Channels)
}

// Decoder is the generic audio decoder interface, for use with the
// RegisterFormat() function.
type Decoder interface {
	ReadSeeker

	// Config returns the audio stream configuration of this decoder. It may
	// block until at least the configuration part of the stream has been read.
	Config() Config
}
