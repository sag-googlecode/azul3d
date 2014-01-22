// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.
// +build tests

// Test - Opens and decodes a wav file.
package main

import(
	_ "azul3d.org/audio/wav"
	"azul3d.org/audio"
	"log"
	"os"
)

func test(fileName string) {
	log.Println(fileName)

	file, err := os.Open("src/azul3d.org/assets/audio/" + fileName)
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

	// Create an buffer that can hold 1 second of audio samples
	bufSize := 1 * config.SampleRate * config.Channels
	buf := make(audio.F64Samples, bufSize)

	// Fill the buffer with as many audio samples as we can
	read, err := decoder.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Read", read, "audio samples.")
	log.Println("")

	// readBuf := buf.Slice(0, read)
	// for i := 0; i < readBuf.Len(); i++ {
	//     sample := readBuf.At(i)
	// }
}

func main() {
	test("tune_stereo_44100hz_alaw.wav")
	test("tune_stereo_44100hz_float32.wav")
	test("tune_stereo_44100hz_float64.wav")
	test("tune_stereo_44100hz_int16.wav")
	test("tune_stereo_44100hz_int24.wav")
	test("tune_stereo_44100hz_int32.wav")
	test("tune_stereo_44100hz_mulaw.wav")
	test("tune_stereo_44100hz_uint8.wav")
}
