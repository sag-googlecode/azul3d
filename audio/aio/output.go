package aio

import (
	"azul3d.org/audio"
	"fmt"
	"sync"
)

// Output represents a generic audio output device.
type Output struct {
	*nativeOutput

	access     sync.RWMutex
	bufferSize int
	config     *audio.Config

	// Name is the descriptive name of the audio output device.
	Name string

	// The type of audio data the output is able to play.
	Type audio.Type
}

func newOutput(name string, audioType audio.Type) *Output {
	return &Output{
		nativeOutput: new(nativeOutput),
		Name:         name,
		Type:         audioType,
	}
}

// String returns a string representation of this output.
func (o *Output) String() string {
	return fmt.Sprintf("Output(%q, %v)", o.Name, o.Type)
}

// Equals tells if this output is equal to the other output. This is useful due
// to the fact that pointer comparison between outputs is not valid (e.g. after
// a call to the Refresh() function).
func (o *Output) Equals(other *Output) bool {
	return o.nativeOutput.equals(other.nativeOutput)
}

// SetConfig specifies the buffer size used to write audio samples to the
// output device, as well as the audio configuration that will be used.
func (o *Output) SetConfig(bufferSize int, c *audio.Config) {
	o.access.Lock()
	o.bufferSize = bufferSize
	o.config = c
	o.nativeOutput.setConfig(bufferSize, c)
	o.access.Unlock()
}

// Config returns the buffer size and audio configuration previously set in the
// last call to the SetConfig() function.
func (o *Output) Config() (bufferSize int, c *audio.Config) {
	o.access.RLock()
	defer o.access.RUnlock()

	return o.bufferSize, o.config
}

// Write implements the audio.Writer interface.
func (o *Output) Write(b audio.Buffer) (wrote int, err error) {
	return o.nativeOutput.write(b)
}
