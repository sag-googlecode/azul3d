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
func (c *Config) String() string {
	return fmt.Sprintf("Config(SampleRate=%v, Channels=%v)", c.SampleRate, c.Channels)
}

// Decoder is the generic audio decoder interface, for use with the
// RegisterFormat() function.
type Decoder interface {
	ReadSeeker

	// Config returns the audio stream configuration of this decoder.
	//
	// This function must never return nil at any point in time. Instead, at
	// creation time of a new decoder, the decoder should block untill at least
	// the stream configuration has been read (or ErrInvalidData could be
	// returned if the decoder does not understand the data in the stream).
	Config() *Config
}
