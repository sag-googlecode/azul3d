// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.
// +build tests

// Test application - Opens and decodes a wav file.
package main

import(
	_ "code.google.com/p/azul3d/audio/wav"
	"code.google.com/p/azul3d/audio"
	"log"
	"os"
)

func main() {
	file, err := os.Open("src/code.google.com/p/azul3d/assets/audio/tune_stereo_44100hz_uint8.wav")
	if err != nil {
		log.Fatal(err)
	}

	// Create an decoder for the audio source
	decoder, format, err := audio.NewDecoder(file)
	if err != nil {
		log.Fatal(err)
	}

	// Grab the decoder's configuration
	config := decoder.Config()
	log.Println("Decoding an", format, "file.")
	log.Println(config)

	// Create an buffer that can hold 3 seconds of audio samples
	buf := make([]audio.Sample, 3 * config.SampleRate * config.Channels)

	// Fill the buffer with as many audio samples as we can
	n, err := decoder.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Read", n, "audio samples.")
}
