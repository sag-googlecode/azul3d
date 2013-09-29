// Test - Plays audio data using aio package.
package main

import(
	_ "code.google.com/p/azul3d/audio/wav"
	"code.google.com/p/azul3d/audio/aio"
	"code.google.com/p/azul3d/audio"
	"time"
	"math"
	"log"
	"os"
)

var sineCount int

func SineWave(b audio.Buffer, frequency float64, c *audio.Config) {
	length := b.Len()

	for i := 0; i < length; i++ {
		sineCount++
		s := math.Sin(frequency * 2.0 * math.Pi * float64(sineCount) / float64(c.SampleRate))
		b.Set(i, audio.F64(s))
	}
}

func main() {
	aio.SetDebugOutput(os.Stdout)
	outputs := aio.Outputs()

	for _, output := range outputs {
		if output.Equals(aio.DefaultOutput()) {
			log.Println("Default -", output)
		} else {
			log.Println(output)
		}
	}

	output := aio.DefaultOutput()
	if output == nil {
		log.Println("No default output device was found.")
		return
	}

	//file, err := os.Open("src/code.google.com/p/azul3d/assets/audio/tune_stereo_44100hz_int16.wav")
	file, err := os.Open("src/code.google.com/p/azul3d/assets/audio/kakariko16.wav")
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

	// To create an buffer that can hold 1 second of audio samples:
	//
	bufSize := 1 * config.SampleRate * config.Channels
	// We'll use a small number here though, to show off low-latency.
	//bufSize := 4000
	output.SetConfig(bufSize, config)
	buf := output.Type.NewBuffer(bufSize / 2)

	playPause := make(chan bool)
	go func() {
		for{
			select{
			case <-playPause:
				log.Println("Pause")
				// We got paused, wait again untill play
				<-playPause
				log.Println("Play")
			default:
				break
			}

			n, rerr := decoder.Read(buf)
			//SineWave(buf, 200, config)

			// n may be <= buf.Len(); so we should take care here.
			//
			// Like: bufToWrite = buf[:n]
			bufToWrite := buf.Slice(0, n)

			// Write bufToWrite into the output until there is no more data left.
			for bufToWrite.Len() > 0 {
				wrote, err := output.Write(bufToWrite)
				if err != nil {
					log.Fatal(err)
				}

				// Like: bufToWrite = bufToWrite[wrote:]
				bufToWrite = bufToWrite.Slice(wrote, bufToWrite.Len())
			}

			if rerr != nil {
				log.Fatal(rerr)
			}
		}
	}()

	for{
		// Wait 6 seconds
		time.Sleep(6 * time.Second)

		// Pause
		playPause <- true

		// Wait 3 second
		time.Sleep(3 * time.Second)

		// Play
		playPause <- true

		// Wait 6 seconds
		time.Sleep(6 * time.Second)
	}
}

