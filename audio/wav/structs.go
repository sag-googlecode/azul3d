// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package wav

type chunkHeader struct {
	// Chunk ID, like "RIFF" or "fmt " in ascii
	ChunkID [4]byte

	// Chunk Size, varies depending on which chunk
	ChunkSize [4]byte
}

type factChunk struct {
	SampleLength [4]byte
}

// the 16-byte 'fmt' chunk
type fmtChunk16 struct {
	// Format code
	FormatTag [2]byte

	// Number of interleaved channels
	Channels [2]byte

	// Sampling rate (blocks per second)
	SamplesPerSec [4]byte

	// Data rate
	AvgBytesPerSec [4]byte

	// Data block size (bytes)
	BlockAlign [2]byte

	// Bits per sample
	BitsPerSample [2]byte
}

// the 18-byte 'fmt' chunk
type fmtChunk18 struct {
	// Size of the extension (0 or 22)
	Size [2]byte
}

// the 40-byte 'fmt' chunk
type fmtChunk40 struct {
	// Number of valid bits
	ValidBitsPerSample [2]byte

	// Speaker position mask
	ChannelMask [4]byte

	// GUID, including the data format code
	SubFormat [16]byte
}