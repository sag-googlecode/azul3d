package aio

import (
	"azul3d.org/audio"
	"fmt"
)

// Input represents a generic audio input device.
type Input struct {
	*nativeInput

	// Name is the descriptive name of the audio input device.
	Name string

	// The type of audio data the input is able to record.
	AudioType audio.Type
}

func newInput(name string, audioType audio.Type) *Input {
	return &Input{
		nativeInput: new(nativeInput),
		Name:        name,
		AudioType:   audioType,
	}
}

// String returns a string representation of this input.
func (i *Input) String() string {
	return fmt.Sprintf("Input(%q, %v)", i.Name, i.AudioType)
}

// Equals tells if this input is equal to the other input. This is useful due
// to the fact that pointer comparison between inputs is not valid (e.g. after
// a call to the Refresh() function).
func (i *Input) Equals(other *Input) bool {
	return i.nativeInput.equals(other.nativeInput)
}
